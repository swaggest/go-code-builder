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
	MapOfAnything        map[string]interface{} `json:"-"`                            // Key must match pattern: `^x-`.
	AdditionalProperties map[string]Property    `json:"-"`                            // All unmatched properties.
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
func (p *Properties) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalProperties(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysProperties {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mp.MapOfAnything == nil {
				mp.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if mp.AdditionalProperties == nil {
			mp.AdditionalProperties = make(map[string]Property, 1)
		}

		var val Property

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mp.AdditionalProperties[key] = val
	}

	*p = Properties(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p Properties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProperties(p), p.MapOfAnything, p.AdditionalProperties)
}

// Reference structure is generated from "#/definitions/reference".
type Reference struct {
	// Format: uri-reference.
	// Required.
	Ref                  string                 `json:"$ref"`
	AdditionalProperties map[string]interface{} `json:"-"`    // All unmatched properties.
}

type marshalReference Reference

var ignoreKeysReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (r *Reference) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalReference(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysReference {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mr.AdditionalProperties == nil {
			mr.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mr.AdditionalProperties[key] = val
	}

	*r = Reference(mr)

	return nil
}

// MarshalJSON encodes JSON.
func (r Reference) MarshalJSON() ([]byte, error) {
	if len(r.AdditionalProperties) == 0 {
		return json.Marshal(marshalReference(r))
	}

	return marshalUnion(marshalReference(r), r.AdditionalProperties)
}

// RefOrSchema structure is generated from "#/definitions/refOrSchema".
type RefOrSchema struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (r *RefOrSchema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &r.Reference)
	if err != nil {
		r.Reference = nil
	}

	err = json.Unmarshal(data, &r.Schema)
	if err != nil {
		r.Schema = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (r RefOrSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(r.Reference, r.Schema)
}

// NoTypeWithExamples structure is generated from "#/definitions/noTypeWithExamples".
type NoTypeWithExamples struct {
	FirstName            string                 `json:"firstName,omitempty"`
	LastName             string                 `json:"lastName,omitempty"`
	Age                  *int64                 `json:"age"`
	Gender               string                 `json:"gender,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties.
}

type marshalNoTypeWithExamples NoTypeWithExamples

var ignoreKeysNoTypeWithExamples = []string{
	"firstName",
	"lastName",
	"age",
	"gender",
}

// UnmarshalJSON decodes JSON.
func (n *NoTypeWithExamples) UnmarshalJSON(data []byte) error {
	var err error

	mn := marshalNoTypeWithExamples(*n)

	err = json.Unmarshal(data, &mn)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysNoTypeWithExamples {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mn.AdditionalProperties == nil {
			mn.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mn.AdditionalProperties[key] = val
	}

	*n = NoTypeWithExamples(mn)

	return nil
}

// MarshalJSON encodes JSON.
func (n NoTypeWithExamples) MarshalJSON() ([]byte, error) {
	if len(n.AdditionalProperties) == 0 {
		return json.Marshal(marshalNoTypeWithExamples(n))
	}

	return marshalUnion(marshalNoTypeWithExamples(n), n.AdditionalProperties)
}

// NoTypeWithExample structure is generated from "#/definitions/noTypeWithExample".
type NoTypeWithExample struct {
	FirstName            string                 `json:"firstName,omitempty"`
	LastName             string                 `json:"lastName,omitempty"`
	Age                  int64                  `json:"age,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties.
}

type marshalNoTypeWithExample NoTypeWithExample

var ignoreKeysNoTypeWithExample = []string{
	"firstName",
	"lastName",
	"age",
}

// UnmarshalJSON decodes JSON.
func (n *NoTypeWithExample) UnmarshalJSON(data []byte) error {
	var err error

	mn := marshalNoTypeWithExample(*n)

	err = json.Unmarshal(data, &mn)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysNoTypeWithExample {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mn.AdditionalProperties == nil {
			mn.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mn.AdditionalProperties[key] = val
	}

	*n = NoTypeWithExample(mn)

	return nil
}

// MarshalJSON encodes JSON.
func (n NoTypeWithExample) MarshalJSON() ([]byte, error) {
	if len(n.AdditionalProperties) == 0 {
		return json.Marshal(marshalNoTypeWithExample(n))
	}

	return marshalUnion(marshalNoTypeWithExample(n), n.AdditionalProperties)
}

// Address structure is generated from "#/definitions/address".
type Address struct {
	Stripped             string                 `json:"addressStripped,omitempty"`
	Address1             string                 `json:"address1,omitempty"`
	Address2             string                 `json:"address2,omitempty"`
	City                 string                 `json:"city,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                         // All unmatched properties.
}

type marshalAddress Address

var ignoreKeysAddress = []string{
	"addressStripped",
	"address1",
	"address2",
	"city",
}

// UnmarshalJSON decodes JSON.
func (a *Address) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAddress(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysAddress {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if ma.AdditionalProperties == nil {
			ma.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ma.AdditionalProperties[key] = val
	}

	*a = Address(ma)

	return nil
}

// MarshalJSON encodes JSON.
func (a Address) MarshalJSON() ([]byte, error) {
	if len(a.AdditionalProperties) == 0 {
		return json.Marshal(marshalAddress(a))
	}

	return marshalUnion(marshalAddress(a), a.AdditionalProperties)
}

// Table structure is generated from "#/definitions/table".
type Table struct {
	Value map[string]Property `json:"value"` // Required.
}

type marshalTable Table

// UnmarshalJSON decodes JSON.
func (t *Table) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTable(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"table"` {
		return fmt.Errorf(`bad or missing const value for "type" ("table" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	*t = Table(mt)

	return nil
}

var (
	// constTable is unconditionally added to JSON.
	constTable = json.RawMessage(`{"type":"table"}`)
)

// MarshalJSON encodes JSON.
func (t Table) MarshalJSON() ([]byte, error) {
	return marshalUnion(constTable, marshalTable(t))
}

// Scalar structure is generated from "#/definitions/scalar".
type Scalar struct {
	Type  StringedType `json:"type"`  // Required.
	Value string       `json:"value"` // Required.
}

// Property structure is generated from "#/definitions/property".
type Property struct {
	Scalar *Scalar `json:"-"`
	Table  *Table  `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (p *Property) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &p.Scalar)
	if err != nil {
		p.Scalar = nil
	}

	err = json.Unmarshal(data, &p.Table)
	if err != nil {
		p.Table = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (p Property) MarshalJSON() ([]byte, error) {
	return marshalUnion(p.Scalar, p.Table)
}

// ShortStr structure is generated from "#/definitions/shortStr".
type ShortStr struct {
	Value string `json:"value"` // Required.
}

type marshalShortStr ShortStr

// UnmarshalJSON decodes JSON.
func (s *ShortStr) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalShortStr(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"shortstr"` {
		return fmt.Errorf(`bad or missing const value for "type" ("shortstr" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	*s = ShortStr(ms)

	return nil
}

var (
	// constShortStr is unconditionally added to JSON.
	constShortStr = json.RawMessage(`{"type":"shortstr"}`)
)

// MarshalJSON encodes JSON.
func (s ShortStr) MarshalJSON() ([]byte, error) {
	return marshalUnion(constShortStr, marshalShortStr(s))
}

// PropertyOctet structure is generated from "#/definitions/propertyOctet".
type PropertyOctet struct {
	Value string `json:"value"` // Required.
}

type marshalPropertyOctet PropertyOctet

// UnmarshalJSON decodes JSON.
func (p *PropertyOctet) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalPropertyOctet(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"octet"` {
		return fmt.Errorf(`bad or missing const value for "type" ("octet" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	*p = PropertyOctet(mp)

	return nil
}

var (
	// constPropertyOctet is unconditionally added to JSON.
	constPropertyOctet = json.RawMessage(`{"type":"octet"}`)
)

// MarshalJSON encodes JSON.
func (p PropertyOctet) MarshalJSON() ([]byte, error) {
	return marshalUnion(constPropertyOctet, marshalPropertyOctet(p))
}

// Timestamp structure is generated from "#/definitions/timestamp".
type Timestamp struct {
	Value string `json:"value"` // Required.
}

type marshalTimestamp Timestamp

// UnmarshalJSON decodes JSON.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTimestamp(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"timestamp"` {
		return fmt.Errorf(`bad or missing const value for "type" ("timestamp" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	*t = Timestamp(mt)

	return nil
}

var (
	// constTimestamp is unconditionally added to JSON.
	constTimestamp = json.RawMessage(`{"type":"timestamp"}`)
)

// MarshalJSON encodes JSON.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return marshalUnion(constTimestamp, marshalTimestamp(t))
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
