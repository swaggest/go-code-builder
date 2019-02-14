package entities

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// Schema structure is generated from #
// Core schema meta-schema
type Schema struct {
	ID                   string                                      `json:"$id,omitempty"`
	*Schema                                     `json:"-"`
	Ref                  string                                      `json:"$ref,omitempty"`
	Comment              string                                      `json:"$comment,omitempty"`
	Title                string                                      `json:"title,omitempty"`
	Description          string                                      `json:"description,omitempty"`
	Default              interface{}                                 `json:"default,omitempty"`
	ReadOnly             bool                                        `json:"readOnly,omitempty"`
	Examples             []interface{}                               `json:"examples,omitempty"`
	MultipleOf           float64                                     `json:"multipleOf,omitempty"`
	Maximum              float64                                     `json:"maximum,omitempty"`
	ExclusiveMaximum     float64                                     `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                     `json:"minimum,omitempty"`
	ExclusiveMinimum     float64                                     `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                       `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefault0                 `json:"minLength,omitempty"`
	Pattern              string                                      `json:"pattern,omitempty"`
	AdditionalItems      *Schema                                     `json:"additionalItems,omitempty"`      // Core schema meta-schema
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             int64                                       `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefault0                 `json:"minItems,omitempty"`
	UniqueItems          bool                                        `json:"uniqueItems,omitempty"`
	Contains             *Schema                                     `json:"contains,omitempty"`             // Core schema meta-schema
	MaxProperties        int64                                       `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefault0                 `json:"minProperties,omitempty"`
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
	Format               string                                      `json:"format,omitempty"`
	ContentMediaType     string                                      `json:"contentMediaType,omitempty"`
	ContentEncoding      string                                      `json:"contentEncoding,omitempty"`
	If                   *Schema                                     `json:"if,omitempty"`                   // Core schema meta-schema
	Then                 *Schema                                     `json:"then,omitempty"`                 // Core schema meta-schema
	Else                 *Schema                                     `json:"else,omitempty"`                 // Core schema meta-schema
	AllOf                []Schema                                    `json:"allOf,omitempty"`
	AnyOf                []Schema                                    `json:"anyOf,omitempty"`
	OneOf                []Schema                                    `json:"oneOf,omitempty"`
	Not                  *Schema                                     `json:"not,omitempty"`                  // Core schema meta-schema
	Type0                *Type0                                      `json:"-"`
	Type1                *Type1                                      `json:"-"`
}

type marshalSchema Schema

// UnmarshalJSON decodes JSON
func (i *Schema) UnmarshalJSON(data []byte) error {
	ii := marshalSchema(*i)
	mayUnmarshal := []interface{}{&ii., &ii.Type0, &ii.Type1}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
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
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		ii. = nil
	}
	if mayUnmarshal[1] == nil {
		ii.Type0 = nil
	}
	if mayUnmarshal[2] == nil {
		ii.Type1 = nil
	}
	if err != nil {
		return err
	}
	*i = Schema(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchema(i), i., i.Type0, i.Type1)
}

// NonNegativeIntegerDefault0 structure is generated from #/definitions/nonNegativeIntegerDefault0
type NonNegativeIntegerDefault0 struct {
	Int64 *int64 `json:"-"`
}

type marshalNonNegativeIntegerDefault0 NonNegativeIntegerDefault0

// UnmarshalJSON decodes JSON
func (i *NonNegativeIntegerDefault0) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.Int64},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i NonNegativeIntegerDefault0) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNonNegativeIntegerDefault0(i), i.Int64)
}

// Items structure is generated from #->items
type Items struct {
	Schema *Schema  `json:"-"`
	AnyOf1 []Schema `json:"-"`
}

type marshalItems Items

// UnmarshalJSON decodes JSON
func (i *Items) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}
	if mayUnmarshal[1] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Items) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalItems(i), i.Schema, i.AnyOf1)
}

// DependenciesAdditionalProperties structure is generated from #->dependencies->additionalProperties
type DependenciesAdditionalProperties struct {
	Schema *Schema  `json:"-"`
	AnyOf1 []string `json:"-"`
}

type marshalDependenciesAdditionalProperties DependenciesAdditionalProperties

// UnmarshalJSON decodes JSON
func (i *DependenciesAdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}
	if mayUnmarshal[1] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i DependenciesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalDependenciesAdditionalProperties(i), i.Schema, i.AnyOf1)
}

// Type structure is generated from #->type
type Type struct {
	AnyOf1 []interface{} `json:"-"`
}

type marshalType Type

// UnmarshalJSON decodes JSON
func (i *Type) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Type) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalType(i), i.AnyOf1)
}

// Type0 structure is generated from #/type/0
// Core schema meta-schema
type Type0 struct {
	ID                   string                                      `json:"$id,omitempty"`
	Schema               string                                      `json:"$schema,omitempty"`
	Ref                  string                                      `json:"$ref,omitempty"`
	Comment              string                                      `json:"$comment,omitempty"`
	Title                string                                      `json:"title,omitempty"`
	Description          string                                      `json:"description,omitempty"`
	Default              interface{}                                 `json:"default,omitempty"`
	ReadOnly             bool                                        `json:"readOnly,omitempty"`
	Examples             []interface{}                               `json:"examples,omitempty"`
	MultipleOf           float64                                     `json:"multipleOf,omitempty"`
	Maximum              float64                                     `json:"maximum,omitempty"`
	ExclusiveMaximum     float64                                     `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                     `json:"minimum,omitempty"`
	ExclusiveMinimum     float64                                     `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                       `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefault0                 `json:"minLength,omitempty"`
	Pattern              string                                      `json:"pattern,omitempty"`
	AdditionalItems      *Schema                                     `json:"additionalItems,omitempty"`      // Core schema meta-schema
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             int64                                       `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefault0                 `json:"minItems,omitempty"`
	UniqueItems          bool                                        `json:"uniqueItems,omitempty"`
	Contains             *Schema                                     `json:"contains,omitempty"`             // Core schema meta-schema
	MaxProperties        int64                                       `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefault0                 `json:"minProperties,omitempty"`
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
	Format               string                                      `json:"format,omitempty"`
	ContentMediaType     string                                      `json:"contentMediaType,omitempty"`
	ContentEncoding      string                                      `json:"contentEncoding,omitempty"`
	If                   *Schema                                     `json:"if,omitempty"`                   // Core schema meta-schema
	Then                 *Schema                                     `json:"then,omitempty"`                 // Core schema meta-schema
	Else                 *Schema                                     `json:"else,omitempty"`                 // Core schema meta-schema
	AllOf                []Schema                                    `json:"allOf,omitempty"`
	AnyOf                []Schema                                    `json:"anyOf,omitempty"`
	OneOf                []Schema                                    `json:"oneOf,omitempty"`
	Not                  *Schema                                     `json:"not,omitempty"`                  // Core schema meta-schema
}

// Type1 structure is generated from #/type/1
// Core schema meta-schema
type Type1 struct {
	ID                   string                                      `json:"$id,omitempty"`
	Schema               string                                      `json:"$schema,omitempty"`
	Ref                  string                                      `json:"$ref,omitempty"`
	Comment              string                                      `json:"$comment,omitempty"`
	Title                string                                      `json:"title,omitempty"`
	Description          string                                      `json:"description,omitempty"`
	Default              interface{}                                 `json:"default,omitempty"`
	ReadOnly             bool                                        `json:"readOnly,omitempty"`
	Examples             []interface{}                               `json:"examples,omitempty"`
	MultipleOf           float64                                     `json:"multipleOf,omitempty"`
	Maximum              float64                                     `json:"maximum,omitempty"`
	ExclusiveMaximum     float64                                     `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                     `json:"minimum,omitempty"`
	ExclusiveMinimum     float64                                     `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                       `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefault0                 `json:"minLength,omitempty"`
	Pattern              string                                      `json:"pattern,omitempty"`
	AdditionalItems      *Schema                                     `json:"additionalItems,omitempty"`      // Core schema meta-schema
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             int64                                       `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefault0                 `json:"minItems,omitempty"`
	UniqueItems          bool                                        `json:"uniqueItems,omitempty"`
	Contains             *Schema                                     `json:"contains,omitempty"`             // Core schema meta-schema
	MaxProperties        int64                                       `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefault0                 `json:"minProperties,omitempty"`
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
	Format               string                                      `json:"format,omitempty"`
	ContentMediaType     string                                      `json:"contentMediaType,omitempty"`
	ContentEncoding      string                                      `json:"contentEncoding,omitempty"`
	If                   *Schema                                     `json:"if,omitempty"`                   // Core schema meta-schema
	Then                 *Schema                                     `json:"then,omitempty"`                 // Core schema meta-schema
	Else                 *Schema                                     `json:"else,omitempty"`                 // Core schema meta-schema
	AllOf                []Schema                                    `json:"allOf,omitempty"`
	AnyOf                []Schema                                    `json:"anyOf,omitempty"`
	OneOf                []Schema                                    `json:"oneOf,omitempty"`
	Not                  *Schema                                     `json:"not,omitempty"`                  // Core schema meta-schema
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
