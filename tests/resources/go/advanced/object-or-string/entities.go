// Package entities contains generated structures.
package entities

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ElementObject structure is generated from "#/definitions/element[object]".
type ElementObject struct {
	ID                   string                 `json:"id,omitempty"`
	Val                  int64                  `json:"val,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`             // All unmatched properties.
}

type marshalElementObject ElementObject

var ignoreKeysElementObject = []string{
	"id",
	"val",
}

// UnmarshalJSON decodes JSON.
func (e *ElementObject) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalElementObject(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysElementObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if me.AdditionalProperties == nil {
			me.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		me.AdditionalProperties[key] = val
	}

	*e = ElementObject(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e ElementObject) MarshalJSON() ([]byte, error) {
	if len(e.AdditionalProperties) == 0 {
		return json.Marshal(marshalElementObject(e))
	}

	return marshalUnion(marshalElementObject(e), e.AdditionalProperties)
}

// Element structure is generated from "#/definitions/element".
type Element struct {
	Element    *ElementObject `json:"-"`
	TypeString *string        `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (e *Element) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &e.Element)
	if err != nil {
		e.Element = nil
	}

	err = json.Unmarshal(data, &e.TypeString)
	if err != nil {
		e.TypeString = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (e Element) MarshalJSON() ([]byte, error) {
	return marshalUnion(e.Element, e.TypeString)
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	isObject := true

	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		if string(j) == "{}" {
			continue
		}

		if string(j) == "null" {
			continue
		}

		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false

				continue
			}

			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}

		result = append(result, j[1:]...)
	}

	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
