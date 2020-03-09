package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestAsyncAPI_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"asyncapi":"2.0.0","id":"bbfff","info":{"title":"cdef","version":"aadce","description":"dc","termsOfService":"daecfe","contact":{"name":"cddbd","url":"accdf","email":"fd"},"license":{"name":"ad","url":"ead"}},"servers":{"adbfc":{"url":"aebd","description":"ddbfe","protocol":"bfcce","protocolVersion":"fdfed","variables":{"cdba":{"enum":["bedceb"],"default":"dadf","description":"bffd","examples":["eae"]}},"security":[{"db":["bfdbc"]}],"bindings":{"ec":"eb"}}},"defaultContentType":"fffbce","channels":{"dca":{"$ref":"ecaddd","parameters":{"bcf":{"description":"fa","schema":{"abdfad":"aeab"},"location":"bafd","$ref":"eaceae"}},"description":"cc","publish":{"traits":[{"adfae":"daeaab"}],"summary":"acc","description":"fcadff","tags":[{"name":"eca","description":"bdb","externalDocs":{}}],"externalDocs":{"description":"dfe","url":"cfdcf"},"operationId":"baf","bindings":{"amqp":{"expiration":4725,"userId":"dad","priority":6968,"deliveryMode":660,"mandatory":true,"replyTo":"fe","timestamp":true,"ack":true,"bindingVersion":"latest"},"efbef":"fadcba"},"message":{}},"subscribe":{"traits":[[]],"summary":"bbecdf","description":"caf","tags":[{"name":"eeacab","description":"bd","externalDocs":{}}],"externalDocs":{"description":"db","url":"ebae"},"operationId":"fcbcf","bindings":{"amqp":{"expiration":599,"userId":"eebbee","priority":2257,"deliveryMode":7130,"mandatory":true,"replyTo":"baabed","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"cdcb":"bd"},"message":{}},"deprecated":true,"bindings":{"amqp":{"is":"queue","exchange":{"name":"db","type":"fanout","durable":true,"autoDelete":true,"eefcce":"db"},"bindingVersion":"0.1.0"},"feebff":"dbdeae"}}},"components":{"schemas":{"aeade":{"cc":"bd"}},"messages":{"fcccfe":{"oneOf":[]}},"securitySchemes":{},"parameters":{"ec":{"description":"cae","schema":{"ecafef":"eebd"},"location":"fcf","$ref":"fdecbe"}},"correlationIds":{},"operationTraits":{"eea":{"summary":"dccbf","description":"dcbaf","tags":[{"name":"abcdaf","description":"ac","externalDocs":{}}],"externalDocs":{"description":"dcefd","url":"dfecaf"},"operationId":"ffaf","bindings":{"amqp":{"expiration":150,"userId":"adbcdf","priority":8095,"deliveryMode":3447,"mandatory":true,"replyTo":"eceec","timestamp":true,"ack":true,"bindingVersion":"latest"},"eea":"aef"}}},"messageTraits":{"fdafad":{"schemaFormat":"dc","contentType":"aadabb","headers":{"ba":"eb"},"correlationId":{"$ref":"ccfde","acfacf":"dfcdfa"},"tags":[{"name":"ac","description":"cac","externalDocs":{}}],"summary":"ddd","name":"aceba","title":"bc","description":"dfa","externalDocs":{"description":"bcacab","url":"ae"},"deprecated":true,"examples":[{"feb":"adfd"}],"bindings":{"amqp":{"contentEncoding":"efe","messageType":"ec","bindingVersion":"0.1.0"},"ffdfa":"ffee"}}},"serverBindings":{"afc":{"eda":"cacafa"}},"channelBindings":{"fde":{"amqp":{"is":"queue","exchange":{"name":"bb","type":"headers","durable":true,"autoDelete":true,"ec":"acecb"},"bindingVersion":"0.1.0"},"ef":"befe"}},"operationBindings":{"dbfbf":{"amqp":{"expiration":1393,"userId":"cebee","priority":1109,"deliveryMode":1,"mandatory":true,"replyTo":"faeefd","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"dc":"eebc"}},"messageBindings":{"ba":{"amqp":{"contentEncoding":"ce","messageType":"cdbcc","bindingVersion":"0.1.0"},"cabba":"ccfc"}}},"tags":[{"name":"dfce","description":"dbbfd","externalDocs":{"description":"fabcce","url":"caee"}}],"externalDocs":{"description":"cbed","url":"edbe"}}`)
		v AsyncAPI
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestInfo_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"title":"bddcd","version":"eb","description":"bbab","termsOfService":"ddf","contact":{"name":"ebdb","url":"cbe","email":"becfee"},"license":{"name":"aaab","url":"fe"}}`)
		v Info
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestContact_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"da","url":"ce","email":"bfcfa"}`)
		v Contact
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestLicense_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"af","url":"bfe"}`)
		v License
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestServer_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"url":"cf","description":"accb","protocol":"aa","protocolVersion":"dfba","variables":{"ab":{"enum":["ac"],"default":"faceb","description":"cffd","examples":["bbeffa"]}},"security":[{"bfd":["dde"]}],"bindings":{"ecf":"fbffc"}}`)
		v Server
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestServerVariable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"enum":["ffde"],"default":"df","description":"aeedb","examples":["ccafeb"]}`)
		v ServerVariable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestServerBindingsObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"eea":"efaabe"}`)
		v ServerBindingsObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestChannelItem_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"cbe","parameters":{"caa":{"description":"bf","schema":{"aacdcd":"bdc"},"location":"ebaebb","$ref":"bc"}},"description":"daaf","publish":{"traits":[[{},{"bebfdd":"dfead"}]],"summary":"aad","description":"ebb","tags":[{"name":"aac","description":"cfd","externalDocs":{"description":"ad","url":"aebd"}}],"externalDocs":{"description":"dea","url":"ac"},"operationId":"baed","bindings":{"amqp":{"expiration":9108,"userId":"dfedae","priority":9389,"deliveryMode":2,"mandatory":true,"replyTo":"dd","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"bcfdec":"eadbbf"},"message":{"schemaFormat":"dadfa","contentType":"bbee","headers":{"aeceb":"daedb"},"correlationId":{"cceec":"aa"},"tags":[{}],"summary":"ceadd","name":"bcdde","title":"bf","description":"fdb","externalDocs":{"description":"efaafd","url":"ccfe"},"deprecated":true,"examples":[{"cca":"ea"}],"bindings":{"amqp":{},"eceb":"dcbef"},"traits":[null]}},"subscribe":{"traits":[{"$ref":"abbc","fedf":"dccf"}],"summary":"bfc","description":"efaae","tags":[{"name":"adbedd","description":"ffe","externalDocs":{"description":"bfd","url":"decaab"}}],"externalDocs":{"description":"cec","url":"ededcf"},"operationId":"caf","bindings":{"amqp":{"expiration":7702,"userId":"dfd","priority":8629,"deliveryMode":1,"mandatory":true,"replyTo":"ccefda","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"dbcbb":"ccccf"},"message":{"oneOf":[null]}},"deprecated":true,"bindings":{"amqp":{"is":"routingKey","exchange":{"name":"fbe","type":"fanout","durable":true,"autoDelete":true,"eabbff":"efccf"},"bindingVersion":"latest"},"bbbd":"ccdfa"}}`)
		v ChannelItem
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestParameter_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"fb","schema":{"dbbff":"afbeaf"},"location":"cdd","$ref":"cea"}`)
		v Parameter
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperation_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"traits":[{"summary":"bac","description":"faf","tags":[{"name":"beda","description":"fdc","externalDocs":{}}],"externalDocs":{"description":"bdbb","url":"bebdb"},"operationId":"cdaef","bindings":{"amqp":{"expiration":4312,"userId":"cf","priority":1468,"deliveryMode":7293,"mandatory":true,"replyTo":"adedf","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"deaa":"fbbbb"}}],"summary":"aebbfb","description":"efbafb","tags":[{"name":"befffe","description":"decac","externalDocs":{"description":"fbd","url":"bdcbf"}}],"externalDocs":{"description":"eafece","url":"bbccdf"},"operationId":"dcb","bindings":{"amqp":{"expiration":7871,"userId":"bd","priority":5279,"deliveryMode":2,"mandatory":true,"replyTo":"cadd","timestamp":true,"ack":true,"bindingVersion":"latest"},"ffee":"afdfca"},"message":{"schemaFormat":"fe","contentType":"aadade","headers":{"fbfdc":"dffeda"},"correlationId":{"$ref":"dcfdfd","acbeeb":"aaaea"},"tags":[{"name":"fdecdf","description":"bdeed","externalDocs":{}}],"summary":"ab","name":"dd","title":"faafe","description":"db","externalDocs":{"description":"fbaab","url":"fccbcb"},"deprecated":true,"examples":[{"ea":"bc"}],"bindings":{"amqp":{"contentEncoding":"dedbac","messageType":"ceccca","bindingVersion":"0.1.0"},"edc":"adecb"},"traits":[[]]}}`)
		v Operation
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestReference_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"db","dff":"efedc"}`)
		v Reference
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperationTrait_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"summary":"feab","description":"bdda","tags":[{"name":"edcf","description":"afefb","externalDocs":{"description":"abeac","url":"efb"}}],"externalDocs":{"description":"eea","url":"ef"},"operationId":"bc","bindings":{"amqp":{"expiration":6088,"userId":"ccfeb","priority":1896,"deliveryMode":1,"mandatory":true,"replyTo":"bdabf","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"baee":"adafef"}}`)
		v OperationTrait
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTag_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"cecca","description":"ddfa","externalDocs":{"description":"cdc","url":"fbdeaf"}}`)
		v Tag
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestExternalDocs_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"fadffe","url":"debf"}`)
		v ExternalDocs
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperationBindingsObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"amqp":{"expiration":4570,"userId":"edefca","priority":6137,"deliveryMode":1,"mandatory":true,"replyTo":"aaeedb","timestamp":true,"ack":true,"bindingVersion":"latest"},"bebfa":"bafccc"}`)
		v OperationBindingsObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAMQP091OperationBindingObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"expiration":8982,"userId":"dcd","priority":4772,"deliveryMode":2,"mandatory":true,"replyTo":"ccd","timestamp":true,"ack":true,"bindingVersion":"latest"}`)
		v AMQP091OperationBindingObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperationTraitsItems_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`[{"$ref":"fbc","ba":"dade"},{"fdddef":"cab"}]`)
		v OperationTraitsItems
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageOneOf1OneOf0_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"oneOf":[{"oneOf":[null]}]}`)
		v MessageOneOf1OneOf0
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageOneOf1OneOf1_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemaFormat":"bbf","contentType":"fab","headers":{"ffadb":"dbfb"},"correlationId":{"description":"ebe","location":"cff"},"tags":[{"name":"cfcf","description":"cdffb","externalDocs":{"description":"aeabef","url":"bc"}}],"summary":"fdfa","name":"fcced","title":"edbadd","description":"cebffe","externalDocs":{"description":"aab","url":"caa"},"deprecated":true,"examples":[{"ceeffd":"bc"}],"bindings":{"amqp":{"contentEncoding":"bbaeb","messageType":"cdfa","bindingVersion":"0.1.0"},"fbeaea":"fabcf"},"traits":[{"schemaFormat":"cfe","contentType":"efb","headers":{"cffcdf":"bda"},"correlationId":{"$ref":"faa","aebafd":"dbeea"},"tags":[{"name":"efbada","description":"cfecef","externalDocs":{}}],"summary":"fdea","name":"bcaceb","title":"bbed","description":"bef","externalDocs":{"description":"ccff","url":"fbedff"},"deprecated":true,"examples":[{"bdfa":"ccb"}],"bindings":{"amqp":{"contentEncoding":"afadf","messageType":"dff","bindingVersion":"0.1.0"},"afe":"faa"}}]}`)
		v MessageOneOf1OneOf1
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestCorrelationID_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"da","location":"faeff"}`)
		v CorrelationID
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageOneOf1OneOf1CorrelationID_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"cf","ecb":"ec"}`)
		v MessageOneOf1OneOf1CorrelationID
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageBindingsObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"amqp":{"contentEncoding":"ceb","messageType":"cf","bindingVersion":"0.1.0"},"abecce":"ffaae"}`)
		v MessageBindingsObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAMQP091MessageBindingObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"contentEncoding":"be","messageType":"dfddca","bindingVersion":"latest"}`)
		v AMQP091MessageBindingObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageTrait_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemaFormat":"eb","contentType":"aec","headers":{"$ref":"cebbbb","fab":"faaf"},"correlationId":{"$ref":"be","ecec":"acfc"},"tags":[{"name":"acdbfe","description":"bfaea","externalDocs":{"description":"affeea","url":"bc"}}],"summary":"ebf","name":"ad","title":"acf","description":"cf","externalDocs":{"description":"edaddd","url":"bdab"},"deprecated":true,"examples":[{"cdadb":"ecbacf"}],"bindings":{"amqp":{"contentEncoding":"caeae","messageType":"ce","bindingVersion":"0.1.0"},"eaea":"afbb"}}`)
		v MessageTrait
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageTraitHeaders_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"fa":"cbdead"}`)
		v MessageTraitHeaders
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageTraitCorrelationID_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"bdbd","location":"fa"}`)
		v MessageTraitCorrelationID
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageOneOf1OneOf1TraitsItems_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemaFormat":"dff","contentType":"dcbdfe","headers":{"cfba":"dacee"},"correlationId":{"$ref":"eca","bdbcc":"fbeae"},"tags":[{"name":"bff","description":"fc","externalDocs":{"description":"decc","url":"bfeed"}}],"summary":"ac","name":"ea","title":"cecbeb","description":"cefa","externalDocs":{"description":"ffad","url":"bdbdc"},"deprecated":true,"examples":[{"cfbfac":"eedf"}],"bindings":{"amqp":{"contentEncoding":"bc","messageType":"ccb","bindingVersion":"0.1.0"},"dcbe":"cdeec"}}`)
		v MessageOneOf1OneOf1TraitsItems
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessageOneOf1_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemaFormat":"cd","contentType":"afcd","headers":{"accceb":"caeed"},"correlationId":{"description":"caecb","location":"edd"},"tags":[{"name":"caafcb","description":"ef","externalDocs":{"description":"bfab","url":"bacfe"}}],"summary":"eabc","name":"ebfc","title":"af","description":"defc","externalDocs":{"description":"fbbf","url":"afcaf"},"deprecated":true,"examples":[{"eefe":"aafddd"}],"bindings":{"amqp":{"contentEncoding":"bf","messageType":"bdaaec","bindingVersion":"latest"},"aa":"fbe"},"traits":[[{"ffb":"be"},{"bfc":"ebead"}]]}`)
		v MessageOneOf1
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessage_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemaFormat":"fdc","contentType":"bbdfe","headers":{"fbbdc":"eebb"},"correlationId":{"$ref":"db","ecf":"cc"},"tags":[{"name":"fae","description":"bdccfc","externalDocs":{"description":"deeca","url":"fafe"}}],"summary":"aaf","name":"ae","title":"cabb","description":"aafcb","externalDocs":{"description":"deed","url":"cce"},"deprecated":true,"examples":[{"cf":"babb"}],"bindings":{"amqp":{"contentEncoding":"eaefb","messageType":"bc","bindingVersion":"latest"},"aafa":"dfffea"},"traits":[{"schemaFormat":"ecc","contentType":"befdb","tags":[],"summary":"aefff","name":"cc","title":"aea","description":"ebada","externalDocs":{},"deprecated":true,"examples":[],"bindings":{"cab":"afbbc"}}]}`)
		v Message
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestChannelBindingsObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"amqp":{"is":"routingKey","exchange":{"name":"ee","type":"fanout","durable":true,"autoDelete":true,"afc":"bdcc"},"bindingVersion":"latest"},"dbefde":"dae"}`)
		v ChannelBindingsObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAMQP091ChannelBindingObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"is":"queue","exchange":{"name":"dccdf","type":"topic","durable":true,"autoDelete":true,"fbfad":"fcbee"},"bindingVersion":"latest"}`)
		v AMQP091ChannelBindingObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestExchange_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"af","type":"direct","durable":true,"autoDelete":true,"acdd":"fefc"}`)
		v Exchange
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestQueue_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"fcfa","durable":true,"exclusive":true,"autoDelete":true,"ebb":"dfca"}`)
		v Queue
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponents_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemas":{"dffab":{"bc":"fc"}},"messages":{"af":{"oneOf":[null]}},"securitySchemes":{},"parameters":{"ede":{"description":"cecc","schema":{"dfb":"fa"},"location":"bd","$ref":"ccdfe"}},"correlationIds":{},"operationTraits":{"bebfd":{"summary":"afae","description":"dedf","tags":[{"name":"acaecc","description":"decdca","externalDocs":{"description":"ddb","url":"bebfde"}}],"externalDocs":{"description":"dffef","url":"ed"},"operationId":"becbd","bindings":{"amqp":{"expiration":9392,"userId":"da","priority":3637,"deliveryMode":2,"mandatory":true,"replyTo":"bcdc","timestamp":true,"ack":true,"bindingVersion":"0.1.0"},"becc":"ea"}}},"messageTraits":{"fdcade":{"schemaFormat":"bebfbb","contentType":"cccdd","headers":{"$ref":"babc","afbefe":"eeccb"},"correlationId":{"$ref":"bddea","bafaae":"ffdcfb"},"tags":[{"name":"eecee","description":"bbccef","externalDocs":{"description":"fabfcb","url":"cf"}}],"summary":"fbae","name":"ccac","title":"abdafa","description":"fec","externalDocs":{"description":"fcddbe","url":"cbc"},"deprecated":true,"examples":[{"eb":"cbea"}],"bindings":{"amqp":{"contentEncoding":"cfe","messageType":"aedcf","bindingVersion":"latest"},"ab":"ab"}}},"serverBindings":{"aea":{"dacb":"ead"}},"channelBindings":{"fde":{"amqp":{"is":"routingKey","exchange":{"name":"afdbae","type":"topic","durable":true,"autoDelete":true,"bbefdc":"ecbbff"},"bindingVersion":"latest"},"bc":"febbcd"}},"operationBindings":{"caeb":{"amqp":{"expiration":8014,"userId":"bb","priority":8148,"deliveryMode":2,"mandatory":true,"replyTo":"eee","timestamp":true,"ack":true,"bindingVersion":"latest"},"ae":"aa"}},"messageBindings":{"cbede":{"amqp":{"contentEncoding":"edde","messageType":"daea","bindingVersion":"0.1.0"},"fcfc":"facbd"}}}`)
		v Components
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestUserPassword_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"userPassword","description":"dffd"}`)
		v UserPassword
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAPIKey_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"apiKey","in":"user","description":"cea"}`)
		v APIKey
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestX509_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"X509","description":"bcdd"}`)
		v X509
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestSymmetricEncryption_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"symmetricEncryption","description":"dbb"}`)
		v SymmetricEncryption
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAsymmetricEncryption_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"asymmetricEncryption","description":"fcccd"}`)
		v AsymmetricEncryption
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNonBearerHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"de","description":"ff","type":"http"}`)
		v NonBearerHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNonBearerHTTPSecuritySchemeNot_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","eebfcf":"bdfbb"}`)
		v NonBearerHTTPSecuritySchemeNot
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestBearerHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","bearerFormat":"cefcd","type":"http","description":"abbe"}`)
		v BearerHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAPIKeyHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"httpApiKey","name":"cdacca","in":"header","description":"bcdeba"}`)
		v APIKeyHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"httpApiKey","name":"dbc","in":"query","description":"acbdc"}`)
		v HTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOauth2Flows_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"oauth2","description":"ceeff","flows":{"implicit":{"authorizationUrl":"affaac","tokenUrl":"dfadb","refreshUrl":"cf","scopes":{"baca":"ca"}},"password":{"authorizationUrl":"bf","tokenUrl":"dbedae","refreshUrl":"ce","scopes":{"abba":"eefac"}},"clientCredentials":{"authorizationUrl":"fcddb","tokenUrl":"ddcc","refreshUrl":"efdb","scopes":{"efae":"ecf"}},"authorizationCode":{"authorizationUrl":"dd","tokenUrl":"eaeac","refreshUrl":"efeba","scopes":{"deec":"acaab"}}},"aecce":"dcbbf"}`)
		v Oauth2Flows
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOauth2FlowsFlows_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"implicit":{"authorizationUrl":"eda","tokenUrl":"ceb","refreshUrl":"be","scopes":{"bdcd":"ec"}},"password":{"authorizationUrl":"badda","tokenUrl":"cadbdf","refreshUrl":"bcecec","scopes":{"fdbaa":"eefed"}},"clientCredentials":{"authorizationUrl":"ccfaf","tokenUrl":"dacffd","refreshUrl":"af","scopes":{"ffe":"eb"}},"authorizationCode":{"authorizationUrl":"cad","tokenUrl":"cfb","refreshUrl":"cc","scopes":{"cfaca":"fb"}}}`)
		v Oauth2FlowsFlows
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOauth2Flow_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"authorizationUrl":"faeba","tokenUrl":"feedf","refreshUrl":"afed","scopes":{"fbefb":"beccbb"}}`)
		v Oauth2Flow
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOpenIDConnect_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"openIdConnect","description":"bfbed","openIdConnectUrl":"df"}`)
		v OpenIDConnect
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"X509","description":"ddecda"}`)
		v SecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsSecuritySchemesWD_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","bearerFormat":"add","type":"http","description":"bee"}`)
		v ComponentsSecuritySchemesWD
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsSecuritySchemes_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{}`)
		v ComponentsSecuritySchemes
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsCorrelationIdsWD_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"aaf","location":"deebef"}`)
		v ComponentsCorrelationIdsWD
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsCorrelationIds_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{}`)
		v ComponentsCorrelationIds
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}
