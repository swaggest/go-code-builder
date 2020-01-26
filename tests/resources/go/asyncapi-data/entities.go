package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

var ignoreKeysMessagingReaderReads = []string{
	"reads",
	"country",
	"reader_id",
	"week",
	"subscription_id",
}

// UnmarshalJSON decodes JSON.
func (i *MessagingReaderReads) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMessagingReaderReads(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysMessagingReaderReads {
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

	*i = MessagingReaderReads(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i MessagingReaderReads) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalMessagingReaderReads(i))
	}
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

var ignoreKeysBook = []string{
	"amount",
	"entity_id",
	"strategy",
	"type",
	"entity_type",
	"reason",
}

// UnmarshalJSON decodes JSON.
func (i *Book) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalBook(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["entity_type"]; !ok || string(v) != `"book"` {
		return fmt.Errorf(`bad or missing const value for "entity_type" ("book" expected, %s received)`, v)
	}

	delete(m, "entity_type")

	if v, ok := m["reason"]; !ok || string(v) != `"premium"` {
		return fmt.Errorf(`bad or missing const value for "reason" ("premium" expected, %s received)`, v)
	}

	delete(m, "reason")

	for _, key := range ignoreKeysBook {
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

	*i = Book(ii)

	return nil
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
