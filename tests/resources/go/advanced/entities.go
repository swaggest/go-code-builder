package entities

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// Properties structure is generated from #
type Properties struct {
	Headers              *Table              `json:"headers,omitempty"`
	ContentType          *ShortStr           `json:"content-type,omitempty"`
	ContentEncoding      *ShortStr           `json:"content-encoding,omitempty"`
	DeliveryMode         *PropertyOctet      `json:"delivery-mode,omitempty"`
	Priority             *PropertyOctet      `json:"priority,omitempty"`
	CorrelationID        *ShortStr           `json:"correlation-id,omitempty"`
	ReplyTo              *ShortStr           `json:"reply-to,omitempty"`
	Expiration           *ShortStr           `json:"expiration,omitempty"`
	MessageID            *ShortStr           `json:"message-id,omitempty"`
	Timestamp            *Timestamp          `json:"timestamp,omitempty"`
	Type                 *ShortStr           `json:"type,omitempty"`
	UserID               *ShortStr           `json:"user-id,omitempty"`
	AppID                *ShortStr           `json:"app-id,omitempty"`
	AdditionalProperties map[string]Property `json:"-"`
}

type marshalProperties Properties

// UnmarshalJSON decodes JSON
func (i *Properties) UnmarshalJSON(data []byte) error {
	ii := marshalProperties(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
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
		map[string]interface{}{
			"": &ii.AdditionalProperties,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Properties(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Properties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProperties(i), i.AdditionalProperties)
}

// Table structure is generated from #/definitions/table
type Table struct {
	Value map[string]Property `json:"value,omitempty"`
}

type marshalTable Table

// UnmarshalJSON decodes JSON
func (i *Table) UnmarshalJSON(data []byte) error {
	ii := marshalTable(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"value",
		},
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"table"` {
	    return errors.New(`bad or missing const value for "type" ("table" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Table(ii)
	return err
}

var (
	// constTable is unconditionally added to JSON
	constTable = json.RawMessage(`{"type":"table"}`)
)

// MarshalJSON encodes JSON
func (i Table) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTable(i), constTable)
}

// Scalar structure is generated from #/definitions/scalar
type Scalar struct {
	Type  StringedType `json:"type,omitempty"`
	Value string       `json:"value,omitempty"`
}

// Property structure is generated from #/definitions/property
type Property struct {
	Scalar *Scalar `json:"-"`
	Table  *Table  `json:"-"`
}

type marshalProperty Property

// UnmarshalJSON decodes JSON
func (i *Property) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Scalar, &i.Table}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Scalar = nil
	}
	if mayUnmarshal[1] == nil {
		i.Table = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Property) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProperty(i), i.Scalar, i.Table)
}

// ShortStr structure is generated from #/definitions/shortStr
type ShortStr struct {
	Value string `json:"value,omitempty"`
}

type marshalShortStr ShortStr

// UnmarshalJSON decodes JSON
func (i *ShortStr) UnmarshalJSON(data []byte) error {
	ii := marshalShortStr(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"value",
		},
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"shortstr"` {
	    return errors.New(`bad or missing const value for "type" ("shortstr" expected)`)
	}
	if err != nil {
		return err
	}
	*i = ShortStr(ii)
	return err
}

var (
	// constShortStr is unconditionally added to JSON
	constShortStr = json.RawMessage(`{"type":"shortstr"}`)
)

// MarshalJSON encodes JSON
func (i ShortStr) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalShortStr(i), constShortStr)
}

// PropertyOctet structure is generated from #/definitions/propertyOctet
type PropertyOctet struct {
	Value string `json:"value,omitempty"`
}

type marshalPropertyOctet PropertyOctet

// UnmarshalJSON decodes JSON
func (i *PropertyOctet) UnmarshalJSON(data []byte) error {
	ii := marshalPropertyOctet(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"value",
		},
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"octet"` {
	    return errors.New(`bad or missing const value for "type" ("octet" expected)`)
	}
	if err != nil {
		return err
	}
	*i = PropertyOctet(ii)
	return err
}

var (
	// constPropertyOctet is unconditionally added to JSON
	constPropertyOctet = json.RawMessage(`{"type":"octet"}`)
)

// MarshalJSON encodes JSON
func (i PropertyOctet) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPropertyOctet(i), constPropertyOctet)
}

// Timestamp structure is generated from #/definitions/timestamp
type Timestamp struct {
	Value string `json:"value,omitempty"`
}

type marshalTimestamp Timestamp

// UnmarshalJSON decodes JSON
func (i *Timestamp) UnmarshalJSON(data []byte) error {
	ii := marshalTimestamp(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"value",
		},
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"timestamp"` {
	    return errors.New(`bad or missing const value for "type" ("timestamp" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Timestamp(ii)
	return err
}

var (
	// constTimestamp is unconditionally added to JSON
	constTimestamp = json.RawMessage(`{"type":"timestamp"}`)
)

// MarshalJSON encodes JSON
func (i Timestamp) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTimestamp(i), constTimestamp)
}

// StringedType is a constant type
type StringedType string

// StringedType values enumeration
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

// MarshalJSON encodes JSON
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
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
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
		return errors.New("unexpected value")
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
	// closing empty result
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
func unmarshalUnion(
	mustUnmarshal []interface{},
	mayUnmarshal []interface{},
	ignoreKeys []string,
	regexMaps map[string]interface{},
	j []byte,
) error {
	for _, item := range mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			return err
		}
	}

	for i, item := range mayUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			mayUnmarshal[i] = nil
		}
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// removing ignored keys (defined in struct)
	for _, i := range ignoreKeys {
		delete(m, i)
	}

	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// preparing regexp matchers
	var reg = make(map[string]*regexp.Regexp, len(regexMaps))
	for regex := range regexMaps {
		if regex != "" {
			reg[regex], err = regexp.Compile(regex)
			if err != nil {
				return err //todo use errors.Wrap
			}
		}
	}

	subMapsRaw := make(map[string][]byte, len(regexMaps))
	// iterating map and feeding subMaps
	for key, val := range m {
		matched := false
		var ok bool
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex, exp := range reg {
			if exp.MatchString(key) {
				matched = true
				var subMap []byte
				if subMap, ok = subMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				subMap = append(subMap, []byte(*val)...)
				subMap = append(subMap, '}')

				subMapsRaw[regex] = subMap
			}
		}

		// empty regex for additionalProperties
		if !matched {
			var subMap []byte
			if subMap, ok = subMapsRaw[""]; !ok {
				subMap = make([]byte, 1, 100)
				subMap[0] = '{'
			} else {
				subMap = append(subMap[:len(subMap)-1], ',')
			}
			subMap = append(subMap, []byte(keyEscaped)...)
			subMap = append(subMap, []byte(*val)...)
			subMap = append(subMap, '}')

			subMapsRaw[""] = subMap
		}
	}

	for regex := range regexMaps {
		if subMap, ok := subMapsRaw[regex]; !ok {
			continue
		} else {
			err = json.Unmarshal(subMap, regexMaps[regex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
