// Code generated by github.com/swaggest/go-code-builder, DO NOT EDIT.

package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// MessagingReaderReads structure is generated from "#/components/schemas/MessagingReaderReads".
type MessagingReaderReads struct {
	Reads                []Book                 `json:"reads,omitempty"`
	Country              string                 `json:"country,omitempty"`         // Country
	ReaderID             int64                  `json:"reader_id,omitempty"`
	Week                 string                 `json:"week,omitempty"`            // Week
	SubscriptionID       int64                  `json:"subscription_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                         // All unmatched properties
}

type marshalMessagingReaderReads MessagingReaderReads

// UnmarshalJSON decodes JSON.
func (i *MessagingReaderReads) UnmarshalJSON(data []byte) error {
	ii := marshalMessagingReaderReads(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"reads",
			"country",
			"reader_id",
			"week",
			"subscription_id",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = MessagingReaderReads(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i MessagingReaderReads) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessagingReaderReads(i), i.AdditionalProperties)
}

// Book structure is generated from "#/components/schemas/Book".
type Book struct {
	Amount               int64                  `json:"amount,omitempty"`
	EntityID             string                 `json:"entity_id,omitempty"` // ID of the charged entity
	Strategy             PlotStrategy           `json:"strategy,omitempty"`  // Read strategy
	Type                 ReadType               `json:"type,omitempty"`      // Read type
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties
}

type marshalBook Book

// UnmarshalJSON decodes JSON.
func (i *Book) UnmarshalJSON(data []byte) error {
	ii := marshalBook(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"amount",
			"entity_id",
			"strategy",
			"type",
			"entity_type",
			"reason",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if v, ok := constValues["entity_type"]; !ok || string(v) != `"book"` {
		return fmt.Errorf(`bad or missing const value for "entity_type" ("book" expected, %v received)`, v)
	}
	if v, ok := constValues["reason"]; !ok || string(v) != `"premium"` {
		return fmt.Errorf(`bad or missing const value for "reason" ("premium" expected, %v received)`, v)
	}
	if err != nil {
		return err
	}
	*i = Book(ii)
	return err
}

var (
	// constBook is unconditionally added to JSON.
	constBook = json.RawMessage(`{"entity_type":"book","reason":"premium"}`)
)

// MarshalJSON encodes JSON.
func (i Book) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBook, marshalBook(i), i.AdditionalProperties)
}

// PlotStrategy is an enum type.
type PlotStrategy string

// PlotStrategy values enumeration.
const (
	PlotStrategyFast = PlotStrategy("fast")
	PlotStrategySlow = PlotStrategy("slow")
)

// MarshalJSON encodes JSON.
func (i PlotStrategy) MarshalJSON() ([]byte, error) {
	switch i {
	case PlotStrategyFast:
	case PlotStrategySlow:

	default:
		return nil, fmt.Errorf("unexpected PlotStrategy value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *PlotStrategy) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := PlotStrategy(ii)
	switch v {
	case PlotStrategyFast:
	case PlotStrategySlow:

	default:
		return fmt.Errorf("unexpected PlotStrategy value: %v", v)
	}

	*i = v
	return nil
}

// ReadType is an enum type.
type ReadType string

// ReadType values enumeration.
const (
	ReadTypeNight = ReadType("night")
	ReadTypeDay = ReadType("day")
)

// MarshalJSON encodes JSON.
func (i ReadType) MarshalJSON() ([]byte, error) {
	switch i {
	case ReadTypeNight:
	case ReadTypeDay:

	default:
		return nil, fmt.Errorf("unexpected ReadType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ReadType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := ReadType(ii)
	switch v {
	case ReadTypeNight:
	case ReadTypeDay:

	default:
		return fmt.Errorf("unexpected ReadType value: %v", v)
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
type unionMap struct {
	mustUnmarshal        []interface{}
	mayUnmarshal         []interface{}
	ignoreKeys           []string
	additionalProperties interface{}
	jsonData             []byte
}

func (u unionMap) unmarshal() error {
	for _, item := range u.mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			return err
		}
	}
	for i, item := range u.mayUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			u.mayUnmarshal[i] = nil
		}
	}

	if u.additionalProperties == nil {
		return nil
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(u.jsonData, &m)
	if err != nil {
		return err
	}
	// removing ignored keys (defined in struct)
	for _, i := range u.ignoreKeys {
		delete(m, i)
	}
	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// Returning early on empty map.
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
		subMap = append(subMap, []byte(*val)...)
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
