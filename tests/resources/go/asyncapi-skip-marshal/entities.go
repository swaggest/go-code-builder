package entities

import (
	"encoding/json"
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


// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Message *Message         `json:"-"`
	OneOf1  *OperationOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Operation) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Message)
	if err != nil {
		i.Message = nil
	}

	err = json.Unmarshal(data, &i.OneOf1)
	if err != nil {
		i.OneOf1 = nil
	}

	return nil
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


// StreamFramingOneOf0 structure is generated from "#/definitions/stream->framing/oneOf/0".
type StreamFramingOneOf0 struct {
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

type marshalStreamFramingOneOf0 StreamFramingOneOf0

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf0) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalStreamFramingOneOf0(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"chunked"` {
		return fmt.Errorf(`bad or missing const value for "type" ("chunked" expected, %s received)`, v)
	}

	delete(m, "type")

	*i = StreamFramingOneOf0(ii)

	return nil
}


// StreamFramingOneOf1 structure is generated from "#/definitions/stream->framing/oneOf/1".
type StreamFramingOneOf1 struct {

}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"sse"` {
		return fmt.Errorf(`bad or missing const value for "type" ("sse" expected, %s received)`, v)
	}

	delete(m, "type")

	if v, ok := m["delimiter"]; !ok || string(v) != `"\\n\\n"` {
		return fmt.Errorf(`bad or missing const value for "delimiter" ("\\n\\n" expected, %s received)`, v)
	}

	delete(m, "delimiter")

	return nil
}


// StreamFraming structure is generated from "#/definitions/stream->framing".
//
// Stream Framing Object.
type StreamFraming struct {
	OneOf0               *StreamFramingOneOf0   `json:"-"`
	OneOf1               *StreamFramingOneOf1   `json:"-"`
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *StreamFraming) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.OneOf0)
	if err != nil {
		i.OneOf0 = nil
	}

	err = json.Unmarshal(data, &i.OneOf1)
	if err != nil {
		i.OneOf1 = nil
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


// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalUserPassword UserPassword

var ignoreKeysUserPassword = []string{
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"userPassword"` {
		return fmt.Errorf(`bad or missing const value for "type" ("userPassword" expected, %s received)`, v)
	}

	delete(m, "type")

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


// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	In            APIKeyIn               `json:"in,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKey APIKey

var ignoreKeysAPIKey = []string{
	"in",
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"apiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("apiKey" expected, %s received)`, v)
	}

	delete(m, "type")

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


// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalX509 X509

var ignoreKeysX509 = []string{
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"X509"` {
		return fmt.Errorf(`bad or missing const value for "type" ("X509" expected, %s received)`, v)
	}

	delete(m, "type")

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


// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalSymmetricEncryption SymmetricEncryption

var ignoreKeysSymmetricEncryption = []string{
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"symmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("symmetricEncryption" expected, %s received)`, v)
	}

	delete(m, "type")

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


// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAsymmetricEncryption AsymmetricEncryption

var ignoreKeysAsymmetricEncryption = []string{
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"asymmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("asymmetricEncryption" expected, %s received)`, v)
	}

	delete(m, "type")

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


// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
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

	if v, ok := m["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %s received)`, v)
	}

	delete(m, "type")

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


// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

var ignoreKeysBearerHTTPSecurityScheme = []string{
	"bearerFormat",
	"description",
	"scheme",
	"type",
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

	if v, ok := m["scheme"]; !ok || string(v) != `"bearer"` {
		return fmt.Errorf(`bad or missing const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(m, "scheme")

	if v, ok := m["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %s received)`, v)
	}

	delete(m, "type")

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


// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name,omitempty"`
	In            APIKeyHTTPSecuritySchemeIn `json:"in,omitempty"`
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

var ignoreKeysAPIKeyHTTPSecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
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

	if v, ok := m["type"]; !ok || string(v) != `"httpApiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("httpApiKey" expected, %s received)`, v)
	}

	delete(m, "type")

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


// AsyncAPIAsyncapi is an enum type.
type AsyncAPIAsyncapi string

// AsyncAPIAsyncapi values enumeration.
const (
	AsyncAPIAsyncapi100 = AsyncAPIAsyncapi("1.0.0")
	AsyncAPIAsyncapi110 = AsyncAPIAsyncapi("1.1.0")
	AsyncAPIAsyncapi120 = AsyncAPIAsyncapi("1.2.0")
)

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

// StreamFramingOneOf0Delimiter is an enum type.
type StreamFramingOneOf0Delimiter string

// StreamFramingOneOf0Delimiter values enumeration.
const (
	StreamFramingOneOf0DelimiterRN = StreamFramingOneOf0Delimiter(`\r\n`)
	StreamFramingOneOf0DelimiterN = StreamFramingOneOf0Delimiter(`\n`)
)

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

// APIKeyIn is an enum type.
type APIKeyIn string

// APIKeyIn values enumeration.
const (
	APIKeyInUser = APIKeyIn("user")
	APIKeyInPassword = APIKeyIn("password")
)

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

// APIKeyHTTPSecuritySchemeIn is an enum type.
type APIKeyHTTPSecuritySchemeIn string

// APIKeyHTTPSecuritySchemeIn values enumeration.
const (
	APIKeyHTTPSecuritySchemeInHeader = APIKeyHTTPSecuritySchemeIn("header")
	APIKeyHTTPSecuritySchemeInQuery = APIKeyHTTPSecuritySchemeIn("query")
	APIKeyHTTPSecuritySchemeInCookie = APIKeyHTTPSecuritySchemeIn("cookie")
)

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

// Regular expressions for pattern properties.
var (
	regexX = regexp.MustCompile("^x-")
	regex = regexp.MustCompile("^[^.]")
	regexAZAZ09 = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
