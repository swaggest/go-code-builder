package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CoreSchemaMetaSchema structure is generated from "#[object]".
//
// Core schema meta-schema.
type CoreSchemaMetaSchema struct {
	ID                   *string                                     `json:"$id,omitempty"`
	Schema               *string                                     `json:"$schema,omitempty"`
	Ref                  *string                                     `json:"$ref,omitempty"`
	Comment              *string                                     `json:"$comment,omitempty"`
	Title                *string                                     `json:"title,omitempty"`
	Description          *string                                     `json:"description,omitempty"`
	Default              interface{}                                 `json:"default,omitempty"`
	ReadOnly             *bool                                       `json:"readOnly,omitempty"`
	Examples             []interface{}                               `json:"examples,omitempty"`
	MultipleOf           *float64                                    `json:"multipleOf,omitempty"`
	Maximum              *float64                                    `json:"maximum,omitempty"`
	ExclusiveMaximum     *float64                                    `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64                                    `json:"minimum,omitempty"`
	ExclusiveMinimum     *float64                                    `json:"exclusiveMinimum,omitempty"`
	MaxLength            *int64                                      `json:"maxLength,omitempty"`
	MinLength            int64                                       `json:"minLength,omitempty"`
	Pattern              *string                                     `json:"pattern,omitempty"`
	AdditionalItems      *Schema                                     `json:"additionalItems,omitempty"`      // Core schema meta-schema
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             *int64                                      `json:"maxItems,omitempty"`
	MinItems             int64                                       `json:"minItems,omitempty"`
	UniqueItems          *bool                                       `json:"uniqueItems,omitempty"`
	Contains             *Schema                                     `json:"contains,omitempty"`             // Core schema meta-schema
	MaxProperties        *int64                                      `json:"maxProperties,omitempty"`
	MinProperties        int64                                       `json:"minProperties,omitempty"`
	Required             []string                                    `json:"required,omitempty"`
	AdditionalProperties *Schema                                     `json:"additionalProperties,omitempty"` // Core schema meta-schema
	Definitions          map[string]Schema                           `json:"definitions,omitempty"`
	Properties           map[string]Schema                           `json:"properties,omitempty"`
	PatternProperties    map[string]Schema                           `json:"patternProperties,omitempty"`
	Dependencies         map[string]DependenciesAdditionalProperties `json:"dependencies,omitempty"`
	PropertyNames        *Schema                                     `json:"propertyNames,omitempty"`        // Core schema meta-schema
	Const                interface{}                                 `json:"const,omitempty"`
	Enum                 []interface{}                               `json:"enum,omitempty"`
	Type                 *Type                                       `json:"type,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	ContentMediaType     *string                                     `json:"contentMediaType,omitempty"`
	ContentEncoding      *string                                     `json:"contentEncoding,omitempty"`
	If                   *Schema                                     `json:"if,omitempty"`                   // Core schema meta-schema
	Then                 *Schema                                     `json:"then,omitempty"`                 // Core schema meta-schema
	Else                 *Schema                                     `json:"else,omitempty"`                 // Core schema meta-schema
	AllOf                []Schema                                    `json:"allOf,omitempty"`
	AnyOf                []Schema                                    `json:"anyOf,omitempty"`
	OneOf                []Schema                                    `json:"oneOf,omitempty"`
	Not                  *Schema                                     `json:"not,omitempty"`                  // Core schema meta-schema
	ExtraProperties      map[string]interface{}                      `json:"-"`                              // All unmatched properties
}

type marshalCoreSchemaMetaSchema CoreSchemaMetaSchema

// UnmarshalJSON decodes JSON.
func (i *CoreSchemaMetaSchema) UnmarshalJSON(data []byte) error {
	ii := marshalCoreSchemaMetaSchema(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$id",
			"$schema",
			"$ref",
			"$comment",
			"title",
			"description",
			"default",
			"readOnly",
			"examples",
			"multipleOf",
			"maximum",
			"exclusiveMaximum",
			"minimum",
			"exclusiveMinimum",
			"maxLength",
			"minLength",
			"pattern",
			"additionalItems",
			"items",
			"maxItems",
			"minItems",
			"uniqueItems",
			"contains",
			"maxProperties",
			"minProperties",
			"required",
			"additionalProperties",
			"definitions",
			"properties",
			"patternProperties",
			"dependencies",
			"propertyNames",
			"const",
			"enum",
			"type",
			"format",
			"contentMediaType",
			"contentEncoding",
			"if",
			"then",
			"else",
			"allOf",
			"anyOf",
			"oneOf",
			"not",
		},
		additionalProperties: &ii.ExtraProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = CoreSchemaMetaSchema(ii)

	return err
}

// MarshalJSON encodes JSON.
func (i CoreSchemaMetaSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalCoreSchemaMetaSchema(i), i.ExtraProperties)
}

// Schema structure is generated from "#".
//
// Core schema meta-schema.
type Schema struct {
	TypeObject  *CoreSchemaMetaSchema `json:"-"`
	TypeBoolean *bool                 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Schema) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.TypeObject, &i.TypeBoolean}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.TypeObject = nil
	}

	if mayUnmarshal[1] == nil {
		i.TypeBoolean = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.TypeObject, i.TypeBoolean)
}

// Items structure is generated from "#[object]->items".
type Items struct {
	Schema      *Schema  `json:"-"`
	SchemaArray []Schema `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Items) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.SchemaArray}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.SchemaArray = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i Items) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.SchemaArray)
}

// DependenciesAdditionalProperties structure is generated from "#[object]->dependencies->additionalProperties".
type DependenciesAdditionalProperties struct {
	Schema      *Schema  `json:"-"`
	StringArray []string `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *DependenciesAdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.StringArray}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.StringArray = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i DependenciesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.StringArray)
}

// Type structure is generated from "#[object]->type".
type Type struct {
	SimpleTypes              *SimpleTypes  `json:"-"`
	SliceOfSimpleTypesValues []SimpleTypes `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Type) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.SimpleTypes, &i.SliceOfSimpleTypesValues}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.SimpleTypes = nil
	}

	if mayUnmarshal[1] == nil {
		i.SliceOfSimpleTypesValues = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i Type) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.SimpleTypes, i.SliceOfSimpleTypesValues)
}

// SimpleTypes is an enum type.
type SimpleTypes string

// SimpleTypes values enumeration.
const (
	SimpleTypesArray = SimpleTypes("array")
	SimpleTypesBoolean = SimpleTypes("boolean")
	SimpleTypesInteger = SimpleTypes("integer")
	SimpleTypesNull = SimpleTypes("null")
	SimpleTypesNumber = SimpleTypes("number")
	SimpleTypesObject = SimpleTypes("object")
	SimpleTypesString = SimpleTypes("string")
)

// MarshalJSON encodes JSON.
func (i SimpleTypes) MarshalJSON() ([]byte, error) {
	switch i {
	case SimpleTypesArray:
	case SimpleTypesBoolean:
	case SimpleTypesInteger:
	case SimpleTypesNull:
	case SimpleTypesNumber:
	case SimpleTypesObject:
	case SimpleTypesString:

	default:
		return nil, fmt.Errorf("unexpected SimpleTypes value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SimpleTypes) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SimpleTypes(ii)

	switch v {
	case SimpleTypesArray:
	case SimpleTypesBoolean:
	case SimpleTypesInteger:
	case SimpleTypesNull:
	case SimpleTypesNumber:
	case SimpleTypesObject:
	case SimpleTypesString:

	default:
		return fmt.Errorf("unexpected SimpleTypes value: %v", v)
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

	if u.additionalProperties == nil {
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
