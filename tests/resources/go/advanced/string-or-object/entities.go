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
	AdditionalProperties map[string]interface{} `json:"-"`             // All unmatched properties.
}

type marshalElement Element

var ignoreKeysElement = []string{
	"id",
	"val",
}

// UnmarshalJSON decodes JSON.
func (v *Element) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalElement(*v)

	err = json.Unmarshal(data, &vv)
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
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Element(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Element) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalElement(v))
	}

	return marshalUnion(marshalElement(v), v.AdditionalProperties)
}

// ElementConditional structure is generated from "#/definitions/element".
type ElementConditional struct {
	TypeString *string  `json:"-"`
	Element    *Element `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *ElementConditional) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.TypeString)
	if err != nil {
		v.TypeString = nil
	}

	err = json.Unmarshal(data, &v.Element)
	if err != nil {
		v.Element = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ElementConditional) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.TypeString, v.Element)
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
