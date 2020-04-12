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
	// Country.
	// Value must match pattern: `^[a-zA-Z]{2}$`.
	Country              string                 `json:"country,omitempty"`
	ReaderID             int64                  `json:"reader_id,omitempty"`       // Format: int32.
	// Week.
	// Value must match pattern: `^[0-9]{4}-W(0[1-9]|[1-4][0-9]|5[0-2])$`.
	Week                 string                 `json:"week,omitempty"`
	SubscriptionID       int64                  `json:"subscription_id,omitempty"` // Format: int32.
	AdditionalProperties map[string]interface{} `json:"-"`                         // All unmatched properties.
}

type marshalMessagingReaderReads MessagingReaderReads

var knownKeysMessagingReaderReads = []string{
	"reads",
	"country",
	"reader_id",
	"week",
	"subscription_id",
}

// UnmarshalJSON decodes JSON.
func (m *MessagingReaderReads) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessagingReaderReads(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysMessagingReaderReads {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mm.AdditionalProperties == nil {
			mm.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mm.AdditionalProperties[key] = val
	}

	*m = MessagingReaderReads(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessagingReaderReads) MarshalJSON() ([]byte, error) {
	if len(m.AdditionalProperties) == 0 {
		return json.Marshal(marshalMessagingReaderReads(m))
	}

	return marshalUnion(marshalMessagingReaderReads(m), m.AdditionalProperties)
}

// Book structure is generated from "#/components/schemas/Book".
type Book struct {
	Amount               int64                  `json:"amount,omitempty"`    // Format: int64.
	EntityID             string                 `json:"entity_id,omitempty"` // ID of the charged entity.
	Strategy             PlotStrategy           `json:"strategy,omitempty"`  // Read strategy.
	Type                 ReadType               `json:"type,omitempty"`      // Read type.
	AdditionalProperties map[string]interface{} `json:"-"`                   // All unmatched properties.
}

type marshalBook Book

var knownKeysBook = []string{
	"amount",
	"entity_id",
	"strategy",
	"type",
	"entity_type",
	"reason",
}

// UnmarshalJSON decodes JSON.
func (b *Book) UnmarshalJSON(data []byte) error {
	var err error

	mb := marshalBook(*b)

	err = json.Unmarshal(data, &mb)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["entity_type"]; !ok || string(v) != `"book"` {
		return fmt.Errorf(`bad or missing const value for "entity_type" ("book" expected, %s received)`, v)
	}

	delete(rawMap, "entity_type")

	if v, ok := rawMap["reason"]; !ok || string(v) != `"premium"` {
		return fmt.Errorf(`bad or missing const value for "reason" ("premium" expected, %s received)`, v)
	}

	delete(rawMap, "reason")

	for _, key := range knownKeysBook {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mb.AdditionalProperties == nil {
			mb.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mb.AdditionalProperties[key] = val
	}

	*b = Book(mb)

	return nil
}

var (
	// constBook is unconditionally added to JSON.
	constBook = json.RawMessage(`{"entity_type":"book","reason":"premium"}`)
)

// MarshalJSON encodes JSON.
func (b Book) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBook, marshalBook(b), b.AdditionalProperties)
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
	result := []byte("{")
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
