package model

import (
	"encoding/json"
)

/*
DataExpression is a data model to for feed filter data expression
*/
type DataExpressionWrap struct {
	ID   string         `json:"id"`
	Expr ExpressionNode `json:"expr"`
}

// ExpressionNode is a abstract container, it takes/generate the "expr"
type ExpressionNode interface{}

// AllOf is a type of ExpressionNode
type AllOf struct {
	AllOf []DataExpressionWrap `json:"allOf"`
}

// AnyOf is a type of ExpressionNode
type AnyOf struct {
	AnyOf []DataExpressionWrap `json:"anyOf"`
}

// NotTrue is a type of ExpressionNode
type NotTrue struct {
	NotTrue DataExpressionWrap `json:"notTrue"`
}

// PredicateWrap is a type of ExpressionNode
type PredicateWrap struct {
	Predicate Predicate `json:"pred"`
}

// Predicate is a type of ExpressionNode
type Predicate struct {
	Type  string  `json:"type"`
	Param Literal `json:"param"`
}

type Literal struct {
	Text string `json:"text"`
}

type DataExpressionRoot struct {
	Root DataExpressionWrap `json:"dataExpression"`
}

// Custom unmarshal function for DataExpressionWrap
// since DataExpressionWrap contains interface ExpressionNode
// which needs "look-ahead" into next level
// in order to decide what type to unmarshal
func (target *DataExpressionWrap) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(*objMap["id"], &target.ID); err != nil {
		return err
	}

	// Look ahead into the next level keys
	// Then check if key includes certain type
	// Then use this type to unmarshal
	var expr map[string]*json.RawMessage
	if err = json.Unmarshal(*objMap["expr"], &expr); err != nil {
		return err
	}

	if val, ok := expr["allOf"]; ok {
		var tmp []*json.RawMessage
		if err = json.Unmarshal(*val, &tmp); err != nil {
			return err
		}
		node := AllOf{AllOf: []DataExpressionWrap{}}
		for _, t := range tmp {
			var tt DataExpressionWrap
			if err = json.Unmarshal(*t, &tt); err != nil {
				return err
			}
			node.AllOf = append(node.AllOf, tt)
		}
		target.Expr = node
	} else if val, ok := expr["anyOf"]; ok {
		var tmp []*json.RawMessage
		if err = json.Unmarshal(*val, &tmp); err != nil {
			return err
		}
		node := AnyOf{AnyOf: []DataExpressionWrap{}}
		for _, t := range tmp {
			var tt DataExpressionWrap
			if err = json.Unmarshal(*t, &tt); err != nil {
				return err
			}
			node.AnyOf = append(node.AnyOf, tt)
		}
		target.Expr = node
	} else if val, ok := expr["notTrue"]; ok {
		var node NotTrue
		if err = json.Unmarshal(*val, &node.NotTrue); err != nil {
			return err
		}
		target.Expr = node
	} else if val, ok := expr["pred"]; ok {
		var node PredicateWrap
		if err = json.Unmarshal(*val, &node.Predicate); err != nil {
			return err
		}
		target.Expr = node
	}
	return nil
}
