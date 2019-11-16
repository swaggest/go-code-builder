package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

// MarshalJSON encodes JSON.
func (i OperationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationOneOf1(i), i.MapOfAnything)
}

// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Message *Message         `json:"-"`
	OneOf1  *OperationOneOf1 `json:"-"`
}

// MarshalJSON encodes JSON.
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Message, i.OneOf1)
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

// MarshalJSON encodes JSON.
func (i Stream) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStream(i), i.MapOfAnything)
}

// StreamFramingOneOf0 structure is generated from "#/definitions/stream->framing/oneOf/0".
type StreamFramingOneOf0 struct {
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

type marshalStreamFramingOneOf0 StreamFramingOneOf0

var (
	// constStreamFramingOneOf0 is unconditionally added to JSON.
	constStreamFramingOneOf0 = json.RawMessage(`{"type":"chunked"}`)
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constStreamFramingOneOf0, marshalStreamFramingOneOf0(i))
}

// StreamFramingOneOf1 structure is generated from "#/definitions/stream->framing/oneOf/1".
type StreamFramingOneOf1 struct {

}

var (
	// constStreamFramingOneOf1 is unconditionally added to JSON.
	constStreamFramingOneOf1 = json.RawMessage(`{"type":"sse","delimiter":"\\n\\n"}`)
)

// MarshalJSON encodes JSON.
func (i StreamFramingOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constStreamFramingOneOf1)
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

// MarshalJSON encodes JSON.
func (i StreamFraming) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfAnything, i.AdditionalProperties, i.OneOf0, i.OneOf1)
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

// MarshalJSON encodes JSON.
func (i Reference) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalReference(i), i.AdditionalProperties)
}

// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalUserPassword UserPassword

var (
	// constUserPassword is unconditionally added to JSON.
	constUserPassword = json.RawMessage(`{"type":"userPassword"}`)
)

// MarshalJSON encodes JSON.
func (i UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(constUserPassword, marshalUserPassword(i), i.MapOfAnything)
}

// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	In            APIKeyIn               `json:"in,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKey APIKey

var (
	// constAPIKey is unconditionally added to JSON.
	constAPIKey = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON.
func (i APIKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKey, marshalAPIKey(i), i.MapOfAnything)
}

// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalX509 X509

var (
	// constX509 is unconditionally added to JSON.
	constX509 = json.RawMessage(`{"type":"X509"}`)
)

// MarshalJSON encodes JSON.
func (i X509) MarshalJSON() ([]byte, error) {
	return marshalUnion(constX509, marshalX509(i), i.MapOfAnything)
}

// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalSymmetricEncryption SymmetricEncryption

var (
	// constSymmetricEncryption is unconditionally added to JSON.
	constSymmetricEncryption = json.RawMessage(`{"type":"symmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (i SymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constSymmetricEncryption, marshalSymmetricEncryption(i), i.MapOfAnything)
}

// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAsymmetricEncryption AsymmetricEncryption

var (
	// constAsymmetricEncryption is unconditionally added to JSON.
	constAsymmetricEncryption = json.RawMessage(`{"type":"asymmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (i AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsymmetricEncryption, marshalAsymmetricEncryption(i), i.MapOfAnything)
}

// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

var (
	// constNonBearerHTTPSecurityScheme is unconditionally added to JSON.
	constNonBearerHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON.
func (i NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constNonBearerHTTPSecurityScheme, marshalNonBearerHTTPSecurityScheme(i), i.MapOfAnything)
}

// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

var (
	// constBearerHTTPSecurityScheme is unconditionally added to JSON.
	constBearerHTTPSecurityScheme = json.RawMessage(`{"scheme":"bearer","type":"http"}`)
)

// MarshalJSON encodes JSON.
func (i BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBearerHTTPSecurityScheme, marshalBearerHTTPSecurityScheme(i), i.MapOfAnything)
}

// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name,omitempty"`
	In            APIKeyHTTPSecuritySchemeIn `json:"in,omitempty"`
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

var (
	// constAPIKeyHTTPSecurityScheme is unconditionally added to JSON.
	constAPIKeyHTTPSecurityScheme = json.RawMessage(`{"type":"httpApiKey"}`)
)

// MarshalJSON encodes JSON.
func (i APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeyHTTPSecurityScheme, marshalAPIKeyHTTPSecurityScheme(i), i.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
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

// MarshalJSON encodes JSON.
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.UserPassword, i.APIKey, i.X509, i.SymmetricEncryption, i.AsymmetricEncryption, i.HTTPSecurityScheme)
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from "#/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
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
