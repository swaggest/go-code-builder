// Package entities contains generated structures.
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
	ID                   *string                                     `json:"$id,omitempty"`                  // Format: uri-reference.
	Schema               *string                                     `json:"$schema,omitempty"`              // Format: uri.
	Ref                  *string                                     `json:"$ref,omitempty"`                 // Format: uri-reference.
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
	Pattern              *string                                     `json:"pattern,omitempty"`              // Format: regex.
	AdditionalItems      *Schema                                     `json:"additionalItems,omitempty"`      // Core schema meta-schema.
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             *int64                                      `json:"maxItems,omitempty"`
	MinItems             int64                                       `json:"minItems,omitempty"`
	UniqueItems          *bool                                       `json:"uniqueItems,omitempty"`
	Contains             *Schema                                     `json:"contains,omitempty"`             // Core schema meta-schema.
	MaxProperties        *int64                                      `json:"maxProperties,omitempty"`
	MinProperties        int64                                       `json:"minProperties,omitempty"`
	Required             []string                                    `json:"required,omitempty"`
	AdditionalProperties *Schema                                     `json:"additionalProperties,omitempty"` // Core schema meta-schema.
	Definitions          map[string]Schema                           `json:"definitions,omitempty"`
	Properties           map[string]Schema                           `json:"properties,omitempty"`
	PatternProperties    map[string]Schema                           `json:"patternProperties,omitempty"`
	Dependencies         map[string]DependenciesAdditionalProperties `json:"dependencies,omitempty"`
	PropertyNames        *Schema                                     `json:"propertyNames,omitempty"`        // Core schema meta-schema.
	Const                *interface{}                                `json:"const,omitempty"`
	Enum                 []interface{}                               `json:"enum,omitempty"`
	Type                 *Type                                       `json:"type,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	ContentMediaType     *string                                     `json:"contentMediaType,omitempty"`
	ContentEncoding      *string                                     `json:"contentEncoding,omitempty"`
	If                   *Schema                                     `json:"if,omitempty"`                   // Core schema meta-schema.
	Then                 *Schema                                     `json:"then,omitempty"`                 // Core schema meta-schema.
	Else                 *Schema                                     `json:"else,omitempty"`                 // Core schema meta-schema.
	AllOf                []Schema                                    `json:"allOf,omitempty"`
	AnyOf                []Schema                                    `json:"anyOf,omitempty"`
	OneOf                []Schema                                    `json:"oneOf,omitempty"`
	Not                  *Schema                                     `json:"not,omitempty"`                  // Core schema meta-schema.
	ExtraProperties      map[string]interface{}                      `json:"-"`                              // All unmatched properties.
}

// WithID sets ID value.
func (c *CoreSchemaMetaSchema) WithID(val string) *CoreSchemaMetaSchema {
	c.ID = &val
	return c
}

// WithSchema sets Schema value.
func (c *CoreSchemaMetaSchema) WithSchema(val string) *CoreSchemaMetaSchema {
	c.Schema = &val
	return c
}

// WithRef sets Ref value.
func (c *CoreSchemaMetaSchema) WithRef(val string) *CoreSchemaMetaSchema {
	c.Ref = &val
	return c
}

// WithComment sets Comment value.
func (c *CoreSchemaMetaSchema) WithComment(val string) *CoreSchemaMetaSchema {
	c.Comment = &val
	return c
}

// WithTitle sets Title value.
func (c *CoreSchemaMetaSchema) WithTitle(val string) *CoreSchemaMetaSchema {
	c.Title = &val
	return c
}

// WithDescription sets Description value.
func (c *CoreSchemaMetaSchema) WithDescription(val string) *CoreSchemaMetaSchema {
	c.Description = &val
	return c
}

// WithDefault sets Default value.
func (c *CoreSchemaMetaSchema) WithDefault(val interface{}) *CoreSchemaMetaSchema {
	c.Default = &val
	return c
}

// WithReadOnly sets ReadOnly value.
func (c *CoreSchemaMetaSchema) WithReadOnly(val bool) *CoreSchemaMetaSchema {
	c.ReadOnly = &val
	return c
}

// WithExamples sets Examples value.
func (c *CoreSchemaMetaSchema) WithExamples(val ...interface{}) *CoreSchemaMetaSchema {
	c.Examples = val
	return c
}

// WithMultipleOf sets MultipleOf value.
func (c *CoreSchemaMetaSchema) WithMultipleOf(val float64) *CoreSchemaMetaSchema {
	c.MultipleOf = &val
	return c
}

// WithMaximum sets Maximum value.
func (c *CoreSchemaMetaSchema) WithMaximum(val float64) *CoreSchemaMetaSchema {
	c.Maximum = &val
	return c
}

// WithExclusiveMaximum sets ExclusiveMaximum value.
func (c *CoreSchemaMetaSchema) WithExclusiveMaximum(val float64) *CoreSchemaMetaSchema {
	c.ExclusiveMaximum = &val
	return c
}

// WithMinimum sets Minimum value.
func (c *CoreSchemaMetaSchema) WithMinimum(val float64) *CoreSchemaMetaSchema {
	c.Minimum = &val
	return c
}

// WithExclusiveMinimum sets ExclusiveMinimum value.
func (c *CoreSchemaMetaSchema) WithExclusiveMinimum(val float64) *CoreSchemaMetaSchema {
	c.ExclusiveMinimum = &val
	return c
}

// WithMaxLength sets MaxLength value.
func (c *CoreSchemaMetaSchema) WithMaxLength(val int64) *CoreSchemaMetaSchema {
	c.MaxLength = &val
	return c
}

// WithMinLength sets MinLength value.
func (c *CoreSchemaMetaSchema) WithMinLength(val int64) *CoreSchemaMetaSchema {
	c.MinLength = val
	return c
}

// WithPattern sets Pattern value.
func (c *CoreSchemaMetaSchema) WithPattern(val string) *CoreSchemaMetaSchema {
	c.Pattern = &val
	return c
}

// WithAdditionalItems sets AdditionalItems value.
func (c *CoreSchemaMetaSchema) WithAdditionalItems(val Schema) *CoreSchemaMetaSchema {
	c.AdditionalItems = &val
	return c
}

// WithItems sets Items value.
func (c *CoreSchemaMetaSchema) WithItems(val Items) *CoreSchemaMetaSchema {
	c.Items = &val
	return c
}

// WithMaxItems sets MaxItems value.
func (c *CoreSchemaMetaSchema) WithMaxItems(val int64) *CoreSchemaMetaSchema {
	c.MaxItems = &val
	return c
}

// WithMinItems sets MinItems value.
func (c *CoreSchemaMetaSchema) WithMinItems(val int64) *CoreSchemaMetaSchema {
	c.MinItems = val
	return c
}

// WithUniqueItems sets UniqueItems value.
func (c *CoreSchemaMetaSchema) WithUniqueItems(val bool) *CoreSchemaMetaSchema {
	c.UniqueItems = &val
	return c
}

// WithContains sets Contains value.
func (c *CoreSchemaMetaSchema) WithContains(val Schema) *CoreSchemaMetaSchema {
	c.Contains = &val
	return c
}

// WithMaxProperties sets MaxProperties value.
func (c *CoreSchemaMetaSchema) WithMaxProperties(val int64) *CoreSchemaMetaSchema {
	c.MaxProperties = &val
	return c
}

// WithMinProperties sets MinProperties value.
func (c *CoreSchemaMetaSchema) WithMinProperties(val int64) *CoreSchemaMetaSchema {
	c.MinProperties = val
	return c
}

// WithRequired sets Required value.
func (c *CoreSchemaMetaSchema) WithRequired(val ...string) *CoreSchemaMetaSchema {
	c.Required = val
	return c
}

// WithAdditionalProperties sets AdditionalProperties value.
func (c *CoreSchemaMetaSchema) WithAdditionalProperties(val Schema) *CoreSchemaMetaSchema {
	c.AdditionalProperties = &val
	return c
}

// WithDefinitions sets Definitions value.
func (c *CoreSchemaMetaSchema) WithDefinitions(val map[string]Schema) *CoreSchemaMetaSchema {
	c.Definitions = val
	return c
}

// WithDefinitionsItem sets Definitions item value.
func (c *CoreSchemaMetaSchema) WithDefinitionsItem(key string, val Schema) *CoreSchemaMetaSchema {
	if c.Definitions == nil {
		c.Definitions = make(map[string]Schema, 1)
	}

	c.Definitions[key] = val

	return c
}

// WithProperties sets Properties value.
func (c *CoreSchemaMetaSchema) WithProperties(val map[string]Schema) *CoreSchemaMetaSchema {
	c.Properties = val
	return c
}

// WithPropertiesItem sets Properties item value.
func (c *CoreSchemaMetaSchema) WithPropertiesItem(key string, val Schema) *CoreSchemaMetaSchema {
	if c.Properties == nil {
		c.Properties = make(map[string]Schema, 1)
	}

	c.Properties[key] = val

	return c
}

// WithPatternProperties sets PatternProperties value.
func (c *CoreSchemaMetaSchema) WithPatternProperties(val map[string]Schema) *CoreSchemaMetaSchema {
	c.PatternProperties = val
	return c
}

// WithPatternPropertiesItem sets PatternProperties item value.
func (c *CoreSchemaMetaSchema) WithPatternPropertiesItem(key string, val Schema) *CoreSchemaMetaSchema {
	if c.PatternProperties == nil {
		c.PatternProperties = make(map[string]Schema, 1)
	}

	c.PatternProperties[key] = val

	return c
}

// WithDependencies sets Dependencies value.
func (c *CoreSchemaMetaSchema) WithDependencies(val map[string]DependenciesAdditionalProperties) *CoreSchemaMetaSchema {
	c.Dependencies = val
	return c
}

// WithDependenciesItem sets Dependencies item value.
func (c *CoreSchemaMetaSchema) WithDependenciesItem(key string, val DependenciesAdditionalProperties) *CoreSchemaMetaSchema {
	if c.Dependencies == nil {
		c.Dependencies = make(map[string]DependenciesAdditionalProperties, 1)
	}

	c.Dependencies[key] = val

	return c
}

// WithPropertyNames sets PropertyNames value.
func (c *CoreSchemaMetaSchema) WithPropertyNames(val Schema) *CoreSchemaMetaSchema {
	c.PropertyNames = &val
	return c
}

// WithConst sets Const value.
func (c *CoreSchemaMetaSchema) WithConst(val interface{}) *CoreSchemaMetaSchema {
	c.Const = &val
	return c
}

// WithEnum sets Enum value.
func (c *CoreSchemaMetaSchema) WithEnum(val ...interface{}) *CoreSchemaMetaSchema {
	c.Enum = val
	return c
}

// WithType sets Type value.
func (c *CoreSchemaMetaSchema) WithType(val Type) *CoreSchemaMetaSchema {
	c.Type = &val
	return c
}

// WithFormat sets Format value.
func (c *CoreSchemaMetaSchema) WithFormat(val string) *CoreSchemaMetaSchema {
	c.Format = &val
	return c
}

// WithContentMediaType sets ContentMediaType value.
func (c *CoreSchemaMetaSchema) WithContentMediaType(val string) *CoreSchemaMetaSchema {
	c.ContentMediaType = &val
	return c
}

// WithContentEncoding sets ContentEncoding value.
func (c *CoreSchemaMetaSchema) WithContentEncoding(val string) *CoreSchemaMetaSchema {
	c.ContentEncoding = &val
	return c
}

// WithIf sets If value.
func (c *CoreSchemaMetaSchema) WithIf(val Schema) *CoreSchemaMetaSchema {
	c.If = &val
	return c
}

// WithThen sets Then value.
func (c *CoreSchemaMetaSchema) WithThen(val Schema) *CoreSchemaMetaSchema {
	c.Then = &val
	return c
}

// WithElse sets Else value.
func (c *CoreSchemaMetaSchema) WithElse(val Schema) *CoreSchemaMetaSchema {
	c.Else = &val
	return c
}

// WithAllOf sets AllOf value.
func (c *CoreSchemaMetaSchema) WithAllOf(val ...Schema) *CoreSchemaMetaSchema {
	c.AllOf = val
	return c
}

// WithAnyOf sets AnyOf value.
func (c *CoreSchemaMetaSchema) WithAnyOf(val ...Schema) *CoreSchemaMetaSchema {
	c.AnyOf = val
	return c
}

// WithOneOf sets OneOf value.
func (c *CoreSchemaMetaSchema) WithOneOf(val ...Schema) *CoreSchemaMetaSchema {
	c.OneOf = val
	return c
}

// WithNot sets Not value.
func (c *CoreSchemaMetaSchema) WithNot(val Schema) *CoreSchemaMetaSchema {
	c.Not = &val
	return c
}

// WithExtraProperties sets ExtraProperties value.
func (c *CoreSchemaMetaSchema) WithExtraProperties(val map[string]interface{}) *CoreSchemaMetaSchema {
	c.ExtraProperties = val
	return c
}

// WithExtraPropertiesItem sets ExtraProperties item value.
func (c *CoreSchemaMetaSchema) WithExtraPropertiesItem(key string, val interface{}) *CoreSchemaMetaSchema {
	if c.ExtraProperties == nil {
		c.ExtraProperties = make(map[string]interface{}, 1)
	}

	c.ExtraProperties[key] = val

	return c
}

type marshalCoreSchemaMetaSchema CoreSchemaMetaSchema

var knownKeysCoreSchemaMetaSchema = []string{
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
func (c *CoreSchemaMetaSchema) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalCoreSchemaMetaSchema(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mc.Default == nil {
		if _, ok := rawMap["default"]; ok {
			var v interface{}
			mc.Default = &v
		}
	}

	if mc.Const == nil {
		if _, ok := rawMap["const"]; ok {
			var v interface{}
			mc.Const = &v
		}
	}

	for _, key := range knownKeysCoreSchemaMetaSchema {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mc.ExtraProperties == nil {
			mc.ExtraProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mc.ExtraProperties[key] = val
	}

	*c = CoreSchemaMetaSchema(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c CoreSchemaMetaSchema) MarshalJSON() ([]byte, error) {
	if len(c.ExtraProperties) == 0 {
		return json.Marshal(marshalCoreSchemaMetaSchema(c))
	}

	return marshalUnion(marshalCoreSchemaMetaSchema(c), c.ExtraProperties)
}

// Schema structure is generated from "#".
//
// Core schema meta-schema.
type Schema struct {
	TypeObject  *CoreSchemaMetaSchema `json:"-"`
	TypeBoolean *bool                 `json:"-"`
}

// WithTypeObject sets TypeObject value.
func (s *Schema) WithTypeObject(val CoreSchemaMetaSchema) *Schema {
	s.TypeObject = &val
	return s
}

// WithTypeBoolean sets TypeBoolean value.
func (s *Schema) WithTypeBoolean(val bool) *Schema {
	s.TypeBoolean = &val
	return s
}

// UnmarshalJSON decodes JSON.
func (s *Schema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &s.TypeObject)
	if err != nil {
		s.TypeObject = nil
	}

	err = json.Unmarshal(data, &s.TypeBoolean)
	if err != nil {
		s.TypeBoolean = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.TypeObject, s.TypeBoolean)
}

// Items structure is generated from "#[object]->items".
type Items struct {
	Schema      *Schema  `json:"-"`
	SchemaArray []Schema `json:"-"`
}

// WithSchema sets Schema value.
func (i *Items) WithSchema(val Schema) *Items {
	i.Schema = &val
	return i
}

// WithSchemaArray sets SchemaArray value.
func (i *Items) WithSchemaArray(val ...Schema) *Items {
	i.SchemaArray = val
	return i
}

// UnmarshalJSON decodes JSON.
func (i *Items) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		anyOfErrors["Schema"] = err
		i.Schema = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &i.SchemaArray)
	if err != nil {
		anyOfErrors["SchemaArray"] = err
		i.SchemaArray = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for Items failed with %d valid results: %v", anyOfValid, anyOfErrors)
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
func (d *DependenciesAdditionalProperties) WithSchema(val Schema) *DependenciesAdditionalProperties {
	d.Schema = &val
	return d
}

// WithStringArray sets StringArray value.
func (d *DependenciesAdditionalProperties) WithStringArray(val ...string) *DependenciesAdditionalProperties {
	d.StringArray = val
	return d
}

// UnmarshalJSON decodes JSON.
func (d *DependenciesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &d.Schema)
	if err != nil {
		anyOfErrors["Schema"] = err
		d.Schema = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &d.StringArray)
	if err != nil {
		anyOfErrors["StringArray"] = err
		d.StringArray = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for DependenciesAdditionalProperties failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (d DependenciesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(d.Schema, d.StringArray)
}

// Type structure is generated from "#[object]->type".
type Type struct {
	SimpleTypes              *SimpleTypes  `json:"-"`
	SliceOfSimpleTypesValues []SimpleTypes `json:"-"`
}

// WithSimpleTypes sets SimpleTypes value.
func (t *Type) WithSimpleTypes(val SimpleTypes) *Type {
	t.SimpleTypes = &val
	return t
}

// WithSliceOfSimpleTypesValues sets SliceOfSimpleTypesValues value.
func (t *Type) WithSliceOfSimpleTypesValues(val ...SimpleTypes) *Type {
	t.SliceOfSimpleTypesValues = val
	return t
}

// UnmarshalJSON decodes JSON.
func (t *Type) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &t.SimpleTypes)
	if err != nil {
		anyOfErrors["SimpleTypes"] = err
		t.SimpleTypes = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &t.SliceOfSimpleTypesValues)
	if err != nil {
		anyOfErrors["SliceOfSimpleTypesValues"] = err
		t.SliceOfSimpleTypesValues = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for Type failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (t Type) MarshalJSON() ([]byte, error) {
	return marshalUnion(t.SimpleTypes, t.SliceOfSimpleTypesValues)
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
