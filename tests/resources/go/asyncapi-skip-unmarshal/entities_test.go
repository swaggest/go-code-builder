package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestAsyncAPI_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"asyncapi":"1.1.0","info":{"title":"acbbff","version":"acdefb","description":"ad","termsOfService":"eadc","contact":{"name":"daecfe","url":"cddbd","email":"accdf"},"license":{"name":"fd","url":"ad"}},"baseTopic":"ead","servers":[{"url":"adbfc","description":"aebd","scheme":"mqtt","schemeVersion":"dbfee","variables":{"fcceaf":{"enum":["fedacd"],"default":"ab","description":"edc"}}}],"topics":{},"stream":{"framing":{"type":"chunked","delimiter":"\\n"},"read":[{"$ref":"adfcb","headers":{"$ref":"fdeea","format":"cd","title":"bb","description":"db","multipleOf":9565.357,"maximum":5893.055,"exclusiveMaximum":true,"minimum":2609.79,"exclusiveMinimum":true,"maxLength":6998,"minLength":2310,"pattern":"bfff","maxItems":4931,"minItems":4142,"uniqueItems":true,"maxProperties":5406,"minProperties":500,"required":["dca"],"enum":["ecaddd"],"additionalProperties":{}},"payload":{"$ref":"bcf","format":"fa","title":"abdfad","description":"aeab","multipleOf":1263.149,"maximum":8833.061,"exclusiveMaximum":true,"minimum":6902.046,"exclusiveMinimum":true,"maxLength":6237,"minLength":6477,"pattern":"eaceae","maxItems":7528,"minItems":4281,"uniqueItems":true,"maxProperties":8813,"minProperties":9648,"required":["adfae"],"enum":["daeaab"],"additionalProperties":{}},"tags":[{"name":"acc","description":"fcadff","externalDocs":{}}],"summary":"eca","description":"bdb","externalDocs":{"description":"dfe","url":"cfdcf"},"deprecated":true}],"write":[{"$ref":"baf","headers":{"$ref":"dda","format":"acdcfe","title":"de","description":"befdf","multipleOf":3963.133,"maximum":262.11,"exclusiveMaximum":true,"minimum":1073.653,"exclusiveMinimum":true,"maxLength":283,"minLength":5098,"pattern":"bffbb","maxItems":5528,"minItems":9677,"uniqueItems":true,"maxProperties":8420,"minProperties":2783,"required":["caf"],"enum":["eeacab"],"additionalProperties":{}},"payload":{"$ref":"bd","format":"db","title":"ebae","description":"fcbcf","multipleOf":599.177,"maximum":1533.385,"exclusiveMaximum":true,"minimum":1213.435,"exclusiveMinimum":true,"maxLength":9372,"minLength":446,"pattern":"ee","maxItems":2257,"minItems":6163,"uniqueItems":true,"maxProperties":7130,"minProperties":360,"required":["aab"],"enum":["daac"],"additionalProperties":{}},"tags":[{"name":"cba","description":"df","externalDocs":{}}],"summary":"bed","description":"dde","externalDocs":{"description":"fc","url":"ec"},"deprecated":true}]},"events":{"receive":[{"$ref":"babfe","headers":{"$ref":"bffb","format":"bdea","title":"aaead","description":"ac","multipleOf":9164.437,"maximum":6692.329,"exclusiveMaximum":true,"minimum":4621.525,"exclusiveMinimum":true,"maxLength":2649,"minLength":1323,"pattern":"cccf","maxItems":743,"minItems":2102,"uniqueItems":true,"maxProperties":954,"minProperties":7528,"required":["cec"],"enum":["efecaf"],"additionalProperties":{}},"payload":{"$ref":"fdeeb","format":"afcfaf","title":"ec","description":"eaeeaf","multipleOf":3976.113,"maximum":9501.761,"exclusiveMaximum":true,"minimum":9252.001,"exclusiveMinimum":true,"maxLength":5567,"minLength":2035,"pattern":"dcbaf","maxItems":4860,"minItems":1670,"uniqueItems":true,"maxProperties":6043,"minProperties":241,"required":["afe"],"enum":["ce"],"additionalProperties":{}},"tags":[{"name":"cef","description":"bdfec","externalDocs":{}}],"summary":"fcf","description":"afd","externalDocs":{"description":"adbcdf","url":"debec"},"deprecated":true}],"send":[{"$ref":"ecb","headers":{"$ref":"eea","format":"aef","title":"fdafad","description":"dc","multipleOf":5934.814,"maximum":2183.932,"exclusiveMaximum":true,"minimum":4336.764,"exclusiveMinimum":true,"maxLength":2487,"minLength":8074,"pattern":"bbfbad","maxItems":1840,"minItems":1913,"uniqueItems":true,"maxProperties":7879,"minProperties":5888,"required":["cfd"],"enum":["eac"],"additionalProperties":{}},"payload":{"$ref":"ac","format":"fdfc","title":"fafacf","description":"acfdd","multipleOf":6975.095,"maximum":7500.218,"exclusiveMaximum":true,"minimum":747.239,"exclusiveMinimum":true,"maxLength":8583,"minLength":365,"pattern":"adb","maxItems":8645,"minItems":47,"uniqueItems":true,"maxProperties":3227,"minProperties":2642,"required":["db"],"enum":["aca"],"additionalProperties":{}},"tags":[{"name":"eaecfe","description":"cadfd","externalDocs":{}}],"summary":"efe","description":"ec","externalDocs":{"description":"dff","url":"facf"},"deprecated":true}]},"components":{"schemas":{"eeea":{"$ref":"ccedaf","format":"acaf","title":"cfd","description":"ddb","multipleOf":7056.821,"maximum":796.261,"exclusiveMaximum":true,"minimum":2382.109,"exclusiveMinimum":true,"maxLength":9828,"minLength":444,"pattern":"acecb","maxItems":6884,"minItems":6603,"uniqueItems":true,"maxProperties":5695,"minProperties":1874,"required":["befe"],"enum":["dbfbf"],"additionalProperties":{"$ref":"aceb","format":"ed","title":"ffa","description":"efdc","multipleOf":2480.336,"maximum":6731.473,"exclusiveMaximum":true,"minimum":8516.719,"exclusiveMinimum":true,"maxLength":555,"pattern":"ebceba","maxItems":8905,"uniqueItems":true,"maxProperties":5159,"required":[],"enum":["ccdbc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}}},"messages":{"cecabb":{"$ref":"bccfc","headers":{"$ref":"dfce","format":"dbbfd","title":"fabcce","description":"caee","multipleOf":7233.648,"maximum":7035.17,"exclusiveMaximum":true,"minimum":7592.52,"exclusiveMinimum":true,"maxLength":6102,"minLength":4063,"pattern":"edbe","maxItems":9696,"minItems":7317,"uniqueItems":true,"maxProperties":3039,"minProperties":2597,"required":["ddeb"],"enum":["bbab"],"additionalProperties":{}},"payload":{"$ref":"ddf","format":"ebdb","title":"cbe","description":"becfee","multipleOf":6592,"maximum":5264.259,"exclusiveMaximum":true,"minimum":8090.772,"exclusiveMinimum":true,"maxLength":1354,"minLength":4582,"pattern":"fe","maxItems":5339,"minItems":1604,"uniqueItems":true,"maxProperties":5773,"minProperties":7456,"required":["eb"],"enum":["fc"],"additionalProperties":{}},"tags":[{"name":"ac","description":"fd","externalDocs":{}}],"summary":"fefcfc","description":"ccba","externalDocs":{"description":"ad","url":"fbac"},"deprecated":true}},"securitySchemes":{"baac":"faceb"},"parameters":{"cffd":{"description":"bbeffa","name":"bfd","schema":{"$ref":"dde","format":"ecf","title":"fbffc","description":"ffde","multipleOf":3878.278,"maximum":6968.005,"exclusiveMaximum":true,"minimum":4585.232,"exclusiveMinimum":true,"maxLength":7662,"minLength":7357,"pattern":"edbac","maxItems":8240,"minItems":6563,"uniqueItems":true,"maxProperties":6265,"minProperties":113,"required":["deeae"],"enum":["faabea"],"additionalProperties":{}},"$ref":"befcaa"}}},"tags":[{"name":"bf","description":"aacdcd","externalDocs":{"description":"bdc","url":"ebaebb"}}],"security":[{"bc":["daaf"]}],"externalDocs":{"description":"febebf","url":"dfdfea"}}`)
		v AsyncAPI
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestInfo_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"title":"ba","version":"dbebb","description":"aac","termsOfService":"cfd","contact":{"name":"ad","url":"aebd","email":"dea"},"license":{"name":"ac","url":"baed"}}`)
		v Info
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestContact_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"bdfeda","url":"ef","email":"dd"}`)
		v Contact
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestLicense_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"ab","url":"fdecf"}`)
		v License
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestServer_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"url":"adbbf","description":"bdd","scheme":"secure-mqtt","schemeVersion":"fa","variables":{"bbee":{"enum":["aeceb"],"default":"daedb","description":"ccce"}}}`)
		v Server
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestServerVariable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"enum":["ceaa"],"default":"ceadd","description":"bcdde"}`)
		v ServerVariable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTopics_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{}`)
		v Topics
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTopicItem_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"bf","parameters":[{"description":"fdb","name":"efaafd","schema":{"$ref":"ccfe","format":"cca","title":"ea","description":"eceb","multipleOf":9015.626,"maximum":7641.239,"exclusiveMaximum":true,"minimum":5956.079,"exclusiveMinimum":true,"maxLength":6269,"minLength":324,"pattern":"baf","maxItems":684,"minItems":4517,"uniqueItems":true,"maxProperties":8797,"minProperties":3754,"required":["fedf"],"enum":["dccf"],"additionalProperties":{"$ref":"bfc","format":"efaae","title":"adbedd","description":"ffe","multipleOf":108.887,"maximum":1118.998,"exclusiveMaximum":true,"minimum":7582.783,"exclusiveMinimum":true,"maxLength":3130,"pattern":"decaab","maxItems":6355,"uniqueItems":true,"maxProperties":5239,"required":[],"enum":["cc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"$ref":"de"}],"publish":{"oneOf":[{"$ref":"fd","headers":{"$ref":"afc","format":"dfd","title":"ea","description":"ce","multipleOf":9749.913,"maximum":6724.698,"exclusiveMaximum":true,"minimum":6979.371,"exclusiveMinimum":true,"maxLength":6868,"pattern":"dbcbb","maxItems":8077,"uniqueItems":true,"maxProperties":52,"required":[],"enum":["ccf"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"payload":{"$ref":"ada","format":"fbe","title":"aeab","description":"ffdefc","multipleOf":2469.205,"maximum":7473.208,"exclusiveMaximum":true,"minimum":2498.697,"exclusiveMinimum":true,"maxLength":2896,"pattern":"bbde","maxItems":8077,"uniqueItems":true,"maxProperties":3803,"required":[],"enum":["fadfbd"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"tags":[{}],"summary":"bbff","description":"afbeaf","externalDocs":{"description":"cdd","url":"cea"},"deprecated":true},{"$ref":"bbac","headers":{"$ref":"faf","format":"beda","title":"fdc","description":"bdbb","multipleOf":7060.518,"maximum":282.695,"exclusiveMaximum":true,"minimum":4780.313,"exclusiveMinimum":true,"maxLength":680,"pattern":"bfcdae","maxItems":9350,"uniqueItems":true,"maxProperties":4312,"required":[],"enum":["cf"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"payload":{"$ref":"fecad","format":"dfcbd","title":"aaa","description":"bbbb","multipleOf":449.278,"maximum":7316.065,"exclusiveMaximum":true,"minimum":3601.327,"exclusiveMinimum":true,"maxLength":9453,"pattern":"fb","maxItems":7822,"uniqueItems":true,"maxProperties":3380,"required":[],"enum":["baf"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"tags":[{}],"summary":"bbe","description":"ffeed","externalDocs":{"description":"cacc","url":"bdc"},"deprecated":true}]},"subscribe":{"oneOf":[{"$ref":"cbfa","headers":{"$ref":"afecef","format":"bc","title":"dfcd","description":"bcfbd","multipleOf":5279.04,"maximum":5840.998,"exclusiveMaximum":true,"minimum":7400.775,"exclusiveMinimum":true,"maxLength":3866,"pattern":"ddb","maxItems":2166,"uniqueItems":true,"maxProperties":5476,"required":[],"enum":["eedaf"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"payload":{"$ref":"fcabb","format":"fe","title":"aadade","description":"fbfdc","multipleOf":8578.971,"maximum":2288.546,"exclusiveMaximum":true,"minimum":7288.509,"exclusiveMinimum":true,"maxLength":8314,"pattern":"da","maxItems":327,"uniqueItems":true,"maxProperties":7805,"required":[],"enum":["cfdfdb"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"tags":[{}],"summary":"cb","description":"ebcaaa","externalDocs":{"description":"acfde","url":"dfcb"},"deprecated":true},{"$ref":"eeda","headers":{"$ref":"bad","format":"bf","title":"afeed","description":"ef","multipleOf":3456.988,"maximum":3504.542,"exclusiveMaximum":true,"minimum":6749.716,"exclusiveMinimum":true,"maxLength":2620,"pattern":"fccbcb","maxItems":7064,"uniqueItems":true,"maxProperties":8774,"required":[],"enum":["abcade"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"payload":{"$ref":"bac","format":"ceccca","title":"eedcc","description":"decb","multipleOf":9427.626,"maximum":9180.391,"exclusiveMaximum":true,"minimum":5178.601,"exclusiveMinimum":true,"maxLength":877,"pattern":"dff","maxItems":2516,"uniqueItems":true,"maxProperties":9629,"required":[],"enum":["edc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}},"tags":[{}],"summary":"feab","description":"bdda","externalDocs":{"description":"edcf","url":"afefb"},"deprecated":true}]},"deprecated":true}`)
		v TopicItem
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestParameter_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"abeac","name":"efb","schema":{"$ref":"eea","format":"ef","title":"bc","description":"accfeb","multipleOf":1895.517,"maximum":8929.009,"exclusiveMaximum":true,"minimum":9649.889,"exclusiveMinimum":true,"maxLength":9393,"minLength":2942,"pattern":"bf","maxItems":1443,"minItems":1342,"uniqueItems":true,"maxProperties":141,"minProperties":7088,"required":["ec"],"enum":["daf"],"additionalProperties":{"$ref":"fececc","format":"ad","title":"faccdc","description":"fbdeaf","multipleOf":967.74,"maximum":3203.193,"exclusiveMaximum":true,"minimum":1463.248,"exclusiveMinimum":true,"maxLength":9825,"minLength":4151,"pattern":"ebde","maxItems":2622,"minItems":6196,"uniqueItems":true,"maxProperties":4570,"minProperties":2520,"required":["def"],"enum":["aecaaa"],"additionalProperties":true,"type":"boolean","items":{"$ref":"dbbdb","format":"bfacb","title":"fccc","description":"bdcddb","multipleOf":4498.315,"maximum":6088.312,"exclusiveMaximum":true,"minimum":9358.635,"exclusiveMinimum":true,"maxLength":6825,"pattern":"ceefb","maxItems":4326,"uniqueItems":true,"maxProperties":9385,"required":[],"enum":["aa"],"additionalProperties":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}}},"$ref":"adeaf"}`)
		v Parameter
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestMessage_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"dd","headers":{"$ref":"facabd","format":"debb","title":"bf","description":"bdffa","multipleOf":7073.166,"maximum":3886.164,"exclusiveMaximum":true,"minimum":6135.887,"exclusiveMinimum":true,"maxLength":5027,"minLength":6483,"pattern":"bdee","maxItems":5142,"minItems":1363,"uniqueItems":true,"maxProperties":8753,"minProperties":7754,"required":["fccfc"],"enum":["ecdffb"],"additionalProperties":{"$ref":"aeabef","format":"bc","title":"fdfa","description":"fcced","multipleOf":1247.513,"maximum":603.315,"exclusiveMaximum":true,"minimum":4705.879,"exclusiveMinimum":true,"maxLength":107,"minLength":9866,"pattern":"dace","maxItems":1517,"minItems":4710,"uniqueItems":true,"maxProperties":6391,"minProperties":5390,"required":["aab"],"enum":["caa"],"additionalProperties":true,"type":"string","items":{"$ref":"eef","format":"deb","title":"ebb","description":"ebfcdf","multipleOf":6236.704,"maximum":5849.378,"exclusiveMaximum":true,"minimum":479.926,"exclusiveMinimum":true,"maxLength":9075,"pattern":"eaeaaf","maxItems":814,"uniqueItems":true,"maxProperties":3786,"required":[],"enum":["fb"],"additionalProperties":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}}},"payload":{"$ref":"cfe","format":"efb","title":"fcffcd","description":"db","multipleOf":4338.834,"maximum":9721.037,"exclusiveMaximum":true,"minimum":3824.675,"exclusiveMinimum":true,"maxLength":1488,"minLength":4698,"pattern":"afa","maxItems":7378,"minItems":7744,"uniqueItems":true,"maxProperties":6365,"minProperties":5528,"required":["fdb"],"enum":["ead"],"additionalProperties":{"$ref":"fbada","format":"cfecef","title":"fdea","description":"bcaceb","multipleOf":8749.026,"maximum":6451.294,"exclusiveMaximum":true,"minimum":4518.362,"exclusiveMinimum":true,"maxLength":8253,"minLength":7085,"pattern":"bef","maxItems":4755,"minItems":5285,"uniqueItems":true,"maxProperties":2235,"minProperties":1508,"required":["dfb"],"enum":["dff"],"additionalProperties":true,"type":"array","items":{"$ref":"df","format":"cccb","title":"afadf","description":"dff","multipleOf":5035.095,"maximum":4382.454,"exclusiveMaximum":true,"minimum":9659.331,"exclusiveMinimum":true,"maxLength":3870,"pattern":"dfaaf","maxItems":7771,"uniqueItems":true,"maxProperties":7662,"required":[],"enum":["faeff"],"additionalProperties":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}}},"tags":[{"name":"bcfcec","description":"fec","externalDocs":{"description":"ceb","url":"cf"}}],"summary":"cabe","description":"cedffa","externalDocs":{"description":"eebe","url":"dfddca"},"deprecated":true}`)
		v Message
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTag_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"name":"eebb","description":"eccc","externalDocs":{"description":"eb","url":"bb"}}`)
		v Tag
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestExternalDocs_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"description":"fab","url":"faaf"}`)
		v ExternalDocs
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperationOneOf1_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"oneOf":[{"$ref":"cbea","headers":{"$ref":"cecdac","format":"cfac","title":"bfe","description":"bfaea","multipleOf":8125.073,"maximum":3497.294,"exclusiveMaximum":true,"minimum":2838.018,"exclusiveMinimum":true,"maxLength":4888,"minLength":5278,"pattern":"aab","maxItems":2498,"minItems":9744,"uniqueItems":true,"maxProperties":7528,"minProperties":3117,"required":["aadc"],"enum":["cfd"],"additionalProperties":{"$ref":"fdedad","format":"dd","title":"dabb","description":"dadbfe","multipleOf":7677.89,"maximum":1647.228,"exclusiveMaximum":true,"minimum":5360.085,"exclusiveMinimum":true,"maxLength":3526,"pattern":"bcae","maxItems":6337,"uniqueItems":true,"maxProperties":3357,"required":[],"enum":["ce"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"de","format":"eaca","title":"bbdff","description":"dcbd","multipleOf":2529.294,"maximum":9517.757,"exclusiveMaximum":true,"minimum":846.29,"exclusiveMinimum":true,"maxLength":509,"minLength":4427,"pattern":"dbdbf","maxItems":8638,"minItems":2894,"uniqueItems":true,"maxProperties":816,"minProperties":5103,"required":["fe"],"enum":["cb"],"additionalProperties":{"$ref":"febcc","format":"babd","title":"ceece","description":"ca","multipleOf":4345.95,"maximum":9311.106,"exclusiveMaximum":true,"minimum":2358.146,"exclusiveMinimum":true,"maxLength":8271,"pattern":"cef","maxItems":649,"uniqueItems":true,"maxProperties":5051,"required":[],"enum":["edb"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"fdfcfd","description":"cc","externalDocs":{"description":"bfeed","url":"ac"}}],"summary":"ea","description":"cecbeb","externalDocs":{"description":"cefa","url":"ffad"},"deprecated":true},{"$ref":"bdbdc","headers":{"$ref":"cfbfac","format":"eedf","title":"bc","description":"ccb","multipleOf":1618.477,"maximum":4716.947,"exclusiveMaximum":true,"minimum":8448.725,"exclusiveMinimum":true,"maxLength":8472,"minLength":6022,"pattern":"acde","maxItems":5956,"minItems":4115,"uniqueItems":true,"maxProperties":5990,"minProperties":5914,"required":["df"],"enum":["fcdea"],"additionalProperties":{"$ref":"ccebac","format":"eedbfc","title":"ecb","description":"edd","multipleOf":9076.292,"maximum":1229.548,"exclusiveMaximum":true,"minimum":844.924,"exclusiveMinimum":true,"maxLength":7213,"pattern":"cbc","maxItems":9375,"uniqueItems":true,"maxProperties":5898,"required":[],"enum":["bfab"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"bacfe","format":"eabc","title":"ebfc","description":"af","multipleOf":7213.431,"maximum":5878.715,"exclusiveMaximum":true,"minimum":9310.118,"exclusiveMinimum":true,"maxLength":2641,"minLength":4852,"pattern":"fbbf","maxItems":898,"minItems":5845,"uniqueItems":true,"maxProperties":9668,"minProperties":6958,"required":["ffeef"],"enum":["baafdd"],"additionalProperties":{"$ref":"abfeb","format":"aaecfe","title":"adf","description":"ec","multipleOf":7810.458,"maximum":7187.649,"exclusiveMaximum":true,"minimum":9041.03,"exclusiveMinimum":true,"maxLength":6474,"pattern":"ab","maxItems":6137,"uniqueItems":true,"maxProperties":3911,"required":[],"enum":["fca"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"beadb","description":"bf","externalDocs":{"description":"cfb","url":"dfe"}}],"summary":"fbbdc","description":"eebb","externalDocs":{"description":"bdbf","url":"cfe"},"deprecated":true}]}`)
		v OperationOneOf1
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestOperation_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"dfaebb","headers":{"$ref":"ccfc","format":"deeca","title":"fafe","description":"aaf","multipleOf":8992.209,"maximum":8207.469,"exclusiveMaximum":true,"minimum":756.969,"exclusiveMinimum":true,"maxLength":4863,"minLength":8786,"pattern":"bbeaa","maxItems":1253,"minItems":3443,"uniqueItems":true,"maxProperties":4747,"minProperties":1514,"required":["eedb"],"enum":["cefcf"],"additionalProperties":{"$ref":"babb","format":"eaefb","title":"bc","description":"baaf","multipleOf":4110.837,"maximum":3703.057,"exclusiveMaximum":true,"minimum":8837.95,"exclusiveMinimum":true,"maxLength":1234,"minLength":9567,"pattern":"ea","maxItems":2679,"minItems":9909,"uniqueItems":true,"maxProperties":8372,"minProperties":9300,"required":["fbefdb"],"enum":["ccaef"],"additionalProperties":true,"type":"string","items":{}}},"payload":{"$ref":"eccda","format":"ab","title":"bad","description":"ccabe","multipleOf":1875.699,"maximum":8282.899,"exclusiveMaximum":true,"minimum":8894.425,"exclusiveMinimum":true,"maxLength":2435,"minLength":5834,"pattern":"deecba","maxItems":1454,"minItems":4214,"uniqueItems":true,"maxProperties":2034,"minProperties":7053,"required":["ccb"],"enum":["dbefde"],"additionalProperties":{"$ref":"dae","format":"fdc","title":"dfd","description":"fbfad","multipleOf":1996.751,"maximum":7175.95,"exclusiveMaximum":true,"minimum":4500.258,"exclusiveMinimum":true,"maxLength":1909,"minLength":1161,"pattern":"feafba","maxItems":2425,"minItems":2429,"uniqueItems":true,"maxProperties":3421,"minProperties":6232,"required":["fefc"],"enum":["fcfa"],"additionalProperties":true,"type":"boolean","items":{}}},"tags":[{"name":"bb","description":"dfca","externalDocs":{"description":"dffab","url":"bc"}}],"summary":"fc","description":"af","externalDocs":{"description":"eefed","url":"ece"},"deprecated":true}`)
		v Operation
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestStream_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"framing":{"type":"chunked","delimiter":"\\r\\n"},"read":[{"$ref":"fbd","headers":{"$ref":"aabdec","format":"df","title":"cbebf","description":"dafa","multipleOf":2758.374,"maximum":6043.068,"exclusiveMaximum":true,"minimum":5720.872,"exclusiveMinimum":true,"maxLength":9588,"minLength":2495,"pattern":"caca","maxItems":9523,"minItems":6451,"uniqueItems":true,"maxProperties":1071,"minProperties":3939,"required":["ecd"],"enum":["aeddba"],"additionalProperties":{"$ref":"ebfd","format":"ad","title":"fe","description":"cedb","multipleOf":8626.167,"maximum":8713.61,"exclusiveMaximum":true,"minimum":1674.419,"exclusiveMinimum":true,"maxLength":6866,"pattern":"aa","maxItems":81,"uniqueItems":true,"maxProperties":9787,"required":[],"enum":["dab"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"dc","format":"cbecc","title":"ea","description":"fdcade","multipleOf":3906.105,"maximum":6734.144,"exclusiveMaximum":true,"minimum":6962.701,"exclusiveMinimum":true,"maxLength":6282,"minLength":6785,"pattern":"baccc","maxItems":5013,"minItems":6131,"uniqueItems":true,"maxProperties":6617,"minProperties":9357,"required":["abce"],"enum":["fbefe"],"additionalProperties":{"$ref":"eeccb","format":"fbdde","title":"abafa","description":"efff","multipleOf":7600.382,"maximum":2768.065,"exclusiveMaximum":true,"minimum":5059.677,"exclusiveMinimum":true,"maxLength":9192,"pattern":"eecee","maxItems":7622,"uniqueItems":true,"maxProperties":8960,"required":[],"enum":["cc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"fcfab","description":"cbc","externalDocs":{"description":"fafbae","url":"ccac"}}],"summary":"abdafa","description":"fec","externalDocs":{"description":"fcddbe","url":"cbc"},"deprecated":true}],"write":[{"$ref":"eb","headers":{"$ref":"cbea","format":"cfe","title":"aedcf","description":"eabcab","multipleOf":7397.399,"maximum":9904.062,"exclusiveMaximum":true,"minimum":5436.893,"exclusiveMinimum":true,"maxLength":7739,"minLength":7314,"pattern":"acbf","maxItems":2217,"minItems":5269,"uniqueItems":true,"maxProperties":7747,"minProperties":4171,"required":["deaba"],"enum":["db"],"additionalProperties":{"$ref":"ee","format":"bbefdc","title":"ecbbff","description":"bbcaf","multipleOf":5840.498,"maximum":5030.052,"exclusiveMaximum":true,"minimum":4519.36,"exclusiveMinimum":true,"maxLength":6434,"pattern":"ac","maxItems":7581,"uniqueItems":true,"maxProperties":4385,"required":[],"enum":["cfbb"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"da","format":"eedcae","title":"aa","description":"cbede","multipleOf":1561.496,"maximum":7674.331,"exclusiveMaximum":true,"minimum":8797.221,"exclusiveMinimum":true,"maxLength":3720,"minLength":6320,"pattern":"daea","maxItems":1599,"minItems":4690,"uniqueItems":true,"maxProperties":4443,"minProperties":1334,"required":["ccfa"],"enum":["bd"],"additionalProperties":{"$ref":"bd","format":"fdceb","title":"ea","description":"ab","multipleOf":7847.258,"maximum":98.73,"exclusiveMaximum":true,"minimum":3212.74,"exclusiveMinimum":true,"maxLength":2528,"pattern":"dbb","maxItems":8918,"uniqueItems":true,"maxProperties":102,"required":[],"enum":["ccc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"fdef","description":"fafa","externalDocs":{"description":"ebfcf","url":"bdfbb"}}],"summary":"ccefc","description":"ccabb","externalDocs":{"description":"eecdac","url":"aae"},"deprecated":true}]}`)
		v Stream
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestStreamFramingOneOf0_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"chunked","delimiter":"\\r\\n"}`)
		v StreamFramingOneOf0
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestStreamFramingOneOf1_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"sse","delimiter":"\\n\\n"}`)
		v StreamFramingOneOf1
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestStreamFraming_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"sse","delimiter":"\\n\\n"}`)
		v StreamFraming
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestEvents_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"receive":[{"$ref":"bdb","headers":{"$ref":"ea","format":"cb","title":"ce","description":"ceeff","multipleOf":4927.226,"maximum":627.551,"exclusiveMaximum":true,"minimum":4651.339,"exclusiveMinimum":true,"maxLength":1147,"minLength":8067,"pattern":"cddfa","maxItems":9820,"minItems":6864,"uniqueItems":true,"maxProperties":5850,"minProperties":9899,"required":["abacab"],"enum":["afbfad"],"additionalProperties":{"$ref":"eda","format":"eceaa","title":"badee","description":"acdf","multipleOf":6744.988,"maximum":4020.843,"exclusiveMaximum":true,"minimum":5514.688,"exclusiveMinimum":true,"maxLength":2056,"pattern":"ddcc","maxItems":3531,"uniqueItems":true,"maxProperties":2449,"required":[],"enum":["dbaef"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"eb","format":"cfdd","title":"fe","description":"eacb","multipleOf":9324.086,"maximum":3384.514,"exclusiveMaximum":true,"minimum":8148.068,"exclusiveMinimum":true,"maxLength":272,"minLength":3673,"pattern":"deec","maxItems":7652,"minItems":6696,"uniqueItems":true,"maxProperties":1982,"minProperties":8078,"required":["bdae"],"enum":["ce"],"additionalProperties":{"$ref":"dcbbf","format":"eda","title":"ceb","description":"be","multipleOf":9698.706,"maximum":4057.317,"exclusiveMaximum":true,"minimum":7419.395,"exclusiveMinimum":true,"maxLength":4201,"pattern":"deceb","maxItems":808,"uniqueItems":true,"maxProperties":337,"required":[],"enum":["adc"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"dbdf","description":"bcecec","externalDocs":{"description":"fdbaa","url":"eefed"}}],"summary":"ccfaf","description":"dacffd","externalDocs":{"description":"af","url":"ffe"},"deprecated":true}],"send":[{"$ref":"eb","headers":{"$ref":"cad","format":"cfb","title":"cc","description":"cfaca","multipleOf":3643.655,"maximum":8680.401,"exclusiveMaximum":true,"minimum":8922.739,"exclusiveMinimum":true,"maxLength":5411,"minLength":277,"pattern":"eb","maxItems":9814,"minItems":9489,"uniqueItems":true,"maxProperties":6095,"minProperties":6889,"required":["dfca"],"enum":["edcf"],"additionalProperties":{"$ref":"efbf","format":"ecc","title":"ba","description":"bfbed","multipleOf":1540.792,"maximum":3754.606,"exclusiveMaximum":true,"minimum":8950.868,"exclusiveMinimum":true,"maxLength":8196,"pattern":"fd","maxItems":4478,"uniqueItems":true,"maxProperties":6821,"required":[],"enum":["dad"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"bafa","format":"dde","title":"eed","description":"aaf","multipleOf":848.018,"maximum":8335.138,"exclusiveMaximum":true,"minimum":2024.63,"exclusiveMinimum":true,"maxLength":6522,"minLength":8542,"pattern":"fadb","maxItems":3291,"minItems":5411,"uniqueItems":true,"maxProperties":5280,"minProperties":8375,"required":["aa"],"enum":["efee"],"additionalProperties":{"$ref":"aebc","format":"ad","title":"fdbed","description":"ddad","multipleOf":8521.845,"maximum":2948.209,"exclusiveMaximum":true,"minimum":6184.13,"exclusiveMinimum":true,"maxLength":9118,"pattern":"dfebf","maxItems":2723,"uniqueItems":true,"maxProperties":9394,"required":[],"enum":["ac"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"aefb","description":"da","externalDocs":{"description":"fbdeeb","url":"eedb"}}],"summary":"fededd","description":"bdff","externalDocs":{"description":"dfb","url":"febd"},"deprecated":true}]}`)
		v Events
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponents_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"schemas":{"deccce":{"$ref":"accfa","format":"cdfb","title":"fddbe","description":"bf","multipleOf":9854.449,"maximum":5530.484,"exclusiveMaximum":true,"minimum":326.439,"exclusiveMinimum":true,"maxLength":5226,"minLength":7252,"pattern":"bae","maxItems":8544,"minItems":2592,"uniqueItems":true,"maxProperties":8813,"minProperties":3268,"required":["afdec"],"enum":["beaade"],"additionalProperties":{"$ref":"fae","format":"acbc","title":"ff","description":"fae","multipleOf":3909.187,"maximum":2904.132,"exclusiveMaximum":true,"minimum":2144.649,"exclusiveMinimum":true,"maxLength":7393,"minLength":6192,"pattern":"ce","maxItems":8032,"minItems":9537,"uniqueItems":true,"maxProperties":2539,"minProperties":1982,"required":["efdcb"],"enum":["ebab"],"additionalProperties":true,"type":"boolean","items":{}}}},"messages":{"ececb":{"$ref":"cdc","headers":{"$ref":"deeed","format":"bbebbb","title":"bcf","description":"ad","multipleOf":6021.466,"maximum":2171.303,"exclusiveMaximum":true,"minimum":3964.627,"exclusiveMinimum":true,"maxLength":4956,"minLength":6280,"pattern":"dceaf","maxItems":4811,"minItems":419,"uniqueItems":true,"maxProperties":3398,"minProperties":494,"required":["aec"],"enum":["dadbd"],"additionalProperties":{"$ref":"edeaad","format":"ceccff","title":"ff","description":"dd","multipleOf":1562.015,"maximum":9387.192,"exclusiveMaximum":true,"minimum":6796.16,"exclusiveMinimum":true,"maxLength":8476,"pattern":"fbfbe","maxItems":9288,"uniqueItems":true,"maxProperties":4822,"required":[],"enum":["fcda"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"payload":{"$ref":"afccbe","format":"ebbe","title":"bdbdaf","description":"cdedcb","multipleOf":1306.928,"maximum":9565.204,"exclusiveMaximum":true,"minimum":5360.026,"exclusiveMinimum":true,"maxLength":7947,"minLength":9008,"pattern":"affbae","maxItems":2508,"minItems":3320,"uniqueItems":true,"maxProperties":3336,"minProperties":9535,"required":["cccfd"],"enum":["dd"],"additionalProperties":{"$ref":"bc","format":"dcad","title":"ddadff","description":"abff","multipleOf":1211.038,"maximum":3956.483,"exclusiveMaximum":true,"minimum":1908.758,"exclusiveMinimum":true,"maxLength":7717,"pattern":"da","maxItems":6558,"uniqueItems":true,"maxProperties":3795,"required":[],"enum":["baaad"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"tags":[{"name":"fdbdf","description":"ffb","externalDocs":{"description":"bbce","url":"fb"}}],"summary":"fb","description":"adb","externalDocs":{"description":"ac","url":"cecab"},"deprecated":true}},"securitySchemes":{"afeef":"adfae"},"parameters":{"cab":{"description":"dbcb","name":"dfbad","schema":{"$ref":"aa","format":"becdd","title":"bc","description":"dc","multipleOf":9810.798,"maximum":35.843,"exclusiveMaximum":true,"minimum":1559.752,"exclusiveMinimum":true,"maxLength":4034,"minLength":5397,"pattern":"fdcc","maxItems":4663,"minItems":777,"uniqueItems":true,"maxProperties":4910,"minProperties":2088,"required":["eef"],"enum":["df"],"additionalProperties":{"$ref":"edfcbd","format":"aadb","title":"adf","description":"de","multipleOf":8956.418,"maximum":8467.222,"exclusiveMaximum":true,"minimum":6014.334,"exclusiveMinimum":true,"maxLength":7931,"pattern":"bea","maxItems":6843,"uniqueItems":true,"maxProperties":6907,"required":[],"enum":["bcedff"],"additionalProperties":{},"items":{},"allOf":[],"oneOf":[],"anyOf":[],"not":{}}},"$ref":"ea"}}}`)
		v Components
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestReference_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"bbaf","cbdd":"dea"}`)
		v Reference
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestUserPassword_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"userPassword","description":"cffbb"}`)
		v UserPassword
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAPIKey_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"apiKey","in":"user","description":"dddcf"}`)
		v APIKey
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestX509_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"X509","description":"cbbaf"}`)
		v X509
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestSymmetricEncryption_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"symmetricEncryption","description":"baadd"}`)
		v SymmetricEncryption
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAsymmetricEncryption_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"asymmetricEncryption","description":"effead"}`)
		v AsymmetricEncryption
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNonBearerHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"fdb","description":"cfda","type":"http"}`)
		v NonBearerHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNonBearerHTTPSecuritySchemeNot_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","fcccbc":"fdfcfc"}`)
		v NonBearerHTTPSecuritySchemeNot
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestBearerHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","bearerFormat":"bf","type":"http","description":"baffcf"}`)
		v BearerHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAPIKeyHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"httpApiKey","name":"cf","in":"cookie","description":"eeed"}`)
		v APIKeyHTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestHTTPSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"scheme":"bearer","bearerFormat":"deccb","type":"http","description":"bccdc"}`)
		v HTTPSecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestSecurityScheme_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"asymmetricEncryption","description":"dcea"}`)
		v SecurityScheme
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsSecuritySchemesAZAZ09_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"symmetricEncryption","description":"cef"}`)
		v ComponentsSecuritySchemesAZAZ09
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestComponentsSecuritySchemes_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"fffa":"fcfbd"}`)
		v ComponentsSecuritySchemes
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}
