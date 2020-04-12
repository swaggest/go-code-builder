package jsonschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Schema structure is generated from "#[object]".
//
// Core schema meta-schema.
type Schema struct {
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
	AdditionalItems      *SchemaOrBool                               `json:"additionalItems,omitempty"`      // Core schema meta-schema.
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             *int64                                      `json:"maxItems,omitempty"`
	MinItems             int64                                       `json:"minItems,omitempty"`
	UniqueItems          *bool                                       `json:"uniqueItems,omitempty"`
	Contains             *SchemaOrBool                               `json:"contains,omitempty"`             // Core schema meta-schema.
	MaxProperties        *int64                                      `json:"maxProperties,omitempty"`
	MinProperties        int64                                       `json:"minProperties,omitempty"`
	Required             []string                                    `json:"required,omitempty"`
	AdditionalProperties *SchemaOrBool                               `json:"additionalProperties,omitempty"` // Core schema meta-schema.
	Definitions          map[string]SchemaOrBool                     `json:"definitions,omitempty"`
	Properties           map[string]SchemaOrBool                     `json:"properties,omitempty"`
	PatternProperties    map[string]SchemaOrBool                     `json:"patternProperties,omitempty"`
	Dependencies         map[string]DependenciesAdditionalProperties `json:"dependencies,omitempty"`
	PropertyNames        *SchemaOrBool                               `json:"propertyNames,omitempty"`        // Core schema meta-schema.
	Const                *interface{}                                `json:"const,omitempty"`
	Enum                 []interface{}                               `json:"enum,omitempty"`
	Type                 *Type                                       `json:"type,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	ContentMediaType     *string                                     `json:"contentMediaType,omitempty"`
	ContentEncoding      *string                                     `json:"contentEncoding,omitempty"`
	If                   *SchemaOrBool                               `json:"if,omitempty"`                   // Core schema meta-schema.
	Then                 *SchemaOrBool                               `json:"then,omitempty"`                 // Core schema meta-schema.
	Else                 *SchemaOrBool                               `json:"else,omitempty"`                 // Core schema meta-schema.
	AllOf                []SchemaOrBool                              `json:"allOf,omitempty"`
	AnyOf                []SchemaOrBool                              `json:"anyOf,omitempty"`
	OneOf                []SchemaOrBool                              `json:"oneOf,omitempty"`
	Not                  *SchemaOrBool                               `json:"not,omitempty"`                  // Core schema meta-schema.
	ExtraProperties      map[string]interface{}                      `json:"-"`                              // All unmatched properties.
}

// WithID sets ID value.
func (s *Schema) WithID(val string) *Schema {
	s.ID = &val
	return s
}

// WithSchema sets Schema value.
func (s *Schema) WithSchema(val string) *Schema {
	s.Schema = &val
	return s
}

// WithRef sets Ref value.
func (s *Schema) WithRef(val string) *Schema {
	s.Ref = &val
	return s
}

// WithComment sets Comment value.
func (s *Schema) WithComment(val string) *Schema {
	s.Comment = &val
	return s
}

// WithTitle sets Title value.
func (s *Schema) WithTitle(val string) *Schema {
	s.Title = &val
	return s
}

// WithDescription sets Description value.
func (s *Schema) WithDescription(val string) *Schema {
	s.Description = &val
	return s
}

// WithDefault sets Default value.
func (s *Schema) WithDefault(val interface{}) *Schema {
	s.Default = &val
	return s
}

// WithReadOnly sets ReadOnly value.
func (s *Schema) WithReadOnly(val bool) *Schema {
	s.ReadOnly = &val
	return s
}

// WithExamples sets Examples value.
func (s *Schema) WithExamples(val ...interface{}) *Schema {
	s.Examples = val
	return s
}

// WithMultipleOf sets MultipleOf value.
func (s *Schema) WithMultipleOf(val float64) *Schema {
	s.MultipleOf = &val
	return s
}

// WithMaximum sets Maximum value.
func (s *Schema) WithMaximum(val float64) *Schema {
	s.Maximum = &val
	return s
}

// WithExclusiveMaximum sets ExclusiveMaximum value.
func (s *Schema) WithExclusiveMaximum(val float64) *Schema {
	s.ExclusiveMaximum = &val
	return s
}

// WithMinimum sets Minimum value.
func (s *Schema) WithMinimum(val float64) *Schema {
	s.Minimum = &val
	return s
}

// WithExclusiveMinimum sets ExclusiveMinimum value.
func (s *Schema) WithExclusiveMinimum(val float64) *Schema {
	s.ExclusiveMinimum = &val
	return s
}

// WithMaxLength sets MaxLength value.
func (s *Schema) WithMaxLength(val int64) *Schema {
	s.MaxLength = &val
	return s
}

// WithMinLength sets MinLength value.
func (s *Schema) WithMinLength(val int64) *Schema {
	s.MinLength = val
	return s
}

// WithPattern sets Pattern value.
func (s *Schema) WithPattern(val string) *Schema {
	s.Pattern = &val
	return s
}

// WithAdditionalItems sets AdditionalItems value.
func (s *Schema) WithAdditionalItems(val SchemaOrBool) *Schema {
	s.AdditionalItems = &val
	return s
}

// AdditionalItemsEns ensures returned AdditionalItems is not nil.
func (s *Schema) AdditionalItemsEns() *SchemaOrBool {
	if s.AdditionalItems == nil {
		s.AdditionalItems = new(SchemaOrBool)
	}

	return s.AdditionalItems
}

// WithItems sets Items value.
func (s *Schema) WithItems(val Items) *Schema {
	s.Items = &val
	return s
}

// ItemsEns ensures returned Items is not nil.
func (s *Schema) ItemsEns() *Items {
	if s.Items == nil {
		s.Items = new(Items)
	}

	return s.Items
}

// WithMaxItems sets MaxItems value.
func (s *Schema) WithMaxItems(val int64) *Schema {
	s.MaxItems = &val
	return s
}

// WithMinItems sets MinItems value.
func (s *Schema) WithMinItems(val int64) *Schema {
	s.MinItems = val
	return s
}

// WithUniqueItems sets UniqueItems value.
func (s *Schema) WithUniqueItems(val bool) *Schema {
	s.UniqueItems = &val
	return s
}

// WithContains sets Contains value.
func (s *Schema) WithContains(val SchemaOrBool) *Schema {
	s.Contains = &val
	return s
}

// ContainsEns ensures returned Contains is not nil.
func (s *Schema) ContainsEns() *SchemaOrBool {
	if s.Contains == nil {
		s.Contains = new(SchemaOrBool)
	}

	return s.Contains
}

// WithMaxProperties sets MaxProperties value.
func (s *Schema) WithMaxProperties(val int64) *Schema {
	s.MaxProperties = &val
	return s
}

// WithMinProperties sets MinProperties value.
func (s *Schema) WithMinProperties(val int64) *Schema {
	s.MinProperties = val
	return s
}

// WithRequired sets Required value.
func (s *Schema) WithRequired(val ...string) *Schema {
	s.Required = val
	return s
}

// WithAdditionalProperties sets AdditionalProperties value.
func (s *Schema) WithAdditionalProperties(val SchemaOrBool) *Schema {
	s.AdditionalProperties = &val
	return s
}

// AdditionalPropertiesEns ensures returned AdditionalProperties is not nil.
func (s *Schema) AdditionalPropertiesEns() *SchemaOrBool {
	if s.AdditionalProperties == nil {
		s.AdditionalProperties = new(SchemaOrBool)
	}

	return s.AdditionalProperties
}

// WithDefinitions sets Definitions value.
func (s *Schema) WithDefinitions(val map[string]SchemaOrBool) *Schema {
	s.Definitions = val
	return s
}

// WithDefinitionsItem sets Definitions item value.
func (s *Schema) WithDefinitionsItem(key string, val SchemaOrBool) *Schema {
	if s.Definitions == nil {
		s.Definitions = make(map[string]SchemaOrBool, 1)
	}

	s.Definitions[key] = val

	return s
}

// WithProperties sets Properties value.
func (s *Schema) WithProperties(val map[string]SchemaOrBool) *Schema {
	s.Properties = val
	return s
}

// WithPropertiesItem sets Properties item value.
func (s *Schema) WithPropertiesItem(key string, val SchemaOrBool) *Schema {
	if s.Properties == nil {
		s.Properties = make(map[string]SchemaOrBool, 1)
	}

	s.Properties[key] = val

	return s
}

// WithPatternProperties sets PatternProperties value.
func (s *Schema) WithPatternProperties(val map[string]SchemaOrBool) *Schema {
	s.PatternProperties = val
	return s
}

// WithPatternPropertiesItem sets PatternProperties item value.
func (s *Schema) WithPatternPropertiesItem(key string, val SchemaOrBool) *Schema {
	if s.PatternProperties == nil {
		s.PatternProperties = make(map[string]SchemaOrBool, 1)
	}

	s.PatternProperties[key] = val

	return s
}

// WithDependencies sets Dependencies value.
func (s *Schema) WithDependencies(val map[string]DependenciesAdditionalProperties) *Schema {
	s.Dependencies = val
	return s
}

// WithDependenciesItem sets Dependencies item value.
func (s *Schema) WithDependenciesItem(key string, val DependenciesAdditionalProperties) *Schema {
	if s.Dependencies == nil {
		s.Dependencies = make(map[string]DependenciesAdditionalProperties, 1)
	}

	s.Dependencies[key] = val

	return s
}

// WithPropertyNames sets PropertyNames value.
func (s *Schema) WithPropertyNames(val SchemaOrBool) *Schema {
	s.PropertyNames = &val
	return s
}

// PropertyNamesEns ensures returned PropertyNames is not nil.
func (s *Schema) PropertyNamesEns() *SchemaOrBool {
	if s.PropertyNames == nil {
		s.PropertyNames = new(SchemaOrBool)
	}

	return s.PropertyNames
}

// WithConst sets Const value.
func (s *Schema) WithConst(val interface{}) *Schema {
	s.Const = &val
	return s
}

// WithEnum sets Enum value.
func (s *Schema) WithEnum(val ...interface{}) *Schema {
	s.Enum = val
	return s
}

// WithType sets Type value.
func (s *Schema) WithType(val Type) *Schema {
	s.Type = &val
	return s
}

// TypeEns ensures returned Type is not nil.
func (s *Schema) TypeEns() *Type {
	if s.Type == nil {
		s.Type = new(Type)
	}

	return s.Type
}

// WithFormat sets Format value.
func (s *Schema) WithFormat(val string) *Schema {
	s.Format = &val
	return s
}

// WithContentMediaType sets ContentMediaType value.
func (s *Schema) WithContentMediaType(val string) *Schema {
	s.ContentMediaType = &val
	return s
}

// WithContentEncoding sets ContentEncoding value.
func (s *Schema) WithContentEncoding(val string) *Schema {
	s.ContentEncoding = &val
	return s
}

// WithIf sets If value.
func (s *Schema) WithIf(val SchemaOrBool) *Schema {
	s.If = &val
	return s
}

// IfEns ensures returned If is not nil.
func (s *Schema) IfEns() *SchemaOrBool {
	if s.If == nil {
		s.If = new(SchemaOrBool)
	}

	return s.If
}

// WithThen sets Then value.
func (s *Schema) WithThen(val SchemaOrBool) *Schema {
	s.Then = &val
	return s
}

// ThenEns ensures returned Then is not nil.
func (s *Schema) ThenEns() *SchemaOrBool {
	if s.Then == nil {
		s.Then = new(SchemaOrBool)
	}

	return s.Then
}

// WithElse sets Else value.
func (s *Schema) WithElse(val SchemaOrBool) *Schema {
	s.Else = &val
	return s
}

// ElseEns ensures returned Else is not nil.
func (s *Schema) ElseEns() *SchemaOrBool {
	if s.Else == nil {
		s.Else = new(SchemaOrBool)
	}

	return s.Else
}

// WithAllOf sets AllOf value.
func (s *Schema) WithAllOf(val ...SchemaOrBool) *Schema {
	s.AllOf = val
	return s
}

// WithAnyOf sets AnyOf value.
func (s *Schema) WithAnyOf(val ...SchemaOrBool) *Schema {
	s.AnyOf = val
	return s
}

// WithOneOf sets OneOf value.
func (s *Schema) WithOneOf(val ...SchemaOrBool) *Schema {
	s.OneOf = val
	return s
}

// WithNot sets Not value.
func (s *Schema) WithNot(val SchemaOrBool) *Schema {
	s.Not = &val
	return s
}

// NotEns ensures returned Not is not nil.
func (s *Schema) NotEns() *SchemaOrBool {
	if s.Not == nil {
		s.Not = new(SchemaOrBool)
	}

	return s.Not
}

// WithExtraProperties sets ExtraProperties value.
func (s *Schema) WithExtraProperties(val map[string]interface{}) *Schema {
	s.ExtraProperties = val
	return s
}

// WithExtraPropertiesItem sets ExtraProperties item value.
func (s *Schema) WithExtraPropertiesItem(key string, val interface{}) *Schema {
	if s.ExtraProperties == nil {
		s.ExtraProperties = make(map[string]interface{}, 1)
	}

	s.ExtraProperties[key] = val

	return s
}

type marshalSchema Schema

var knownKeysSchema = []string{
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
func (s *Schema) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSchema(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if ms.Default == nil {
		if _, ok := rawMap["default"]; ok {
			var v interface{}
			ms.Default = &v
		}
	}

	if ms.Const == nil {
		if _, ok := rawMap["const"]; ok {
			var v interface{}
			ms.Const = &v
		}
	}

	for _, key := range knownKeysSchema {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if ms.ExtraProperties == nil {
			ms.ExtraProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ms.ExtraProperties[key] = val
	}

	*s = Schema(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Schema) MarshalJSON() ([]byte, error) {
	if len(s.ExtraProperties) == 0 {
		return json.Marshal(marshalSchema(s))
	}

	return marshalUnion(marshalSchema(s), s.ExtraProperties)
}

// SchemaOrBool structure is generated from "#".
//
// Core schema meta-schema.
type SchemaOrBool struct {
	TypeObject  *Schema `json:"-"`
	TypeBoolean *bool   `json:"-"`
}

// WithTypeObject sets TypeObject value.
func (s *SchemaOrBool) WithTypeObject(val Schema) *SchemaOrBool {
	s.TypeObject = &val
	return s
}

// TypeObjectEns ensures returned TypeObject is not nil.
func (s *SchemaOrBool) TypeObjectEns() *Schema {
	if s.TypeObject == nil {
		s.TypeObject = new(Schema)
	}

	return s.TypeObject
}

// WithTypeBoolean sets TypeBoolean value.
func (s *SchemaOrBool) WithTypeBoolean(val bool) *SchemaOrBool {
	s.TypeBoolean = &val
	return s
}

// UnmarshalJSON decodes JSON.
func (s *SchemaOrBool) UnmarshalJSON(data []byte) error {
	var err error

	typeValid := false

	if !typeValid {
		err = json.Unmarshal(data, &s.TypeObject)
		if err != nil {
			s.TypeObject = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		err = json.Unmarshal(data, &s.TypeBoolean)
		if err != nil {
			s.TypeBoolean = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		return err
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaOrBool) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.TypeObject, s.TypeBoolean)
}

// Items structure is generated from "#[object]->items".
type Items struct {
	SchemaOrBool *SchemaOrBool  `json:"-"`
	SchemaArray  []SchemaOrBool `json:"-"`
}

// WithSchemaOrBool sets SchemaOrBool value.
func (i *Items) WithSchemaOrBool(val SchemaOrBool) *Items {
	i.SchemaOrBool = &val
	return i
}

// SchemaOrBoolEns ensures returned SchemaOrBool is not nil.
func (i *Items) SchemaOrBoolEns() *SchemaOrBool {
	if i.SchemaOrBool == nil {
		i.SchemaOrBool = new(SchemaOrBool)
	}

	return i.SchemaOrBool
}

// WithSchemaArray sets SchemaArray value.
func (i *Items) WithSchemaArray(val ...SchemaOrBool) *Items {
	i.SchemaArray = val
	return i
}

// UnmarshalJSON decodes JSON.
func (i *Items) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &i.SchemaOrBool)
	if err != nil {
		anyOfErrors["SchemaOrBool"] = err
		i.SchemaOrBool = nil
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
	return marshalUnion(i.SchemaOrBool, i.SchemaArray)
}

// DependenciesAdditionalProperties structure is generated from "#[object]->dependencies->additionalProperties".
type DependenciesAdditionalProperties struct {
	SchemaOrBool *SchemaOrBool `json:"-"`
	StringArray  []string      `json:"-"`
}

// WithSchemaOrBool sets SchemaOrBool value.
func (d *DependenciesAdditionalProperties) WithSchemaOrBool(val SchemaOrBool) *DependenciesAdditionalProperties {
	d.SchemaOrBool = &val
	return d
}

// SchemaOrBoolEns ensures returned SchemaOrBool is not nil.
func (d *DependenciesAdditionalProperties) SchemaOrBoolEns() *SchemaOrBool {
	if d.SchemaOrBool == nil {
		d.SchemaOrBool = new(SchemaOrBool)
	}

	return d.SchemaOrBool
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

	err = json.Unmarshal(data, &d.SchemaOrBool)
	if err != nil {
		anyOfErrors["SchemaOrBool"] = err
		d.SchemaOrBool = nil
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
	return marshalUnion(d.SchemaOrBool, d.StringArray)
}

// Type structure is generated from "#[object]->type".
type Type struct {
	SimpleTypes             *SimpleType  `json:"-"`
	SliceOfSimpleTypeValues []SimpleType `json:"-"`
}

// WithSimpleTypes sets SimpleTypes value.
func (t *Type) WithSimpleTypes(val SimpleType) *Type {
	t.SimpleTypes = &val
	return t
}

// WithSliceOfSimpleTypeValues sets SliceOfSimpleTypeValues value.
func (t *Type) WithSliceOfSimpleTypeValues(val ...SimpleType) *Type {
	t.SliceOfSimpleTypeValues = val
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

	err = json.Unmarshal(data, &t.SliceOfSimpleTypeValues)
	if err != nil {
		anyOfErrors["SliceOfSimpleTypeValues"] = err
		t.SliceOfSimpleTypeValues = nil
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
	return marshalUnion(t.SimpleTypes, t.SliceOfSimpleTypeValues)
}

// SimpleType is an enum type.
type SimpleType string

// SimpleType values enumeration.
const (
	Array = SimpleType("array")
	Boolean = SimpleType("boolean")
	Integer = SimpleType("integer")
	Null = SimpleType("null")
	Number = SimpleType("number")
	Object = SimpleType("object")
	String = SimpleType("string")
)

// MarshalJSON encodes JSON.
func (i SimpleType) MarshalJSON() ([]byte, error) {
	switch i {
	case Array:
	case Boolean:
	case Integer:
	case Null:
	case Number:
	case Object:
	case String:

	default:
		return nil, fmt.Errorf("unexpected SimpleType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SimpleType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SimpleType(ii)

	switch v {
	case Array:
	case Boolean:
	case Integer:
	case Null:
	case Number:
	case Object:
	case String:

	default:
		return fmt.Errorf("unexpected SimpleType value: %v", v)
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
