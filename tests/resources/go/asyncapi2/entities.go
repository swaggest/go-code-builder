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
func (a *AsyncAPI) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAsyncAPI(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["asyncapi"]; !ok || string(v) != `"2.0.0"` {
		return fmt.Errorf(`bad or missing const value for "asyncapi" ("2.0.0" expected, %s received)`, v)
	}

	delete(rawMap, "asyncapi")

	for _, key := range ignoreKeysAsyncAPI {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*a = AsyncAPI(ma)

	return nil
}

var (
	// constAsyncAPI is unconditionally added to JSON.
	constAsyncAPI = json.RawMessage(`{"asyncapi":"2.0.0"}`)
)

// MarshalJSON encodes JSON.
func (a AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsyncAPI, marshalAsyncAPI(a), a.MapOfAnything)
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
func (i *Info) UnmarshalJSON(data []byte) error {
	var err error

	mi := marshalInfo(*i)

	err = json.Unmarshal(data, &mi)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysInfo {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mi.MapOfAnything == nil {
				mi.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mi.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*i = Info(mi)

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
func (c *Contact) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalContact(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysContact {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*c = Contact(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(c), c.MapOfAnything)
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
func (l *License) UnmarshalJSON(data []byte) error {
	var err error

	ml := marshalLicense(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysLicense {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ml.MapOfAnything == nil {
				ml.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ml.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*l = License(ml)

	return nil
}

// MarshalJSON encodes JSON.
func (l License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(l), l.MapOfAnything)
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
func (s *Server) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalServer(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysServer {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*s = Server(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(s), s.MapOfAnything)
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
func (s *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalServerVariable(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysServerVariable {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*s = ServerVariable(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(s), s.MapOfAnything)
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
func (s *ServerBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalServerBindingsObject(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if ms.HTTP == nil {
		if _, ok := rawMap["http"]; ok {
			var v interface{}
			ms.HTTP = &v
		}
	}

	if ms.Ws == nil {
		if _, ok := rawMap["ws"]; ok {
			var v interface{}
			ms.Ws = &v
		}
	}

	if ms.Amqp == nil {
		if _, ok := rawMap["amqp"]; ok {
			var v interface{}
			ms.Amqp = &v
		}
	}

	if ms.Amqp1 == nil {
		if _, ok := rawMap["amqp1"]; ok {
			var v interface{}
			ms.Amqp1 = &v
		}
	}

	if ms.Mqtt == nil {
		if _, ok := rawMap["mqtt"]; ok {
			var v interface{}
			ms.Mqtt = &v
		}
	}

	if ms.Mqtt5 == nil {
		if _, ok := rawMap["mqtt5"]; ok {
			var v interface{}
			ms.Mqtt5 = &v
		}
	}

	if ms.Kafka == nil {
		if _, ok := rawMap["kafka"]; ok {
			var v interface{}
			ms.Kafka = &v
		}
	}

	if ms.Nats == nil {
		if _, ok := rawMap["nats"]; ok {
			var v interface{}
			ms.Nats = &v
		}
	}

	if ms.Jms == nil {
		if _, ok := rawMap["jms"]; ok {
			var v interface{}
			ms.Jms = &v
		}
	}

	if ms.Sns == nil {
		if _, ok := rawMap["sns"]; ok {
			var v interface{}
			ms.Sns = &v
		}
	}

	if ms.Sqs == nil {
		if _, ok := rawMap["sqs"]; ok {
			var v interface{}
			ms.Sqs = &v
		}
	}

	if ms.Stomp == nil {
		if _, ok := rawMap["stomp"]; ok {
			var v interface{}
			ms.Stomp = &v
		}
	}

	if ms.Redis == nil {
		if _, ok := rawMap["redis"]; ok {
			var v interface{}
			ms.Redis = &v
		}
	}

	for _, key := range ignoreKeysServerBindingsObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if ms.AdditionalProperties == nil {
			ms.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ms.AdditionalProperties[key] = val
	}

	*s = ServerBindingsObject(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s ServerBindingsObject) MarshalJSON() ([]byte, error) {
	if len(s.AdditionalProperties) == 0 {
		return json.Marshal(marshalServerBindingsObject(s))
	}

	return marshalUnion(marshalServerBindingsObject(s), s.AdditionalProperties)
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
func (c *ChannelItem) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalChannelItem(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysChannelItem {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*c = ChannelItem(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c ChannelItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalChannelItem(c), c.MapOfAnything)
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
func (p *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalParameter(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysParameter {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mp.MapOfAnything == nil {
				mp.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*p = Parameter(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(p), p.MapOfAnything)
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
func (o *Operation) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOperation(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysOperation {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*o = Operation(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(o), o.MapOfAnything)
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
func (r *Reference) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalReference(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysReference {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mr.AdditionalProperties == nil {
			mr.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mr.AdditionalProperties[key] = val
	}

	*r = Reference(mr)

	return nil
}

// MarshalJSON encodes JSON.
func (r Reference) MarshalJSON() ([]byte, error) {
	if len(r.AdditionalProperties) == 0 {
		return json.Marshal(marshalReference(r))
	}

	return marshalUnion(marshalReference(r), r.AdditionalProperties)
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
func (o *OperationTrait) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOperationTrait(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysOperationTrait {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*o = OperationTrait(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationTrait(o), o.MapOfAnything)
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
func (t *Tag) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTag(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysTag {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mt.MapOfAnything == nil {
				mt.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mt.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*t = Tag(mt)

	return nil
}

// MarshalJSON encodes JSON.
func (t Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(t), t.MapOfAnything)
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
func (e *ExternalDocs) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalExternalDocs(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysExternalDocs {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if me.MapOfAnything == nil {
				me.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			me.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*e = ExternalDocs(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(e), e.MapOfAnything)
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
func (o *OperationBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOperationBindingsObject(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mo.HTTP == nil {
		if _, ok := rawMap["http"]; ok {
			var v interface{}
			mo.HTTP = &v
		}
	}

	if mo.Ws == nil {
		if _, ok := rawMap["ws"]; ok {
			var v interface{}
			mo.Ws = &v
		}
	}

	if mo.Amqp1 == nil {
		if _, ok := rawMap["amqp1"]; ok {
			var v interface{}
			mo.Amqp1 = &v
		}
	}

	if mo.Mqtt == nil {
		if _, ok := rawMap["mqtt"]; ok {
			var v interface{}
			mo.Mqtt = &v
		}
	}

	if mo.Mqtt5 == nil {
		if _, ok := rawMap["mqtt5"]; ok {
			var v interface{}
			mo.Mqtt5 = &v
		}
	}

	if mo.Kafka == nil {
		if _, ok := rawMap["kafka"]; ok {
			var v interface{}
			mo.Kafka = &v
		}
	}

	if mo.Nats == nil {
		if _, ok := rawMap["nats"]; ok {
			var v interface{}
			mo.Nats = &v
		}
	}

	if mo.Jms == nil {
		if _, ok := rawMap["jms"]; ok {
			var v interface{}
			mo.Jms = &v
		}
	}

	if mo.Sns == nil {
		if _, ok := rawMap["sns"]; ok {
			var v interface{}
			mo.Sns = &v
		}
	}

	if mo.Sqs == nil {
		if _, ok := rawMap["sqs"]; ok {
			var v interface{}
			mo.Sqs = &v
		}
	}

	if mo.Stomp == nil {
		if _, ok := rawMap["stomp"]; ok {
			var v interface{}
			mo.Stomp = &v
		}
	}

	if mo.Redis == nil {
		if _, ok := rawMap["redis"]; ok {
			var v interface{}
			mo.Redis = &v
		}
	}

	for _, key := range ignoreKeysOperationBindingsObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mo.AdditionalProperties == nil {
			mo.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mo.AdditionalProperties[key] = val
	}

	*o = OperationBindingsObject(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationBindingsObject) MarshalJSON() ([]byte, error) {
	if len(o.AdditionalProperties) == 0 {
		return json.Marshal(marshalOperationBindingsObject(o))
	}

	return marshalUnion(marshalOperationBindingsObject(o), o.AdditionalProperties)
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
func (o *OperationTraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &o.Reference)
	if err != nil {
		o.Reference = nil
	}

	err = json.Unmarshal(data, &o.OperationTrait)
	if err != nil {
		o.OperationTrait = nil
	}

	err = json.Unmarshal(data, &o.SliceOfAnything)
	if err != nil {
		o.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationTraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(o.Reference, o.OperationTrait, o.SliceOfAnything)
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
func (m *MessageOneOf1OneOf1) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessageOneOf1OneOf1(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mm.Payload == nil {
		if _, ok := rawMap["payload"]; ok {
			var v interface{}
			mm.Payload = &v
		}
	}

	for _, key := range ignoreKeysMessageOneOf1OneOf1 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mm.MapOfAnything == nil {
				mm.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mm.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*m = MessageOneOf1OneOf1(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageOneOf1OneOf1(m), m.MapOfAnything)
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
func (c *CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalCorrelationID(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysCorrelationID {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*c = CorrelationID(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalCorrelationID(c), c.MapOfAnything)
}

// MessageOneOf1OneOf1CorrelationID structure is generated from "#/definitions/message/oneOf/1/oneOf/1->correlationId".
type MessageOneOf1OneOf1CorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		m.Reference = nil
	}

	err = json.Unmarshal(data, &m.CorrelationID)
	if err != nil {
		m.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.CorrelationID)
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
func (m *MessageBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessageBindingsObject(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mm.HTTP == nil {
		if _, ok := rawMap["http"]; ok {
			var v interface{}
			mm.HTTP = &v
		}
	}

	if mm.Ws == nil {
		if _, ok := rawMap["ws"]; ok {
			var v interface{}
			mm.Ws = &v
		}
	}

	if mm.Amqp1 == nil {
		if _, ok := rawMap["amqp1"]; ok {
			var v interface{}
			mm.Amqp1 = &v
		}
	}

	if mm.Mqtt == nil {
		if _, ok := rawMap["mqtt"]; ok {
			var v interface{}
			mm.Mqtt = &v
		}
	}

	if mm.Mqtt5 == nil {
		if _, ok := rawMap["mqtt5"]; ok {
			var v interface{}
			mm.Mqtt5 = &v
		}
	}

	if mm.Kafka == nil {
		if _, ok := rawMap["kafka"]; ok {
			var v interface{}
			mm.Kafka = &v
		}
	}

	if mm.Nats == nil {
		if _, ok := rawMap["nats"]; ok {
			var v interface{}
			mm.Nats = &v
		}
	}

	if mm.Jms == nil {
		if _, ok := rawMap["jms"]; ok {
			var v interface{}
			mm.Jms = &v
		}
	}

	if mm.Sns == nil {
		if _, ok := rawMap["sns"]; ok {
			var v interface{}
			mm.Sns = &v
		}
	}

	if mm.Sqs == nil {
		if _, ok := rawMap["sqs"]; ok {
			var v interface{}
			mm.Sqs = &v
		}
	}

	if mm.Stomp == nil {
		if _, ok := rawMap["stomp"]; ok {
			var v interface{}
			mm.Stomp = &v
		}
	}

	if mm.Redis == nil {
		if _, ok := rawMap["redis"]; ok {
			var v interface{}
			mm.Redis = &v
		}
	}

	for _, key := range ignoreKeysMessageBindingsObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mm.AdditionalProperties == nil {
			mm.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mm.AdditionalProperties[key] = val
	}

	*m = MessageBindingsObject(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageBindingsObject) MarshalJSON() ([]byte, error) {
	if len(m.AdditionalProperties) == 0 {
		return json.Marshal(marshalMessageBindingsObject(m))
	}

	return marshalUnion(marshalMessageBindingsObject(m), m.AdditionalProperties)
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
func (m *MessageTrait) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessageTrait(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysMessageTrait {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mm.MapOfAnything == nil {
				mm.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mm.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*m = MessageTrait(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageTrait(m), m.MapOfAnything)
}

// MessageTraitHeaders structure is generated from "#/definitions/messageTrait->headers".
type MessageTraitHeaders struct {
	Reference *Reference             `json:"-"`
	Schema    map[string]interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageTraitHeaders) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		m.Reference = nil
	}

	err = json.Unmarshal(data, &m.Schema)
	if err != nil {
		m.Schema = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTraitHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.Schema)
}

// MessageTraitCorrelationID structure is generated from "#/definitions/messageTrait->correlationId".
type MessageTraitCorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageTraitCorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		m.Reference = nil
	}

	err = json.Unmarshal(data, &m.CorrelationID)
	if err != nil {
		m.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTraitCorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.CorrelationID)
}

// MessageOneOf1OneOf1TraitsItems structure is generated from "#/definitions/message/oneOf/1/oneOf/1->traits->items".
type MessageOneOf1OneOf1TraitsItems struct {
	Reference       *Reference    `json:"-"`
	MessageTrait    *MessageTrait `json:"-"`
	SliceOfAnything []interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1TraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		m.Reference = nil
	}

	err = json.Unmarshal(data, &m.MessageTrait)
	if err != nil {
		m.MessageTrait = nil
	}

	err = json.Unmarshal(data, &m.SliceOfAnything)
	if err != nil {
		m.SliceOfAnything = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1TraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.MessageTrait, m.SliceOfAnything)
}

// MessageOneOf1 structure is generated from "#/definitions/message/oneOf/1".
type MessageOneOf1 struct {
	OneOf0 *MessageOneOf1OneOf0 `json:"-"`
	OneOf1 *MessageOneOf1OneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.OneOf0)
	if err != nil {
		m.OneOf0 = nil
	}

	err = json.Unmarshal(data, &m.OneOf1)
	if err != nil {
		m.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.OneOf0, m.OneOf1)
}

// Message structure is generated from "#/definitions/message".
type Message struct {
	Reference *Reference     `json:"-"`
	OneOf1    *MessageOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *Message) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		m.Reference = nil
	}

	err = json.Unmarshal(data, &m.OneOf1)
	if err != nil {
		m.OneOf1 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.OneOf1)
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
func (c *ChannelBindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalChannelBindingsObject(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mc.HTTP == nil {
		if _, ok := rawMap["http"]; ok {
			var v interface{}
			mc.HTTP = &v
		}
	}

	if mc.Ws == nil {
		if _, ok := rawMap["ws"]; ok {
			var v interface{}
			mc.Ws = &v
		}
	}

	if mc.Amqp1 == nil {
		if _, ok := rawMap["amqp1"]; ok {
			var v interface{}
			mc.Amqp1 = &v
		}
	}

	if mc.Mqtt == nil {
		if _, ok := rawMap["mqtt"]; ok {
			var v interface{}
			mc.Mqtt = &v
		}
	}

	if mc.Mqtt5 == nil {
		if _, ok := rawMap["mqtt5"]; ok {
			var v interface{}
			mc.Mqtt5 = &v
		}
	}

	if mc.Kafka == nil {
		if _, ok := rawMap["kafka"]; ok {
			var v interface{}
			mc.Kafka = &v
		}
	}

	if mc.Nats == nil {
		if _, ok := rawMap["nats"]; ok {
			var v interface{}
			mc.Nats = &v
		}
	}

	if mc.Jms == nil {
		if _, ok := rawMap["jms"]; ok {
			var v interface{}
			mc.Jms = &v
		}
	}

	if mc.Sns == nil {
		if _, ok := rawMap["sns"]; ok {
			var v interface{}
			mc.Sns = &v
		}
	}

	if mc.Sqs == nil {
		if _, ok := rawMap["sqs"]; ok {
			var v interface{}
			mc.Sqs = &v
		}
	}

	if mc.Stomp == nil {
		if _, ok := rawMap["stomp"]; ok {
			var v interface{}
			mc.Stomp = &v
		}
	}

	if mc.Redis == nil {
		if _, ok := rawMap["redis"]; ok {
			var v interface{}
			mc.Redis = &v
		}
	}

	for _, key := range ignoreKeysChannelBindingsObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mc.AdditionalProperties == nil {
			mc.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mc.AdditionalProperties[key] = val
	}

	*c = ChannelBindingsObject(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c ChannelBindingsObject) MarshalJSON() ([]byte, error) {
	if len(c.AdditionalProperties) == 0 {
		return json.Marshal(marshalChannelBindingsObject(c))
	}

	return marshalUnion(marshalChannelBindingsObject(c), c.AdditionalProperties)
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
func (e *Exchange) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalExchange(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysExchange {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if me.AdditionalProperties == nil {
			me.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		me.AdditionalProperties[key] = val
	}

	*e = Exchange(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e Exchange) MarshalJSON() ([]byte, error) {
	if len(e.AdditionalProperties) == 0 {
		return json.Marshal(marshalExchange(e))
	}

	return marshalUnion(marshalExchange(e), e.AdditionalProperties)
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
func (q *Queue) UnmarshalJSON(data []byte) error {
	var err error

	mq := marshalQueue(*q)

	err = json.Unmarshal(data, &mq)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysQueue {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mq.AdditionalProperties == nil {
			mq.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mq.AdditionalProperties[key] = val
	}

	*q = Queue(mq)

	return nil
}

// MarshalJSON encodes JSON.
func (q Queue) MarshalJSON() ([]byte, error) {
	if len(q.AdditionalProperties) == 0 {
		return json.Marshal(marshalQueue(q))
	}

	return marshalUnion(marshalQueue(q), q.AdditionalProperties)
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
func (u *UserPassword) UnmarshalJSON(data []byte) error {
	var err error

	mu := marshalUserPassword(*u)

	err = json.Unmarshal(data, &mu)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"userPassword"` {
		return fmt.Errorf(`bad or missing const value for "type" ("userPassword" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysUserPassword {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mu.MapOfAnything == nil {
				mu.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mu.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*u = UserPassword(mu)

	return nil
}

var (
	// constUserPassword is unconditionally added to JSON.
	constUserPassword = json.RawMessage(`{"type":"userPassword"}`)
)

// MarshalJSON encodes JSON.
func (u UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(constUserPassword, marshalUserPassword(u), u.MapOfAnything)
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
func (a *APIKey) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAPIKey(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"apiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("apiKey" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysAPIKey {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*a = APIKey(ma)

	return nil
}

var (
	// constAPIKey is unconditionally added to JSON.
	constAPIKey = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON.
func (a APIKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKey, marshalAPIKey(a), a.MapOfAnything)
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
func (x *X509) UnmarshalJSON(data []byte) error {
	var err error

	mx := marshalX509(*x)

	err = json.Unmarshal(data, &mx)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"X509"` {
		return fmt.Errorf(`bad or missing const value for "type" ("X509" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysX509 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mx.MapOfAnything == nil {
				mx.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mx.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*x = X509(mx)

	return nil
}

var (
	// constX509 is unconditionally added to JSON.
	constX509 = json.RawMessage(`{"type":"X509"}`)
)

// MarshalJSON encodes JSON.
func (x X509) MarshalJSON() ([]byte, error) {
	return marshalUnion(constX509, marshalX509(x), x.MapOfAnything)
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
func (s *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSymmetricEncryption(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"symmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("symmetricEncryption" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysSymmetricEncryption {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*s = SymmetricEncryption(ms)

	return nil
}

var (
	// constSymmetricEncryption is unconditionally added to JSON.
	constSymmetricEncryption = json.RawMessage(`{"type":"symmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (s SymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constSymmetricEncryption, marshalSymmetricEncryption(s), s.MapOfAnything)
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
func (a *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAsymmetricEncryption(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"asymmetricEncryption"` {
		return fmt.Errorf(`bad or missing const value for "type" ("asymmetricEncryption" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysAsymmetricEncryption {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*a = AsymmetricEncryption(ma)

	return nil
}

var (
	// constAsymmetricEncryption is unconditionally added to JSON.
	constAsymmetricEncryption = json.RawMessage(`{"type":"asymmetricEncryption"}`)
)

// MarshalJSON encodes JSON.
func (a AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsymmetricEncryption, marshalAsymmetricEncryption(a), a.MapOfAnything)
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
func (n *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	mn := marshalNonBearerHTTPSecurityScheme(*n)

	err = json.Unmarshal(data, &mn)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysNonBearerHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mn.MapOfAnything == nil {
				mn.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mn.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*n = NonBearerHTTPSecurityScheme(mn)

	return nil
}

var (
	// constNonBearerHTTPSecurityScheme is unconditionally added to JSON.
	constNonBearerHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON.
func (n NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constNonBearerHTTPSecurityScheme, marshalNonBearerHTTPSecurityScheme(n), n.MapOfAnything)
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
func (b *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	mb := marshalBearerHTTPSecurityScheme(*b)

	err = json.Unmarshal(data, &mb)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["scheme"]; !ok || string(v) != `"bearer"` {
		return fmt.Errorf(`bad or missing const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(rawMap, "scheme")

	if v, ok := rawMap["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysBearerHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mb.MapOfAnything == nil {
				mb.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mb.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*b = BearerHTTPSecurityScheme(mb)

	return nil
}

var (
	// constBearerHTTPSecurityScheme is unconditionally added to JSON.
	constBearerHTTPSecurityScheme = json.RawMessage(`{"scheme":"bearer","type":"http"}`)
)

// MarshalJSON encodes JSON.
func (b BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBearerHTTPSecurityScheme, marshalBearerHTTPSecurityScheme(b), b.MapOfAnything)
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
func (a *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAPIKeyHTTPSecurityScheme(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"httpApiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("httpApiKey" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysAPIKeyHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*a = APIKeyHTTPSecurityScheme(ma)

	return nil
}

var (
	// constAPIKeyHTTPSecurityScheme is unconditionally added to JSON.
	constAPIKeyHTTPSecurityScheme = json.RawMessage(`{"type":"httpApiKey"}`)
)

// MarshalJSON encodes JSON.
func (a APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeyHTTPSecurityScheme, marshalAPIKeyHTTPSecurityScheme(a), a.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (h *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &h.NonBearerHTTPSecurityScheme)
	if err != nil {
		h.NonBearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &h.BearerHTTPSecurityScheme)
	if err != nil {
		h.BearerHTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &h.APIKeyHTTPSecurityScheme)
	if err != nil {
		h.APIKeyHTTPSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (h HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(h.NonBearerHTTPSecurityScheme, h.BearerHTTPSecurityScheme, h.APIKeyHTTPSecurityScheme)
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
func (o *Oauth2Flows) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOauth2Flows(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"oauth2"` {
		return fmt.Errorf(`bad or missing const value for "type" ("oauth2" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysOauth2Flows {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if mo.AdditionalProperties == nil {
			mo.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mo.AdditionalProperties[key] = val
	}

	*o = Oauth2Flows(mo)

	return nil
}

var (
	// constOauth2Flows is unconditionally added to JSON.
	constOauth2Flows = json.RawMessage(`{"type":"oauth2"}`)
)

// MarshalJSON encodes JSON.
func (o Oauth2Flows) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOauth2Flows, marshalOauth2Flows(o), o.MapOfAnything, o.AdditionalProperties)
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
func (o *Oauth2Flow) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOauth2Flow(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range ignoreKeysOauth2Flow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*o = Oauth2Flow(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o Oauth2Flow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2Flow(o), o.MapOfAnything)
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
func (o *OpenIDConnect) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOpenIDConnect(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, ok := rawMap["type"]; !ok || string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad or missing const value for "type" ("openIdConnect" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range ignoreKeysOpenIDConnect {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWD.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	*o = OpenIDConnect(mo)

	return nil
}

var (
	// constOpenIDConnect is unconditionally added to JSON.
	constOpenIDConnect = json.RawMessage(`{"type":"openIdConnect"}`)
)

// MarshalJSON encodes JSON.
func (o OpenIDConnect) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnect, marshalOpenIDConnect(o), o.MapOfAnything)
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
func (s *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &s.UserPassword)
	if err != nil {
		s.UserPassword = nil
	}

	err = json.Unmarshal(data, &s.APIKey)
	if err != nil {
		s.APIKey = nil
	}

	err = json.Unmarshal(data, &s.X509)
	if err != nil {
		s.X509 = nil
	}

	err = json.Unmarshal(data, &s.SymmetricEncryption)
	if err != nil {
		s.SymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &s.AsymmetricEncryption)
	if err != nil {
		s.AsymmetricEncryption = nil
	}

	err = json.Unmarshal(data, &s.HTTPSecurityScheme)
	if err != nil {
		s.HTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &s.Oauth2Flows)
	if err != nil {
		s.Oauth2Flows = nil
	}

	err = json.Unmarshal(data, &s.OpenIDConnect)
	if err != nil {
		s.OpenIDConnect = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.UserPassword, s.APIKey, s.X509, s.SymmetricEncryption, s.AsymmetricEncryption, s.HTTPSecurityScheme, s.Oauth2Flows, s.OpenIDConnect)
}

// ComponentsSecuritySchemesWD structure is generated from "#/definitions/components->securitySchemes->^[\w\d\.\-_]+$".
type ComponentsSecuritySchemesWD struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSecuritySchemesWD) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &c.Reference)
	if err != nil {
		c.Reference = nil
	}

	err = json.Unmarshal(data, &c.SecurityScheme)
	if err != nil {
		c.SecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSecuritySchemesWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.Reference, c.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesWDValues map[string]ComponentsSecuritySchemesWD `json:"-"` // Key must match pattern: `^[\w\d\.\-_]+$`.
	AdditionalProperties                   map[string]interface{}                 `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexWD.MatchString(key) {
			matched = true

			if c.MapOfComponentsSecuritySchemesWDValues == nil {
				c.MapOfComponentsSecuritySchemesWDValues = make(map[string]ComponentsSecuritySchemesWD, 1)
			}

			var val ComponentsSecuritySchemesWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfComponentsSecuritySchemesWDValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if c.AdditionalProperties == nil {
			c.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		c.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfComponentsSecuritySchemesWDValues, c.AdditionalProperties)
}

// ComponentsCorrelationIdsWD structure is generated from "#/definitions/components->correlationIds->^[\w\d\.\-_]+$".
type ComponentsCorrelationIdsWD struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsCorrelationIdsWD) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &c.Reference)
	if err != nil {
		c.Reference = nil
	}

	err = json.Unmarshal(data, &c.CorrelationID)
	if err != nil {
		c.CorrelationID = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsCorrelationIdsWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.Reference, c.CorrelationID)
}

// ComponentsCorrelationIds structure is generated from "#/definitions/components->correlationIds".
type ComponentsCorrelationIds struct {
	MapOfComponentsCorrelationIdsWDValues map[string]ComponentsCorrelationIdsWD `json:"-"` // Key must match pattern: `^[\w\d\.\-_]+$`.
	AdditionalProperties                  map[string]interface{}                `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsCorrelationIds) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexWD.MatchString(key) {
			matched = true

			if c.MapOfComponentsCorrelationIdsWDValues == nil {
				c.MapOfComponentsCorrelationIdsWDValues = make(map[string]ComponentsCorrelationIdsWD, 1)
			}

			var val ComponentsCorrelationIdsWD

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfComponentsCorrelationIdsWDValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if c.AdditionalProperties == nil {
			c.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		c.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsCorrelationIds) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfComponentsCorrelationIdsWDValues, c.AdditionalProperties)
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
