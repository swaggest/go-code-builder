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
	// The AsyncAPI specification version of this document.
	// Required.
	Asyncapi      AsyncAPIAsyncapi       `json:"asyncapi"`
	// General information about the API.
	// Required.
	Info          Info                   `json:"info"`
	// The base topic to the API. Example: 'hitch'.
	// Value must match pattern: `^[^/.]`.
	BaseTopic     string                 `json:"baseTopic,omitempty"`
	Servers       []Server               `json:"servers,omitempty"`
	Topics        *Topics                `json:"topics,omitempty"`       // Relative paths to the individual topics. They must be relative to the 'baseTopic'.
	Stream        *Stream                `json:"stream,omitempty"`       // Stream Object.
	Events        *Events                `json:"events,omitempty"`       // Events Object.
	Components    *Components            `json:"components,omitempty"`   // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags          []Tag                  `json:"tags,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
}

type marshalAsyncAPI AsyncAPI

var knownKeysAsyncAPI = []string{
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

var requireKeysAsyncAPI = []string{
	"asyncapi",
	"info",
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

	for _, key := range knownKeysAsyncAPI {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

// MarshalJSON encodes JSON.
func (a AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAsyncAPI(a), a.MapOfAnything)
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
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// License structure is generated from "#/definitions/license".
type License struct {
	// The name of the license type. It's encouraged to use an OSI compatible license.
	// Required.
	Name          string                 `json:"name"`
	// The URL pointing to the license.
	// Format: uri.
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`             // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// Server structure is generated from "#/definitions/server".
//
// An object representing a Server.
type Server struct {
	URL           string                    `json:"url"`                     // Required.
	Description   string                    `json:"description,omitempty"`
	// The transfer protocol.
	// Required.
	Scheme        ServerScheme              `json:"scheme"`
	SchemeVersion string                    `json:"schemeVersion,omitempty"`
	Variables     map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnything map[string]interface{}    `json:"-"`                       // Key must match pattern: `^x-`.
}

type marshalServer Server

var knownKeysServer = []string{
	"url",
	"description",
	"scheme",
	"schemeVersion",
	"variables",
}

var requireKeysServer = []string{
	"url",
	"scheme",
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

		if regexX.MatchString(key) {
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

// ServerVariable structure is generated from "#/definitions/serverVariable".
//
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

type marshalServerVariable ServerVariable

var knownKeysServerVariable = []string{
	"enum",
	"default",
	"description",
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

		if regexX.MatchString(key) {
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

// Topics structure is generated from "#/definitions/topics".
//
// Relative paths to the individual topics. They must be relative to the 'baseTopic'.
type Topics struct {
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
	MapOfTopicItemValues map[string]TopicItem   `json:"-"` // Key must match pattern: `^[^.]`.
}

// UnmarshalJSON decodes JSON.
func (t *Topics) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if t.MapOfAnything == nil {
				t.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			t.MapOfAnything[key] = val
		}

		if regex.MatchString(key) {
			matched = true

			if t.MapOfTopicItemValues == nil {
				t.MapOfTopicItemValues = make(map[string]TopicItem, 1)
			}

			var val TopicItem

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			t.MapOfTopicItemValues[key] = val
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

		return fmt.Errorf("additional properties not allowed in Topics: %v", offendingKeys)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (t Topics) MarshalJSON() ([]byte, error) {
	return marshalUnion(t.MapOfAnything, t.MapOfTopicItemValues)
}

// TopicItem structure is generated from "#/definitions/topicItem".
type TopicItem struct {
	Ref           string                 `json:"$ref,omitempty"`
	Parameters    []Parameter            `json:"parameters,omitempty"`
	Publish       *Operation             `json:"publish,omitempty"`
	Subscribe     *Operation             `json:"subscribe,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

type marshalTopicItem TopicItem

var knownKeysTopicItem = []string{
	"$ref",
	"parameters",
	"publish",
	"subscribe",
	"deprecated",
}

// UnmarshalJSON decodes JSON.
func (t *TopicItem) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTopicItem(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysTopicItem {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

		return fmt.Errorf("additional properties not allowed in TopicItem: %v", offendingKeys)
	}

	*t = TopicItem(mt)

	return nil
}

// MarshalJSON encodes JSON.
func (t TopicItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTopicItem(t), t.MapOfAnything)
}

// Parameter structure is generated from "#/definitions/parameter".
type Parameter struct {
	Description   string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name          string                 `json:"name,omitempty"`        // The name of the parameter.
	Schema        map[string]interface{} `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	Ref           string                 `json:"$ref,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

type marshalParameter Parameter

var knownKeysParameter = []string{
	"description",
	"name",
	"schema",
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

		if regexX.MatchString(key) {
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

// Message structure is generated from "#/definitions/message".
type Message struct {
	Ref           string                 `json:"$ref,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty"`      // A deterministic version of a JSON Schema object.
	Payload       map[string]interface{} `json:"payload,omitempty"`      // A deterministic version of a JSON Schema object.
	Tags          []Tag                  `json:"tags,omitempty"`
	Summary       string                 `json:"summary,omitempty"`      // A brief summary of the message.
	Description   string                 `json:"description,omitempty"`  // A longer description of the message. CommonMark is allowed.
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	Deprecated    bool                   `json:"deprecated,omitempty"`
	Example       *interface{}           `json:"example,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
}

type marshalMessage Message

var knownKeysMessage = []string{
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
func (m *Message) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMessage(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mm.Example == nil {
		if _, ok := rawMap["example"]; ok {
			var v interface{}
			mm.Example = &v
		}
	}

	for _, key := range knownKeysMessage {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

		return fmt.Errorf("additional properties not allowed in Message: %v", offendingKeys)
	}

	*m = Message(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessage(m), m.MapOfAnything)
}

// Tag structure is generated from "#/definitions/tag".
type Tag struct {
	Name          string                 `json:"name"`                   // Required.
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs          `json:"externalDocs,omitempty"` // Information about external documentation.
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// ExternalDocs structure is generated from "#/definitions/externalDocs".
//
// information about external documentation.
type ExternalDocs struct {
	Description   string                 `json:"description,omitempty"`
	// Format: uri.
	// Required.
	URL           string                 `json:"url"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// OperationOneOf1 structure is generated from "#/definitions/operation/oneOf/1".
type OperationOneOf1 struct {
	OneOf         []Message              `json:"oneOf,omitempty"` // Required.
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-`.
}

type marshalOperationOneOf1 OperationOneOf1

var knownKeysOperationOneOf1 = []string{
	"oneOf",
}

var requireKeysOperationOneOf1 = []string{
	"oneOf",
}

// UnmarshalJSON decodes JSON.
func (o *OperationOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOperationOneOf1(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysOperationOneOf1 {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysOperationOneOf1 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

		return fmt.Errorf("additional properties not allowed in OperationOneOf1: %v", offendingKeys)
	}

	*o = OperationOneOf1(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o OperationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationOneOf1(o), o.MapOfAnything)
}

// Operation structure is generated from "#/definitions/operation".
type Operation struct {
	Message *Message         `json:"-"`
	OneOf1  *OperationOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &o.Message)
	if err != nil {
		oneOfErrors["Message"] = err
		o.Message = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &o.OneOf1)
	if err != nil {
		oneOfErrors["OneOf1"] = err
		o.OneOf1 = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for Operation with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (o Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(o.Message, o.OneOf1)
}

// Stream structure is generated from "#/definitions/stream".
//
// Stream Object.
type Stream struct {
	Framing       *StreamFraming         `json:"framing,omitempty"` // Stream Framing Object.
	Read          []Message              `json:"read,omitempty"`    // Stream Read Object.
	Write         []Message              `json:"write,omitempty"`   // Stream Write Object.
	MapOfAnything map[string]interface{} `json:"-"`                 // Key must match pattern: `^x-`.
}

type marshalStream Stream

var knownKeysStream = []string{
	"framing",
	"read",
	"write",
}

// UnmarshalJSON decodes JSON.
func (s *Stream) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalStream(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysStream {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

		return fmt.Errorf("additional properties not allowed in Stream: %v", offendingKeys)
	}

	*s = Stream(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Stream) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStream(s), s.MapOfAnything)
}

// StreamFramingOneOf0 structure is generated from "#/definitions/stream->framing/oneOf/0".
type StreamFramingOneOf0 struct {
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

type marshalStreamFramingOneOf0 StreamFramingOneOf0

var knownKeysStreamFramingOneOf0 = []string{
	"delimiter",
	"type",
}

// UnmarshalJSON decodes JSON.
func (s *StreamFramingOneOf0) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalStreamFramingOneOf0(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"chunked"` {
		return fmt.Errorf(`bad const value for "type" ("chunked" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysStreamFramingOneOf0 {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in StreamFramingOneOf0: %v", offendingKeys)
	}

	*s = StreamFramingOneOf0(ms)

	return nil
}

var (
	// constStreamFramingOneOf0 is unconditionally added to JSON.
	constStreamFramingOneOf0 = json.RawMessage(`{"type":"chunked"}`)
)

// MarshalJSON encodes JSON.
func (s StreamFramingOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constStreamFramingOneOf0, marshalStreamFramingOneOf0(s))
}

// StreamFramingOneOf1 structure is generated from "#/definitions/stream->framing/oneOf/1".
type StreamFramingOneOf1 struct {

}

// UnmarshalJSON decodes JSON.
func (s *StreamFramingOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"sse"` {
		return fmt.Errorf(`bad const value for "type" ("sse" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	if v, exists := rawMap["delimiter"]; exists && string(v) != `"\\n\\n"` {
		return fmt.Errorf(`bad const value for "delimiter" ("\\n\\n" expected, %s received)`, v)
	}

	delete(rawMap, "delimiter")

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in StreamFramingOneOf1: %v", offendingKeys)
	}

	return nil
}

var (
	// constStreamFramingOneOf1 is unconditionally added to JSON.
	constStreamFramingOneOf1 = json.RawMessage(`{"type":"sse","delimiter":"\\n\\n"}`)
)

// MarshalJSON encodes JSON.
func (s StreamFramingOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constStreamFramingOneOf1)
}

// StreamFraming structure is generated from "#/definitions/stream->framing".
//
// Stream Framing Object.
type StreamFraming struct {
	OneOf0               *StreamFramingOneOf0   `json:"-"`
	OneOf1               *StreamFramingOneOf1   `json:"-"`
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// UnmarshalJSON decodes JSON.
func (s *StreamFraming) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.OneOf0)
	if err != nil {
		oneOfErrors["OneOf0"] = err
		s.OneOf0 = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.OneOf1)
	if err != nil {
		oneOfErrors["OneOf1"] = err
		s.OneOf1 = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for StreamFraming with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if s.MapOfAnything == nil {
				s.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			s.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if s.AdditionalProperties == nil {
			s.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		s.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s StreamFraming) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.MapOfAnything, s.AdditionalProperties, s.OneOf0, s.OneOf1)
}

// Events structure is generated from "#/definitions/events".
//
// Events Object.
type Events struct {
	Receive       []Message              `json:"receive,omitempty"` // Events Receive Object.
	Send          []Message              `json:"send,omitempty"`    // Events Send Object.
	MapOfAnything map[string]interface{} `json:"-"`                 // Key must match pattern: `^x-`.
}

type marshalEvents Events

var knownKeysEvents = []string{
	"receive",
	"send",
}

// UnmarshalJSON decodes JSON.
func (e *Events) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalEvents(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysEvents {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
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

		return fmt.Errorf("additional properties not allowed in Events: %v", offendingKeys)
	}

	*e = Events(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e Events) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalEvents(e), e.MapOfAnything)
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

type marshalComponents Components

var knownKeysComponents = []string{
	"schemas",
	"messages",
	"securitySchemes",
	"parameters",
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


// Reference structure is generated from "#/definitions/Reference".
type Reference struct {
	// Format: uri.
	// Required.
	Ref                  string                 `json:"$ref,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties.
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

// UserPassword structure is generated from "#/definitions/userPassword".
type UserPassword struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// APIKey structure is generated from "#/definitions/apiKey".
type APIKey struct {
	In            APIKeyIn               `json:"in"`                    // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// X509 structure is generated from "#/definitions/X509".
type X509 struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// SymmetricEncryption structure is generated from "#/definitions/symmetricEncryption".
type SymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// AsymmetricEncryption structure is generated from "#/definitions/asymmetricEncryption".
type AsymmetricEncryption struct {
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// NonBearerHTTPSecurityScheme structure is generated from "#/definitions/NonBearerHTTPSecurityScheme".
type NonBearerHTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme"`                // Required.
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// NonBearerHTTPSecuritySchemeNot structure is generated from "#/definitions/NonBearerHTTPSecurityScheme->not".
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

// BearerHTTPSecurityScheme structure is generated from "#/definitions/BearerHTTPSecurityScheme".
type BearerHTTPSecurityScheme struct {
	BearerFormat  string                 `json:"bearerFormat,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// APIKeyHTTPSecurityScheme structure is generated from "#/definitions/APIKeyHTTPSecurityScheme".
type APIKeyHTTPSecurityScheme struct {
	Name          string                     `json:"name"`                  // Required.
	In            APIKeyHTTPSecuritySchemeIn `json:"in"`                    // Required.
	Description   string                     `json:"description,omitempty"`
	MapOfAnything map[string]interface{}     `json:"-"`                     // Key must match pattern: `^x-`.
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

		if regexX.MatchString(key) {
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

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
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
func (s *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 6)
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

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.UserPassword, s.APIKey, s.X509, s.SymmetricEncryption, s.AsymmetricEncryption, s.HTTPSecurityScheme)
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from "#/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("oneOf constraint failed for ComponentsSecuritySchemesAZAZ09 with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSecuritySchemesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.Reference, c.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties                       map[string]interface{}                     `json:"-"` // All unmatched properties.
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

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfComponentsSecuritySchemesAZAZ09Values == nil {
				c.MapOfComponentsSecuritySchemesAZAZ09Values = make(map[string]ComponentsSecuritySchemesAZAZ09, 1)
			}

			var val ComponentsSecuritySchemesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfComponentsSecuritySchemesAZAZ09Values[key] = val
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
	return marshalUnion(c.MapOfComponentsSecuritySchemesAZAZ09Values, c.AdditionalProperties)
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
	regexX = regexp.MustCompile("^x-")
	regex = regexp.MustCompile("^[^.]")
	regexAZAZ09 = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
