package jsonschema

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

// WithID sets ID value.
func (v *CoreSchemaMetaSchema) WithID(val string) *CoreSchemaMetaSchema {
	v.ID = &val
	return v
}

// WithSchema sets Schema value.
func (v *CoreSchemaMetaSchema) WithSchema(val string) *CoreSchemaMetaSchema {
	v.Schema = &val
	return v
}

// WithRef sets Ref value.
func (v *CoreSchemaMetaSchema) WithRef(val string) *CoreSchemaMetaSchema {
	v.Ref = &val
	return v
}

// WithComment sets Comment value.
func (v *CoreSchemaMetaSchema) WithComment(val string) *CoreSchemaMetaSchema {
	v.Comment = &val
	return v
}

// WithTitle sets Title value.
func (v *CoreSchemaMetaSchema) WithTitle(val string) *CoreSchemaMetaSchema {
	v.Title = &val
	return v
}

// WithDescription sets Description value.
func (v *CoreSchemaMetaSchema) WithDescription(val string) *CoreSchemaMetaSchema {
	v.Description = &val
	return v
}

// WithDefault sets Default value.
func (v *CoreSchemaMetaSchema) WithDefault(val interface{}) *CoreSchemaMetaSchema {
	v.Default = &val
	return v
}

// WithReadOnly sets ReadOnly value.
func (v *CoreSchemaMetaSchema) WithReadOnly(val bool) *CoreSchemaMetaSchema {
	v.ReadOnly = &val
	return v
}

// WithExamples sets Examples value.
func (v *CoreSchemaMetaSchema) WithExamples(val ...interface{}) *CoreSchemaMetaSchema {
	v.Examples = val
	return v
}

// WithMultipleOf sets MultipleOf value.
func (v *CoreSchemaMetaSchema) WithMultipleOf(val float64) *CoreSchemaMetaSchema {
	v.MultipleOf = &val
	return v
}

// WithMaximum sets Maximum value.
func (v *CoreSchemaMetaSchema) WithMaximum(val float64) *CoreSchemaMetaSchema {
	v.Maximum = &val
	return v
}

// WithExclusiveMaximum sets ExclusiveMaximum value.
func (v *CoreSchemaMetaSchema) WithExclusiveMaximum(val float64) *CoreSchemaMetaSchema {
	v.ExclusiveMaximum = &val
	return v
}

// WithMinimum sets Minimum value.
func (v *CoreSchemaMetaSchema) WithMinimum(val float64) *CoreSchemaMetaSchema {
	v.Minimum = &val
	return v
}

// WithExclusiveMinimum sets ExclusiveMinimum value.
func (v *CoreSchemaMetaSchema) WithExclusiveMinimum(val float64) *CoreSchemaMetaSchema {
	v.ExclusiveMinimum = &val
	return v
}

// WithMaxLength sets MaxLength value.
func (v *CoreSchemaMetaSchema) WithMaxLength(val int64) *CoreSchemaMetaSchema {
	v.MaxLength = &val
	return v
}

// WithMinLength sets MinLength value.
func (v *CoreSchemaMetaSchema) WithMinLength(val int64) *CoreSchemaMetaSchema {
	v.MinLength = val
	return v
}

// WithPattern sets Pattern value.
func (v *CoreSchemaMetaSchema) WithPattern(val string) *CoreSchemaMetaSchema {
	v.Pattern = &val
	return v
}

// WithAdditionalItems sets AdditionalItems value.
func (v *CoreSchemaMetaSchema) WithAdditionalItems(val Schema) *CoreSchemaMetaSchema {
	v.AdditionalItems = &val
	return v
}

// WithItems sets Items value.
func (v *CoreSchemaMetaSchema) WithItems(val Items) *CoreSchemaMetaSchema {
	v.Items = &val
	return v
}

// WithMaxItems sets MaxItems value.
func (v *CoreSchemaMetaSchema) WithMaxItems(val int64) *CoreSchemaMetaSchema {
	v.MaxItems = &val
	return v
}

// WithMinItems sets MinItems value.
func (v *CoreSchemaMetaSchema) WithMinItems(val int64) *CoreSchemaMetaSchema {
	v.MinItems = val
	return v
}

// WithUniqueItems sets UniqueItems value.
func (v *CoreSchemaMetaSchema) WithUniqueItems(val bool) *CoreSchemaMetaSchema {
	v.UniqueItems = &val
	return v
}

// WithContains sets Contains value.
func (v *CoreSchemaMetaSchema) WithContains(val Schema) *CoreSchemaMetaSchema {
	v.Contains = &val
	return v
}

// WithMaxProperties sets MaxProperties value.
func (v *CoreSchemaMetaSchema) WithMaxProperties(val int64) *CoreSchemaMetaSchema {
	v.MaxProperties = &val
	return v
}

// WithMinProperties sets MinProperties value.
func (v *CoreSchemaMetaSchema) WithMinProperties(val int64) *CoreSchemaMetaSchema {
	v.MinProperties = val
	return v
}

// WithRequired sets Required value.
func (v *CoreSchemaMetaSchema) WithRequired(val ...string) *CoreSchemaMetaSchema {
	v.Required = val
	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *CoreSchemaMetaSchema) WithAdditionalProperties(val Schema) *CoreSchemaMetaSchema {
	v.AdditionalProperties = &val
	return v
}

// WithDefinitions sets Definitions value.
func (v *CoreSchemaMetaSchema) WithDefinitions(val map[string]Schema) *CoreSchemaMetaSchema {
	v.Definitions = val
	return v
}

// WithProperties sets Properties value.
func (v *CoreSchemaMetaSchema) WithProperties(val map[string]Schema) *CoreSchemaMetaSchema {
	v.Properties = val
	return v
}

// WithPatternProperties sets PatternProperties value.
func (v *CoreSchemaMetaSchema) WithPatternProperties(val map[string]Schema) *CoreSchemaMetaSchema {
	v.PatternProperties = val
	return v
}

// WithDependencies sets Dependencies value.
func (v *CoreSchemaMetaSchema) WithDependencies(val map[string]DependenciesAdditionalProperties) *CoreSchemaMetaSchema {
	v.Dependencies = val
	return v
}

// WithPropertyNames sets PropertyNames value.
func (v *CoreSchemaMetaSchema) WithPropertyNames(val Schema) *CoreSchemaMetaSchema {
	v.PropertyNames = &val
	return v
}

// WithConst sets Const value.
func (v *CoreSchemaMetaSchema) WithConst(val interface{}) *CoreSchemaMetaSchema {
	v.Const = &val
	return v
}

// WithEnum sets Enum value.
func (v *CoreSchemaMetaSchema) WithEnum(val ...interface{}) *CoreSchemaMetaSchema {
	v.Enum = val
	return v
}

// WithType sets Type value.
func (v *CoreSchemaMetaSchema) WithType(val Type) *CoreSchemaMetaSchema {
	v.Type = &val
	return v
}

// WithFormat sets Format value.
func (v *CoreSchemaMetaSchema) WithFormat(val string) *CoreSchemaMetaSchema {
	v.Format = &val
	return v
}

// WithContentMediaType sets ContentMediaType value.
func (v *CoreSchemaMetaSchema) WithContentMediaType(val string) *CoreSchemaMetaSchema {
	v.ContentMediaType = &val
	return v
}

// WithContentEncoding sets ContentEncoding value.
func (v *CoreSchemaMetaSchema) WithContentEncoding(val string) *CoreSchemaMetaSchema {
	v.ContentEncoding = &val
	return v
}

// WithIf sets If value.
func (v *CoreSchemaMetaSchema) WithIf(val Schema) *CoreSchemaMetaSchema {
	v.If = &val
	return v
}

// WithThen sets Then value.
func (v *CoreSchemaMetaSchema) WithThen(val Schema) *CoreSchemaMetaSchema {
	v.Then = &val
	return v
}

// WithElse sets Else value.
func (v *CoreSchemaMetaSchema) WithElse(val Schema) *CoreSchemaMetaSchema {
	v.Else = &val
	return v
}

// WithAllOf sets AllOf value.
func (v *CoreSchemaMetaSchema) WithAllOf(val ...Schema) *CoreSchemaMetaSchema {
	v.AllOf = val
	return v
}

// WithAnyOf sets AnyOf value.
func (v *CoreSchemaMetaSchema) WithAnyOf(val ...Schema) *CoreSchemaMetaSchema {
	v.AnyOf = val
	return v
}

// WithOneOf sets OneOf value.
func (v *CoreSchemaMetaSchema) WithOneOf(val ...Schema) *CoreSchemaMetaSchema {
	v.OneOf = val
	return v
}

// WithNot sets Not value.
func (v *CoreSchemaMetaSchema) WithNot(val Schema) *CoreSchemaMetaSchema {
	v.Not = &val
	return v
}

// WithExtraProperties sets ExtraProperties value.
func (v *CoreSchemaMetaSchema) WithExtraProperties(val map[string]interface{}) *CoreSchemaMetaSchema {
	v.ExtraProperties = val
	return v
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
	if len(i.ExtraProperties) == 0 {
		return json.Marshal(marshalCoreSchemaMetaSchema(i))
	}
	return marshalUnion(marshalCoreSchemaMetaSchema(i), i.ExtraProperties)
}

// Schema structure is generated from "#".
//
// Core schema meta-schema.
type Schema struct {
	TypeObject  *CoreSchemaMetaSchema `json:"-"`
	TypeBoolean *bool                 `json:"-"`
}

// WithTypeObject sets TypeObject value.
func (v *Schema) WithTypeObject(val CoreSchemaMetaSchema) *Schema {
	v.TypeObject = &val
	return v
}

// WithTypeBoolean sets TypeBoolean value.
func (v *Schema) WithTypeBoolean(val bool) *Schema {
	v.TypeBoolean = &val
	return v
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

// WithSchema sets Schema value.
func (v *Items) WithSchema(val Schema) *Items {
	v.Schema = &val
	return v
}

// WithSchemaArray sets SchemaArray value.
func (v *Items) WithSchemaArray(val ...Schema) *Items {
	v.SchemaArray = val
	return v
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

// WithSchema sets Schema value.
func (v *DependenciesAdditionalProperties) WithSchema(val Schema) *DependenciesAdditionalProperties {
	v.Schema = &val
	return v
}

// WithStringArray sets StringArray value.
func (v *DependenciesAdditionalProperties) WithStringArray(val ...string) *DependenciesAdditionalProperties {
	v.StringArray = val
	return v
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

// WithSimpleTypes sets SimpleTypes value.
func (v *Type) WithSimpleTypes(val SimpleTypes) *Type {
	v.SimpleTypes = &val
	return v
}

// WithSliceOfSimpleTypesValues sets SliceOfSimpleTypesValues value.
func (v *Type) WithSliceOfSimpleTypesValues(val ...SimpleTypes) *Type {
	v.SliceOfSimpleTypesValues = val
	return v
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
