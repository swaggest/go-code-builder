// Package entities contains generated structures.
package entities

import (
	"encoding/json"
)

// ObjectOrString structure is generated from "#".
type ObjectOrString struct {
	A string `json:"a"` // Required.
}

type marshalObjectOrString ObjectOrString

// UnmarshalJSON decodes JSON.
func (o *ObjectOrString) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalObjectOrString(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	*o = ObjectOrString(mo)

	return nil
}
