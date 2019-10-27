package entities

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
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

// UnmarshalJSON decodes JSON.
func (i *AsyncAPI) UnmarshalJSON(data []byte) error {
	ii := marshalAsyncAPI(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = AsyncAPI(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Info) UnmarshalJSON(data []byte) error {
	ii := marshalInfo(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"title",
			"version",
			"description",
			"termsOfService",
			"contact",
			"license",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Info(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Contact) UnmarshalJSON(data []byte) error {
	ii := marshalContact(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"url",
			"email",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Contact(ii)

	return err
}


// License structure is generated from "#/definitions/license".
type License struct {
	Name          string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL           string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnything map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

type marshalLicense License

// UnmarshalJSON decodes JSON.
func (i *License) UnmarshalJSON(data []byte) error {
	ii := marshalLicense(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"url",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = License(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Server) UnmarshalJSON(data []byte) error {
	ii := marshalServer(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"url",
			"description",
			"scheme",
			"schemeVersion",
			"variables",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Server(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	ii := marshalServerVariable(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"enum",
			"default",
			"description",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ServerVariable(ii)

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &i.MapOfAnything, // ^x-
			regex: &i.MapOfTopicItemValues, // ^[^.]
		},
		jsonData: data,
	}.unmarshal()

	return err
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

// UnmarshalJSON decodes JSON.
func (i *TopicItem) UnmarshalJSON(data []byte) error {
	ii := marshalTopicItem(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$ref",
			"parameters",
			"publish",
			"subscribe",
			"deprecated",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = TopicItem(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Parameter) UnmarshalJSON(data []byte) error {
	ii := marshalParameter(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"description",
			"name",
			"schema",
			"$ref",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Parameter(ii)

	return err
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
	Example       interface{}            `json:"example,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalMessage Message

// UnmarshalJSON decodes JSON.
func (i *Message) UnmarshalJSON(data []byte) error {
	ii := marshalMessage(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$ref",
			"headers",
			"payload",
			"tags",
			"summary",
			"description",
			"externalDocs",
			"deprecated",
			"example",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Message(ii)

	return err
}


// Tag structure is generated from "#/definitions/tag".
type Tag struct {
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalTag Tag

// UnmarshalJSON decodes JSON.
func (i *Tag) UnmarshalJSON(data []byte) error {
	ii := marshalTag(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"description",
			"externalDocs",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Tag(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	ii := marshalExternalDocs(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"description",
			"url",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ExternalDocs(ii)

	return err
}


// OperationOneOf1 structure is generated from "#/definitions/operation/oneOf/1".
type OperationOneOf1 struct {
	OneOf         []Message              `json:"oneOf,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalOperationOneOf1 OperationOneOf1

// UnmarshalJSON decodes JSON.
func (i *OperationOneOf1) UnmarshalJSON(data []byte) error {
	ii := marshalOperationOneOf1(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"oneOf",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = OperationOneOf1(ii)

	return err
}


// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Message *Message         `json:"-"`
	OneOf1  *OperationOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Operation) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Message, &i.OneOf1}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Message = nil
	}

	if mayUnmarshal[1] == nil {
		i.OneOf1 = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Stream) UnmarshalJSON(data []byte) error {
	ii := marshalStream(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"framing",
			"read",
			"write",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Stream(ii)

	return err
}


// StreamFramingOneOf0 structure is generated from "#/definitions/stream->framing/oneOf/0".
type StreamFramingOneOf0 struct {
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

type marshalStreamFramingOneOf0 StreamFramingOneOf0

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf0) UnmarshalJSON(data []byte) error {
	ii := marshalStreamFramingOneOf0(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"chunked"` {
		return fmt.Errorf(`bad or missing const value for "type" ("chunked" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = StreamFramingOneOf0(ii)

	return err
}


// StreamFramingOneOf1 structure is generated from "#/definitions/stream->framing/oneOf/1".
type StreamFramingOneOf1 struct {

}

// UnmarshalJSON decodes JSON.
func (i *StreamFramingOneOf1) UnmarshalJSON(data []byte) error {
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"sse"` {
		return fmt.Errorf(`bad or missing const value for "type" ("sse" expected, %v received)`, v)
	}

	if v, ok := constValues["delimiter"]; !ok || string(v) != `"\\n\\n"` {
		return fmt.Errorf(`bad or missing const value for "delimiter" ("\\n\\n" expected, %v received)`, v)
	}

	return err
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
	mayUnmarshal := []interface{}{&i.OneOf0, &i.OneOf1}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &i.MapOfAnything, // ^x-
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.OneOf0 = nil
	}

	if mayUnmarshal[1] == nil {
		i.OneOf1 = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Events) UnmarshalJSON(data []byte) error {
	ii := marshalEvents(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"receive",
			"send",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Events(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Reference) UnmarshalJSON(data []byte) error {
	ii := marshalReference(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$ref",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Reference(ii)

	return err
}


// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalUserPassword UserPassword

// UnmarshalJSON decodes JSON.
func (i *UserPassword) UnmarshalJSON(data []byte) error {
	ii := marshalUserPassword(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"userPassword"` {
		return fmt.Errorf(`bad or missing const value for "type" ("userPassword" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = UserPassword(ii)

	return err
}


// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	In            APIKeyIn               `json:"in,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKey APIKey

// UnmarshalJSON decodes JSON.
func (i *APIKey) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKey(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"in",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"apiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("apiKey" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = APIKey(ii)

	return err
}


// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalX509 X509

// UnmarshalJSON decodes JSON.
func (i *X509) UnmarshalJSON(data []byte) error {
	ii := marshalX509(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"X509"` {
		return fmt.Errorf(`bad or missing const value for "type" ("X509" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = X509(ii)

	return err
}


// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalSymmetricEncryption SymmetricEncryption

// UnmarshalJSON decodes JSON.
func (i *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	ii := marshalSymmetricEncryption(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"symmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("symmetricEncryption" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = SymmetricEncryption(ii)

	return err
}


// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAsymmetricEncryption AsymmetricEncryption

// UnmarshalJSON decodes JSON.
func (i *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	ii := marshalAsymmetricEncryption(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"asymmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("asymmetricEncryption" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = AsymmetricEncryption(ii)

	return err
}


// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

// UnmarshalJSON decodes JSON.
func (i *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalNonBearerHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"scheme",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = NonBearerHTTPSecurityScheme(ii)

	return err
}


// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

// UnmarshalJSON decodes JSON.
func (i *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalBearerHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"bearerFormat",
			"description",
			"scheme",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["scheme"]; !ok || string(v) != `"bearer"` {
		return fmt.Errorf(`bad or missing const value for "scheme" ("bearer" expected, %v received)`, v)
	}

	if v, ok := constValues["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = BearerHTTPSecurityScheme(ii)

	return err
}


// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name,omitempty"`
	In            APIKeyHTTPSecuritySchemeIn `json:"in,omitempty"`
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

// UnmarshalJSON decodes JSON.
func (i *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKeyHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"name",
			"in",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"httpApiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("httpApiKey" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = APIKeyHTTPSecurityScheme(ii)

	return err
}


// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.NonBearerHTTPSecurityScheme, &i.BearerHTTPSecurityScheme, &i.APIKeyHTTPSecurityScheme}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.NonBearerHTTPSecurityScheme = nil
	}

	if mayUnmarshal[1] == nil {
		i.BearerHTTPSecurityScheme = nil
	}

	if mayUnmarshal[2] == nil {
		i.APIKeyHTTPSecurityScheme = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.UserPassword, &i.APIKey, &i.X509, &i.SymmetricEncryption, &i.AsymmetricEncryption, &i.HTTPSecurityScheme}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.UserPassword = nil
	}

	if mayUnmarshal[1] == nil {
		i.APIKey = nil
	}

	if mayUnmarshal[2] == nil {
		i.X509 = nil
	}

	if mayUnmarshal[3] == nil {
		i.SymmetricEncryption = nil
	}

	if mayUnmarshal[4] == nil {
		i.AsymmetricEncryption = nil
	}

	if mayUnmarshal[5] == nil {
		i.HTTPSecurityScheme = nil
	}

	return err
}


// ComponentsSecuritySchemesAZAZ09 structure is generated from "#/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Reference, &i.SecurityScheme}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.SecurityScheme = nil
	}

	return err
}


// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                       map[string]interface{}                     `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsSecuritySchemesAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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

type unionMap struct {
	mustUnmarshal        []interface{}
	mayUnmarshal         []interface{}
	ignoreKeys           []string
	patternProperties    map[*regexp.Regexp]interface{}
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

	if len(u.patternProperties) == 0 && u.additionalProperties == nil {
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

	if len(u.patternProperties) != 0 {
		err = u.unmarshalPatternProperties(m)
		if err != nil {
			return err
		}
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
func (u unionMap) unmarshalPatternProperties(m map[string]*json.RawMessage) error {
	patternMapsRaw := make(map[*regexp.Regexp][]byte, len(u.patternProperties))

	// Iterating map and filling pattern properties sub maps.
	for key, val := range m {
		matched := false
		ok := false
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex := range u.patternProperties {
			if regex.MatchString(key) {
				matched = true

				var subMap []byte

				if subMap, ok = patternMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)

				if val != nil {
					subMap = append(subMap, []byte(*val)...)
				} else {
					subMap = append(subMap, []byte("null")...)
				}

				subMap = append(subMap, '}')

				patternMapsRaw[regex] = subMap
			}
		}

		// Remove from properties map if matched to at least one regex.
		if matched {
			delete(m, key)
		}
	}

	for regex := range u.patternProperties {
		if subMap, ok := patternMapsRaw[regex]; !ok {
			continue
		} else {
			err := json.Unmarshal(subMap, u.patternProperties[regex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
