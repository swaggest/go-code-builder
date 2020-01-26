// Package entities contains generated structures.
package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

// Properties structure is generated from "#".
type Properties struct {
	RefOrSchema          *RefOrSchema           `json:"refOrSchema,omitempty"`
	NoTypeWithExamples   *NoTypeWithExamples    `json:"noTypeWithExamples,omitempty"`
	NoTypeWithExample    *NoTypeWithExample     `json:"noTypeWithExample,omitempty"`
	Address              *Address               `json:"address,omitempty"`
	Headers              *Table                 `json:"headers,omitempty"`
	ContentType          *ShortStr              `json:"content-type,omitempty"`
	ContentEncoding      *ShortStr              `json:"content-encoding,omitempty"`
	DeliveryMode         *PropertyOctet         `json:"delivery-mode,omitempty"`
	Priority             *PropertyOctet         `json:"priority,omitempty"`
	CorrelationID        *ShortStr              `json:"correlation-id,omitempty"`
	ReplyTo              *ShortStr              `json:"reply-to,omitempty"`
	Expiration           *ShortStr              `json:"expiration,omitempty"`
	MessageID            *ShortStr              `json:"message-id,omitempty"`
	Timestamp            *Timestamp             `json:"timestamp,omitempty"`
	Type                 *ShortStr              `json:"type,omitempty"`
	UserID               *ShortStr              `json:"user-id,omitempty"`
	AppID                *ShortStr              `json:"app-id,omitempty"`
	MapOfAnything        map[string]interface{} `json:"-"`                            // Key must match pattern: ^x-
	AdditionalProperties map[string]Property    `json:"-"`                            // All unmatched properties
}

type marshalProperties Properties

var ignoreKeysProperties = []string{
	"refOrSchema",
	"noTypeWithExamples",
	"noTypeWithExample",
	"address",
	"headers",
	"content-type",
	"content-encoding",
	"delivery-mode",
	"priority",
	"correlation-id",
	"reply-to",
	"expiration",
	"message-id",
	"timestamp",
	"type",
	"user-id",
	"app-id",
}

// UnmarshalJSON decodes JSON.
func (i *Properties) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalProperties(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysProperties {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]Property, 1)
		}

		var val Property

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = Properties(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Properties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProperties(i), i.MapOfAnything, i.AdditionalProperties)
}

// Reference structure is generated from "#/definitions/reference".
type Reference struct {
	Ref                  string                 `json:"$ref,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties
}

type marshalReference Reference

var ignoreKeysReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (i *Reference) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalReference(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysReference {
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

	*i = Reference(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Reference) MarshalJSON() ([]byte, error) {
    if len(i.AdditionalProperties) == 0 {
        return json.Marshal(marshalReference(i))
    }
	return marshalUnion(marshalReference(i), i.AdditionalProperties)
}

// RefOrSchema structure is generated from "#/definitions/refOrSchema".
type RefOrSchema struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *RefOrSchema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i RefOrSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Schema)
}

// NoTypeWithExamples structure is generated from "#/definitions/noTypeWithExamples".
type NoTypeWithExamples struct {
	FirstName            string                 `json:"firstName,omitempty"`
	LastName             string                 `json:"lastName,omitempty"`
	Age                  *int64                 `json:"age"`
	Gender               string                 `json:"gender,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties
}

type marshalNoTypeWithExamples NoTypeWithExamples

var ignoreKeysNoTypeWithExamples = []string{
	"firstName",
	"lastName",
	"age",
	"gender",
}

// UnmarshalJSON decodes JSON.
func (i *NoTypeWithExamples) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalNoTypeWithExamples(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysNoTypeWithExamples {
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

	*i = NoTypeWithExamples(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i NoTypeWithExamples) MarshalJSON() ([]byte, error) {
    if len(i.AdditionalProperties) == 0 {
        return json.Marshal(marshalNoTypeWithExamples(i))
    }
	return marshalUnion(marshalNoTypeWithExamples(i), i.AdditionalProperties)
}

// NoTypeWithExample structure is generated from "#/definitions/noTypeWithExample".
type NoTypeWithExample struct {
	FirstName            string                 `json:"firstName,omitempty"`
	LastName             string                 `json:"lastName,omitempty"`
	Age                  int64                  `json:"age,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties
}

type marshalNoTypeWithExample NoTypeWithExample

var ignoreKeysNoTypeWithExample = []string{
	"firstName",
	"lastName",
	"age",
}

// UnmarshalJSON decodes JSON.
func (i *NoTypeWithExample) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalNoTypeWithExample(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysNoTypeWithExample {
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

	*i = NoTypeWithExample(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i NoTypeWithExample) MarshalJSON() ([]byte, error) {
    if len(i.AdditionalProperties) == 0 {
        return json.Marshal(marshalNoTypeWithExample(i))
    }
	return marshalUnion(marshalNoTypeWithExample(i), i.AdditionalProperties)
}

// Address structure is generated from "#/definitions/address".
type Address struct {
	Stripped             string                 `json:"addressStripped,omitempty"`
	Address1             string                 `json:"address1,omitempty"`
	Address2             string                 `json:"address2,omitempty"`
	City                 string                 `json:"city,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                         // All unmatched properties
}

type marshalAddress Address

var ignoreKeysAddress = []string{
	"addressStripped",
	"address1",
	"address2",
	"city",
}

// UnmarshalJSON decodes JSON.
func (i *Address) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAddress(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAddress {
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

	*i = Address(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Address) MarshalJSON() ([]byte, error) {
    if len(i.AdditionalProperties) == 0 {
        return json.Marshal(marshalAddress(i))
    }
	return marshalUnion(marshalAddress(i), i.AdditionalProperties)
}

// Table structure is generated from "#/definitions/table".
type Table struct {
	Value map[string]Property `json:"value,omitempty"`
}

type marshalTable Table

// UnmarshalJSON decodes JSON.
func (i *Table) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalTable(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"table"` {
		return fmt.Errorf(`bad or missing const value for "type" ("table" expected, %s received)`, v)
	}

	delete(m, "type")

	*i = Table(ii)

	return nil
}

var (
	// constTable is unconditionally added to JSON.
	constTable = json.RawMessage(`{"type":"table"}`)
)

// MarshalJSON encodes JSON.
func (i Table) MarshalJSON() ([]byte, error) {
	return marshalUnion(constTable, marshalTable(i))
}

// Scalar structure is generated from "#/definitions/scalar".
type Scalar struct {
	Type  StringedType `json:"type,omitempty"`
	Value string       `json:"value,omitempty"`
}

// Property structure is generated from "#/definitions/property".
type Property struct {
	Scalar *Scalar `json:"-"`
	Table  *Table  `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Property) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Scalar)
	if err != nil {
		i.Scalar = nil
	}

	err = json.Unmarshal(data, &i.Table)
	if err != nil {
		i.Table = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Property) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Scalar, i.Table)
}

// ShortStr structure is generated from "#/definitions/shortStr".
type ShortStr struct {
	Value string `json:"value,omitempty"`
}

type marshalShortStr ShortStr

// UnmarshalJSON decodes JSON.
func (i *ShortStr) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalShortStr(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"shortstr"` {
		return fmt.Errorf(`bad or missing const value for "type" ("shortstr" expected, %s received)`, v)
	}

	delete(m, "type")

	*i = ShortStr(ii)

	return nil
}

var (
	// constShortStr is unconditionally added to JSON.
	constShortStr = json.RawMessage(`{"type":"shortstr"}`)
)

// MarshalJSON encodes JSON.
func (i ShortStr) MarshalJSON() ([]byte, error) {
	return marshalUnion(constShortStr, marshalShortStr(i))
}

// PropertyOctet structure is generated from "#/definitions/propertyOctet".
type PropertyOctet struct {
	Value string `json:"value,omitempty"`
}

type marshalPropertyOctet PropertyOctet

// UnmarshalJSON decodes JSON.
func (i *PropertyOctet) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalPropertyOctet(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"octet"` {
		return fmt.Errorf(`bad or missing const value for "type" ("octet" expected, %s received)`, v)
	}

	delete(m, "type")

	*i = PropertyOctet(ii)

	return nil
}

var (
	// constPropertyOctet is unconditionally added to JSON.
	constPropertyOctet = json.RawMessage(`{"type":"octet"}`)
)

// MarshalJSON encodes JSON.
func (i PropertyOctet) MarshalJSON() ([]byte, error) {
	return marshalUnion(constPropertyOctet, marshalPropertyOctet(i))
}

// Timestamp structure is generated from "#/definitions/timestamp".
type Timestamp struct {
	Value string `json:"value,omitempty"`
}

type marshalTimestamp Timestamp

// UnmarshalJSON decodes JSON.
func (i *Timestamp) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalTimestamp(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"timestamp"` {
		return fmt.Errorf(`bad or missing const value for "type" ("timestamp" expected, %s received)`, v)
	}

	delete(m, "type")

	*i = Timestamp(ii)

	return nil
}

var (
	// constTimestamp is unconditionally added to JSON.
	constTimestamp = json.RawMessage(`{"type":"timestamp"}`)
)

// MarshalJSON encodes JSON.
func (i Timestamp) MarshalJSON() ([]byte, error) {
	return marshalUnion(constTimestamp, marshalTimestamp(i))
}

// StringedType is an enum type.
type StringedType string

// StringedType values enumeration.
const (
	StringedTypeBit = StringedType("bit")
	StringedTypeOctet = StringedType("octet")
	StringedTypeShort = StringedType("short")
	StringedTypeLong = StringedType("long")
	StringedTypeLonglong = StringedType("longlong")
	StringedTypeShortstr = StringedType("shortstr")
	StringedTypeLongstr = StringedType("longstr")
	StringedTypeTimestamp = StringedType("timestamp")
)

// MarshalJSON encodes JSON.
func (i StringedType) MarshalJSON() ([]byte, error) {
	switch i {
	case StringedTypeBit:
	case StringedTypeOctet:
	case StringedTypeShort:
	case StringedTypeLong:
	case StringedTypeLonglong:
	case StringedTypeShortstr:
	case StringedTypeLongstr:
	case StringedTypeTimestamp:

	default:
		return nil, fmt.Errorf("unexpected StringedType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *StringedType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := StringedType(ii)

	switch v {
	case StringedTypeBit:
	case StringedTypeOctet:
	case StringedTypeShort:
	case StringedTypeLong:
	case StringedTypeLonglong:
	case StringedTypeShortstr:
	case StringedTypeLongstr:
	case StringedTypeTimestamp:

	default:
		return fmt.Errorf("unexpected StringedType value: %v", v)
	}

	*i = v

	return nil
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
// Regular expressions for pattern properties.
var (
	regexX = regexp.MustCompile("^x-")
)
