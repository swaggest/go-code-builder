package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

// AsyncAPI structure is generated from "#".
//
// AsyncAPI 1.2.0 schema.
type AsyncAPI struct {
	Asyncapi      AsyncAPIAsyncapi       `json:"asyncapi,omitempty"`     // The AsyncAPI specification version of this document.
	Info          *Info                  `json:"info,omitempty"`         // General information about the API.
	BaseTopic     string                 `json:"baseTopic,omitempty"`    // The base topic to the API. Example: 'hitch'.
	Servers       []Server               `json:"servers,omitempty"`
	Topics        *Topics                `json:"topics,omitempty"`       // Relative paths to the individual topics. They must be relative to the 'baseTopic'.
	Stream        *Stream                `json:"stream,omitempty"`       // Stream Object
	Events        *Events                `json:"events,omitempty"`       // Events Object
	Components    *Components            `json:"components,omitempty"`   // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags          []Tag                  `json:"tags,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalAsyncAPI AsyncAPI

var ignoreKeysAsyncAPI = []string{
	"asyncapi",
	"info",
	"baseTopic",
	"servers",
	"topics",
	"stream",
	"events",
	"components",
	"tags",
	"security",
	"externalDocs",
}

// UnmarshalJSON decodes JSON.
func (i *AsyncAPI) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAsyncAPI(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAsyncAPI {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = AsyncAPI(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAsyncAPI(i), i.MapOfAnything)
}

// Info structure is generated from "#/definitions/info".
//
// General information about the API.
type Info struct {
	Title          string                 `json:"title,omitempty"`          // A unique and precise title of the API.
	Version        string                 `json:"version,omitempty"`        // A semantic version number of the API.
	Description    string                 `json:"description,omitempty"`    // A longer description of the API. Should be different from the title. CommonMark is allowed.
	TermsOfService string                 `json:"termsOfService,omitempty"` // A URL to the Terms of Service for the API. MUST be in the format of a URL.
	Contact        *Contact               `json:"contact,omitempty"`        // Contact information for the owners of the API.
	License        *License               `json:"license,omitempty"`
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: ^x-
}

type marshalInfo Info

var ignoreKeysInfo = []string{
	"title",
	"version",
	"description",
	"termsOfService",
	"contact",
	"license",
}

// UnmarshalJSON decodes JSON.
func (i *Info) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalInfo(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysInfo {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Info(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnything)
}

// Contact structure is generated from "#/definitions/contact".
//
// Contact information for the owners of the API.
type Contact struct {
	Name          string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	URL           string                 `json:"url,omitempty"`   // The URL pointing to the contact information.
	Email         string                 `json:"email,omitempty"` // The email address of the contact person/organization.
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalContact Contact

var ignoreKeysContact = []string{
	"name",
	"url",
	"email",
}

// UnmarshalJSON decodes JSON.
func (i *Contact) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalContact(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysContact {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Contact(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(i), i.MapOfAnything)
}

// License structure is generated from "#/definitions/license".
type License struct {
	Name          string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL           string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnything map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

type marshalLicense License

var ignoreKeysLicense = []string{
	"name",
	"url",
}

// UnmarshalJSON decodes JSON.
func (i *License) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalLicense(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysLicense {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = License(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(i), i.MapOfAnything)
}

// Server structure is generated from "#/definitions/server".
//
// An object representing a Server.
type Server struct {
	URL           string                    `json:"url,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Scheme        ServerScheme              `json:"scheme,omitempty"`        // The transfer protocol.
	SchemeVersion string                    `json:"schemeVersion,omitempty"`
	Variables     map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnything map[string]interface{}    `json:"-"`                       // Key must match pattern: ^x-
}

type marshalServer Server

var ignoreKeysServer = []string{
	"url",
	"description",
	"scheme",
	"schemeVersion",
	"variables",
}

// UnmarshalJSON decodes JSON.
func (i *Server) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalServer(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysServer {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Server(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(i), i.MapOfAnything)
}

// ServerVariable structure is generated from "#/definitions/serverVariable".
//
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalServerVariable ServerVariable

var ignoreKeysServerVariable = []string{
	"enum",
	"default",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalServerVariable(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysServerVariable {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ServerVariable(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(i), i.MapOfAnything)
}

// Topics structure is generated from "#/definitions/topics".
//
// Relative paths to the individual topics. They must be relative to the 'baseTopic'.
type Topics struct {
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	MapOfTopicItemValues map[string]TopicItem   `json:"-"` // Key must match pattern: ^[^.]
}

// UnmarshalJSON decodes JSON.
func (i *Topics) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if i.MapOfAnything == nil {
				i.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfAnything[key] = val
		}

		if regex.MatchString(key) {
			matched = true

			if i.MapOfTopicItemValues == nil {
				i.MapOfTopicItemValues = make(map[string]TopicItem, 1)
			}

			var val TopicItem

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfTopicItemValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Topics) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfAnything, i.MapOfTopicItemValues)
}

// TopicItem structure is generated from "#/definitions/topicItem".
type TopicItem struct {
	Ref           string                 `json:"$ref,omitempty"`
	Parameters    []Parameter            `json:"parameters,omitempty"`
	Publish       *Operation             `json:"publish,omitempty"`
	Subscribe     *Operation             `json:"subscribe,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: ^x-
}

type marshalTopicItem TopicItem

var ignoreKeysTopicItem = []string{
	"$ref",
	"parameters",
	"publish",
	"subscribe",
	"deprecated",
}

// UnmarshalJSON decodes JSON.
func (i *TopicItem) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalTopicItem(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysTopicItem {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = TopicItem(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i TopicItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTopicItem(i), i.MapOfAnything)
}

// Parameter structure is generated from "#/definitions/parameter".
type Parameter struct {
	Description   string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name          string                 `json:"name,omitempty"`        // The name of the parameter.
	Schema        map[string]interface{} `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	Ref           string                 `json:"$ref,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalParameter Parameter

var ignoreKeysParameter = []string{
	"description",
	"name",
	"schema",
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (i *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalParameter(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysParameter {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Parameter(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(i), i.MapOfAnything)
}

// Message structure is generated from "#/definitions/message".
type Message struct {
	Ref           string                 `json:"$ref,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty"`      // A deterministic version of a JSON Schema object.
	Payload       map[string]interface{} `json:"payload,omitempty"`      // A deterministic version of a JSON Schema object.
	Tags          []Tag                  `json:"tags,omitempty"`
	Summary       string                 `json:"summary,omitempty"`      // A brief summary of the message.
	Description   string                 `json:"description,omitempty"`  // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	Deprecated    bool                   `json:"deprecated,omitempty"`
	Example       *interface{}           `json:"example,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalMessage Message

var ignoreKeysMessage = []string{
	"$ref",
	"headers",
	"payload",
	"tags",
	"summary",
	"description",
	"externalDocs",
	"deprecated",
	"example",
}

// UnmarshalJSON decodes JSON.
func (i *Message) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMessage(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			ii.Example = &v
		}
	}

	for _, key := range ignoreKeysMessage {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Message(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessage(i), i.MapOfAnything)
}

// Tag structure is generated from "#/definitions/tag".
type Tag struct {
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalTag Tag

var ignoreKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

// UnmarshalJSON decodes JSON.
func (i *Tag) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalTag(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysTag {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Tag(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(i), i.MapOfAnything)
}

// ExternalDocs structure is generated from "#/definitions/externalDocs".
//
// information about external documentation.
type ExternalDocs struct {
	Description   string                 `json:"description,omitempty"`
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalExternalDocs ExternalDocs

var ignoreKeysExternalDocs = []string{
	"description",
	"url",
}

// UnmarshalJSON decodes JSON.
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalExternalDocs(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysExternalDocs {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ExternalDocs(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(i), i.MapOfAnything)
}

// OperationOneOf1 structure is generated from "#/definitions/operation/oneOf/1".
type OperationOneOf1 struct {
	OneOf         []Message              `json:"oneOf,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalOperationOneOf1 OperationOneOf1

var ignoreKeysOperationOneOf1 = []string{
	"oneOf",
}

// UnmarshalJSON decodes JSON.
func (i *OperationOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOperationOneOf1(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOperationOneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = OperationOneOf1(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationOneOf1(i), i.MapOfAnything)
}

// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Message         *Message         `json:"-"`
	OperationOneOf1 *OperationOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Operation) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Message)
	if err != nil {
		i.Message = nil
	}

	err = json.Unmarshal(data, &i.OperationOneOf1)
	if err != nil {
		i.OperationOneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Message, i.OperationOneOf1)
}

// Stream structure is generated from "#/definitions/stream".
//
// Stream Object.
type Stream struct {
	Framing       *StreamFraming         `json:"framing,omitempty"` // Stream Framing Object
	Read          []Message              `json:"read,omitempty"`    // Stream Read Object
	Write         []Message              `json:"write,omitempty"`   // Stream Write Object
	MapOfAnything map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

type marshalStream Stream

var ignoreKeysStream = []string{
	"framing",
	"read",
	"write",
}

// UnmarshalJSON decodes JSON.
func (i *Stream) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalStream(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysStream {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Stream(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Stream) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStream(i), i.MapOfAnything)
}

// StreamFramingOneOf0 structure is generated from "#/definitions/stream->framing/oneOf/0".
type StreamFramingOneOf0 struct {
	Type      StreamFramingOneOf0Type      `json:"type,omitempty"`
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

// StreamFramingOneOf1 structure is generated from "#/definitions/stream->framing/oneOf/1".
type StreamFramingOneOf1 struct {
	Type      StreamFramingOneOf1Type      `json:"type,omitempty"`
	Delimiter StreamFramingOneOf1Delimiter `json:"delimiter,omitempty"`
}

// StreamFraming structure is generated from "#/definitions/stream->framing".
//
// Stream Framing Object.
type StreamFraming struct {
	StreamFramingOneOf0  *StreamFramingOneOf0   `json:"-"`
	StreamFramingOneOf1  *StreamFramingOneOf1   `json:"-"`
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *StreamFraming) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.StreamFramingOneOf0)
	if err != nil {
		i.StreamFramingOneOf0 = nil
	}

	err = json.Unmarshal(data, &i.StreamFramingOneOf1)
	if err != nil {
		i.StreamFramingOneOf1 = nil
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if i.MapOfAnything == nil {
				i.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i StreamFraming) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfAnything, i.AdditionalProperties, i.StreamFramingOneOf0, i.StreamFramingOneOf1)
}

// Events structure is generated from "#/definitions/events".
//
// Events Object.
type Events struct {
	Receive       []Message              `json:"receive,omitempty"` // Events Receive Object
	Send          []Message              `json:"send,omitempty"`    // Events Send Object
	MapOfAnything map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

type marshalEvents Events

var ignoreKeysEvents = []string{
	"receive",
	"send",
}

// UnmarshalJSON decodes JSON.
func (i *Events) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalEvents(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysEvents {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Events(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Events) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalEvents(i), i.MapOfAnything)
}

// Components structure is generated from "#/definitions/components".
//
// An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
type Components struct {
	Schemas         map[string]map[string]interface{} `json:"schemas,omitempty"`         // JSON objects describing schemas the API uses.
	Messages        map[string]Message                `json:"messages,omitempty"`        // JSON objects describing the messages being consumed and produced by the API.
	SecuritySchemes *ComponentsSecuritySchemes        `json:"securitySchemes,omitempty"`
	Parameters      map[string]Parameter              `json:"parameters,omitempty"`      // JSON objects describing re-usable topic parameters.
}

// Reference structure is generated from "#/definitions/Reference".
type Reference struct {
	Ref                  string                 `json:"$ref,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties
}

type marshalReference Reference

var ignoreKeysReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (i *Reference) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalReference(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysReference {
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

	*i = Reference(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Reference) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalReference(i))
	}
	return marshalUnion(marshalReference(i), i.AdditionalProperties)
}

// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Type          UserPasswordType       `json:"type,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalUserPassword UserPassword

var ignoreKeysUserPassword = []string{
	"type",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *UserPassword) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalUserPassword(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysUserPassword {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = UserPassword(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalUserPassword(i), i.MapOfAnything)
}

// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	Type          APIKeyType             `json:"type,omitempty"`
	In            APIKeyIn               `json:"in,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKey APIKey

var ignoreKeysAPIKey = []string{
	"type",
	"in",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *APIKey) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAPIKey(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAPIKey {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = APIKey(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i APIKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAPIKey(i), i.MapOfAnything)
}

// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Type          X509Type               `json:"type,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalX509 X509

var ignoreKeysX509 = []string{
	"type",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *X509) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalX509(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysX509 {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = X509(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i X509) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalX509(i), i.MapOfAnything)
}

// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Type          SymmetricEncryptionType `json:"type,omitempty"`
	Description   string                  `json:"description,omitempty"`
	MapOfAnything map[string]interface{}  `json:"-"`                     // Key must match pattern: ^x-
}

type marshalSymmetricEncryption SymmetricEncryption

var ignoreKeysSymmetricEncryption = []string{
	"type",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalSymmetricEncryption(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysSymmetricEncryption {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = SymmetricEncryption(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i SymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSymmetricEncryption(i), i.MapOfAnything)
}

// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Type          AsymmetricEncryptionType `json:"type,omitempty"`
	Description   string                   `json:"description,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAsymmetricEncryption AsymmetricEncryption

var ignoreKeysAsymmetricEncryption = []string{
	"type",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAsymmetricEncryption(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAsymmetricEncryption {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = AsymmetricEncryption(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAsymmetricEncryption(i), i.MapOfAnything)
}

// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                          `json:"scheme,omitempty"`
	Description   string                          `json:"description,omitempty"`
	Type          NonBearerHTTPSecuritySchemeType `json:"type,omitempty"`
	MapOfAnything map[string]interface{}          `json:"-"`                     // Key must match pattern: ^x-
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

var ignoreKeysNonBearerHTTPSecurityScheme = []string{
	"scheme",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalNonBearerHTTPSecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysNonBearerHTTPSecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = NonBearerHTTPSecurityScheme(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNonBearerHTTPSecurityScheme(i), i.MapOfAnything)
}

// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	Scheme        BearerHTTPSecuritySchemeScheme `json:"scheme,omitempty"`
	BearerFormat  string                         `json:"bearerFormat,omitempty"`
	Type          BearerHTTPSecuritySchemeType   `json:"type,omitempty"`
	Description   string                         `json:"description,omitempty"`
	MapOfAnything map[string]interface{}         `json:"-"`                      // Key must match pattern: ^x-
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

var ignoreKeysBearerHTTPSecurityScheme = []string{
	"scheme",
	"bearerFormat",
	"type",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalBearerHTTPSecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysBearerHTTPSecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = BearerHTTPSecurityScheme(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalBearerHTTPSecurityScheme(i), i.MapOfAnything)
}

// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Type          APIKeyHTTPSecuritySchemeType `json:"type,omitempty"`
	Name          string                       `json:"name,omitempty"`
	In            APIKeyHTTPSecuritySchemeIn   `json:"in,omitempty"`
	Description   string                       `json:"description,omitempty"`
	MapOfAnything map[string]interface{}       `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

var ignoreKeysAPIKeyHTTPSecurityScheme = []string{
	"type",
	"name",
	"in",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAPIKeyHTTPSecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAPIKeyHTTPSecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = APIKeyHTTPSecurityScheme(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAPIKeyHTTPSecurityScheme(i), i.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.NonBearerHTTPSecurityScheme)
	if err != nil {
		i.NonBearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &i.BearerHTTPSecurityScheme)
	if err != nil {
		i.BearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &i.APIKeyHTTPSecurityScheme)
	if err != nil {
		i.APIKeyHTTPSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.NonBearerHTTPSecurityScheme, i.BearerHTTPSecurityScheme, i.APIKeyHTTPSecurityScheme)
}

// SecurityScheme structure is generated from "#/definitions/SecurityScheme".
type SecurityScheme struct {
	UserPassword         *UserPassword         `json:"-"`
	APIKey               *APIKey               `json:"-"`
	X509                 *X509                 `json:"-"`
	SymmetricEncryption  *SymmetricEncryption  `json:"-"`
	AsymmetricEncryption *AsymmetricEncryption `json:"-"`
	HTTPSecurityScheme   *HTTPSecurityScheme   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.UserPassword)
	if err != nil {
		i.UserPassword = nil
	}

	err = json.Unmarshal(data, &i.APIKey)
	if err != nil {
		i.APIKey = nil
	}

	err = json.Unmarshal(data, &i.X509)
	if err != nil {
		i.X509 = nil
	}

	err = json.Unmarshal(data, &i.SymmetricEncryption)
	if err != nil {
		i.SymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &i.AsymmetricEncryption)
	if err != nil {
		i.AsymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &i.HTTPSecurityScheme)
	if err != nil {
		i.HTTPSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.UserPassword, i.APIKey, i.X509, i.SymmetricEncryption, i.AsymmetricEncryption, i.HTTPSecurityScheme)
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from "#/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.SecurityScheme)
	if err != nil {
		i.SecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSecuritySchemesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                       map[string]interface{}                     `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if i.MapOfComponentsSecuritySchemesAZAZ09Values == nil {
				i.MapOfComponentsSecuritySchemesAZAZ09Values = make(map[string]ComponentsSecuritySchemesAZAZ09, 1)
			}

			var val ComponentsSecuritySchemesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsSecuritySchemesAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsSecuritySchemesAZAZ09Values, i.AdditionalProperties)
}

// AsyncAPIAsyncapi is an enum type.
type AsyncAPIAsyncapi string

// AsyncAPIAsyncapi values enumeration.
const (
	AsyncAPIAsyncapi100 = AsyncAPIAsyncapi("1.0.0")
	AsyncAPIAsyncapi110 = AsyncAPIAsyncapi("1.1.0")
	AsyncAPIAsyncapi120 = AsyncAPIAsyncapi("1.2.0")
)

// MarshalJSON encodes JSON.
func (i AsyncAPIAsyncapi) MarshalJSON() ([]byte, error) {
	switch i {
	case AsyncAPIAsyncapi100:
	case AsyncAPIAsyncapi110:
	case AsyncAPIAsyncapi120:

	default:
		return nil, fmt.Errorf("unexpected AsyncAPIAsyncapi value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AsyncAPIAsyncapi) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AsyncAPIAsyncapi(ii)

	switch v {
	case AsyncAPIAsyncapi100:
	case AsyncAPIAsyncapi110:
	case AsyncAPIAsyncapi120:

	default:
		return fmt.Errorf("unexpected AsyncAPIAsyncapi value: %v", v)
	}

	*i = v

	return nil
}

// ServerScheme is an enum type.
type ServerScheme string

// ServerScheme values enumeration.
const (
	ServerSchemeKafka = ServerScheme("kafka")
	ServerSchemeKafkaSecure = ServerScheme("kafka-secure")
	ServerSchemeAmqp = ServerScheme("amqp")
	ServerSchemeAmqps = ServerScheme("amqps")
	ServerSchemeMqtt = ServerScheme("mqtt")
	ServerSchemeMqtts = ServerScheme("mqtts")
	ServerSchemeSecureMqtt = ServerScheme("secure-mqtt")
	ServerSchemeWs = ServerScheme("ws")
	ServerSchemeWss = ServerScheme("wss")
	ServerSchemeStomp = ServerScheme("stomp")
	ServerSchemeStomps = ServerScheme("stomps")
	ServerSchemeJms = ServerScheme("jms")
	ServerSchemeHTTP = ServerScheme("http")
	ServerSchemeHTTPS = ServerScheme("https")
)

// MarshalJSON encodes JSON.
func (i ServerScheme) MarshalJSON() ([]byte, error) {
	switch i {
	case ServerSchemeKafka:
	case ServerSchemeKafkaSecure:
	case ServerSchemeAmqp:
	case ServerSchemeAmqps:
	case ServerSchemeMqtt:
	case ServerSchemeMqtts:
	case ServerSchemeSecureMqtt:
	case ServerSchemeWs:
	case ServerSchemeWss:
	case ServerSchemeStomp:
	case ServerSchemeStomps:
	case ServerSchemeJms:
	case ServerSchemeHTTP:
	case ServerSchemeHTTPS:

	default:
		return nil, fmt.Errorf("unexpected ServerScheme value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ServerScheme) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ServerScheme(ii)

	switch v {
	case ServerSchemeKafka:
	case ServerSchemeKafkaSecure:
	case ServerSchemeAmqp:
	case ServerSchemeAmqps:
	case ServerSchemeMqtt:
	case ServerSchemeMqtts:
	case ServerSchemeSecureMqtt:
	case ServerSchemeWs:
	case ServerSchemeWss:
	case ServerSchemeStomp:
	case ServerSchemeStomps:
	case ServerSchemeJms:
	case ServerSchemeHTTP:
	case ServerSchemeHTTPS:

	default:
		return fmt.Errorf("unexpected ServerScheme value: %v", v)
	}

	*i = v

	return nil
}

// StreamFramingOneOf0Type is an enum type.
type StreamFramingOneOf0Type string

// StreamFramingOneOf0Type values enumeration.
const (
	StreamFramingOneOf0TypeChunked = StreamFramingOneOf0Type("chunked")
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf0Type) MarshalJSON() ([]byte, error) {
	switch i {
	case StreamFramingOneOf0TypeChunked:

	default:
		return nil, fmt.Errorf("unexpected StreamFramingOneOf0Type value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf0Type) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := StreamFramingOneOf0Type(ii)

	switch v {
	case StreamFramingOneOf0TypeChunked:

	default:
		return fmt.Errorf("unexpected StreamFramingOneOf0Type value: %v", v)
	}

	*i = v

	return nil
}

// StreamFramingOneOf0Delimiter is an enum type.
type StreamFramingOneOf0Delimiter string

// StreamFramingOneOf0Delimiter values enumeration.
const (
	StreamFramingOneOf0DelimiterRN = StreamFramingOneOf0Delimiter(`\r\n`)
	StreamFramingOneOf0DelimiterN = StreamFramingOneOf0Delimiter(`\n`)
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf0Delimiter) MarshalJSON() ([]byte, error) {
	switch i {
	case StreamFramingOneOf0DelimiterRN:
	case StreamFramingOneOf0DelimiterN:

	default:
		return nil, fmt.Errorf("unexpected StreamFramingOneOf0Delimiter value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf0Delimiter) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := StreamFramingOneOf0Delimiter(ii)

	switch v {
	case StreamFramingOneOf0DelimiterRN:
	case StreamFramingOneOf0DelimiterN:

	default:
		return fmt.Errorf("unexpected StreamFramingOneOf0Delimiter value: %v", v)
	}

	*i = v

	return nil
}

// StreamFramingOneOf1Type is an enum type.
type StreamFramingOneOf1Type string

// StreamFramingOneOf1Type values enumeration.
const (
	StreamFramingOneOf1TypeSse = StreamFramingOneOf1Type("sse")
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf1Type) MarshalJSON() ([]byte, error) {
	switch i {
	case StreamFramingOneOf1TypeSse:

	default:
		return nil, fmt.Errorf("unexpected StreamFramingOneOf1Type value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf1Type) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := StreamFramingOneOf1Type(ii)

	switch v {
	case StreamFramingOneOf1TypeSse:

	default:
		return fmt.Errorf("unexpected StreamFramingOneOf1Type value: %v", v)
	}

	*i = v

	return nil
}

// StreamFramingOneOf1Delimiter is an enum type.
type StreamFramingOneOf1Delimiter string

// StreamFramingOneOf1Delimiter values enumeration.
const (
	StreamFramingOneOf1DelimiterNN = StreamFramingOneOf1Delimiter(`\n\n`)
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf1Delimiter) MarshalJSON() ([]byte, error) {
	switch i {
	case StreamFramingOneOf1DelimiterNN:

	default:
		return nil, fmt.Errorf("unexpected StreamFramingOneOf1Delimiter value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf1Delimiter) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := StreamFramingOneOf1Delimiter(ii)

	switch v {
	case StreamFramingOneOf1DelimiterNN:

	default:
		return fmt.Errorf("unexpected StreamFramingOneOf1Delimiter value: %v", v)
	}

	*i = v

	return nil
}

// UserPasswordType is an enum type.
type UserPasswordType string

// UserPasswordType values enumeration.
const (
	UserPasswordTypeUserPassword = UserPasswordType("userPassword")
)

// MarshalJSON encodes JSON.
func (i UserPasswordType) MarshalJSON() ([]byte, error) {
	switch i {
	case UserPasswordTypeUserPassword:

	default:
		return nil, fmt.Errorf("unexpected UserPasswordType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *UserPasswordType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := UserPasswordType(ii)

	switch v {
	case UserPasswordTypeUserPassword:

	default:
		return fmt.Errorf("unexpected UserPasswordType value: %v", v)
	}

	*i = v

	return nil
}

// APIKeyType is an enum type.
type APIKeyType string

// APIKeyType values enumeration.
const (
	APIKeyTypeAPIKey = APIKeyType("apiKey")
)

// MarshalJSON encodes JSON.
func (i APIKeyType) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyTypeAPIKey:

	default:
		return nil, fmt.Errorf("unexpected APIKeyType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeyType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeyType(ii)

	switch v {
	case APIKeyTypeAPIKey:

	default:
		return fmt.Errorf("unexpected APIKeyType value: %v", v)
	}

	*i = v

	return nil
}

// APIKeyIn is an enum type.
type APIKeyIn string

// APIKeyIn values enumeration.
const (
	APIKeyInUser = APIKeyIn("user")
	APIKeyInPassword = APIKeyIn("password")
)

// MarshalJSON encodes JSON.
func (i APIKeyIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyInUser:
	case APIKeyInPassword:

	default:
		return nil, fmt.Errorf("unexpected APIKeyIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeyIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeyIn(ii)

	switch v {
	case APIKeyInUser:
	case APIKeyInPassword:

	default:
		return fmt.Errorf("unexpected APIKeyIn value: %v", v)
	}

	*i = v

	return nil
}

// X509Type is an enum type.
type X509Type string

// X509Type values enumeration.
const (
	X509TypeX509 = X509Type("X509")
)

// MarshalJSON encodes JSON.
func (i X509Type) MarshalJSON() ([]byte, error) {
	switch i {
	case X509TypeX509:

	default:
		return nil, fmt.Errorf("unexpected X509Type value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *X509Type) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := X509Type(ii)

	switch v {
	case X509TypeX509:

	default:
		return fmt.Errorf("unexpected X509Type value: %v", v)
	}

	*i = v

	return nil
}

// SymmetricEncryptionType is an enum type.
type SymmetricEncryptionType string

// SymmetricEncryptionType values enumeration.
const (
	SymmetricEncryptionTypeSymmetricEncryption = SymmetricEncryptionType("symmetricEncryption")
)

// MarshalJSON encodes JSON.
func (i SymmetricEncryptionType) MarshalJSON() ([]byte, error) {
	switch i {
	case SymmetricEncryptionTypeSymmetricEncryption:

	default:
		return nil, fmt.Errorf("unexpected SymmetricEncryptionType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SymmetricEncryptionType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SymmetricEncryptionType(ii)

	switch v {
	case SymmetricEncryptionTypeSymmetricEncryption:

	default:
		return fmt.Errorf("unexpected SymmetricEncryptionType value: %v", v)
	}

	*i = v

	return nil
}

// AsymmetricEncryptionType is an enum type.
type AsymmetricEncryptionType string

// AsymmetricEncryptionType values enumeration.
const (
	AsymmetricEncryptionTypeAsymmetricEncryption = AsymmetricEncryptionType("asymmetricEncryption")
)

// MarshalJSON encodes JSON.
func (i AsymmetricEncryptionType) MarshalJSON() ([]byte, error) {
	switch i {
	case AsymmetricEncryptionTypeAsymmetricEncryption:

	default:
		return nil, fmt.Errorf("unexpected AsymmetricEncryptionType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AsymmetricEncryptionType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AsymmetricEncryptionType(ii)

	switch v {
	case AsymmetricEncryptionTypeAsymmetricEncryption:

	default:
		return fmt.Errorf("unexpected AsymmetricEncryptionType value: %v", v)
	}

	*i = v

	return nil
}

// NonBearerHTTPSecuritySchemeType is an enum type.
type NonBearerHTTPSecuritySchemeType string

// NonBearerHTTPSecuritySchemeType values enumeration.
const (
	NonBearerHTTPSecuritySchemeTypeHTTP = NonBearerHTTPSecuritySchemeType("http")
)

// MarshalJSON encodes JSON.
func (i NonBearerHTTPSecuritySchemeType) MarshalJSON() ([]byte, error) {
	switch i {
	case NonBearerHTTPSecuritySchemeTypeHTTP:

	default:
		return nil, fmt.Errorf("unexpected NonBearerHTTPSecuritySchemeType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *NonBearerHTTPSecuritySchemeType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := NonBearerHTTPSecuritySchemeType(ii)

	switch v {
	case NonBearerHTTPSecuritySchemeTypeHTTP:

	default:
		return fmt.Errorf("unexpected NonBearerHTTPSecuritySchemeType value: %v", v)
	}

	*i = v

	return nil
}

// BearerHTTPSecuritySchemeScheme is an enum type.
type BearerHTTPSecuritySchemeScheme string

// BearerHTTPSecuritySchemeScheme values enumeration.
const (
	BearerHTTPSecuritySchemeSchemeBearer = BearerHTTPSecuritySchemeScheme("bearer")
)

// MarshalJSON encodes JSON.
func (i BearerHTTPSecuritySchemeScheme) MarshalJSON() ([]byte, error) {
	switch i {
	case BearerHTTPSecuritySchemeSchemeBearer:

	default:
		return nil, fmt.Errorf("unexpected BearerHTTPSecuritySchemeScheme value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *BearerHTTPSecuritySchemeScheme) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := BearerHTTPSecuritySchemeScheme(ii)

	switch v {
	case BearerHTTPSecuritySchemeSchemeBearer:

	default:
		return fmt.Errorf("unexpected BearerHTTPSecuritySchemeScheme value: %v", v)
	}

	*i = v

	return nil
}

// BearerHTTPSecuritySchemeType is an enum type.
type BearerHTTPSecuritySchemeType string

// BearerHTTPSecuritySchemeType values enumeration.
const (
	BearerHTTPSecuritySchemeTypeHTTP = BearerHTTPSecuritySchemeType("http")
)

// MarshalJSON encodes JSON.
func (i BearerHTTPSecuritySchemeType) MarshalJSON() ([]byte, error) {
	switch i {
	case BearerHTTPSecuritySchemeTypeHTTP:

	default:
		return nil, fmt.Errorf("unexpected BearerHTTPSecuritySchemeType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *BearerHTTPSecuritySchemeType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := BearerHTTPSecuritySchemeType(ii)

	switch v {
	case BearerHTTPSecuritySchemeTypeHTTP:

	default:
		return fmt.Errorf("unexpected BearerHTTPSecuritySchemeType value: %v", v)
	}

	*i = v

	return nil
}

// APIKeyHTTPSecuritySchemeType is an enum type.
type APIKeyHTTPSecuritySchemeType string

// APIKeyHTTPSecuritySchemeType values enumeration.
const (
	APIKeyHTTPSecuritySchemeTypeHTTPAPIKey = APIKeyHTTPSecuritySchemeType("httpApiKey")
)

// MarshalJSON encodes JSON.
func (i APIKeyHTTPSecuritySchemeType) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyHTTPSecuritySchemeTypeHTTPAPIKey:

	default:
		return nil, fmt.Errorf("unexpected APIKeyHTTPSecuritySchemeType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeyHTTPSecuritySchemeType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeyHTTPSecuritySchemeType(ii)

	switch v {
	case APIKeyHTTPSecuritySchemeTypeHTTPAPIKey:

	default:
		return fmt.Errorf("unexpected APIKeyHTTPSecuritySchemeType value: %v", v)
	}

	*i = v

	return nil
}

// APIKeyHTTPSecuritySchemeIn is an enum type.
type APIKeyHTTPSecuritySchemeIn string

// APIKeyHTTPSecuritySchemeIn values enumeration.
const (
	APIKeyHTTPSecuritySchemeInHeader = APIKeyHTTPSecuritySchemeIn("header")
	APIKeyHTTPSecuritySchemeInQuery = APIKeyHTTPSecuritySchemeIn("query")
	APIKeyHTTPSecuritySchemeInCookie = APIKeyHTTPSecuritySchemeIn("cookie")
)

// MarshalJSON encodes JSON.
func (i APIKeyHTTPSecuritySchemeIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyHTTPSecuritySchemeInHeader:
	case APIKeyHTTPSecuritySchemeInQuery:
	case APIKeyHTTPSecuritySchemeInCookie:

	default:
		return nil, fmt.Errorf("unexpected APIKeyHTTPSecuritySchemeIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeyHTTPSecuritySchemeIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeyHTTPSecuritySchemeIn(ii)

	switch v {
	case APIKeyHTTPSecuritySchemeInHeader:
	case APIKeyHTTPSecuritySchemeInQuery:
	case APIKeyHTTPSecuritySchemeInCookie:

	default:
		return fmt.Errorf("unexpected APIKeyHTTPSecuritySchemeIn value: %v", v)
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
// Regular expressions for pattern properties.
var (
	regexX = regexp.MustCompile("^x-")
	regex = regexp.MustCompile("^[^.]")
	regexAZAZ09 = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
