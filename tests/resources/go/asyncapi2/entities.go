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
// AsyncAPI 2.0.0 schema.
type AsyncAPI struct {
	// A unique id representing the application.
	// Format: uri.
	ID                 string                 `json:"id,omitempty"`
	// General information about the API.
	// Required.
	Info               Info                   `json:"info"`
	Servers            map[string]Server      `json:"servers,omitempty"`
	DefaultContentType string                 `json:"defaultContentType,omitempty"`
	Channels           map[string]ChannelItem `json:"channels"`                     // Required.
	Components         *Components            `json:"components,omitempty"`         // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags               []Tag                  `json:"tags,omitempty"`
	ExternalDocs       *ExternalDocs          `json:"externalDocs,omitempty"`       // Information about external documentation.
	MapOfAnything      map[string]interface{} `json:"-"`                            // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalAsyncAPI AsyncAPI

var ignoreKeysAsyncAPI = []string{
	"id",
	"info",
	"servers",
	"defaultContentType",
	"channels",
	"components",
	"tags",
	"externalDocs",
	"asyncapi",
}

// UnmarshalJSON decodes JSON.
func (v *AsyncAPI) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAsyncAPI(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["asyncapi"]; !ok || string(v) != `"2.0.0"` {
		return fmt.Errorf(`bad or missing const value for "asyncapi" ("2.0.0" expected, %s received)`, v)
	}

	delete(m, "asyncapi")

	for _, key := range ignoreKeysAsyncAPI {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = AsyncAPI(vv)

	return nil
}

var (
	// constAsyncAPI is unconditionally added to JSON.
	constAsyncAPI = json.RawMessage(`{"asyncapi":"2.0.0"}`)
)

// MarshalJSON encodes JSON.
func (v AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsyncAPI, marshalAsyncAPI(v), v.MapOfAnything)
}

// Info structure is generated from "#/definitions/info".
//
// General information about the API.
type Info struct {
	// A unique and precise title of the API.
	// Required.
	Title          string                 `json:"title"`
	// A semantic version number of the API.
	// Required.
	Version        string                 `json:"version"`
	Description    string                 `json:"description,omitempty"`    // A longer description of the API. Should be different from the title. CommonMark is allowed.
	// A URL to the Terms of Service for the API. MUST be in the format of a URL.
	// Format: uri.
	TermsOfService string                 `json:"termsOfService,omitempty"`
	Contact        *Contact               `json:"contact,omitempty"`        // Contact information for the owners of the API.
	License        *License               `json:"license,omitempty"`
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
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
func (v *Info) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalInfo(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Info(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(v), v.MapOfAnything)
}

// Contact structure is generated from "#/definitions/contact".
//
// Contact information for the owners of the API.
type Contact struct {
	Name          string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	// The URL pointing to the contact information.
	// Format: uri.
	URL           string                 `json:"url,omitempty"`
	// The email address of the contact person/organization.
	// Format: email.
	Email         string                 `json:"email,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalContact Contact

var ignoreKeysContact = []string{
	"name",
	"url",
	"email",
}

// UnmarshalJSON decodes JSON.
func (v *Contact) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalContact(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Contact(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(v), v.MapOfAnything)
}

// License structure is generated from "#/definitions/license".
type License struct {
	// The name of the license type. It's encouraged to use an OSI compatible license.
	// Required.
	Name          string                 `json:"name"`
	// The URL pointing to the license.
	// Format: uri.
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`             // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalLicense License

var ignoreKeysLicense = []string{
	"name",
	"url",
}

// UnmarshalJSON decodes JSON.
func (v *License) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalLicense(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = License(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(v), v.MapOfAnything)
}

// Server structure is generated from "#/definitions/server".
//
// An object representing a Server.
type Server struct {
	URL             string                    `json:"url"`                       // Required.
	Description     string                    `json:"description,omitempty"`
	// The transfer protocol.
	// Required.
	Protocol        string                    `json:"protocol"`
	ProtocolVersion string                    `json:"protocolVersion,omitempty"`
	Variables       map[string]ServerVariable `json:"variables,omitempty"`
	Security        []map[string][]string     `json:"security,omitempty"`
	Bindings        *ServerBindingsObject     `json:"bindings,omitempty"`
	MapOfAnything   map[string]interface{}    `json:"-"`                         // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalServer Server

var ignoreKeysServer = []string{
	"url",
	"description",
	"protocol",
	"protocolVersion",
	"variables",
	"security",
	"bindings",
}

// UnmarshalJSON decodes JSON.
func (v *Server) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalServer(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Server(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(v), v.MapOfAnything)
}

// ServerVariable structure is generated from "#/definitions/serverVariable".
//
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Examples      []string               `json:"examples,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalServerVariable ServerVariable

var ignoreKeysServerVariable = []string{
	"enum",
	"default",
	"description",
	"examples",
}

// UnmarshalJSON decodes JSON.
func (v *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalServerVariable(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ServerVariable(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(v), v.MapOfAnything)
}

// ServerBindingsObject structure is generated from "#/definitions/serverBindingsObject".
type ServerBindingsObject struct {
	HTTP                 *interface{}           `json:"http,omitempty"`
	Ws                   *interface{}           `json:"ws,omitempty"`
	Amqp                 *interface{}           `json:"amqp,omitempty"`
	Amqp1                *interface{}           `json:"amqp1,omitempty"`
	Mqtt                 *interface{}           `json:"mqtt,omitempty"`
	Mqtt5                *interface{}           `json:"mqtt5,omitempty"`
	Kafka                *interface{}           `json:"kafka,omitempty"`
	Nats                 *interface{}           `json:"nats,omitempty"`
	Jms                  *interface{}           `json:"jms,omitempty"`
	Sns                  *interface{}           `json:"sns,omitempty"`
	Sqs                  *interface{}           `json:"sqs,omitempty"`
	Stomp                *interface{}           `json:"stomp,omitempty"`
	Redis                *interface{}           `json:"redis,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`               // All unmatched properties.
}

type marshalServerBindingsObject ServerBindingsObject

var ignoreKeysServerBindingsObject = []string{
	"http",
	"ws",
	"amqp",
	"amqp1",
	"mqtt",
	"mqtt5",
	"kafka",
	"nats",
	"jms",
	"sns",
	"sqs",
	"stomp",
	"redis",
}

// UnmarshalJSON decodes JSON.
func (v *ServerBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalServerBindingsObject(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			vv.HTTP = &v
		}
	}

	if vv.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			vv.Ws = &v
		}
	}

	if vv.Amqp == nil {
		if _, ok := m["amqp"]; ok {
			var v interface{}
			vv.Amqp = &v
		}
	}

	if vv.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			vv.Amqp1 = &v
		}
	}

	if vv.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			vv.Mqtt = &v
		}
	}

	if vv.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			vv.Mqtt5 = &v
		}
	}

	if vv.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			vv.Kafka = &v
		}
	}

	if vv.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			vv.Nats = &v
		}
	}

	if vv.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			vv.Jms = &v
		}
	}

	if vv.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			vv.Sns = &v
		}
	}

	if vv.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			vv.Sqs = &v
		}
	}

	if vv.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			vv.Stomp = &v
		}
	}

	if vv.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			vv.Redis = &v
		}
	}

	for _, key := range ignoreKeysServerBindingsObject {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = ServerBindingsObject(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ServerBindingsObject) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalServerBindingsObject(v))
	}

	return marshalUnion(marshalServerBindingsObject(v), v.AdditionalProperties)
}

// ChannelItem structure is generated from "#/definitions/channelItem".
type ChannelItem struct {
	Ref           string                 `json:"$ref,omitempty"`        // Format: uri-reference.
	Parameters    map[string]Parameter   `json:"parameters,omitempty"`
	Description   string                 `json:"description,omitempty"` // A description of the channel.
	Publish       *Operation             `json:"publish,omitempty"`
	Subscribe     *Operation             `json:"subscribe,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`
	Bindings      *ChannelBindingsObject `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalChannelItem ChannelItem

var ignoreKeysChannelItem = []string{
	"$ref",
	"parameters",
	"description",
	"publish",
	"subscribe",
	"deprecated",
	"bindings",
}

// UnmarshalJSON decodes JSON.
func (v *ChannelItem) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalChannelItem(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysChannelItem {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ChannelItem(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ChannelItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalChannelItem(v), v.MapOfAnything)
}

// Parameter structure is generated from "#/definitions/parameter".
type Parameter struct {
	Description   string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use. GitHub Flavored Markdown is allowed.
	Schema        map[string]interface{} `json:"schema,omitempty"`
	// A runtime expression that specifies the location of the parameter value.
	// Value must match pattern: `^\$message\.(header|payload)\#(\/(([^\/~])|(~[01]))*)*`.
	Location      string                 `json:"location,omitempty"`
	Ref           string                 `json:"$ref,omitempty"`        // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalParameter Parameter

var ignoreKeysParameter = []string{
	"description",
	"schema",
	"location",
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (v *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalParameter(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Parameter(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(v), v.MapOfAnything)
}

// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Traits        []OperationTraitsItems   `json:"traits,omitempty"`
	Summary       string                   `json:"summary,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Tags          []Tag                    `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs            `json:"externalDocs,omitempty"` // Information about external documentation.
	ID            string                   `json:"operationId,omitempty"`
	Bindings      *OperationBindingsObject `json:"bindings,omitempty"`
	Message       *Message                 `json:"message,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalOperation Operation

var ignoreKeysOperation = []string{
	"traits",
	"summary",
	"description",
	"tags",
	"externalDocs",
	"operationId",
	"bindings",
	"message",
}

// UnmarshalJSON decodes JSON.
func (v *Operation) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOperation(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOperation {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Operation(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(v), v.MapOfAnything)
}

// Reference structure is generated from "#/definitions/Reference".
type Reference struct {
	// Format: uri-reference.
	// Required.
	Ref                  string                 `json:"$ref,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties.
}

type marshalReference Reference

var ignoreKeysReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (v *Reference) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalReference(*v)

	err = json.Unmarshal(data, &vv)
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
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Reference(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Reference) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalReference(v))
	}

	return marshalUnion(marshalReference(v), v.AdditionalProperties)
}

// OperationTrait structure is generated from "#/definitions/operationTrait".
type OperationTrait struct {
	Summary       string                   `json:"summary,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Tags          []Tag                    `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs            `json:"externalDocs,omitempty"` // Information about external documentation.
	OperationID   string                   `json:"operationId,omitempty"`
	Bindings      *OperationBindingsObject `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalOperationTrait OperationTrait

var ignoreKeysOperationTrait = []string{
	"summary",
	"description",
	"tags",
	"externalDocs",
	"operationId",
	"bindings",
}

// UnmarshalJSON decodes JSON.
func (v *OperationTrait) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOperationTrait(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOperationTrait {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OperationTrait(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v OperationTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationTrait(v), v.MapOfAnything)
}

// Tag structure is generated from "#/definitions/tag".
type Tag struct {
	Name          string                 `json:"name"`                   // Required.
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalTag Tag

var ignoreKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

// UnmarshalJSON decodes JSON.
func (v *Tag) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalTag(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Tag(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(v), v.MapOfAnything)
}

// ExternalDocs structure is generated from "#/definitions/externalDocs".
//
// information about external documentation.
type ExternalDocs struct {
	Description   string                 `json:"description,omitempty"`
	// Format: uri.
	// Required.
	URL           string                 `json:"url"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalExternalDocs ExternalDocs

var ignoreKeysExternalDocs = []string{
	"description",
	"url",
}

// UnmarshalJSON decodes JSON.
func (v *ExternalDocs) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalExternalDocs(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ExternalDocs(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(v), v.MapOfAnything)
}

// OperationBindingsObject structure is generated from "#/definitions/operationBindingsObject".
type OperationBindingsObject struct {
	HTTP                 *interface{}                   `json:"http,omitempty"`
	Ws                   *interface{}                   `json:"ws,omitempty"`
	// AMQP 0-9-1 Operation Binding Object.
	// This object contains information about the operation representation in AMQP.
	// See https://github.com/asyncapi/bindings/tree/master/amqp#operation-binding-object.
	Amqp                 *AMQP091OperationBindingObject `json:"amqp,omitempty"`
	Amqp1                *interface{}                   `json:"amqp1,omitempty"`
	Mqtt                 *interface{}                   `json:"mqtt,omitempty"`
	Mqtt5                *interface{}                   `json:"mqtt5,omitempty"`
	Kafka                *interface{}                   `json:"kafka,omitempty"`
	Nats                 *interface{}                   `json:"nats,omitempty"`
	Jms                  *interface{}                   `json:"jms,omitempty"`
	Sns                  *interface{}                   `json:"sns,omitempty"`
	Sqs                  *interface{}                   `json:"sqs,omitempty"`
	Stomp                *interface{}                   `json:"stomp,omitempty"`
	Redis                *interface{}                   `json:"redis,omitempty"`
	AdditionalProperties map[string]interface{}         `json:"-"`               // All unmatched properties.
}

type marshalOperationBindingsObject OperationBindingsObject

var ignoreKeysOperationBindingsObject = []string{
	"http",
	"ws",
	"amqp",
	"amqp1",
	"mqtt",
	"mqtt5",
	"kafka",
	"nats",
	"jms",
	"sns",
	"sqs",
	"stomp",
	"redis",
}

// UnmarshalJSON decodes JSON.
func (v *OperationBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOperationBindingsObject(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			vv.HTTP = &v
		}
	}

	if vv.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			vv.Ws = &v
		}
	}

	if vv.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			vv.Amqp1 = &v
		}
	}

	if vv.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			vv.Mqtt = &v
		}
	}

	if vv.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			vv.Mqtt5 = &v
		}
	}

	if vv.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			vv.Kafka = &v
		}
	}

	if vv.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			vv.Nats = &v
		}
	}

	if vv.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			vv.Jms = &v
		}
	}

	if vv.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			vv.Sns = &v
		}
	}

	if vv.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			vv.Sqs = &v
		}
	}

	if vv.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			vv.Stomp = &v
		}
	}

	if vv.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			vv.Redis = &v
		}
	}

	for _, key := range ignoreKeysOperationBindingsObject {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = OperationBindingsObject(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v OperationBindingsObject) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalOperationBindingsObject(v))
	}

	return marshalUnion(marshalOperationBindingsObject(v), v.AdditionalProperties)
}

// AMQP091OperationBindingObject structure is generated from "amqp-operation-binding-object-0.1.0.json".
//
// AMQP 0-9-1 Operation Binding Object.
//
// This object contains information about the operation representation in AMQP.
// See https://github.com/asyncapi/bindings/tree/master/amqp#operation-binding-object.
type AMQP091OperationBindingObject struct {
	Expiration     int64                                       `json:"expiration,omitempty"`     // TTL (Time-To-Live) for the message. It MUST be greater than or equal to zero. Applies to Publish, Subscribe.
	UserID         string                                      `json:"userId,omitempty"`         // Identifies the user who has sent the message. Applies to Publish, Subscribe.
	Cc             []string                                    `json:"cc,omitempty"`             // The routing keys the message should be routed to at the time of publishing. Applies to Publish, Subscribe.
	Priority       int64                                       `json:"priority,omitempty"`       // A priority for the message. Applies to Publish, Subscribe.
	DeliveryMode   AMQP091OperationBindingObjectDeliveryMode   `json:"deliveryMode,omitempty"`   // Delivery mode of the message. Its value MUST be either 1 (transient) or 2 (persistent). Applies to Publish, Subscribe.
	Mandatory      bool                                        `json:"mandatory,omitempty"`      // Whether the message is mandatory or not. Applies to Publish.
	Bcc            []string                                    `json:"bcc,omitempty"`            // Like cc but consumers will not receive this information. Applies to Publish.
	ReplyTo        string                                      `json:"replyTo,omitempty"`        // Name of the queue where the consumer should send the response. Applies to Publish, Subscribe.
	Timestamp      bool                                        `json:"timestamp,omitempty"`      // Whether the message should include a timestamp or not. Applies to Publish, Subscribe.
	Ack            bool                                        `json:"ack,omitempty"`            // Whether the consumer should ack the message or not. Applies to Subscribe.
	BindingVersion AMQP091OperationBindingObjectBindingVersion `json:"bindingVersion,omitempty"` // The version of this binding. If omitted, "latest" MUST be assumed. Applies to Publish, Subscribe.
}

// OperationTraitsItems structure is generated from "#/definitions/operation->traits->items".
type OperationTraitsItems struct {
	Reference       *Reference      `json:"-"`
	OperationTrait  *OperationTrait `json:"-"`
	SliceOfAnything []interface{}   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *OperationTraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.OperationTrait)
	if err != nil {
		v.OperationTrait = nil
	}

	err = json.Unmarshal(data, &v.SliceOfAnything)
	if err != nil {
		v.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v OperationTraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.OperationTrait, v.SliceOfAnything)
}

// MessageOneOf1OneOf0 structure is generated from "#/definitions/message/oneOf/1/oneOf/0".
type MessageOneOf1OneOf0 struct {
	OneOf []Message `json:"oneOf,omitempty"` // Required.
}

// MessageOneOf1OneOf1 structure is generated from "#/definitions/message/oneOf/1/oneOf/1".
type MessageOneOf1OneOf1 struct {
	SchemaFormat  string                            `json:"schemaFormat,omitempty"`
	ContentType   string                            `json:"contentType,omitempty"`
	Headers       map[string]interface{}            `json:"headers,omitempty"`
	Payload       *interface{}                      `json:"payload,omitempty"`
	CorrelationID *MessageOneOf1OneOf1CorrelationID `json:"correlationId,omitempty"`
	Tags          []Tag                             `json:"tags,omitempty"`
	Summary       string                            `json:"summary,omitempty"`       // A brief summary of the message.
	Name          string                            `json:"name,omitempty"`          // Name of the message.
	Title         string                            `json:"title,omitempty"`         // A human-friendly title for the message.
	Description   string                            `json:"description,omitempty"`   // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs                     `json:"externalDocs,omitempty"`  // Information about external documentation.
	Deprecated    bool                              `json:"deprecated,omitempty"`
	Examples      []map[string]interface{}          `json:"examples,omitempty"`
	Bindings      *MessageBindingsObject            `json:"bindings,omitempty"`
	Traits        []MessageOneOf1OneOf1TraitsItems  `json:"traits,omitempty"`
	MapOfAnything map[string]interface{}            `json:"-"`                       // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalMessageOneOf1OneOf1 MessageOneOf1OneOf1

var ignoreKeysMessageOneOf1OneOf1 = []string{
	"schemaFormat",
	"contentType",
	"headers",
	"payload",
	"correlationId",
	"tags",
	"summary",
	"name",
	"title",
	"description",
	"externalDocs",
	"deprecated",
	"examples",
	"bindings",
	"traits",
}

// UnmarshalJSON decodes JSON.
func (v *MessageOneOf1OneOf1) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalMessageOneOf1OneOf1(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Payload == nil {
		if _, ok := m["payload"]; ok {
			var v interface{}
			vv.Payload = &v
		}
	}

	for _, key := range ignoreKeysMessageOneOf1OneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = MessageOneOf1OneOf1(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageOneOf1OneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageOneOf1OneOf1(v), v.MapOfAnything)
}

// CorrelationID structure is generated from "#/definitions/correlationId".
type CorrelationID struct {
	Description   string                 `json:"description,omitempty"` // A optional description of the correlation ID. GitHub Flavored Markdown is allowed.
	// A runtime expression that specifies the location of the correlation ID.
	// Value must match pattern: `^\$message\.(header|payload)\#(\/(([^\/~])|(~[01]))*)*`.
	// Required.
	Location      string                 `json:"location"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalCorrelationID CorrelationID

var ignoreKeysCorrelationID = []string{
	"description",
	"location",
}

// UnmarshalJSON decodes JSON.
func (v *CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalCorrelationID(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysCorrelationID {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = CorrelationID(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalCorrelationID(v), v.MapOfAnything)
}

// MessageOneOf1OneOf1CorrelationID structure is generated from "#/definitions/message/oneOf/1/oneOf/1->correlationId".
type MessageOneOf1OneOf1CorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *MessageOneOf1OneOf1CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.CorrelationID)
	if err != nil {
		v.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageOneOf1OneOf1CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.CorrelationID)
}

// MessageBindingsObject structure is generated from "#/definitions/messageBindingsObject".
type MessageBindingsObject struct {
	HTTP                 *interface{}                 `json:"http,omitempty"`
	Ws                   *interface{}                 `json:"ws,omitempty"`
	// AMQP 0-9-1 Message Binding Object.
	// This object contains information about the message representation in AMQP.
	// See https://github.com/asyncapi/bindings/tree/master/amqp#message-binding-object.
	Amqp                 *AMQP091MessageBindingObject `json:"amqp,omitempty"`
	Amqp1                *interface{}                 `json:"amqp1,omitempty"`
	Mqtt                 *interface{}                 `json:"mqtt,omitempty"`
	Mqtt5                *interface{}                 `json:"mqtt5,omitempty"`
	Kafka                *interface{}                 `json:"kafka,omitempty"`
	Nats                 *interface{}                 `json:"nats,omitempty"`
	Jms                  *interface{}                 `json:"jms,omitempty"`
	Sns                  *interface{}                 `json:"sns,omitempty"`
	Sqs                  *interface{}                 `json:"sqs,omitempty"`
	Stomp                *interface{}                 `json:"stomp,omitempty"`
	Redis                *interface{}                 `json:"redis,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties.
}

type marshalMessageBindingsObject MessageBindingsObject

var ignoreKeysMessageBindingsObject = []string{
	"http",
	"ws",
	"amqp",
	"amqp1",
	"mqtt",
	"mqtt5",
	"kafka",
	"nats",
	"jms",
	"sns",
	"sqs",
	"stomp",
	"redis",
}

// UnmarshalJSON decodes JSON.
func (v *MessageBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalMessageBindingsObject(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			vv.HTTP = &v
		}
	}

	if vv.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			vv.Ws = &v
		}
	}

	if vv.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			vv.Amqp1 = &v
		}
	}

	if vv.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			vv.Mqtt = &v
		}
	}

	if vv.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			vv.Mqtt5 = &v
		}
	}

	if vv.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			vv.Kafka = &v
		}
	}

	if vv.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			vv.Nats = &v
		}
	}

	if vv.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			vv.Jms = &v
		}
	}

	if vv.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			vv.Sns = &v
		}
	}

	if vv.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			vv.Sqs = &v
		}
	}

	if vv.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			vv.Stomp = &v
		}
	}

	if vv.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			vv.Redis = &v
		}
	}

	for _, key := range ignoreKeysMessageBindingsObject {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = MessageBindingsObject(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageBindingsObject) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalMessageBindingsObject(v))
	}

	return marshalUnion(marshalMessageBindingsObject(v), v.AdditionalProperties)
}

// AMQP091MessageBindingObject structure is generated from "amqp-message-binding-object-0.1.0.json".
//
// AMQP 0-9-1 Message Binding Object.
//
// This object contains information about the message representation in AMQP.
// See https://github.com/asyncapi/bindings/tree/master/amqp#message-binding-object.
type AMQP091MessageBindingObject struct {
	ContentEncoding string                                    `json:"contentEncoding,omitempty"` // A MIME encoding for the message content.
	MessageType     string                                    `json:"messageType,omitempty"`     // Application-specific message type.
	BindingVersion  AMQP091MessageBindingObjectBindingVersion `json:"bindingVersion,omitempty"`  // The version of this binding. If omitted, "latest" MUST be assumed.
}

// MessageTrait structure is generated from "#/definitions/messageTrait".
type MessageTrait struct {
	SchemaFormat  string                     `json:"schemaFormat,omitempty"`
	ContentType   string                     `json:"contentType,omitempty"`
	Headers       *MessageTraitHeaders       `json:"headers,omitempty"`
	CorrelationID *MessageTraitCorrelationID `json:"correlationId,omitempty"`
	Tags          []Tag                      `json:"tags,omitempty"`
	Summary       string                     `json:"summary,omitempty"`       // A brief summary of the message.
	Name          string                     `json:"name,omitempty"`          // Name of the message.
	Title         string                     `json:"title,omitempty"`         // A human-friendly title for the message.
	Description   string                     `json:"description,omitempty"`   // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs              `json:"externalDocs,omitempty"`  // Information about external documentation.
	Deprecated    bool                       `json:"deprecated,omitempty"`
	Examples      []map[string]interface{}   `json:"examples,omitempty"`
	Bindings      *MessageBindingsObject     `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                       // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalMessageTrait MessageTrait

var ignoreKeysMessageTrait = []string{
	"schemaFormat",
	"contentType",
	"headers",
	"correlationId",
	"tags",
	"summary",
	"name",
	"title",
	"description",
	"externalDocs",
	"deprecated",
	"examples",
	"bindings",
}

// UnmarshalJSON decodes JSON.
func (v *MessageTrait) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalMessageTrait(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysMessageTrait {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = MessageTrait(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageTrait(v), v.MapOfAnything)
}

// MessageTraitHeaders structure is generated from "#/definitions/messageTrait->headers".
type MessageTraitHeaders struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *MessageTraitHeaders) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.Schema)
	if err != nil {
		v.Schema = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageTraitHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.Schema)
}

// MessageTraitCorrelationID structure is generated from "#/definitions/messageTrait->correlationId".
type MessageTraitCorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *MessageTraitCorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.CorrelationID)
	if err != nil {
		v.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageTraitCorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.CorrelationID)
}

// MessageOneOf1OneOf1TraitsItems structure is generated from "#/definitions/message/oneOf/1/oneOf/1->traits->items".
type MessageOneOf1OneOf1TraitsItems struct {
	Reference       *Reference    `json:"-"`
	MessageTrait    *MessageTrait `json:"-"`
	SliceOfAnything []interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *MessageOneOf1OneOf1TraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.MessageTrait)
	if err != nil {
		v.MessageTrait = nil
	}

	err = json.Unmarshal(data, &v.SliceOfAnything)
	if err != nil {
		v.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageOneOf1OneOf1TraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.MessageTrait, v.SliceOfAnything)
}

// MessageOneOf1 structure is generated from "#/definitions/message/oneOf/1".
type MessageOneOf1 struct {
	OneOf0 *MessageOneOf1OneOf0 `json:"-"`
	OneOf1 *MessageOneOf1OneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *MessageOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.OneOf0)
	if err != nil {
		v.OneOf0 = nil
	}

	err = json.Unmarshal(data, &v.OneOf1)
	if err != nil {
		v.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v MessageOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.OneOf0, v.OneOf1)
}

// Message structure is generated from "#/definitions/message".
type Message struct {
	Reference *Reference     `json:"-"`
	OneOf1    *MessageOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *Message) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.OneOf1)
	if err != nil {
		v.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.OneOf1)
}

// ChannelBindingsObject structure is generated from "#/definitions/channelBindingsObject".
type ChannelBindingsObject struct {
	HTTP                 *interface{}                 `json:"http,omitempty"`
	Ws                   *interface{}                 `json:"ws,omitempty"`
	// AMQP 0-9-1 Channel Binding Object.
	// This object contains information about the channel representation in AMQP.
	// See https://github.com/asyncapi/bindings/tree/master/amqp#channel-binding-object.
	Amqp                 *AMQP091ChannelBindingObject `json:"amqp,omitempty"`
	Amqp1                *interface{}                 `json:"amqp1,omitempty"`
	Mqtt                 *interface{}                 `json:"mqtt,omitempty"`
	Mqtt5                *interface{}                 `json:"mqtt5,omitempty"`
	Kafka                *interface{}                 `json:"kafka,omitempty"`
	Nats                 *interface{}                 `json:"nats,omitempty"`
	Jms                  *interface{}                 `json:"jms,omitempty"`
	Sns                  *interface{}                 `json:"sns,omitempty"`
	Sqs                  *interface{}                 `json:"sqs,omitempty"`
	Stomp                *interface{}                 `json:"stomp,omitempty"`
	Redis                *interface{}                 `json:"redis,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties.
}

type marshalChannelBindingsObject ChannelBindingsObject

var ignoreKeysChannelBindingsObject = []string{
	"http",
	"ws",
	"amqp",
	"amqp1",
	"mqtt",
	"mqtt5",
	"kafka",
	"nats",
	"jms",
	"sns",
	"sqs",
	"stomp",
	"redis",
}

// UnmarshalJSON decodes JSON.
func (v *ChannelBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalChannelBindingsObject(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			vv.HTTP = &v
		}
	}

	if vv.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			vv.Ws = &v
		}
	}

	if vv.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			vv.Amqp1 = &v
		}
	}

	if vv.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			vv.Mqtt = &v
		}
	}

	if vv.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			vv.Mqtt5 = &v
		}
	}

	if vv.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			vv.Kafka = &v
		}
	}

	if vv.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			vv.Nats = &v
		}
	}

	if vv.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			vv.Jms = &v
		}
	}

	if vv.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			vv.Sns = &v
		}
	}

	if vv.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			vv.Sqs = &v
		}
	}

	if vv.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			vv.Stomp = &v
		}
	}

	if vv.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			vv.Redis = &v
		}
	}

	for _, key := range ignoreKeysChannelBindingsObject {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = ChannelBindingsObject(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ChannelBindingsObject) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalChannelBindingsObject(v))
	}

	return marshalUnion(marshalChannelBindingsObject(v), v.AdditionalProperties)
}

// AMQP091ChannelBindingObject structure is generated from "amqp-channel-binding-object-0.1.0.json".
//
// AMQP 0-9-1 Channel Binding Object.
//
// This object contains information about the channel representation in AMQP.
// See https://github.com/asyncapi/bindings/tree/master/amqp#channel-binding-object.
type AMQP091ChannelBindingObject struct {
	Is             AMQP091ChannelBindingObjectIs             `json:"is,omitempty"`             // Defines what type of channel is it. Can be either `queue` or `routingKey` (default).
	Exchange       *Exchange                                 `json:"exchange,omitempty"`       // When `is`=`routingKey`, this object defines the exchange properties.
	Queue          *Queue                                    `json:"queue,omitempty"`          // When `is`=`queue`, this object defines the queue properties.
	BindingVersion AMQP091ChannelBindingObjectBindingVersion `json:"bindingVersion,omitempty"` // The version of this binding. If omitted, "latest" MUST be assumed.
}

// Exchange structure is generated from "#/definitions/exchange".
//
// When `is`=`routingKey`, this object defines the exchange properties.
type Exchange struct {
	Name                 string                 `json:"name,omitempty"`       // The name of the exchange. It MUST NOT exceed 255 characters long.
	Type                 ExchangeType           `json:"type,omitempty"`       // The type of the exchange. Can be either `topic`, `direct`, `fanout`, `default` or `headers`.
	Durable              bool                   `json:"durable,omitempty"`    // Whether the exchange should survive broker restarts or not.
	AutoDelete           bool                   `json:"autoDelete,omitempty"` // Whether the exchange should be deleted when the last queue is unbound from it.
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties.
}

type marshalExchange Exchange

var ignoreKeysExchange = []string{
	"name",
	"type",
	"durable",
	"autoDelete",
}

// UnmarshalJSON decodes JSON.
func (v *Exchange) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalExchange(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysExchange {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Exchange(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Exchange) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalExchange(v))
	}

	return marshalUnion(marshalExchange(v), v.AdditionalProperties)
}

// Queue structure is generated from "#/definitions/queue".
//
// When `is`=`queue`, this object defines the queue properties.
type Queue struct {
	Name                 string                 `json:"name,omitempty"`       // The name of the queue. It MUST NOT exceed 255 characters long.
	Durable              bool                   `json:"durable,omitempty"`    // Whether the queue should survive broker restarts or not.
	Exclusive            bool                   `json:"exclusive,omitempty"`  // Whether the queue should be used only by one connection or not.
	AutoDelete           bool                   `json:"autoDelete,omitempty"` // Whether the queue should be deleted when the last consumer unsubscribes.
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties.
}

type marshalQueue Queue

var ignoreKeysQueue = []string{
	"name",
	"durable",
	"exclusive",
	"autoDelete",
}

// UnmarshalJSON decodes JSON.
func (v *Queue) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalQueue(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysQueue {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Queue(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Queue) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalQueue(v))
	}

	return marshalUnion(marshalQueue(v), v.AdditionalProperties)
}

// Components structure is generated from "#/definitions/components".
//
// An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
type Components struct {
	Schemas           map[string]map[string]interface{}  `json:"schemas,omitempty"`           // JSON objects describing schemas the API uses.
	Messages          map[string]Message                 `json:"messages,omitempty"`          // JSON objects describing the messages being consumed and produced by the API.
	SecuritySchemes   *ComponentsSecuritySchemes         `json:"securitySchemes,omitempty"`
	Parameters        map[string]Parameter               `json:"parameters,omitempty"`        // JSON objects describing re-usable channel parameters.
	CorrelationIds    *ComponentsCorrelationIds          `json:"correlationIds,omitempty"`
	OperationTraits   map[string]OperationTrait          `json:"operationTraits,omitempty"`
	MessageTraits     map[string]MessageTrait            `json:"messageTraits,omitempty"`
	ServerBindings    map[string]ServerBindingsObject    `json:"serverBindings,omitempty"`
	ChannelBindings   map[string]ChannelBindingsObject   `json:"channelBindings,omitempty"`
	OperationBindings map[string]OperationBindingsObject `json:"operationBindings,omitempty"`
	MessageBindings   map[string]MessageBindingsObject   `json:"messageBindings,omitempty"`
}

// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalUserPassword UserPassword

var ignoreKeysUserPassword = []string{
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *UserPassword) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalUserPassword(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = UserPassword(vv)

	return nil
}

var (
	// constUserPassword is unconditionally added to JSON.
	constUserPassword = json.RawMessage(`{"type":"userPassword"}`)
)

// MarshalJSON encodes JSON.
func (v UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(constUserPassword, marshalUserPassword(v), v.MapOfAnything)
}

// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	In            APIKeyIn               `json:"in"`                    // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalAPIKey APIKey

var ignoreKeysAPIKey = []string{
	"in",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *APIKey) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAPIKey(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = APIKey(vv)

	return nil
}

var (
	// constAPIKey is unconditionally added to JSON.
	constAPIKey = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON.
func (v APIKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKey, marshalAPIKey(v), v.MapOfAnything)
}

// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalX509 X509

var ignoreKeysX509 = []string{
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *X509) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalX509(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = X509(vv)

	return nil
}

var (
	// constX509 is unconditionally added to JSON.
	constX509 = json.RawMessage(`{"type":"X509"}`)
)

// MarshalJSON encodes JSON.
func (v X509) MarshalJSON() ([]byte, error) {
	return marshalUnion(constX509, marshalX509(v), v.MapOfAnything)
}

// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalSymmetricEncryption SymmetricEncryption

var ignoreKeysSymmetricEncryption = []string{
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalSymmetricEncryption(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = SymmetricEncryption(vv)

	return nil
}

var (
	// constSymmetricEncryption is unconditionally added to JSON.
	constSymmetricEncryption = json.RawMessage(`{"type":"symmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (v SymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constSymmetricEncryption, marshalSymmetricEncryption(v), v.MapOfAnything)
}

// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalAsymmetricEncryption AsymmetricEncryption

var ignoreKeysAsymmetricEncryption = []string{
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAsymmetricEncryption(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = AsymmetricEncryption(vv)

	return nil
}

var (
	// constAsymmetricEncryption is unconditionally added to JSON.
	constAsymmetricEncryption = json.RawMessage(`{"type":"asymmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (v AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsymmetricEncryption, marshalAsymmetricEncryption(v), v.MapOfAnything)
}

// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme"`                // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

var ignoreKeysNonBearerHTTPSecurityScheme = []string{
	"scheme",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalNonBearerHTTPSecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = NonBearerHTTPSecurityScheme(vv)

	return nil
}

var (
	// constNonBearerHTTPSecurityScheme is unconditionally added to JSON.
	constNonBearerHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON.
func (v NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constNonBearerHTTPSecurityScheme, marshalNonBearerHTTPSecurityScheme(v), v.MapOfAnything)
}

// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

var ignoreKeysBearerHTTPSecurityScheme = []string{
	"bearerFormat",
	"description",
	"scheme",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalBearerHTTPSecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = BearerHTTPSecurityScheme(vv)

	return nil
}

var (
	// constBearerHTTPSecurityScheme is unconditionally added to JSON.
	constBearerHTTPSecurityScheme = json.RawMessage(`{"scheme":"bearer","type":"http"}`)
)

// MarshalJSON encodes JSON.
func (v BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBearerHTTPSecurityScheme, marshalBearerHTTPSecurityScheme(v), v.MapOfAnything)
}

// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name"`                  // Required.
	In            APIKeyHTTPSecuritySchemeIn `json:"in"`                    // Required.
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

var ignoreKeysAPIKeyHTTPSecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAPIKeyHTTPSecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
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

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = APIKeyHTTPSecurityScheme(vv)

	return nil
}

var (
	// constAPIKeyHTTPSecurityScheme is unconditionally added to JSON.
	constAPIKeyHTTPSecurityScheme = json.RawMessage(`{"type":"httpApiKey"}`)
)

// MarshalJSON encodes JSON.
func (v APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeyHTTPSecurityScheme, marshalAPIKeyHTTPSecurityScheme(v), v.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.NonBearerHTTPSecurityScheme)
	if err != nil {
		v.NonBearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.BearerHTTPSecurityScheme)
	if err != nil {
		v.BearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.APIKeyHTTPSecurityScheme)
	if err != nil {
		v.APIKeyHTTPSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.NonBearerHTTPSecurityScheme, v.BearerHTTPSecurityScheme, v.APIKeyHTTPSecurityScheme)
}

// Oauth2Flows structure is generated from "#/definitions/oauth2Flows".
type Oauth2Flows struct {
	Description          string                 `json:"description,omitempty"`
	Flows                Oauth2FlowsFlows       `json:"flows"`                 // Required.
	MapOfAnything        map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
	AdditionalProperties map[string]interface{} `json:"-"`                     // All unmatched properties.
}

type marshalOauth2Flows Oauth2Flows

var ignoreKeysOauth2Flows = []string{
	"description",
	"flows",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *Oauth2Flows) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOauth2Flows(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"oauth2"` {
		return fmt.Errorf(`bad or missing const value for "type" ("oauth2" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysOauth2Flows {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Oauth2Flows(vv)

	return nil
}

var (
	// constOauth2Flows is unconditionally added to JSON.
	constOauth2Flows = json.RawMessage(`{"type":"oauth2"}`)
)

// MarshalJSON encodes JSON.
func (v Oauth2Flows) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOauth2Flows, marshalOauth2Flows(v), v.MapOfAnything, v.AdditionalProperties)
}

// Oauth2FlowsFlows structure is generated from "#/definitions/oauth2Flows->flows".
type Oauth2FlowsFlows struct {
	Implicit          *Oauth2Flow `json:"implicit,omitempty"`
	Password          *Oauth2Flow `json:"password,omitempty"`
	ClientCredentials *Oauth2Flow `json:"clientCredentials,omitempty"`
	AuthorizationCode *Oauth2Flow `json:"authorizationCode,omitempty"`
}

// Oauth2Flow structure is generated from "#/definitions/oauth2Flow".
type Oauth2Flow struct {
	AuthorizationURL string                 `json:"authorizationUrl,omitempty"` // Format: uri.
	TokenURL         string                 `json:"tokenUrl,omitempty"`         // Format: uri.
	RefreshURL       string                 `json:"refreshUrl,omitempty"`       // Format: uri.
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalOauth2Flow Oauth2Flow

var ignoreKeysOauth2Flow = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (v *Oauth2Flow) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOauth2Flow(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOauth2Flow {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Oauth2Flow(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Oauth2Flow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2Flow(v), v.MapOfAnything)
}

// OpenIDConnect structure is generated from "#/definitions/openIdConnect".
type OpenIDConnect struct {
	Description   string                 `json:"description,omitempty"`
	// Format: uri.
	// Required.
	URL           string                 `json:"openIdConnectUrl"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\-\_]+$`.
}

type marshalOpenIDConnect OpenIDConnect

var ignoreKeysOpenIDConnect = []string{
	"description",
	"openIdConnectUrl",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *OpenIDConnect) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOpenIDConnect(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad or missing const value for "type" ("openIdConnect" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysOpenIDConnect {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OpenIDConnect(vv)

	return nil
}

var (
	// constOpenIDConnect is unconditionally added to JSON.
	constOpenIDConnect = json.RawMessage(`{"type":"openIdConnect"}`)
)

// MarshalJSON encodes JSON.
func (v OpenIDConnect) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnect, marshalOpenIDConnect(v), v.MapOfAnything)
}

// SecurityScheme structure is generated from "#/definitions/SecurityScheme".
type SecurityScheme struct {
	UserPassword         *UserPassword         `json:"-"`
	APIKey               *APIKey               `json:"-"`
	X509                 *X509                 `json:"-"`
	SymmetricEncryption  *SymmetricEncryption  `json:"-"`
	AsymmetricEncryption *AsymmetricEncryption `json:"-"`
	HTTPSecurityScheme   *HTTPSecurityScheme   `json:"-"`
	Oauth2Flows          *Oauth2Flows          `json:"-"`
	OpenIDConnect        *OpenIDConnect        `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.UserPassword)
	if err != nil {
		v.UserPassword = nil
	}

	err = json.Unmarshal(data, &v.APIKey)
	if err != nil {
		v.APIKey = nil
	}

	err = json.Unmarshal(data, &v.X509)
	if err != nil {
		v.X509 = nil
	}

	err = json.Unmarshal(data, &v.SymmetricEncryption)
	if err != nil {
		v.SymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &v.AsymmetricEncryption)
	if err != nil {
		v.AsymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &v.HTTPSecurityScheme)
	if err != nil {
		v.HTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.Oauth2Flows)
	if err != nil {
		v.Oauth2Flows = nil
	}

	err = json.Unmarshal(data, &v.OpenIDConnect)
	if err != nil {
		v.OpenIDConnect = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.UserPassword, v.APIKey, v.X509, v.SymmetricEncryption, v.AsymmetricEncryption, v.HTTPSecurityScheme, v.Oauth2Flows, v.OpenIDConnect)
}

// ComponentsSecuritySchemesWD structure is generated from "#/definitions/components->securitySchemes->^[\w\d\.\-_]+$".
type ComponentsSecuritySchemesWD struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsSecuritySchemesWD) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.SecurityScheme)
	if err != nil {
		v.SecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsSecuritySchemesWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesWDValues map[string]ComponentsSecuritySchemesWD `json:"-"` // Key must match pattern: `^[\w\d\.\-_]+$`.
	AdditionalProperties                   map[string]interface{}                 `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexWD.MatchString(key) {
			matched = true

			if v.MapOfComponentsSecuritySchemesWDValues == nil {
				v.MapOfComponentsSecuritySchemesWDValues = make(map[string]ComponentsSecuritySchemesWD, 1)
			}

			var val ComponentsSecuritySchemesWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfComponentsSecuritySchemesWDValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfComponentsSecuritySchemesWDValues, v.AdditionalProperties)
}

// ComponentsCorrelationIdsWD structure is generated from "#/definitions/components->correlationIds->^[\w\d\.\-_]+$".
type ComponentsCorrelationIdsWD struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsCorrelationIdsWD) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Reference)
	if err != nil {
		v.Reference = nil
	}

	err = json.Unmarshal(data, &v.CorrelationID)
	if err != nil {
		v.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsCorrelationIdsWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Reference, v.CorrelationID)
}

// ComponentsCorrelationIds structure is generated from "#/definitions/components->correlationIds".
type ComponentsCorrelationIds struct {
	MapOfComponentsCorrelationIdsWDValues map[string]ComponentsCorrelationIdsWD `json:"-"` // Key must match pattern: `^[\w\d\.\-_]+$`.
	AdditionalProperties                  map[string]interface{}                `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsCorrelationIds) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexWD.MatchString(key) {
			matched = true

			if v.MapOfComponentsCorrelationIdsWDValues == nil {
				v.MapOfComponentsCorrelationIdsWDValues = make(map[string]ComponentsCorrelationIdsWD, 1)
			}

			var val ComponentsCorrelationIdsWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfComponentsCorrelationIdsWDValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsCorrelationIds) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfComponentsCorrelationIdsWDValues, v.AdditionalProperties)
}

// AMQP091OperationBindingObjectDeliveryMode is an enum type.
type AMQP091OperationBindingObjectDeliveryMode int64

// AMQP091OperationBindingObjectDeliveryMode values enumeration.
const (
	AMQP091OperationBindingObjectDeliveryModeTransient = AMQP091OperationBindingObjectDeliveryMode(1)
	AMQP091OperationBindingObjectDeliveryModePersistent = AMQP091OperationBindingObjectDeliveryMode(2)
)

// MarshalJSON encodes JSON.
func (i AMQP091OperationBindingObjectDeliveryMode) MarshalJSON() ([]byte, error) {
	switch i {
	case AMQP091OperationBindingObjectDeliveryModeTransient:
	case AMQP091OperationBindingObjectDeliveryModePersistent:

	default:
		return nil, fmt.Errorf("unexpected AMQP091OperationBindingObjectDeliveryMode value: %v", i)
	}

	return json.Marshal(int64(i))
}

// UnmarshalJSON decodes JSON.
func (i *AMQP091OperationBindingObjectDeliveryMode) UnmarshalJSON(data []byte) error {
	var ii int64

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AMQP091OperationBindingObjectDeliveryMode(ii)

	switch v {
	case AMQP091OperationBindingObjectDeliveryModeTransient:
	case AMQP091OperationBindingObjectDeliveryModePersistent:

	default:
		return fmt.Errorf("unexpected AMQP091OperationBindingObjectDeliveryMode value: %v", v)
	}

	*i = v

	return nil
}

// AMQP091OperationBindingObjectBindingVersion is an enum type.
type AMQP091OperationBindingObjectBindingVersion string

// AMQP091OperationBindingObjectBindingVersion values enumeration.
const (
	AMQP091OperationBindingObjectBindingVersion010 = AMQP091OperationBindingObjectBindingVersion("0.1.0")
	AMQP091OperationBindingObjectBindingVersionLatest = AMQP091OperationBindingObjectBindingVersion("latest")
)

// MarshalJSON encodes JSON.
func (i AMQP091OperationBindingObjectBindingVersion) MarshalJSON() ([]byte, error) {
	switch i {
	case AMQP091OperationBindingObjectBindingVersion010:
	case AMQP091OperationBindingObjectBindingVersionLatest:

	default:
		return nil, fmt.Errorf("unexpected AMQP091OperationBindingObjectBindingVersion value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AMQP091OperationBindingObjectBindingVersion) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AMQP091OperationBindingObjectBindingVersion(ii)

	switch v {
	case AMQP091OperationBindingObjectBindingVersion010:
	case AMQP091OperationBindingObjectBindingVersionLatest:

	default:
		return fmt.Errorf("unexpected AMQP091OperationBindingObjectBindingVersion value: %v", v)
	}

	*i = v

	return nil
}

// AMQP091MessageBindingObjectBindingVersion is an enum type.
type AMQP091MessageBindingObjectBindingVersion string

// AMQP091MessageBindingObjectBindingVersion values enumeration.
const (
	AMQP091MessageBindingObjectBindingVersion010 = AMQP091MessageBindingObjectBindingVersion("0.1.0")
	AMQP091MessageBindingObjectBindingVersionLatest = AMQP091MessageBindingObjectBindingVersion("latest")
)

// MarshalJSON encodes JSON.
func (i AMQP091MessageBindingObjectBindingVersion) MarshalJSON() ([]byte, error) {
	switch i {
	case AMQP091MessageBindingObjectBindingVersion010:
	case AMQP091MessageBindingObjectBindingVersionLatest:

	default:
		return nil, fmt.Errorf("unexpected AMQP091MessageBindingObjectBindingVersion value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AMQP091MessageBindingObjectBindingVersion) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AMQP091MessageBindingObjectBindingVersion(ii)

	switch v {
	case AMQP091MessageBindingObjectBindingVersion010:
	case AMQP091MessageBindingObjectBindingVersionLatest:

	default:
		return fmt.Errorf("unexpected AMQP091MessageBindingObjectBindingVersion value: %v", v)
	}

	*i = v

	return nil
}

// AMQP091ChannelBindingObjectIs is an enum type.
type AMQP091ChannelBindingObjectIs string

// AMQP091ChannelBindingObjectIs values enumeration.
const (
	AMQP091ChannelBindingObjectIsRoutingKey = AMQP091ChannelBindingObjectIs("routingKey")
	AMQP091ChannelBindingObjectIsQueue = AMQP091ChannelBindingObjectIs("queue")
)

// MarshalJSON encodes JSON.
func (i AMQP091ChannelBindingObjectIs) MarshalJSON() ([]byte, error) {
	switch i {
	case AMQP091ChannelBindingObjectIsRoutingKey:
	case AMQP091ChannelBindingObjectIsQueue:

	default:
		return nil, fmt.Errorf("unexpected AMQP091ChannelBindingObjectIs value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AMQP091ChannelBindingObjectIs) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AMQP091ChannelBindingObjectIs(ii)

	switch v {
	case AMQP091ChannelBindingObjectIsRoutingKey:
	case AMQP091ChannelBindingObjectIsQueue:

	default:
		return fmt.Errorf("unexpected AMQP091ChannelBindingObjectIs value: %v", v)
	}

	*i = v

	return nil
}

// ExchangeType is an enum type.
type ExchangeType string

// ExchangeType values enumeration.
const (
	ExchangeTypeTopic = ExchangeType("topic")
	ExchangeTypeDirect = ExchangeType("direct")
	ExchangeTypeFanout = ExchangeType("fanout")
	ExchangeTypeDefault = ExchangeType("default")
	ExchangeTypeHeaders = ExchangeType("headers")
)

// MarshalJSON encodes JSON.
func (i ExchangeType) MarshalJSON() ([]byte, error) {
	switch i {
	case ExchangeTypeTopic:
	case ExchangeTypeDirect:
	case ExchangeTypeFanout:
	case ExchangeTypeDefault:
	case ExchangeTypeHeaders:

	default:
		return nil, fmt.Errorf("unexpected ExchangeType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ExchangeType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ExchangeType(ii)

	switch v {
	case ExchangeTypeTopic:
	case ExchangeTypeDirect:
	case ExchangeTypeFanout:
	case ExchangeTypeDefault:
	case ExchangeTypeHeaders:

	default:
		return fmt.Errorf("unexpected ExchangeType value: %v", v)
	}

	*i = v

	return nil
}

// AMQP091ChannelBindingObjectBindingVersion is an enum type.
type AMQP091ChannelBindingObjectBindingVersion string

// AMQP091ChannelBindingObjectBindingVersion values enumeration.
const (
	AMQP091ChannelBindingObjectBindingVersion010 = AMQP091ChannelBindingObjectBindingVersion("0.1.0")
	AMQP091ChannelBindingObjectBindingVersionLatest = AMQP091ChannelBindingObjectBindingVersion("latest")
)

// MarshalJSON encodes JSON.
func (i AMQP091ChannelBindingObjectBindingVersion) MarshalJSON() ([]byte, error) {
	switch i {
	case AMQP091ChannelBindingObjectBindingVersion010:
	case AMQP091ChannelBindingObjectBindingVersionLatest:

	default:
		return nil, fmt.Errorf("unexpected AMQP091ChannelBindingObjectBindingVersion value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *AMQP091ChannelBindingObjectBindingVersion) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := AMQP091ChannelBindingObjectBindingVersion(ii)

	switch v {
	case AMQP091ChannelBindingObjectBindingVersion010:
	case AMQP091ChannelBindingObjectBindingVersionLatest:

	default:
		return fmt.Errorf("unexpected AMQP091ChannelBindingObjectBindingVersion value: %v", v)
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
	regexXWD = regexp.MustCompile(`^x-[\w\d\.\-\_]+$`)
	regexWD = regexp.MustCompile(`^[\w\d\.\-_]+$`)
)
