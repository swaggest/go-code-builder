// Package entities contains generated structures.
package entities

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Element structure is generated from "#/definitions/element[object]".
type Element struct {
	ID                   string                 `json:"id,omitempty"`
	Val                  int64                  `json:"val,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`             // All unmatched properties
}

type marshalElement Element

var ignoreKeysElement = []string{
	"id",
	"val",
}

// UnmarshalJSON decodes JSON.
func (i *Element) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalElement(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysElement {
		delete(m, key)
	}

	for key, rawValue := range m {
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = Element(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Element) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalElement(i))
	}
	return marshalUnion(marshalElement(i), i.AdditionalProperties)
}

// ElementConditional structure is generated from "#/definitions/element".
type ElementConditional struct {
	Element    *Element `json:"-"`
	TypeString *string  `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ElementConditional) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Element)
	if err != nil {
		i.Element = nil
	}

	err = json.Unmarshal(data, &i.TypeString)
	if err != nil {
		i.TypeString = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ElementConditional) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Element, i.TypeString)
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
