package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestProperties_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"refOrSchema":{"acbbff":"acdefb"},"noTypeWithExamples":{"firstName":"ad","lastName":"eadc","age":5245,"gender":"aecfe","cddbd":"accdf"},"noTypeWithExample":{"firstName":"fd","lastName":"ad","age":8764,"adeadb":"cda"},"address":{"addressStripped":"bdc","address1":"dbfee","address2":"fcceaf","city":"fedacd","ab":"edc"},"headers":{"type":"table","value":{"bddad":{"type":"table","value":{"bffd":{"type":"bit","value":"ecd"}}}}},"content-type":{"type":"shortstr","value":"bb"},"content-encoding":{"type":"shortstr","value":"db"},"delivery-mode":{"type":"octet","value":"aeca"},"priority":{"type":"octet","value":"bfff"},"correlation-id":{"type":"shortstr","value":"bceddc"},"reply-to":{"type":"shortstr","value":"decad"},"expiration":{"type":"shortstr","value":"ddbcfa"},"message-id":{"type":"shortstr","value":"acabdf"},"timestamp":{"type":"timestamp","value":"deae"},"type":{"type":"shortstr","value":"beb"},"user-id":{"type":"shortstr","value":"fd"},"app-id":{"type":"shortstr","value":"eaceae"},"cc":{"type":"table","value":{"adfae":{"type":"timestamp","value":"eaabca"}}}}`)
		v Properties
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestReference_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"$ref":"cd","cad":"faecac"}`)
		v Reference
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestRefOrSchema_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"ba":"fecc"}`)
		v RefOrSchema
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNoTypeWithExamples_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"firstName":"dcf","lastName":"baf","age":4725,"gender":"dad","cd":"fe"}`)
		v NoTypeWithExamples
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestNoTypeWithExample_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"firstName":"de","lastName":"befdf","age":3963,"cbafb":"fbb"}`)
		v NoTypeWithExample
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestAddress_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"addressStripped":"cdf","address1":"caf","address2":"eeacab","city":"bd","db":"ebae"}`)
		v Address
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"table","value":{"fcbcf":{"type":"longstr","value":"ebbeea"}}}`)
		v Table
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestScalar_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"shortstr","value":"fb"}`)
		v Scalar
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestProperty_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"longlong","value":"edaac"}`)
		v Property
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestShortStr_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"shortstr","value":"cba"}`)
		v ShortStr
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestPropertyOctet_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"octet","value":"df"}`)
		v PropertyOctet
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestTimestamp_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"type":"timestamp","value":"bed"}`)
		v Timestamp
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)

	require.NoError(t, err)
	assertjson.Equal(t, jsonValue, marshaled)
}
