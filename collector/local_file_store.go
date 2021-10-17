package collector

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Luismorlan/newsmux/utils"
)

const (
	TmpFileDirPrefix = "_tmp_file_store_"
)

type LocalFileStore struct {
	bucket                    string
	processUrlBeforeFetchFunc ProcessUrlBeforeFetchFuncType
	customizeFileNameFunc     CustomizeFileNameFuncType
	customizeFileExtFunc      CustomizeFileExtFuncType
	folderName                string
}

func NewLocalFileStore(bucket string) (*LocalFileStore, error) {
	folderName, err := CreateFolder(bucket)
	if err != nil {
		return nil, err
	}

	return &LocalFileStore{
		bucket:                    bucket,
		processUrlBeforeFetchFunc: func(s string) string { return s },
		customizeFileNameFunc:     nil,
		customizeFileExtFunc:      nil,
		folderName:                folderName,
	}, nil
}

func CreateFolder(bucket string) (string, error) {
	folderName := TmpFileDirPrefix + bucket
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil && strings.Contains(err.Error(), "file exists") {
		return folderName, nil
	}
	return folderName, err
}

func DeleteFolder(folderName string) error {
	return os.RemoveAll(folderName)
}

func (s *LocalFileStore) CleanUp() {
	DeleteFolder(s.folderName)
}

func (s *LocalFileStore) SetProcessUrlBeforeFetchFunc(f ProcessUrlBeforeFetchFuncType) *LocalFileStore {
	s.processUrlBeforeFetchFunc = f
	return s
}

func (s *LocalFileStore) SetCustomizeFileNameFunc(f CustomizeFileNameFuncType) *LocalFileStore {
	s.customizeFileNameFunc = f
	return s
}

func (s *LocalFileStore) SetCustomizeFileExtFunc(f CustomizeFileExtFuncType) *LocalFileStore {
	s.customizeFileExtFunc = f
	return s
}

func (s *LocalFileStore) GenerateFileNameFromUrl(url string) (key string, err error) {
	if s.customizeFileNameFunc != nil {
		key = s.customizeFileNameFunc(url)
	} else {
		key, err = utils.TextToMd5Hash(url)
	}

	if len(key) == 0 {
		err = errors.New("generate empty s3 key, invalid")
	}

	if s.customizeFileExtFunc != nil {
		key = key + "." + s.customizeFileExtFunc(url)
	} else {
		key = key + utils.GetUrlExtNameWithDot(url)
	}

	return key, err
}

func (s *LocalFileStore) FetchAndStore(url string) (string, error) {
	// Download file to local mainly for testing
	response, err := http.Get(s.processUrlBeforeFetchFunc(url))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	fileName, err := s.GenerateFileNameFromUrl(url)
	localPath := filepath.Join(s.folderName, fileName)

	//open a file for writing
	file, err := os.Create(localPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return fileName, err
}

func (s *LocalFileStore) GetUrlFromKey(key string) string {
	return fmt.Sprintf("local store file : %s", key)
}
