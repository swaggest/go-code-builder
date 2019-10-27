// Package entities contains generated structures.
package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
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

// UnmarshalJSON decodes JSON.
func (i *Properties) UnmarshalJSON(data []byte) error {
	ii := marshalProperties(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Properties(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Reference) UnmarshalJSON(data []byte) error {
	ii := marshalReference(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$ref",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Reference(ii)

	return err
}

// MarshalJSON encodes JSON.
func (i Reference) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalReference(i), i.AdditionalProperties)
}

// RefOrSchema structure is generated from "#/definitions/refOrSchema".
type RefOrSchema struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *RefOrSchema) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Reference, &i.Schema}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Schema = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *NoTypeWithExamples) UnmarshalJSON(data []byte) error {
	ii := marshalNoTypeWithExamples(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"firstName",
			"lastName",
			"age",
			"gender",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = NoTypeWithExamples(ii)

	return err
}

// MarshalJSON encodes JSON.
func (i NoTypeWithExamples) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON decodes JSON.
func (i *NoTypeWithExample) UnmarshalJSON(data []byte) error {
	ii := marshalNoTypeWithExample(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"firstName",
			"lastName",
			"age",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = NoTypeWithExample(ii)

	return err
}

// MarshalJSON encodes JSON.
func (i NoTypeWithExample) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON decodes JSON.
func (i *Address) UnmarshalJSON(data []byte) error {
	ii := marshalAddress(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"addressStripped",
			"address1",
			"address2",
			"city",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Address(ii)

	return err
}

// MarshalJSON encodes JSON.
func (i Address) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAddress(i), i.AdditionalProperties)
}

// Table structure is generated from "#/definitions/table".
type Table struct {
	Value map[string]Property `json:"value,omitempty"`
}

type marshalTable Table

// UnmarshalJSON decodes JSON.
func (i *Table) UnmarshalJSON(data []byte) error {
	ii := marshalTable(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"table"` {
		return fmt.Errorf(`bad or missing const value for "type" ("table" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = Table(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Scalar, &i.Table}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Scalar = nil
	}

	if mayUnmarshal[1] == nil {
		i.Table = nil
	}

	return err
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
	ii := marshalShortStr(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"shortstr"` {
		return fmt.Errorf(`bad or missing const value for "type" ("shortstr" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = ShortStr(ii)

	return err
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
	ii := marshalPropertyOctet(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"octet"` {
		return fmt.Errorf(`bad or missing const value for "type" ("octet" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = PropertyOctet(ii)

	return err
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
	ii := marshalTimestamp(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"timestamp"` {
		return fmt.Errorf(`bad or missing const value for "type" ("timestamp" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = Timestamp(ii)

	return err
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

type unionMap struct {
	mustUnmarshal        []interface{}
	mayUnmarshal         []interface{}
	ignoreKeys           []string
	patternProperties    map[*regexp.Regexp]interface{}
	additionalProperties interface{}
	jsonData             []byte
}

func (u unionMap) unmarshal() error {
	for _, item := range u.mustUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			return err
		}
	}

	for i, item := range u.mayUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			u.mayUnmarshal[i] = nil
		}
	}

	if len(u.patternProperties) == 0 && u.additionalProperties == nil {
		return nil
	}

	// Unmarshal to a generic map.
	var m map[string]*json.RawMessage

	err := json.Unmarshal(u.jsonData, &m)
	if err != nil {
		return err
	}

	// Remove ignored keys (defined in struct).
	for _, i := range u.ignoreKeys {
		delete(m, i)
	}

	// Return early on empty map.
	if len(m) == 0 {
		return nil
	}

	if len(u.patternProperties) != 0 {
		err = u.unmarshalPatternProperties(m)
		if err != nil {
			return err
		}
	}

	// Return early on empty map.
	if len(m) == 0 {
		return nil
	}

	if u.additionalProperties != nil {
		return u.unmarshalAdditionalProperties(m)
	}

	return nil
}

func (u unionMap) unmarshalAdditionalProperties(m map[string]*json.RawMessage) error {
	var err error

	subMap := make([]byte, 1, 100)

	subMap[0] = '{'

	// Iterating map and filling additional properties.
	for key, val := range m {
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		if len(subMap) != 1 {
			subMap = append(subMap[:len(subMap)-1], ',')
		}

		subMap = append(subMap, []byte(keyEscaped)...)

		if val != nil {
			subMap = append(subMap, []byte(*val)...)
		} else {
			subMap = append(subMap, []byte("null")...)
		}

		subMap = append(subMap, '}')
	}

	if len(subMap) > 1 {
		err = json.Unmarshal(subMap, u.additionalProperties)
		if err != nil {
			return err
		}
	}

	return nil
}
func (u unionMap) unmarshalPatternProperties(m map[string]*json.RawMessage) error {
	patternMapsRaw := make(map[*regexp.Regexp][]byte, len(u.patternProperties))

	// Iterating map and filling pattern properties sub maps.
	for key, val := range m {
		matched := false
		ok := false
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex := range u.patternProperties {
			if regex.MatchString(key) {
				matched = true

				var subMap []byte

				if subMap, ok = patternMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)

				if val != nil {
					subMap = append(subMap, []byte(*val)...)
				} else {
					subMap = append(subMap, []byte("null")...)
				}

				subMap = append(subMap, '}')

				patternMapsRaw[regex] = subMap
			}
		}

		// Remove from properties map if matched to at least one regex.
		if matched {
			delete(m, key)
		}
	}

	for regex := range u.patternProperties {
		if subMap, ok := patternMapsRaw[regex]; !ok {
			continue
		} else {
			err := json.Unmarshal(subMap, u.patternProperties[regex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
