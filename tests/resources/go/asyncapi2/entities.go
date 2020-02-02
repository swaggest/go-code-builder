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
	ID                 string                 `json:"id,omitempty"`                 // A unique id representing the application.
	Info               *Info                  `json:"info,omitempty"`               // General information about the API.
	Servers            map[string]Server      `json:"servers,omitempty"`
	DefaultContentType string                 `json:"defaultContentType,omitempty"`
	Channels           map[string]ChannelItem `json:"channels,omitempty"`
	Components         *Components            `json:"components,omitempty"`         // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags               []Tag                  `json:"tags,omitempty"`
	ExternalDocs       *ExternalDocs          `json:"externalDocs,omitempty"`       // information about external documentation
	MapOfAnything      map[string]interface{} `json:"-"`                            // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

var (
	// constAsyncAPI is unconditionally added to JSON.
	constAsyncAPI = json.RawMessage(`{"asyncapi":"2.0.0"}`)
)

// MarshalJSON encodes JSON.
func (i AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsyncAPI, marshalAsyncAPI(i), i.MapOfAnything)
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
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`              // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	URL             string                    `json:"url,omitempty"`
	Description     string                    `json:"description,omitempty"`
	Protocol        string                    `json:"protocol,omitempty"`        // The transfer protocol.
	ProtocolVersion string                    `json:"protocolVersion,omitempty"`
	Variables       map[string]ServerVariable `json:"variables,omitempty"`
	Security        []map[string][]string     `json:"security,omitempty"`
	Bindings        *ServerBindingsObject     `json:"bindings,omitempty"`
	MapOfAnything   map[string]interface{}    `json:"-"`                         // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	Examples      []string               `json:"examples,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
}

type marshalServerVariable ServerVariable

var ignoreKeysServerVariable = []string{
	"enum",
	"default",
	"description",
	"examples",
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

		if regexXWD.MatchString(key) {
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
	AdditionalProperties map[string]interface{} `json:"-"`               // All unmatched properties
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
func (i *ServerBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalServerBindingsObject(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			ii.HTTP = &v
		}
	}

	if ii.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			ii.Ws = &v
		}
	}

	if ii.Amqp == nil {
		if _, ok := m["amqp"]; ok {
			var v interface{}
			ii.Amqp = &v
		}
	}

	if ii.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			ii.Amqp1 = &v
		}
	}

	if ii.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			ii.Mqtt = &v
		}
	}

	if ii.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			ii.Mqtt5 = &v
		}
	}

	if ii.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			ii.Kafka = &v
		}
	}

	if ii.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			ii.Nats = &v
		}
	}

	if ii.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			ii.Jms = &v
		}
	}

	if ii.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			ii.Sns = &v
		}
	}

	if ii.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			ii.Sqs = &v
		}
	}

	if ii.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			ii.Stomp = &v
		}
	}

	if ii.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			ii.Redis = &v
		}
	}

	for _, key := range ignoreKeysServerBindingsObject {
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

	*i = ServerBindingsObject(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ServerBindingsObject) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalServerBindingsObject(i))
	}
	return marshalUnion(marshalServerBindingsObject(i), i.AdditionalProperties)
}

// ChannelItem structure is generated from "#/definitions/channelItem".
type ChannelItem struct {
	Ref           string                 `json:"$ref,omitempty"`
	Parameters    map[string]Parameter   `json:"parameters,omitempty"`
	Description   string                 `json:"description,omitempty"` // A description of the channel.
	Publish       *Operation             `json:"publish,omitempty"`
	Subscribe     *Operation             `json:"subscribe,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`
	Bindings      *ChannelBindingsObject `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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
func (i *ChannelItem) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalChannelItem(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = ChannelItem(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ChannelItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalChannelItem(i), i.MapOfAnything)
}

// Parameter structure is generated from "#/definitions/parameter".
type Parameter struct {
	Description   string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use. GitHub Flavored Markdown is allowed.
	Schema        map[string]interface{} `json:"schema,omitempty"`
	Location      string                 `json:"location,omitempty"`    // A runtime expression that specifies the location of the parameter value
	Ref           string                 `json:"$ref,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
}

type marshalParameter Parameter

var ignoreKeysParameter = []string{
	"description",
	"schema",
	"location",
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

		if regexXWD.MatchString(key) {
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

// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Traits        []OperationTraitsItems   `json:"traits,omitempty"`
	Summary       string                   `json:"summary,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Tags          []Tag                    `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs            `json:"externalDocs,omitempty"` // information about external documentation
	ID            string                   `json:"operationId,omitempty"`
	Bindings      *OperationBindingsObject `json:"bindings,omitempty"`
	Message       *Message                 `json:"message,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                      // Key must match pattern: ^x-[\w\d\.\-\_]+$
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
func (i *Operation) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOperation(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = Operation(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(i), i.MapOfAnything)
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

// OperationTrait structure is generated from "#/definitions/operationTrait".
type OperationTrait struct {
	Summary       string                   `json:"summary,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Tags          []Tag                    `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs            `json:"externalDocs,omitempty"` // information about external documentation
	OperationID   string                   `json:"operationId,omitempty"`
	Bindings      *OperationBindingsObject `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                      // Key must match pattern: ^x-[\w\d\.\-\_]+$
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
func (i *OperationTrait) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOperationTrait(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = OperationTrait(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationTrait(i), i.MapOfAnything)
}

// Tag structure is generated from "#/definitions/tag".
type Tag struct {
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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

// OperationBindingsObject structure is generated from "#/definitions/operationBindingsObject".
type OperationBindingsObject struct {
	HTTP                 *interface{}                   `json:"http,omitempty"`
	Ws                   *interface{}                   `json:"ws,omitempty"`
	// AMQP 0-9-1 Operation Binding Object
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
	AdditionalProperties map[string]interface{}         `json:"-"`               // All unmatched properties
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
func (i *OperationBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOperationBindingsObject(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			ii.HTTP = &v
		}
	}

	if ii.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			ii.Ws = &v
		}
	}

	if ii.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			ii.Amqp1 = &v
		}
	}

	if ii.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			ii.Mqtt = &v
		}
	}

	if ii.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			ii.Mqtt5 = &v
		}
	}

	if ii.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			ii.Kafka = &v
		}
	}

	if ii.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			ii.Nats = &v
		}
	}

	if ii.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			ii.Jms = &v
		}
	}

	if ii.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			ii.Sns = &v
		}
	}

	if ii.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			ii.Sqs = &v
		}
	}

	if ii.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			ii.Stomp = &v
		}
	}

	if ii.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			ii.Redis = &v
		}
	}

	for _, key := range ignoreKeysOperationBindingsObject {
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

	*i = OperationBindingsObject(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationBindingsObject) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalOperationBindingsObject(i))
	}
	return marshalUnion(marshalOperationBindingsObject(i), i.AdditionalProperties)
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
func (i *OperationTraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.OperationTrait)
	if err != nil {
		i.OperationTrait = nil
	}

	err = json.Unmarshal(data, &i.SliceOfAnything)
	if err != nil {
		i.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationTraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.OperationTrait, i.SliceOfAnything)
}

// MessageOneOf1OneOf0 structure is generated from "#/definitions/message/oneOf/1/oneOf/0".
type MessageOneOf1OneOf0 struct {
	OneOf []Message `json:"oneOf,omitempty"`
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
	ExternalDocs  *ExternalDocs                     `json:"externalDocs,omitempty"`  // information about external documentation
	Deprecated    bool                              `json:"deprecated,omitempty"`
	Examples      []map[string]interface{}          `json:"examples,omitempty"`
	Bindings      *MessageBindingsObject            `json:"bindings,omitempty"`
	Traits        []MessageOneOf1OneOf1TraitsItems  `json:"traits,omitempty"`
	MapOfAnything map[string]interface{}            `json:"-"`                       // Key must match pattern: ^x-[\w\d\.\-\_]+$
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
func (i *MessageOneOf1OneOf1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMessageOneOf1OneOf1(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Payload == nil {
		if _, ok := m["payload"]; ok {
			var v interface{}
			ii.Payload = &v
		}
	}

	for _, key := range ignoreKeysMessageOneOf1OneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexXWD.MatchString(key) {
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

	*i = MessageOneOf1OneOf1(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageOneOf1OneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageOneOf1OneOf1(i), i.MapOfAnything)
}

// CorrelationID structure is generated from "#/definitions/correlationId".
type CorrelationID struct {
	Description   string                 `json:"description,omitempty"` // A optional description of the correlation ID. GitHub Flavored Markdown is allowed.
	Location      string                 `json:"location,omitempty"`    // A runtime expression that specifies the location of the correlation ID
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
}

type marshalCorrelationID CorrelationID

var ignoreKeysCorrelationID = []string{
	"description",
	"location",
}

// UnmarshalJSON decodes JSON.
func (i *CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalCorrelationID(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = CorrelationID(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalCorrelationID(i), i.MapOfAnything)
}

// MessageOneOf1OneOf1CorrelationID structure is generated from "#/definitions/message/oneOf/1/oneOf/1->correlationId".
type MessageOneOf1OneOf1CorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MessageOneOf1OneOf1CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.CorrelationID)
	if err != nil {
		i.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageOneOf1OneOf1CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.CorrelationID)
}

// MessageBindingsObject structure is generated from "#/definitions/messageBindingsObject".
type MessageBindingsObject struct {
	HTTP                 *interface{}                 `json:"http,omitempty"`
	Ws                   *interface{}                 `json:"ws,omitempty"`
	// AMQP 0-9-1 Message Binding Object
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
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties
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
func (i *MessageBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMessageBindingsObject(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			ii.HTTP = &v
		}
	}

	if ii.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			ii.Ws = &v
		}
	}

	if ii.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			ii.Amqp1 = &v
		}
	}

	if ii.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			ii.Mqtt = &v
		}
	}

	if ii.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			ii.Mqtt5 = &v
		}
	}

	if ii.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			ii.Kafka = &v
		}
	}

	if ii.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			ii.Nats = &v
		}
	}

	if ii.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			ii.Jms = &v
		}
	}

	if ii.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			ii.Sns = &v
		}
	}

	if ii.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			ii.Sqs = &v
		}
	}

	if ii.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			ii.Stomp = &v
		}
	}

	if ii.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			ii.Redis = &v
		}
	}

	for _, key := range ignoreKeysMessageBindingsObject {
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

	*i = MessageBindingsObject(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageBindingsObject) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalMessageBindingsObject(i))
	}
	return marshalUnion(marshalMessageBindingsObject(i), i.AdditionalProperties)
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
	ExternalDocs  *ExternalDocs              `json:"externalDocs,omitempty"`  // information about external documentation
	Deprecated    bool                       `json:"deprecated,omitempty"`
	Examples      []map[string]interface{}   `json:"examples,omitempty"`
	Bindings      *MessageBindingsObject     `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                       // Key must match pattern: ^x-[\w\d\.\-\_]+$
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
func (i *MessageTrait) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMessageTrait(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = MessageTrait(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageTrait(i), i.MapOfAnything)
}

// MessageTraitHeaders structure is generated from "#/definitions/messageTrait->headers".
type MessageTraitHeaders struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MessageTraitHeaders) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageTraitHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Schema)
}

// MessageTraitCorrelationID structure is generated from "#/definitions/messageTrait->correlationId".
type MessageTraitCorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MessageTraitCorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.CorrelationID)
	if err != nil {
		i.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageTraitCorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.CorrelationID)
}

// MessageOneOf1OneOf1TraitsItems structure is generated from "#/definitions/message/oneOf/1/oneOf/1->traits->items".
type MessageOneOf1OneOf1TraitsItems struct {
	Reference       *Reference    `json:"-"`
	MessageTrait    *MessageTrait `json:"-"`
	SliceOfAnything []interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MessageOneOf1OneOf1TraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.MessageTrait)
	if err != nil {
		i.MessageTrait = nil
	}

	err = json.Unmarshal(data, &i.SliceOfAnything)
	if err != nil {
		i.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageOneOf1OneOf1TraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.MessageTrait, i.SliceOfAnything)
}

// MessageOneOf1 structure is generated from "#/definitions/message/oneOf/1".
type MessageOneOf1 struct {
	OneOf0 *MessageOneOf1OneOf0 `json:"-"`
	OneOf1 *MessageOneOf1OneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MessageOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.OneOf0)
	if err != nil {
		i.OneOf0 = nil
	}

	err = json.Unmarshal(data, &i.OneOf1)
	if err != nil {
		i.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MessageOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.OneOf0, i.OneOf1)
}

// Message structure is generated from "#/definitions/message".
type Message struct {
	Reference *Reference     `json:"-"`
	OneOf1    *MessageOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Message) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.OneOf1)
	if err != nil {
		i.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.OneOf1)
}

// ChannelBindingsObject structure is generated from "#/definitions/channelBindingsObject".
type ChannelBindingsObject struct {
	HTTP                 *interface{}                 `json:"http,omitempty"`
	Ws                   *interface{}                 `json:"ws,omitempty"`
	// AMQP 0-9-1 Channel Binding Object
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
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties
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
func (i *ChannelBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalChannelBindingsObject(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.HTTP == nil {
		if _, ok := m["http"]; ok {
			var v interface{}
			ii.HTTP = &v
		}
	}

	if ii.Ws == nil {
		if _, ok := m["ws"]; ok {
			var v interface{}
			ii.Ws = &v
		}
	}

	if ii.Amqp1 == nil {
		if _, ok := m["amqp1"]; ok {
			var v interface{}
			ii.Amqp1 = &v
		}
	}

	if ii.Mqtt == nil {
		if _, ok := m["mqtt"]; ok {
			var v interface{}
			ii.Mqtt = &v
		}
	}

	if ii.Mqtt5 == nil {
		if _, ok := m["mqtt5"]; ok {
			var v interface{}
			ii.Mqtt5 = &v
		}
	}

	if ii.Kafka == nil {
		if _, ok := m["kafka"]; ok {
			var v interface{}
			ii.Kafka = &v
		}
	}

	if ii.Nats == nil {
		if _, ok := m["nats"]; ok {
			var v interface{}
			ii.Nats = &v
		}
	}

	if ii.Jms == nil {
		if _, ok := m["jms"]; ok {
			var v interface{}
			ii.Jms = &v
		}
	}

	if ii.Sns == nil {
		if _, ok := m["sns"]; ok {
			var v interface{}
			ii.Sns = &v
		}
	}

	if ii.Sqs == nil {
		if _, ok := m["sqs"]; ok {
			var v interface{}
			ii.Sqs = &v
		}
	}

	if ii.Stomp == nil {
		if _, ok := m["stomp"]; ok {
			var v interface{}
			ii.Stomp = &v
		}
	}

	if ii.Redis == nil {
		if _, ok := m["redis"]; ok {
			var v interface{}
			ii.Redis = &v
		}
	}

	for _, key := range ignoreKeysChannelBindingsObject {
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

	*i = ChannelBindingsObject(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ChannelBindingsObject) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalChannelBindingsObject(i))
	}
	return marshalUnion(marshalChannelBindingsObject(i), i.AdditionalProperties)
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
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties
}

type marshalExchange Exchange

var ignoreKeysExchange = []string{
	"name",
	"type",
	"durable",
	"autoDelete",
}

// UnmarshalJSON decodes JSON.
func (i *Exchange) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalExchange(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = Exchange(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Exchange) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalExchange(i))
	}
	return marshalUnion(marshalExchange(i), i.AdditionalProperties)
}

// Queue structure is generated from "#/definitions/queue".
//
// When `is`=`queue`, this object defines the queue properties.
type Queue struct {
	Name                 string                 `json:"name,omitempty"`       // The name of the queue. It MUST NOT exceed 255 characters long.
	Durable              bool                   `json:"durable,omitempty"`    // Whether the queue should survive broker restarts or not.
	Exclusive            bool                   `json:"exclusive,omitempty"`  // Whether the queue should be used only by one connection or not.
	AutoDelete           bool                   `json:"autoDelete,omitempty"` // Whether the queue should be deleted when the last consumer unsubscribes.
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties
}

type marshalQueue Queue

var ignoreKeysQueue = []string{
	"name",
	"durable",
	"exclusive",
	"autoDelete",
}

// UnmarshalJSON decodes JSON.
func (i *Queue) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalQueue(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = Queue(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Queue) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalQueue(i))
	}
	return marshalUnion(marshalQueue(i), i.AdditionalProperties)
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
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

		if regexXWD.MatchString(key) {
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

// Oauth2Flows structure is generated from "#/definitions/oauth2Flows".
type Oauth2Flows struct {
	Description          string                 `json:"description,omitempty"`
	Flows                *Oauth2FlowsFlows      `json:"flows,omitempty"`
	MapOfAnything        map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-[\w\d\.\-\_]+$
	AdditionalProperties map[string]interface{} `json:"-"`                     // All unmatched properties
}

type marshalOauth2Flows Oauth2Flows

var ignoreKeysOauth2Flows = []string{
	"description",
	"flows",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *Oauth2Flows) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOauth2Flows(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = Oauth2Flows(ii)

	return nil
}

var (
	// constOauth2Flows is unconditionally added to JSON.
	constOauth2Flows = json.RawMessage(`{"type":"oauth2"}`)
)

// MarshalJSON encodes JSON.
func (i Oauth2Flows) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOauth2Flows, marshalOauth2Flows(i), i.MapOfAnything, i.AdditionalProperties)
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
	AuthorizationURL string                 `json:"authorizationUrl,omitempty"`
	TokenURL         string                 `json:"tokenUrl,omitempty"`
	RefreshURL       string                 `json:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-[\w\d\.\-\_]+$
}

type marshalOauth2Flow Oauth2Flow

var ignoreKeysOauth2Flow = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *Oauth2Flow) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOauth2Flow(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = Oauth2Flow(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Oauth2Flow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2Flow(i), i.MapOfAnything)
}

// OpenIDConnect structure is generated from "#/definitions/openIdConnect".
type OpenIDConnect struct {
	Description   string                 `json:"description,omitempty"`
	URL           string                 `json:"openIdConnectUrl,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-[\w\d\.\-\_]+$
}

type marshalOpenIDConnect OpenIDConnect

var ignoreKeysOpenIDConnect = []string{
	"description",
	"openIdConnectUrl",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *OpenIDConnect) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOpenIDConnect(*i)

	err = json.Unmarshal(data, &ii)
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

	*i = OpenIDConnect(ii)

	return nil
}

var (
	// constOpenIDConnect is unconditionally added to JSON.
	constOpenIDConnect = json.RawMessage(`{"type":"openIdConnect"}`)
)

// MarshalJSON encodes JSON.
func (i OpenIDConnect) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnect, marshalOpenIDConnect(i), i.MapOfAnything)
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

	err = json.Unmarshal(data, &i.Oauth2Flows)
	if err != nil {
		i.Oauth2Flows = nil
	}

	err = json.Unmarshal(data, &i.OpenIDConnect)
	if err != nil {
		i.OpenIDConnect = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.UserPassword, i.APIKey, i.X509, i.SymmetricEncryption, i.AsymmetricEncryption, i.HTTPSecurityScheme, i.Oauth2Flows, i.OpenIDConnect)
}

// ComponentsSecuritySchemesWD structure is generated from "#/definitions/components->securitySchemes->^[\w\d\.\-_]+$".
type ComponentsSecuritySchemesWD struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemesWD) UnmarshalJSON(data []byte) error {
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
func (i ComponentsSecuritySchemesWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesWDValues map[string]ComponentsSecuritySchemesWD `json:"-"` // Key must match pattern: ^[\w\d\.\-_]+$
	AdditionalProperties                   map[string]interface{}                 `json:"-"` // All unmatched properties
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

		if regexWD.MatchString(key) {
			matched = true

			if i.MapOfComponentsSecuritySchemesWDValues == nil {
				i.MapOfComponentsSecuritySchemesWDValues = make(map[string]ComponentsSecuritySchemesWD, 1)
			}

			var val ComponentsSecuritySchemesWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsSecuritySchemesWDValues[key] = val
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
	return marshalUnion(i.MapOfComponentsSecuritySchemesWDValues, i.AdditionalProperties)
}

// ComponentsCorrelationIdsWD structure is generated from "#/definitions/components->correlationIds->^[\w\d\.\-_]+$".
type ComponentsCorrelationIdsWD struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsCorrelationIdsWD) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.CorrelationID)
	if err != nil {
		i.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsCorrelationIdsWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.CorrelationID)
}

// ComponentsCorrelationIds structure is generated from "#/definitions/components->correlationIds".
type ComponentsCorrelationIds struct {
	MapOfComponentsCorrelationIdsWDValues map[string]ComponentsCorrelationIdsWD `json:"-"` // Key must match pattern: ^[\w\d\.\-_]+$
	AdditionalProperties                  map[string]interface{}                `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsCorrelationIds) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsCorrelationIdsWDValues == nil {
				i.MapOfComponentsCorrelationIdsWDValues = make(map[string]ComponentsCorrelationIdsWD, 1)
			}

			var val ComponentsCorrelationIdsWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsCorrelationIdsWDValues[key] = val
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
func (i ComponentsCorrelationIds) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsCorrelationIdsWDValues, i.AdditionalProperties)
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
