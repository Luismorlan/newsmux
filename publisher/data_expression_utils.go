package publisher

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/Luismorlan/newsmux/model"
)

// TODO(jamie): optimize by first parsing json and match later
// TODO(jamie): should probably create a in-memory cache to avoid constant
// parsing the jsonStr into data expression because such kind of parsing is
// expensive.
func DataExpressionMatchPost(jsonStr string, post model.Post) (bool, error) {
	var dataExpressionWrap model.DataExpressionWrap
	json.Unmarshal([]byte(jsonStr), &dataExpressionWrap)
	return DataExpressionMatch(dataExpressionWrap, post)
}

func DataExpressionMatch(dataExpressionWrap model.DataExpressionWrap, post model.Post) (bool, error) {
	// Empty data expression should match all post.
	if dataExpressionWrap.IsEmpty() {
		return true, nil
	}
	switch expr := dataExpressionWrap.Expr.(type) {
	case model.AllOf:
		if len(expr.AllOf) == 0 {
			return true, nil
		}
		for _, child := range expr.AllOf {
			match, err := DataExpressionMatch(child, post)
			if err != nil {
				return false, err
			}
			if !match {
				return false, nil
			}
		}
		return true, nil
	case model.AnyOf:
		if len(expr.AnyOf) == 0 {
			return true, nil
		}
		for _, child := range expr.AnyOf {
			match, err := DataExpressionMatch(child, post)
			if err != nil {
				return false, err
			}
			if match {
				return true, nil
			}
		}
		return false, nil
	case model.NotTrue:
		match, err := DataExpressionMatch(expr.NotTrue, post)
		if err != nil {
			return false, err
		}
		return !match, nil
	case model.PredicateWrap:
		if expr.Predicate.Type == "LITERAL" {
			return strings.Contains(post.Content, expr.Predicate.Param.Text), nil
		}
	default:
		return false, errors.New("unknown node type when matching data expression")
	}
	return false, nil
}
