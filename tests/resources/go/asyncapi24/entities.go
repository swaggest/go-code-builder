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
// AsyncAPI 2.4.0 schema.
//
// !!Auto generated!!
//  Do not manually edit. .
type AsyncAPI struct {
	// A unique id representing the application.
	// Format: uri.
	ID                 string                                 `json:"id,omitempty"`
	// General information about the API.
	// Required.
	Info               Info                                   `json:"info"`
	Servers            map[string]ServersAdditionalProperties `json:"servers,omitempty"`            // An object representing multiple servers.
	DefaultContentType string                                 `json:"defaultContentType,omitempty"`
	Channels           map[string]ChannelItem                 `json:"channels"`                     // Required.
	Components         *Components                            `json:"components,omitempty"`         // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags               []Tag                                  `json:"tags,omitempty"`
	ExternalDocs       *ExternalDocs                          `json:"externalDocs,omitempty"`       // Information about external documentation.
	MapOfAnything      map[string]interface{}                 `json:"-"`                            // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalAsyncAPI AsyncAPI

var knownKeysAsyncAPI = []string{
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

var requireKeysAsyncAPI = []string{
	"asyncapi",
	"info",
	"channels",
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

	for _, key := range requireKeysAsyncAPI {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["asyncapi"]; exists && string(v) != `"2.4.0"` {
		return fmt.Errorf(`bad const value for "asyncapi" ("2.4.0" expected, %s received)`, v)
	}

	delete(rawMap, "asyncapi")

	for _, key := range knownKeysAsyncAPI {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in AsyncAPI: %v", offendingKeys)
	}

	*a = AsyncAPI(ma)

	return nil
}

var (
	// constAsyncAPI is unconditionally added to JSON.
	constAsyncAPI = json.RawMessage(`{"asyncapi":"2.4.0"}`)
)

// MarshalJSON encodes JSON.
func (a AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAsyncAPI, marshalAsyncAPI(a), a.MapOfAnything)
}

// Info structure is generated from "http://asyncapi.com/definitions/2.4.0/info.json".
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
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalInfo Info

var knownKeysInfo = []string{
	"title",
	"version",
	"description",
	"termsOfService",
	"contact",
	"license",
}

var requireKeysInfo = []string{
	"version",
	"title",
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

	for _, key := range requireKeysInfo {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysInfo {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Info: %v", offendingKeys)
	}

	*i = Info(mi)

	return nil
}

// MarshalJSON encodes JSON.
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnything)
}

// Contact structure is generated from "http://asyncapi.com/definitions/2.4.0/contact.json".
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
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalContact Contact

var knownKeysContact = []string{
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

	for _, key := range knownKeysContact {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Contact: %v", offendingKeys)
	}

	*c = Contact(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(c), c.MapOfAnything)
}

// License structure is generated from "http://asyncapi.com/definitions/2.4.0/license.json".
type License struct {
	// The name of the license type. It's encouraged to use an OSI compatible license.
	// Required.
	Name          string                 `json:"name"`
	// The URL pointing to the license.
	// Format: uri.
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`             // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalLicense License

var knownKeysLicense = []string{
	"name",
	"url",
}

var requireKeysLicense = []string{
	"name",
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

	for _, key := range requireKeysLicense {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysLicense {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in License: %v", offendingKeys)
	}

	*l = License(ml)

	return nil
}

// MarshalJSON encodes JSON.
func (l License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(l), l.MapOfAnything)
}

// Reference structure is generated from "http://asyncapi.com/definitions/2.4.0/Reference.json".
type Reference struct {
	// Format: uri-reference.
	// Required.
	Ref                  string                 `json:"$ref"`
	AdditionalProperties map[string]interface{} `json:"-"`    // All unmatched properties.
}

type marshalReference Reference

var knownKeysReference = []string{
	"$ref",
}

var requireKeysReference = []string{
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

	for _, key := range requireKeysReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysReference {
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

// Server structure is generated from "http://asyncapi.com/definitions/2.4.0/server.json".
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
	Bindings        *BindingsObject           `json:"bindings,omitempty"`
	MapOfAnything   map[string]interface{}    `json:"-"`                         // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalServer Server

var knownKeysServer = []string{
	"url",
	"description",
	"protocol",
	"protocolVersion",
	"variables",
	"security",
	"bindings",
}

var requireKeysServer = []string{
	"url",
	"protocol",
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

	for _, key := range requireKeysServer {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysServer {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Server: %v", offendingKeys)
	}

	*s = Server(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(s), s.MapOfAnything)
}

// ServerVariable structure is generated from "http://asyncapi.com/definitions/2.4.0/serverVariable.json".
//
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Examples      []string               `json:"examples,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalServerVariable ServerVariable

var knownKeysServerVariable = []string{
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

	for _, key := range knownKeysServerVariable {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ServerVariable: %v", offendingKeys)
	}

	*s = ServerVariable(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(s), s.MapOfAnything)
}

// BindingsObject structure is generated from "http://asyncapi.com/definitions/2.4.0/bindingsObject.json".
type BindingsObject struct {
	HTTP                 *interface{}           `json:"http,omitempty"`
	Ws                   *interface{}           `json:"ws,omitempty"`
	Amqp                 *interface{}           `json:"amqp,omitempty"`
	Amqp1                *interface{}           `json:"amqp1,omitempty"`
	Mqtt                 *interface{}           `json:"mqtt,omitempty"`
	Mqtt5                *interface{}           `json:"mqtt5,omitempty"`
	Kafka                *interface{}           `json:"kafka,omitempty"`
	Anypointmq           *interface{}           `json:"anypointmq,omitempty"`
	Nats                 *interface{}           `json:"nats,omitempty"`
	Jms                  *interface{}           `json:"jms,omitempty"`
	Sns                  *interface{}           `json:"sns,omitempty"`
	Sqs                  *interface{}           `json:"sqs,omitempty"`
	Stomp                *interface{}           `json:"stomp,omitempty"`
	Redis                *interface{}           `json:"redis,omitempty"`
	Ibmmq                *interface{}           `json:"ibmmq,omitempty"`
	Solace               *interface{}           `json:"solace,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties.
}

type marshalBindingsObject BindingsObject

var knownKeysBindingsObject = []string{
	"http",
	"ws",
	"amqp",
	"amqp1",
	"mqtt",
	"mqtt5",
	"kafka",
	"anypointmq",
	"nats",
	"jms",
	"sns",
	"sqs",
	"stomp",
	"redis",
	"ibmmq",
	"solace",
}

// UnmarshalJSON decodes JSON.
func (b *BindingsObject) UnmarshalJSON(data []byte) error {
	var err error

	mb := marshalBindingsObject(*b)

	err = json.Unmarshal(data, &mb)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mb.HTTP == nil {
		if _, ok := rawMap["http"]; ok {
			var v interface{}
			mb.HTTP = &v
		}
	}

	if mb.Ws == nil {
		if _, ok := rawMap["ws"]; ok {
			var v interface{}
			mb.Ws = &v
		}
	}

	if mb.Amqp == nil {
		if _, ok := rawMap["amqp"]; ok {
			var v interface{}
			mb.Amqp = &v
		}
	}

	if mb.Amqp1 == nil {
		if _, ok := rawMap["amqp1"]; ok {
			var v interface{}
			mb.Amqp1 = &v
		}
	}

	if mb.Mqtt == nil {
		if _, ok := rawMap["mqtt"]; ok {
			var v interface{}
			mb.Mqtt = &v
		}
	}

	if mb.Mqtt5 == nil {
		if _, ok := rawMap["mqtt5"]; ok {
			var v interface{}
			mb.Mqtt5 = &v
		}
	}

	if mb.Kafka == nil {
		if _, ok := rawMap["kafka"]; ok {
			var v interface{}
			mb.Kafka = &v
		}
	}

	if mb.Anypointmq == nil {
		if _, ok := rawMap["anypointmq"]; ok {
			var v interface{}
			mb.Anypointmq = &v
		}
	}

	if mb.Nats == nil {
		if _, ok := rawMap["nats"]; ok {
			var v interface{}
			mb.Nats = &v
		}
	}

	if mb.Jms == nil {
		if _, ok := rawMap["jms"]; ok {
			var v interface{}
			mb.Jms = &v
		}
	}

	if mb.Sns == nil {
		if _, ok := rawMap["sns"]; ok {
			var v interface{}
			mb.Sns = &v
		}
	}

	if mb.Sqs == nil {
		if _, ok := rawMap["sqs"]; ok {
			var v interface{}
			mb.Sqs = &v
		}
	}

	if mb.Stomp == nil {
		if _, ok := rawMap["stomp"]; ok {
			var v interface{}
			mb.Stomp = &v
		}
	}

	if mb.Redis == nil {
		if _, ok := rawMap["redis"]; ok {
			var v interface{}
			mb.Redis = &v
		}
	}

	if mb.Ibmmq == nil {
		if _, ok := rawMap["ibmmq"]; ok {
			var v interface{}
			mb.Ibmmq = &v
		}
	}

	if mb.Solace == nil {
		if _, ok := rawMap["solace"]; ok {
			var v interface{}
			mb.Solace = &v
		}
	}

	for _, key := range knownKeysBindingsObject {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mb.AdditionalProperties == nil {
			mb.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mb.AdditionalProperties[key] = val
	}

	*b = BindingsObject(mb)

	return nil
}

// MarshalJSON encodes JSON.
func (b BindingsObject) MarshalJSON() ([]byte, error) {
	if len(b.AdditionalProperties) == 0 {
		return json.Marshal(marshalBindingsObject(b))
	}

	return marshalUnion(marshalBindingsObject(b), b.AdditionalProperties)
}

// ServersAdditionalProperties structure is generated from "http://asyncapi.com/definitions/2.4.0/servers.json->additionalProperties".
type ServersAdditionalProperties struct {
	Reference *Reference `json:"-"`
	Server    *Server    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *ServersAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		s.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.Server)
	if err != nil {
		oneOfErrors["Server"] = err
		s.Server = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ServersAdditionalProperties with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s ServersAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.Reference, s.Server)
}

// ChannelItem structure is generated from "http://asyncapi.com/definitions/2.4.0/channelItem.json".
type ChannelItem struct {
	Ref           string                 `json:"$ref,omitempty"`        // Format: uri-reference.
	Parameters    map[string]Parameter   `json:"parameters,omitempty"`
	Description   string                 `json:"description,omitempty"` // A description of the channel.
	Servers       []string               `json:"servers,omitempty"`     // The names of the servers on which this channel is available. If absent or empty then this channel must be available on all servers.
	Publish       *Operation             `json:"publish,omitempty"`
	Subscribe     *Operation             `json:"subscribe,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`
	Bindings      *BindingsObject        `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalChannelItem ChannelItem

var knownKeysChannelItem = []string{
	"$ref",
	"parameters",
	"description",
	"servers",
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

	for _, key := range knownKeysChannelItem {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ChannelItem: %v", offendingKeys)
	}

	*c = ChannelItem(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c ChannelItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalChannelItem(c), c.MapOfAnything)
}

// Parameter structure is generated from "http://asyncapi.com/definitions/2.4.0/parameter.json".
type Parameter struct {
	Description   string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use. GitHub Flavored Markdown is allowed.
	Schema        map[string]interface{} `json:"schema,omitempty"`
	// A runtime expression that specifies the location of the parameter value.
	// Value must match pattern: `^\$message\.(header|payload)#(\/(([^\/~])|(~[01]))*)*`.
	Location      string                 `json:"location,omitempty"`
	Ref           string                 `json:"$ref,omitempty"`        // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalParameter Parameter

var knownKeysParameter = []string{
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

	for _, key := range knownKeysParameter {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Parameter: %v", offendingKeys)
	}

	*p = Parameter(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(p), p.MapOfAnything)
}

// Operation structure is generated from "http://asyncapi.com/definitions/2.4.0/operation.json".
type Operation struct {
	Traits        []OperationTraitsItems `json:"traits,omitempty"`
	Summary       string                 `json:"summary,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	ID            string                 `json:"operationId,omitempty"`
	Bindings      *BindingsObject        `json:"bindings,omitempty"`
	Message       *Message               `json:"message,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalOperation Operation

var knownKeysOperation = []string{
	"traits",
	"summary",
	"description",
	"security",
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

	for _, key := range knownKeysOperation {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Operation: %v", offendingKeys)
	}

	*o = Operation(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(o), o.MapOfAnything)
}

// OperationTrait structure is generated from "http://asyncapi.com/definitions/2.4.0/operationTrait.json".
type OperationTrait struct {
	Summary       string                 `json:"summary,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	OperationID   string                 `json:"operationId,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	Bindings      *BindingsObject        `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalOperationTrait OperationTrait

var knownKeysOperationTrait = []string{
	"summary",
	"description",
	"tags",
	"externalDocs",
	"operationId",
	"security",
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

	for _, key := range knownKeysOperationTrait {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in OperationTrait: %v", offendingKeys)
	}

	*o = OperationTrait(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationTrait(o), o.MapOfAnything)
}

// Tag structure is generated from "http://asyncapi.com/definitions/2.4.0/tag.json".
type Tag struct {
	Name          string                 `json:"name"`                   // Required.
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalTag Tag

var knownKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

var requireKeysTag = []string{
	"name",
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

	for _, key := range requireKeysTag {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysTag {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Tag: %v", offendingKeys)
	}

	*t = Tag(mt)

	return nil
}

// MarshalJSON encodes JSON.
func (t Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(t), t.MapOfAnything)
}

// ExternalDocs structure is generated from "http://asyncapi.com/definitions/2.4.0/externalDocs.json".
//
// information about external documentation.
type ExternalDocs struct {
	Description   string                 `json:"description,omitempty"`
	// Format: uri.
	// Required.
	URL           string                 `json:"url"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalExternalDocs ExternalDocs

var knownKeysExternalDocs = []string{
	"description",
	"url",
}

var requireKeysExternalDocs = []string{
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

	for _, key := range requireKeysExternalDocs {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysExternalDocs {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ExternalDocs: %v", offendingKeys)
	}

	*e = ExternalDocs(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(e), e.MapOfAnything)
}

// OperationTraitsItems structure is generated from "http://asyncapi.com/definitions/2.4.0/operation.json->traits->items".
type OperationTraitsItems struct {
	Reference       *Reference      `json:"-"`
	OperationTrait  *OperationTrait `json:"-"`
	SliceOfAnything []interface{}   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (o *OperationTraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 3)
	oneOfValid := 0

	err = json.Unmarshal(data, &o.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		o.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &o.OperationTrait)
	if err != nil {
		oneOfErrors["OperationTrait"] = err
		o.OperationTrait = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &o.SliceOfAnything)
	if err != nil {
		oneOfErrors["SliceOfAnything"] = err
		o.SliceOfAnything = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for OperationTraitsItems with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationTraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(o.Reference, o.OperationTrait, o.SliceOfAnything)
}

// MessageOneOf1OneOf0 structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/0".
type MessageOneOf1OneOf0 struct {
	OneOf []Message `json:"oneOf"` // Required.
}

type marshalMessageOneOf1OneOf0 MessageOneOf1OneOf0

var knownKeysMessageOneOf1OneOf0 = []string{
	"oneOf",
}

var requireKeysMessageOneOf1OneOf0 = []string{
	"oneOf",
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf0) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessageOneOf1OneOf0(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysMessageOneOf1OneOf0 {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysMessageOneOf1OneOf0 {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in MessageOneOf1OneOf0: %v", offendingKeys)
	}

	*m = MessageOneOf1OneOf0(mm)

	return nil
}


// MessageOneOf1OneOf1 structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1".
type MessageOneOf1OneOf1 struct {
	SchemaFormat  string                             `json:"schemaFormat,omitempty"`
	ContentType   string                             `json:"contentType,omitempty"`
	Headers       *MessageOneOf1OneOf1Headers        `json:"headers,omitempty"`
	MessageID     string                             `json:"messageId,omitempty"`
	Payload       *interface{}                       `json:"payload,omitempty"`
	CorrelationID *MessageOneOf1OneOf1CorrelationID  `json:"correlationId,omitempty"`
	Tags          []Tag                              `json:"tags,omitempty"`
	Summary       string                             `json:"summary,omitempty"`       // A brief summary of the message.
	Name          string                             `json:"name,omitempty"`          // Name of the message.
	Title         string                             `json:"title,omitempty"`         // A human-friendly title for the message.
	Description   string                             `json:"description,omitempty"`   // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs                      `json:"externalDocs,omitempty"`  // Information about external documentation.
	Deprecated    bool                               `json:"deprecated,omitempty"`
	Examples      []MessageOneOf1OneOf1ExamplesItems `json:"examples,omitempty"`
	Bindings      *BindingsObject                    `json:"bindings,omitempty"`
	Traits        []MessageOneOf1OneOf1TraitsItems   `json:"traits,omitempty"`
	MapOfAnything map[string]interface{}             `json:"-"`                       // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalMessageOneOf1OneOf1 MessageOneOf1OneOf1

var knownKeysMessageOneOf1OneOf1 = []string{
	"schemaFormat",
	"contentType",
	"headers",
	"messageId",
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

	for _, key := range knownKeysMessageOneOf1OneOf1 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in MessageOneOf1OneOf1: %v", offendingKeys)
	}

	*m = MessageOneOf1OneOf1(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageOneOf1OneOf1(m), m.MapOfAnything)
}

// JSONSchema structure is generated from "jsonSchema".
//
// Core schema meta-schema.
type JSONSchema struct {
	ID                   string                                                `json:"id,omitempty"`                   // Format: uri.
	Schema               string                                                `json:"$schema,omitempty"`              // Format: uri.
	Title                string                                                `json:"title,omitempty"`
	Description          string                                                `json:"description,omitempty"`
	Default              *interface{}                                          `json:"default,omitempty"`
	MultipleOf           float64                                               `json:"multipleOf,omitempty"`
	Maximum              float64                                               `json:"maximum,omitempty"`
	ExclusiveMaximum     *JSONSchemaExclusiveMaximum                           `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                               `json:"minimum,omitempty"`
	ExclusiveMinimum     *JSONSchemaExclusiveMinimum                           `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                                 `json:"maxLength,omitempty"`
	MinLength            int64                                                 `json:"minLength,omitempty"`
	Pattern              string                                                `json:"pattern,omitempty"`              // Format: regex.
	ExtraProperties      map[string]interface{}                                `json:"-"`                              // All unmatched properties.
	TypeObject           *JSONSchema                                           `json:"-"`
	TypeBoolean          *bool                                                 `json:"-"`
	AdditionalItems      *JSONSchemaAdditionalItems                            `json:"additionalItems,omitempty"`
	Items                *JSONSchemaItems                                      `json:"items,omitempty"`
	MaxItems             int64                                                 `json:"maxItems,omitempty"`
	MinItems             int64                                                 `json:"minItems,omitempty"`
	UniqueItems          bool                                                  `json:"uniqueItems,omitempty"`
	MaxProperties        int64                                                 `json:"maxProperties,omitempty"`
	MinProperties        int64                                                 `json:"minProperties,omitempty"`
	Required             []string                                              `json:"required,omitempty"`
	AdditionalProperties *JSONSchemaAdditionalProperties                       `json:"additionalProperties,omitempty"`
	Definitions          map[string]interface{}                                `json:"definitions,omitempty"`
	Properties           map[string]interface{}                                `json:"properties,omitempty"`
	PatternProperties    map[string]interface{}                                `json:"patternProperties,omitempty"`
	Dependencies         map[string]JSONSchemaDependenciesAdditionalProperties `json:"dependencies,omitempty"`
	Enum                 []interface{}                                         `json:"enum,omitempty"`
	Type                 *JSONSchemaType                                       `json:"type,omitempty"`
	Format               string                                                `json:"format,omitempty"`
	Ref                  string                                                `json:"$ref,omitempty"`                 // Format: uri-reference.
	AllOf                []interface{}                                         `json:"allOf,omitempty"`
	AnyOf                []interface{}                                         `json:"anyOf,omitempty"`
	OneOf                []interface{}                                         `json:"oneOf,omitempty"`
	Not                  *JSONSchema                                           `json:"not,omitempty"`                  // Core schema meta-schema.
	Const                *interface{}                                          `json:"const,omitempty"`
	Contains             *JSONSchema                                           `json:"contains,omitempty"`             // Core schema meta-schema.
	PropertyNames        *JSONSchema                                           `json:"propertyNames,omitempty"`        // Core schema meta-schema.
	If                   *JSONSchema                                           `json:"if,omitempty"`                   // Core schema meta-schema.
	Then                 *JSONSchema                                           `json:"then,omitempty"`                 // Core schema meta-schema.
	Else                 *JSONSchema                                           `json:"else,omitempty"`                 // Core schema meta-schema.
	ContentEncoding      string                                                `json:"contentEncoding,omitempty"`
	ContentMediaType     string                                                `json:"contentMediaType,omitempty"`
}

type marshalJSONSchema JSONSchema

var knownKeysJSONSchema = []string{
	"id",
	"$schema",
	"title",
	"description",
	"default",
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
	"maxProperties",
	"minProperties",
	"required",
	"additionalProperties",
	"definitions",
	"properties",
	"patternProperties",
	"dependencies",
	"enum",
	"type",
	"format",
	"$ref",
	"allOf",
	"anyOf",
	"oneOf",
	"not",
	"const",
	"contains",
	"propertyNames",
	"if",
	"then",
	"else",
	"contentEncoding",
	"contentMediaType",
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchema) UnmarshalJSON(data []byte) error {
	var err error

	mj := marshalJSONSchema(*j)

	err = json.Unmarshal(data, &mj)
	if err != nil {
		return err
	}

	typeValid := false

	if !typeValid {
		err = json.Unmarshal(data, &mj.TypeObject)
		if err != nil {
			mj.TypeObject = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		err = json.Unmarshal(data, &mj.TypeBoolean)
		if err != nil {
			mj.TypeBoolean = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mj.Default == nil {
		if _, ok := rawMap["default"]; ok {
			var v interface{}
			mj.Default = &v
		}
	}

	if mj.Const == nil {
		if _, ok := rawMap["const"]; ok {
			var v interface{}
			mj.Const = &v
		}
	}

	for _, key := range knownKeysJSONSchema {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mj.ExtraProperties == nil {
			mj.ExtraProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mj.ExtraProperties[key] = val
	}

	*j = JSONSchema(mj)

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalJSONSchema(j), j.ExtraProperties, j.TypeObject, j.TypeBoolean)
}

// JSONSchemaExclusiveMaximum structure is generated from "jsonSchema->exclusiveMaximum".
type JSONSchemaExclusiveMaximum struct {
	TypeBoolean *bool    `json:"-"`
	TypeNumber  *float64 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaExclusiveMaximum) UnmarshalJSON(data []byte) error {
	var err error

	typeValid := false

	if !typeValid {
		err = json.Unmarshal(data, &j.TypeBoolean)
		if err != nil {
			j.TypeBoolean = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		err = json.Unmarshal(data, &j.TypeNumber)
		if err != nil {
			j.TypeNumber = nil
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
func (j JSONSchemaExclusiveMaximum) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.TypeBoolean, j.TypeNumber)
}

// JSONSchemaExclusiveMinimum structure is generated from "jsonSchema->exclusiveMinimum".
type JSONSchemaExclusiveMinimum struct {
	TypeBoolean *bool    `json:"-"`
	TypeNumber  *float64 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaExclusiveMinimum) UnmarshalJSON(data []byte) error {
	var err error

	typeValid := false

	if !typeValid {
		err = json.Unmarshal(data, &j.TypeBoolean)
		if err != nil {
			j.TypeBoolean = nil
		} else {
			typeValid = true
		}
	}

	if !typeValid {
		err = json.Unmarshal(data, &j.TypeNumber)
		if err != nil {
			j.TypeNumber = nil
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
func (j JSONSchemaExclusiveMinimum) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.TypeBoolean, j.TypeNumber)
}

// JSONSchemaAdditionalItems structure is generated from "jsonSchema->additionalItems".
type JSONSchemaAdditionalItems struct {
	Bool *bool `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaAdditionalItems) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 1)
	anyOfValid := 0

	err = json.Unmarshal(data, &j.Bool)
	if err != nil {
		anyOfErrors["Bool"] = err
		j.Bool = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for JSONSchemaAdditionalItems failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchemaAdditionalItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.Bool)
}

// JSONSchemaItems structure is generated from "jsonSchema->items".
type JSONSchemaItems struct {
	SliceOfAnything []interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaItems) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 1)
	anyOfValid := 0

	err = json.Unmarshal(data, &j.SliceOfAnything)
	if err != nil {
		anyOfErrors["SliceOfAnything"] = err
		j.SliceOfAnything = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for JSONSchemaItems failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchemaItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.SliceOfAnything)
}

// JSONSchemaAdditionalProperties structure is generated from "jsonSchema->additionalProperties".
type JSONSchemaAdditionalProperties struct {
	Bool *bool `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 1)
	anyOfValid := 0

	err = json.Unmarshal(data, &j.Bool)
	if err != nil {
		anyOfErrors["Bool"] = err
		j.Bool = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for JSONSchemaAdditionalProperties failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchemaAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.Bool)
}

// JSONSchemaDependenciesAdditionalProperties structure is generated from "jsonSchema->dependencies->additionalProperties".
type JSONSchemaDependenciesAdditionalProperties struct {
	SliceOfStringValues []string `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaDependenciesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 1)
	anyOfValid := 0

	err = json.Unmarshal(data, &j.SliceOfStringValues)
	if err != nil {
		anyOfErrors["SliceOfStringValues"] = err
		j.SliceOfStringValues = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for JSONSchemaDependenciesAdditionalProperties failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchemaDependenciesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.SliceOfStringValues)
}

// JSONSchemaType structure is generated from "jsonSchema->type".
type JSONSchemaType struct {
	AnyOf0                                 *JSONSchemaTypeAnyOf0       `json:"-"`
	SliceOfJSONSchemaTypeAnyOf1ItemsValues []JSONSchemaTypeAnyOf1Items `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (j *JSONSchemaType) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &j.AnyOf0)
	if err != nil {
		anyOfErrors["AnyOf0"] = err
		j.AnyOf0 = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &j.SliceOfJSONSchemaTypeAnyOf1ItemsValues)
	if err != nil {
		anyOfErrors["SliceOfJSONSchemaTypeAnyOf1ItemsValues"] = err
		j.SliceOfJSONSchemaTypeAnyOf1ItemsValues = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for JSONSchemaType failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (j JSONSchemaType) MarshalJSON() ([]byte, error) {
	return marshalUnion(j.AnyOf0, j.SliceOfJSONSchemaTypeAnyOf1ItemsValues)
}

// SchemaAllOf1 structure is generated from "http://asyncapi.com/definitions/2.4.0/schema.json/allOf/1".
type SchemaAllOf1 struct {
	AdditionalProperties *SchemaAllOf1AdditionalProperties `json:"additionalProperties,omitempty"`
	Items                *SchemaAllOf1Items                `json:"items,omitempty"`
	AllOf                []Schema                          `json:"allOf,omitempty"`
	OneOf                []Schema                          `json:"oneOf,omitempty"`
	AnyOf                []Schema                          `json:"anyOf,omitempty"`
	Not                  *Schema                           `json:"not,omitempty"`
	Properties           map[string]Schema                 `json:"properties,omitempty"`
	PatternProperties    map[string]Schema                 `json:"patternProperties,omitempty"`
	PropertyNames        *Schema                           `json:"propertyNames,omitempty"`
	Contains             *Schema                           `json:"contains,omitempty"`
	Discriminator        string                            `json:"discriminator,omitempty"`
	ExternalDocs         *ExternalDocs                     `json:"externalDocs,omitempty"`         // Information about external documentation.
	Deprecated           bool                              `json:"deprecated,omitempty"`
	MapOfAnything        map[string]interface{}            `json:"-"`                              // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
	ExtraProperties      map[string]interface{}            `json:"-"`                              // All unmatched properties.
}

type marshalSchemaAllOf1 SchemaAllOf1

var knownKeysSchemaAllOf1 = []string{
	"additionalProperties",
	"items",
	"allOf",
	"oneOf",
	"anyOf",
	"not",
	"properties",
	"patternProperties",
	"propertyNames",
	"contains",
	"discriminator",
	"externalDocs",
	"deprecated",
}

// UnmarshalJSON decodes JSON.
func (s *SchemaAllOf1) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSchemaAllOf1(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysSchemaAllOf1 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	*s = SchemaAllOf1(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaAllOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchemaAllOf1(s), s.MapOfAnything, s.ExtraProperties)
}

// Schema structure is generated from "http://asyncapi.com/definitions/2.4.0/schema.json".
type Schema struct {
	AllOf1 *SchemaAllOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *Schema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &s.AllOf1)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.AllOf1)
}

// SchemaAllOf1AdditionalProperties structure is generated from "http://asyncapi.com/definitions/2.4.0/schema.json/allOf/1->additionalProperties".
type SchemaAllOf1AdditionalProperties struct {
	Schema *Schema `json:"-"`
	Bool   *bool   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *SchemaAllOf1AdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &s.Schema)
	if err != nil {
		anyOfErrors["Schema"] = err
		s.Schema = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &s.Bool)
	if err != nil {
		anyOfErrors["Bool"] = err
		s.Bool = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for SchemaAllOf1AdditionalProperties failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaAllOf1AdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.Schema, s.Bool)
}

// SchemaAllOf1Items structure is generated from "http://asyncapi.com/definitions/2.4.0/schema.json/allOf/1->items".
type SchemaAllOf1Items struct {
	Schema              *Schema  `json:"-"`
	SliceOfSchemaValues []Schema `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *SchemaAllOf1Items) UnmarshalJSON(data []byte) error {
	var err error

	anyOfErrors := make(map[string]error, 2)
	anyOfValid := 0

	err = json.Unmarshal(data, &s.Schema)
	if err != nil {
		anyOfErrors["Schema"] = err
		s.Schema = nil
	} else {
		anyOfValid++
	}

	err = json.Unmarshal(data, &s.SliceOfSchemaValues)
	if err != nil {
		anyOfErrors["SliceOfSchemaValues"] = err
		s.SliceOfSchemaValues = nil
	} else {
		anyOfValid++
	}

	if anyOfValid == 0 {
		return fmt.Errorf("anyOf constraint for SchemaAllOf1Items failed with %d valid results: %v", anyOfValid, anyOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaAllOf1Items) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.Schema, s.SliceOfSchemaValues)
}

// MessageOneOf1OneOf1HeadersAllOf1 structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1->headers/allOf/1".
type MessageOneOf1OneOf1HeadersAllOf1 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1HeadersAllOf1) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"object"` {
		return fmt.Errorf(`bad const value for "type" ("object" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for key, rawValue := range rawMap {
		if m.AdditionalProperties == nil {
			m.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		m.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constMessageOneOf1OneOf1HeadersAllOf1 is unconditionally added to JSON.
	constMessageOneOf1OneOf1HeadersAllOf1 = json.RawMessage(`{"type":"object"}`)
)

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1HeadersAllOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constMessageOneOf1OneOf1HeadersAllOf1, m.AdditionalProperties)
}

// MessageOneOf1OneOf1Headers structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1->headers".
type MessageOneOf1OneOf1Headers struct {
	Schema *Schema                           `json:"-"`
	AllOf1 *MessageOneOf1OneOf1HeadersAllOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1Headers) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Schema)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &m.AllOf1)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1Headers) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Schema, m.AllOf1)
}

// CorrelationID structure is generated from "http://asyncapi.com/definitions/2.4.0/correlationId.json".
type CorrelationID struct {
	Description   string                 `json:"description,omitempty"` // A optional description of the correlation ID. GitHub Flavored Markdown is allowed.
	// A runtime expression that specifies the location of the correlation ID.
	// Value must match pattern: `^\$message\.(header|payload)#(\/(([^\/~])|(~[01]))*)*`.
	// Required.
	Location      string                 `json:"location"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalCorrelationID CorrelationID

var knownKeysCorrelationID = []string{
	"description",
	"location",
}

var requireKeysCorrelationID = []string{
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

	for _, key := range requireKeysCorrelationID {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysCorrelationID {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in CorrelationID: %v", offendingKeys)
	}

	*c = CorrelationID(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalCorrelationID(c), c.MapOfAnything)
}

// MessageOneOf1OneOf1CorrelationID structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1->correlationId".
type MessageOneOf1OneOf1CorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1CorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		m.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.CorrelationID)
	if err != nil {
		oneOfErrors["CorrelationID"] = err
		m.CorrelationID = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for MessageOneOf1OneOf1CorrelationID with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1CorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.CorrelationID)
}

// MessageOneOf1OneOf1ExamplesItems structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1->examples->items".
type MessageOneOf1OneOf1ExamplesItems struct {
	Name    string                 `json:"name,omitempty"`    // Machine readable name of the message example.
	Summary string                 `json:"summary,omitempty"` // A brief summary of the message example.
	Headers map[string]interface{} `json:"headers,omitempty"`
	Payload *interface{}           `json:"payload,omitempty"`
}

type marshalMessageOneOf1OneOf1ExamplesItems MessageOneOf1OneOf1ExamplesItems

var knownKeysMessageOneOf1OneOf1ExamplesItems = []string{
	"name",
	"summary",
	"headers",
	"payload",
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1ExamplesItems) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessageOneOf1OneOf1ExamplesItems(*m)

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

	for _, key := range knownKeysMessageOneOf1OneOf1ExamplesItems {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in MessageOneOf1OneOf1ExamplesItems: %v", offendingKeys)
	}

	*m = MessageOneOf1OneOf1ExamplesItems(mm)

	return nil
}


// MessageTrait structure is generated from "http://asyncapi.com/definitions/2.4.0/messageTrait.json".
type MessageTrait struct {
	SchemaFormat  string                     `json:"schemaFormat,omitempty"`
	ContentType   string                     `json:"contentType,omitempty"`
	Headers       *MessageTraitHeaders       `json:"headers,omitempty"`
	MessageID     string                     `json:"messageId,omitempty"`
	CorrelationID *MessageTraitCorrelationID `json:"correlationId,omitempty"`
	Tags          []Tag                      `json:"tags,omitempty"`
	Summary       string                     `json:"summary,omitempty"`       // A brief summary of the message.
	Name          string                     `json:"name,omitempty"`          // Name of the message.
	Title         string                     `json:"title,omitempty"`         // A human-friendly title for the message.
	Description   string                     `json:"description,omitempty"`   // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs              `json:"externalDocs,omitempty"`  // Information about external documentation.
	Deprecated    bool                       `json:"deprecated,omitempty"`
	Examples      []map[string]interface{}   `json:"examples,omitempty"`
	Bindings      *BindingsObject            `json:"bindings,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                       // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalMessageTrait MessageTrait

var knownKeysMessageTrait = []string{
	"schemaFormat",
	"contentType",
	"headers",
	"messageId",
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

	for _, key := range knownKeysMessageTrait {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in MessageTrait: %v", offendingKeys)
	}

	*m = MessageTrait(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTrait) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessageTrait(m), m.MapOfAnything)
}

// MessageTraitHeadersAllOf1 structure is generated from "http://asyncapi.com/definitions/2.4.0/messageTrait.json->headers/allOf/1".
type MessageTraitHeadersAllOf1 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (m *MessageTraitHeadersAllOf1) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"object"` {
		return fmt.Errorf(`bad const value for "type" ("object" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for key, rawValue := range rawMap {
		if m.AdditionalProperties == nil {
			m.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		m.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constMessageTraitHeadersAllOf1 is unconditionally added to JSON.
	constMessageTraitHeadersAllOf1 = json.RawMessage(`{"type":"object"}`)
)

// MarshalJSON encodes JSON.
func (m MessageTraitHeadersAllOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constMessageTraitHeadersAllOf1, m.AdditionalProperties)
}

// MessageTraitHeaders structure is generated from "http://asyncapi.com/definitions/2.4.0/messageTrait.json->headers".
type MessageTraitHeaders struct {
	Schema *Schema                    `json:"-"`
	AllOf1 *MessageTraitHeadersAllOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageTraitHeaders) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &m.Schema)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &m.AllOf1)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTraitHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Schema, m.AllOf1)
}

// MessageTraitCorrelationID structure is generated from "http://asyncapi.com/definitions/2.4.0/messageTrait.json->correlationId".
type MessageTraitCorrelationID struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageTraitCorrelationID) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		m.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.CorrelationID)
	if err != nil {
		oneOfErrors["CorrelationID"] = err
		m.CorrelationID = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for MessageTraitCorrelationID with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageTraitCorrelationID) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.CorrelationID)
}

// MessageOneOf1OneOf1TraitsItems structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1/oneOf/1->traits->items".
type MessageOneOf1OneOf1TraitsItems struct {
	Reference       *Reference    `json:"-"`
	MessageTrait    *MessageTrait `json:"-"`
	SliceOfAnything []interface{} `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1OneOf1TraitsItems) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 3)
	oneOfValid := 0

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		m.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.MessageTrait)
	if err != nil {
		oneOfErrors["MessageTrait"] = err
		m.MessageTrait = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.SliceOfAnything)
	if err != nil {
		oneOfErrors["SliceOfAnything"] = err
		m.SliceOfAnything = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for MessageOneOf1OneOf1TraitsItems with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1OneOf1TraitsItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.MessageTrait, m.SliceOfAnything)
}

// MessageOneOf1 structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json/oneOf/1".
type MessageOneOf1 struct {
	OneOf0 *MessageOneOf1OneOf0 `json:"-"`
	OneOf1 *MessageOneOf1OneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *MessageOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &m.OneOf0)
	if err != nil {
		oneOfErrors["OneOf0"] = err
		m.OneOf0 = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.OneOf1)
	if err != nil {
		oneOfErrors["OneOf1"] = err
		m.OneOf1 = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for MessageOneOf1 with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m MessageOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.OneOf0, m.OneOf1)
}

// Message structure is generated from "http://asyncapi.com/definitions/2.4.0/message.json".
type Message struct {
	Reference *Reference     `json:"-"`
	OneOf1    *MessageOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (m *Message) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &m.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		m.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &m.OneOf1)
	if err != nil {
		oneOfErrors["OneOf1"] = err
		m.OneOf1 = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for Message with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (m Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(m.Reference, m.OneOf1)
}

// Components structure is generated from "http://asyncapi.com/definitions/2.4.0/components.json".
//
// An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
type Components struct {
	Schemas           map[string]Schema                      `json:"schemas,omitempty"`           // JSON objects describing schemas the API uses.
	Servers           map[string]ServersAdditionalProperties `json:"servers,omitempty"`           // An object representing multiple servers.
	Channels          map[string]ChannelItem                 `json:"channels,omitempty"`
	ServerVariables   map[string]ServerVariable              `json:"serverVariables,omitempty"`
	Messages          map[string]Message                     `json:"messages,omitempty"`          // JSON objects describing the messages being consumed and produced by the API.
	SecuritySchemes   *ComponentsSecuritySchemes             `json:"securitySchemes,omitempty"`
	Parameters        map[string]Parameter                   `json:"parameters,omitempty"`        // JSON objects describing re-usable channel parameters.
	CorrelationIds    *ComponentsCorrelationIds              `json:"correlationIds,omitempty"`
	OperationTraits   map[string]OperationTrait              `json:"operationTraits,omitempty"`
	MessageTraits     map[string]MessageTrait                `json:"messageTraits,omitempty"`
	ServerBindings    map[string]BindingsObject              `json:"serverBindings,omitempty"`
	ChannelBindings   map[string]BindingsObject              `json:"channelBindings,omitempty"`
	OperationBindings map[string]BindingsObject              `json:"operationBindings,omitempty"`
	MessageBindings   map[string]BindingsObject              `json:"messageBindings,omitempty"`
	MapOfAnything     map[string]interface{}                 `json:"-"`                           // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalComponents Components

var knownKeysComponents = []string{
	"schemas",
	"servers",
	"channels",
	"serverVariables",
	"messages",
	"securitySchemes",
	"parameters",
	"correlationIds",
	"operationTraits",
	"messageTraits",
	"serverBindings",
	"channelBindings",
	"operationBindings",
	"messageBindings",
}

// UnmarshalJSON decodes JSON.
func (c *Components) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalComponents(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysComponents {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Components: %v", offendingKeys)
	}

	*c = Components(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c Components) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponents(c), c.MapOfAnything)
}

// UserPassword structure is generated from "http://asyncapi.com/definitions/2.4.0/userPassword.json".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalUserPassword UserPassword

var knownKeysUserPassword = []string{
	"description",
	"type",
}

var requireKeysUserPassword = []string{
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

	for _, key := range requireKeysUserPassword {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"userPassword"` {
		return fmt.Errorf(`bad const value for "type" ("userPassword" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysUserPassword {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in UserPassword: %v", offendingKeys)
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

// APIKey structure is generated from "http://asyncapi.com/definitions/2.4.0/apiKey.json".
type APIKey struct {
	In            APIKeyIn               `json:"in"`                    // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalAPIKey APIKey

var knownKeysAPIKey = []string{
	"in",
	"description",
	"type",
}

var requireKeysAPIKey = []string{
	"type",
	"in",
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

	for _, key := range requireKeysAPIKey {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"apiKey"` {
		return fmt.Errorf(`bad const value for "type" ("apiKey" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysAPIKey {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in APIKey: %v", offendingKeys)
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

// X509 structure is generated from "http://asyncapi.com/definitions/2.4.0/X509.json".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalX509 X509

var knownKeysX509 = []string{
	"description",
	"type",
}

var requireKeysX509 = []string{
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

	for _, key := range requireKeysX509 {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"X509"` {
		return fmt.Errorf(`bad const value for "type" ("X509" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysX509 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in X509: %v", offendingKeys)
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

// SymmetricEncryption structure is generated from "http://asyncapi.com/definitions/2.4.0/symmetricEncryption.json".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalSymmetricEncryption SymmetricEncryption

var knownKeysSymmetricEncryption = []string{
	"description",
	"type",
}

var requireKeysSymmetricEncryption = []string{
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

	for _, key := range requireKeysSymmetricEncryption {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"symmetricEncryption"` {
		return fmt.Errorf(`bad const value for "type" ("symmetricEncryption" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysSymmetricEncryption {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SymmetricEncryption: %v", offendingKeys)
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

// AsymmetricEncryption structure is generated from "http://asyncapi.com/definitions/2.4.0/asymmetricEncryption.json".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalAsymmetricEncryption AsymmetricEncryption

var knownKeysAsymmetricEncryption = []string{
	"description",
	"type",
}

var requireKeysAsymmetricEncryption = []string{
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

	for _, key := range requireKeysAsymmetricEncryption {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"asymmetricEncryption"` {
		return fmt.Errorf(`bad const value for "type" ("asymmetricEncryption" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysAsymmetricEncryption {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in AsymmetricEncryption: %v", offendingKeys)
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

// NonBearerHTTPSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/NonBearerHTTPSecurityScheme.json".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme"`                // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

var knownKeysNonBearerHTTPSecurityScheme = []string{
	"scheme",
	"description",
	"type",
}

var requireKeysNonBearerHTTPSecurityScheme = []string{
	"scheme",
	"type",
}

// UnmarshalJSON decodes JSON.
func (n *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	var not NonBearerHTTPSecuritySchemeNot

	if json.Unmarshal(data, &not) == nil {
		return errors.New("not constraint failed for NonBearerHTTPSecurityScheme")
	}

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

	for _, key := range requireKeysNonBearerHTTPSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"http"` {
		return fmt.Errorf(`bad const value for "type" ("http" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysNonBearerHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in NonBearerHTTPSecurityScheme: %v", offendingKeys)
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

// NonBearerHTTPSecuritySchemeNot structure is generated from "http://asyncapi.com/definitions/2.4.0/NonBearerHTTPSecurityScheme.json->not".
type NonBearerHTTPSecuritySchemeNot struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (n *NonBearerHTTPSecuritySchemeNot) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["scheme"]; exists && string(v) != `"bearer"` {
		return fmt.Errorf(`bad const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(rawMap, "scheme")

	for key, rawValue := range rawMap {
		if n.AdditionalProperties == nil {
			n.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		n.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constNonBearerHTTPSecuritySchemeNot is unconditionally added to JSON.
	constNonBearerHTTPSecuritySchemeNot = json.RawMessage(`{"scheme":"bearer"}`)
)

// MarshalJSON encodes JSON.
func (n NonBearerHTTPSecuritySchemeNot) MarshalJSON() ([]byte, error) {
	return marshalUnion(constNonBearerHTTPSecuritySchemeNot, n.AdditionalProperties)
}

// BearerHTTPSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/BearerHTTPSecurityScheme.json".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

var knownKeysBearerHTTPSecurityScheme = []string{
	"bearerFormat",
	"description",
	"scheme",
	"type",
}

var requireKeysBearerHTTPSecurityScheme = []string{
	"type",
	"scheme",
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

	for _, key := range requireKeysBearerHTTPSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["scheme"]; exists && string(v) != `"bearer"` {
		return fmt.Errorf(`bad const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(rawMap, "scheme")

	if v, exists := rawMap["type"]; exists && string(v) != `"http"` {
		return fmt.Errorf(`bad const value for "type" ("http" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysBearerHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in BearerHTTPSecurityScheme: %v", offendingKeys)
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

// APIKeyHTTPSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/APIKeyHTTPSecurityScheme.json".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name"`                  // Required.
	In            APIKeyHTTPSecuritySchemeIn `json:"in"`                    // Required.
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

var knownKeysAPIKeyHTTPSecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
}

var requireKeysAPIKeyHTTPSecurityScheme = []string{
	"type",
	"name",
	"in",
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

	for _, key := range requireKeysAPIKeyHTTPSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"httpApiKey"` {
		return fmt.Errorf(`bad const value for "type" ("httpApiKey" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysAPIKeyHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in APIKeyHTTPSecurityScheme: %v", offendingKeys)
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

// HTTPSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/HTTPSecurityScheme.json".
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (h *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 3)
	oneOfValid := 0

	err = json.Unmarshal(data, &h.NonBearerHTTPSecurityScheme)
	if err != nil {
		oneOfErrors["NonBearerHTTPSecurityScheme"] = err
		h.NonBearerHTTPSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &h.BearerHTTPSecurityScheme)
	if err != nil {
		oneOfErrors["BearerHTTPSecurityScheme"] = err
		h.BearerHTTPSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &h.APIKeyHTTPSecurityScheme)
	if err != nil {
		oneOfErrors["APIKeyHTTPSecurityScheme"] = err
		h.APIKeyHTTPSecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for HTTPSecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (h HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(h.NonBearerHTTPSecurityScheme, h.BearerHTTPSecurityScheme, h.APIKeyHTTPSecurityScheme)
}

// Oauth2Flows structure is generated from "http://asyncapi.com/definitions/2.4.0/oauth2Flows.json".
type Oauth2Flows struct {
	Description          string                 `json:"description,omitempty"`
	Flows                Oauth2FlowsFlows       `json:"flows"`                 // Required.
	MapOfAnything        map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
	AdditionalProperties map[string]interface{} `json:"-"`                     // All unmatched properties.
}

type marshalOauth2Flows Oauth2Flows

var knownKeysOauth2Flows = []string{
	"description",
	"flows",
	"type",
}

var requireKeysOauth2Flows = []string{
	"type",
	"flows",
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

	for _, key := range requireKeysOauth2Flows {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"oauth2"` {
		return fmt.Errorf(`bad const value for "type" ("oauth2" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysOauth2Flows {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

// Oauth2FlowsFlows structure is generated from "http://asyncapi.com/definitions/2.4.0/oauth2Flows.json->flows".
type Oauth2FlowsFlows struct {
	Implicit          *Oauth2Flow `json:"implicit,omitempty"`
	Password          *Oauth2Flow `json:"password,omitempty"`
	ClientCredentials *Oauth2Flow `json:"clientCredentials,omitempty"`
	AuthorizationCode *Oauth2Flow `json:"authorizationCode,omitempty"`
}

type marshalOauth2FlowsFlows Oauth2FlowsFlows

var knownKeysOauth2FlowsFlows = []string{
	"implicit",
	"password",
	"clientCredentials",
	"authorizationCode",
}

// UnmarshalJSON decodes JSON.
func (o *Oauth2FlowsFlows) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOauth2FlowsFlows(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysOauth2FlowsFlows {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Oauth2FlowsFlows: %v", offendingKeys)
	}

	*o = Oauth2FlowsFlows(mo)

	return nil
}


// Oauth2Flow structure is generated from "http://asyncapi.com/definitions/2.4.0/oauth2Flow.json".
type Oauth2Flow struct {
	AuthorizationURL string                 `json:"authorizationUrl,omitempty"` // Format: uri.
	TokenURL         string                 `json:"tokenUrl,omitempty"`         // Format: uri.
	RefreshURL       string                 `json:"refreshUrl,omitempty"`       // Format: uri.
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalOauth2Flow Oauth2Flow

var knownKeysOauth2Flow = []string{
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

	for _, key := range knownKeysOauth2Flow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Oauth2Flow: %v", offendingKeys)
	}

	*o = Oauth2Flow(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o Oauth2Flow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2Flow(o), o.MapOfAnything)
}

// OpenIDConnect structure is generated from "http://asyncapi.com/definitions/2.4.0/openIdConnect.json".
type OpenIDConnect struct {
	Description   string                 `json:"description,omitempty"`
	// Format: uri.
	// Required.
	URL           string                 `json:"openIdConnectUrl"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalOpenIDConnect OpenIDConnect

var knownKeysOpenIDConnect = []string{
	"description",
	"openIdConnectUrl",
	"type",
}

var requireKeysOpenIDConnect = []string{
	"type",
	"openIdConnectUrl",
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

	for _, key := range requireKeysOpenIDConnect {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad const value for "type" ("openIdConnect" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysOpenIDConnect {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in OpenIDConnect: %v", offendingKeys)
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

// SaslPlainSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/SaslPlainSecurityScheme.json".
type SaslPlainSecurityScheme struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalSaslPlainSecurityScheme SaslPlainSecurityScheme

var knownKeysSaslPlainSecurityScheme = []string{
	"description",
	"type",
}

var requireKeysSaslPlainSecurityScheme = []string{
	"type",
}

// UnmarshalJSON decodes JSON.
func (s *SaslPlainSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSaslPlainSecurityScheme(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSaslPlainSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"plain"` {
		return fmt.Errorf(`bad const value for "type" ("plain" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysSaslPlainSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SaslPlainSecurityScheme: %v", offendingKeys)
	}

	*s = SaslPlainSecurityScheme(ms)

	return nil
}

var (
	// constSaslPlainSecurityScheme is unconditionally added to JSON.
	constSaslPlainSecurityScheme = json.RawMessage(`{"type":"plain"}`)
)

// MarshalJSON encodes JSON.
func (s SaslPlainSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constSaslPlainSecurityScheme, marshalSaslPlainSecurityScheme(s), s.MapOfAnything)
}

// SaslScramSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/SaslScramSecurityScheme.json".
type SaslScramSecurityScheme struct {
	Type          SaslScramSecuritySchemeType `json:"type"`                  // Required.
	Description   string                      `json:"description,omitempty"`
	MapOfAnything map[string]interface{}      `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalSaslScramSecurityScheme SaslScramSecurityScheme

var knownKeysSaslScramSecurityScheme = []string{
	"type",
	"description",
}

var requireKeysSaslScramSecurityScheme = []string{
	"type",
}

// UnmarshalJSON decodes JSON.
func (s *SaslScramSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSaslScramSecurityScheme(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSaslScramSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysSaslScramSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SaslScramSecurityScheme: %v", offendingKeys)
	}

	*s = SaslScramSecurityScheme(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s SaslScramSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSaslScramSecurityScheme(s), s.MapOfAnything)
}

// SaslGssapiSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/SaslGssapiSecurityScheme.json".
type SaslGssapiSecurityScheme struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-[\w\d\.\x2d_]+$`.
}

type marshalSaslGssapiSecurityScheme SaslGssapiSecurityScheme

var knownKeysSaslGssapiSecurityScheme = []string{
	"description",
	"type",
}

var requireKeysSaslGssapiSecurityScheme = []string{
	"type",
}

// UnmarshalJSON decodes JSON.
func (s *SaslGssapiSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSaslGssapiSecurityScheme(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSaslGssapiSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"gssapi"` {
		return fmt.Errorf(`bad const value for "type" ("gssapi" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysSaslGssapiSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexXWDX2D.MatchString(key) {
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

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SaslGssapiSecurityScheme: %v", offendingKeys)
	}

	*s = SaslGssapiSecurityScheme(ms)

	return nil
}

var (
	// constSaslGssapiSecurityScheme is unconditionally added to JSON.
	constSaslGssapiSecurityScheme = json.RawMessage(`{"type":"gssapi"}`)
)

// MarshalJSON encodes JSON.
func (s SaslGssapiSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constSaslGssapiSecurityScheme, marshalSaslGssapiSecurityScheme(s), s.MapOfAnything)
}

// SaslSecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/SaslSecurityScheme.json".
type SaslSecurityScheme struct {
	SaslPlainSecurityScheme  *SaslPlainSecurityScheme  `json:"-"`
	SaslScramSecurityScheme  *SaslScramSecurityScheme  `json:"-"`
	SaslGssapiSecurityScheme *SaslGssapiSecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *SaslSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 3)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.SaslPlainSecurityScheme)
	if err != nil {
		oneOfErrors["SaslPlainSecurityScheme"] = err
		s.SaslPlainSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SaslScramSecurityScheme)
	if err != nil {
		oneOfErrors["SaslScramSecurityScheme"] = err
		s.SaslScramSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SaslGssapiSecurityScheme)
	if err != nil {
		oneOfErrors["SaslGssapiSecurityScheme"] = err
		s.SaslGssapiSecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SaslSecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SaslSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.SaslPlainSecurityScheme, s.SaslScramSecurityScheme, s.SaslGssapiSecurityScheme)
}

// SecurityScheme structure is generated from "http://asyncapi.com/definitions/2.4.0/SecurityScheme.json".
type SecurityScheme struct {
	UserPassword         *UserPassword         `json:"-"`
	APIKey               *APIKey               `json:"-"`
	X509                 *X509                 `json:"-"`
	SymmetricEncryption  *SymmetricEncryption  `json:"-"`
	AsymmetricEncryption *AsymmetricEncryption `json:"-"`
	HTTPSecurityScheme   *HTTPSecurityScheme   `json:"-"`
	Oauth2Flows          *Oauth2Flows          `json:"-"`
	OpenIDConnect        *OpenIDConnect        `json:"-"`
	SaslSecurityScheme   *SaslSecurityScheme   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (s *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 9)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.UserPassword)
	if err != nil {
		oneOfErrors["UserPassword"] = err
		s.UserPassword = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.APIKey)
	if err != nil {
		oneOfErrors["APIKey"] = err
		s.APIKey = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.X509)
	if err != nil {
		oneOfErrors["X509"] = err
		s.X509 = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SymmetricEncryption)
	if err != nil {
		oneOfErrors["SymmetricEncryption"] = err
		s.SymmetricEncryption = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.AsymmetricEncryption)
	if err != nil {
		oneOfErrors["AsymmetricEncryption"] = err
		s.AsymmetricEncryption = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.HTTPSecurityScheme)
	if err != nil {
		oneOfErrors["HTTPSecurityScheme"] = err
		s.HTTPSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.Oauth2Flows)
	if err != nil {
		oneOfErrors["Oauth2Flows"] = err
		s.Oauth2Flows = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.OpenIDConnect)
	if err != nil {
		oneOfErrors["OpenIDConnect"] = err
		s.OpenIDConnect = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SaslSecurityScheme)
	if err != nil {
		oneOfErrors["SaslSecurityScheme"] = err
		s.SaslSecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.UserPassword, s.APIKey, s.X509, s.SymmetricEncryption, s.AsymmetricEncryption, s.HTTPSecurityScheme, s.Oauth2Flows, s.OpenIDConnect, s.SaslSecurityScheme)
}

// ComponentsSecuritySchemesWD structure is generated from "http://asyncapi.com/definitions/2.4.0/components.json->securitySchemes->^[\w\d\.\-_]+$".
type ComponentsSecuritySchemesWD struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSecuritySchemesWD) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &c.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		c.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &c.SecurityScheme)
	if err != nil {
		oneOfErrors["SecurityScheme"] = err
		c.SecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ComponentsSecuritySchemesWD with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSecuritySchemesWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.Reference, c.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "http://asyncapi.com/definitions/2.4.0/components.json->securitySchemes".
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

// ComponentsCorrelationIdsWD structure is generated from "http://asyncapi.com/definitions/2.4.0/components.json->correlationIds->^[\w\d\.\-_]+$".
type ComponentsCorrelationIdsWD struct {
	Reference     *Reference     `json:"-"`
	CorrelationID *CorrelationID `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsCorrelationIdsWD) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &c.Reference)
	if err != nil {
		oneOfErrors["Reference"] = err
		c.Reference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &c.CorrelationID)
	if err != nil {
		oneOfErrors["CorrelationID"] = err
		c.CorrelationID = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ComponentsCorrelationIdsWD with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsCorrelationIdsWD) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.Reference, c.CorrelationID)
}

// ComponentsCorrelationIds structure is generated from "http://asyncapi.com/definitions/2.4.0/components.json->correlationIds".
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

// JSONSchemaTypeAnyOf0 is an enum type.
type JSONSchemaTypeAnyOf0 string

// JSONSchemaTypeAnyOf0 values enumeration.
const (
	JSONSchemaTypeAnyOf0Array = JSONSchemaTypeAnyOf0("array")
	JSONSchemaTypeAnyOf0Boolean = JSONSchemaTypeAnyOf0("boolean")
	JSONSchemaTypeAnyOf0Integer = JSONSchemaTypeAnyOf0("integer")
	JSONSchemaTypeAnyOf0Null = JSONSchemaTypeAnyOf0("null")
	JSONSchemaTypeAnyOf0Number = JSONSchemaTypeAnyOf0("number")
	JSONSchemaTypeAnyOf0Object = JSONSchemaTypeAnyOf0("object")
	JSONSchemaTypeAnyOf0String = JSONSchemaTypeAnyOf0("string")
)

// MarshalJSON encodes JSON.
func (i JSONSchemaTypeAnyOf0) MarshalJSON() ([]byte, error) {
	switch i {
	case JSONSchemaTypeAnyOf0Array:
	case JSONSchemaTypeAnyOf0Boolean:
	case JSONSchemaTypeAnyOf0Integer:
	case JSONSchemaTypeAnyOf0Null:
	case JSONSchemaTypeAnyOf0Number:
	case JSONSchemaTypeAnyOf0Object:
	case JSONSchemaTypeAnyOf0String:

	default:
		return nil, fmt.Errorf("unexpected JSONSchemaTypeAnyOf0 value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *JSONSchemaTypeAnyOf0) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := JSONSchemaTypeAnyOf0(ii)

	switch v {
	case JSONSchemaTypeAnyOf0Array:
	case JSONSchemaTypeAnyOf0Boolean:
	case JSONSchemaTypeAnyOf0Integer:
	case JSONSchemaTypeAnyOf0Null:
	case JSONSchemaTypeAnyOf0Number:
	case JSONSchemaTypeAnyOf0Object:
	case JSONSchemaTypeAnyOf0String:

	default:
		return fmt.Errorf("unexpected JSONSchemaTypeAnyOf0 value: %v", v)
	}

	*i = v

	return nil
}

// JSONSchemaTypeAnyOf1Items is an enum type.
type JSONSchemaTypeAnyOf1Items string

// JSONSchemaTypeAnyOf1Items values enumeration.
const (
	JSONSchemaTypeAnyOf1ItemsArray = JSONSchemaTypeAnyOf1Items("array")
	JSONSchemaTypeAnyOf1ItemsBoolean = JSONSchemaTypeAnyOf1Items("boolean")
	JSONSchemaTypeAnyOf1ItemsInteger = JSONSchemaTypeAnyOf1Items("integer")
	JSONSchemaTypeAnyOf1ItemsNull = JSONSchemaTypeAnyOf1Items("null")
	JSONSchemaTypeAnyOf1ItemsNumber = JSONSchemaTypeAnyOf1Items("number")
	JSONSchemaTypeAnyOf1ItemsObject = JSONSchemaTypeAnyOf1Items("object")
	JSONSchemaTypeAnyOf1ItemsString = JSONSchemaTypeAnyOf1Items("string")
)

// MarshalJSON encodes JSON.
func (i JSONSchemaTypeAnyOf1Items) MarshalJSON() ([]byte, error) {
	switch i {
	case JSONSchemaTypeAnyOf1ItemsArray:
	case JSONSchemaTypeAnyOf1ItemsBoolean:
	case JSONSchemaTypeAnyOf1ItemsInteger:
	case JSONSchemaTypeAnyOf1ItemsNull:
	case JSONSchemaTypeAnyOf1ItemsNumber:
	case JSONSchemaTypeAnyOf1ItemsObject:
	case JSONSchemaTypeAnyOf1ItemsString:

	default:
		return nil, fmt.Errorf("unexpected JSONSchemaTypeAnyOf1Items value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *JSONSchemaTypeAnyOf1Items) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := JSONSchemaTypeAnyOf1Items(ii)

	switch v {
	case JSONSchemaTypeAnyOf1ItemsArray:
	case JSONSchemaTypeAnyOf1ItemsBoolean:
	case JSONSchemaTypeAnyOf1ItemsInteger:
	case JSONSchemaTypeAnyOf1ItemsNull:
	case JSONSchemaTypeAnyOf1ItemsNumber:
	case JSONSchemaTypeAnyOf1ItemsObject:
	case JSONSchemaTypeAnyOf1ItemsString:

	default:
		return fmt.Errorf("unexpected JSONSchemaTypeAnyOf1Items value: %v", v)
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

// SaslScramSecuritySchemeType is an enum type.
type SaslScramSecuritySchemeType string

// SaslScramSecuritySchemeType values enumeration.
const (
	SaslScramSecuritySchemeTypeScramSha256 = SaslScramSecuritySchemeType("scramSha256")
	SaslScramSecuritySchemeTypeScramSha512 = SaslScramSecuritySchemeType("scramSha512")
)

// MarshalJSON encodes JSON.
func (i SaslScramSecuritySchemeType) MarshalJSON() ([]byte, error) {
	switch i {
	case SaslScramSecuritySchemeTypeScramSha256:
	case SaslScramSecuritySchemeTypeScramSha512:

	default:
		return nil, fmt.Errorf("unexpected SaslScramSecuritySchemeType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SaslScramSecuritySchemeType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SaslScramSecuritySchemeType(ii)

	switch v {
	case SaslScramSecuritySchemeTypeScramSha256:
	case SaslScramSecuritySchemeTypeScramSha512:

	default:
		return fmt.Errorf("unexpected SaslScramSecuritySchemeType value: %v", v)
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
// Regular expressions for pattern properties.
var (
	regexXWDX2D = regexp.MustCompile(`^x-[\w\d\.\x2d_]+$`)
	regexWD = regexp.MustCompile(`^[\w\d\.\-_]+$`)
)
