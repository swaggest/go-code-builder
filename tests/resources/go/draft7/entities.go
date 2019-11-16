package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	Default              *interface{}                                `json:"default,omitempty"`
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
	Const                *interface{}                                `json:"const,omitempty"`
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

var ignoreKeysCoreSchemaMetaSchema = []string{
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
}

// UnmarshalJSON decodes JSON.
func (i *CoreSchemaMetaSchema) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalCoreSchemaMetaSchema(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Default == nil {
		if _, ok := m["default"]; ok {
			var v interface{}
			ii.Default = &v
		}
	}

	if ii.Const == nil {
		if _, ok := m["const"]; ok {
			var v interface{}
			ii.Const = &v
		}
	}

	for _, key := range ignoreKeysCoreSchemaMetaSchema {
		delete(m, key)
	}

	for key, rawValue := range m {
		if ii.ExtraProperties == nil {
			ii.ExtraProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.ExtraProperties[key] = val
	}

	*i = CoreSchemaMetaSchema(ii)

	return nil
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
	var err error

	err = json.Unmarshal(data, &i.TypeObject)
	if err != nil {
		i.TypeObject = nil
	}

	err = json.Unmarshal(data, &i.TypeBoolean)
	if err != nil {
		i.TypeBoolean = nil
	}

	return nil
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
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.SchemaArray)
	if err != nil {
		i.SchemaArray = nil
	}

	return nil
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
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.StringArray)
	if err != nil {
		i.StringArray = nil
	}

	return nil
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
	var err error

	err = json.Unmarshal(data, &i.SimpleTypes)
	if err != nil {
		i.SimpleTypes = nil
	}

	err = json.Unmarshal(data, &i.SliceOfSimpleTypesValues)
	if err != nil {
		i.SliceOfSimpleTypesValues = nil
	}

	return nil
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
