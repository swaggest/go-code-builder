package jsonschema_test

import (
	"encoding/json"
	"io/ioutil"
	entities "test/draft7"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/yudai/gojsondiff/formatter"
)

func TestSchema_MarshalJSON_roundtrip_asyncapi2(t *testing.T) {
	data, err := ioutil.ReadFile("../../asyncapi-2.0.0.json")
	require.NoError(t, err)

	s := entities.Schema{}
	require.NoError(t, json.Unmarshal(data, &s))

	marshaled, err := json.Marshal(s)
	require.NoError(t, err)
	assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}.Equal(t, data, marshaled)
}

func TestSchema_MarshalJSON_roundtrip_draft7(t *testing.T) {
	data, err := ioutil.ReadFile("../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json")
	require.NoError(t, err)

	s := entities.Schema{}
	require.NoError(t, json.Unmarshal(data, &s))

	marshaled, err := json.Marshal(s)
	require.NoError(t, err)
	assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}.Equal(t, data, marshaled)
}

func TestSchema_MarshalJSON_roundtrip_sample(t *testing.T) {
	data, err := ioutil.ReadFile("sample.json")
	require.NoError(t, err)

	s := entities.Schema{}
	require.NoError(t, json.Unmarshal(data, &s))

	//if s.TypeObject.Default != nil {
	//	val := (*s.TypeObject.Default).(string)
	//	println(val)
	//}

	marshaled, err := json.Marshal(s)
	require.NoError(t, err)
	assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}.Equal(t, data, marshaled)
}

func BenchmarkSchema_MarshalJSON_roundtrip_sample(b *testing.B) {
	data, err := ioutil.ReadFile("sample.json")
	require.NoError(b, err)

	s := entities.Schema{}
	require.NoError(b, json.Unmarshal(data, &s))

	marshaled, err := json.Marshal(s)
	require.NoError(b, err)
	assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}.Equal(b, data, marshaled)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &s)
		_, _ = json.Marshal(s)
	}
}

// Pointers:
// BenchmarkSchema_UnmarshalJSON-4   	    3699	    311725 ns/op	   98225 B/op	    1256 allocs/op
// BenchmarkSchema_MarshalJSON-4     	    5398	    216649 ns/op	   45012 B/op	     931 allocs/op

// Values:
// BenchmarkSchema_UnmarshalJSON-4   	    3667	    315218 ns/op	  104194 B/op	    1217 allocs/op
// BenchmarkSchema_MarshalJSON-4     	    5157	    212650 ns/op	   44995 B/op	     931 allocs/op

func BenchmarkSchema_UnmarshalJSON_map(b *testing.B) {
	data, err := ioutil.ReadFile("../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json")
	require.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()

	s := map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &s)
	}
}

func BenchmarkSchema_UnmarshalJSON(b *testing.B) {
	data, err := ioutil.ReadFile("../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json")
	require.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()

	s := entities.Schema{}

	//js := jsoniter.ConfigCompatibleWithStandardLibrary

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &s)
	}
}

func BenchmarkSchema_UnmarshalJSON_asyncapi(b *testing.B) {
	data, err := ioutil.ReadFile("../../asyncapi-2.0.0.json")
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	s := entities.Schema{}

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &s)
	}
}

func BenchmarkSchema_MarshalJSON_map(b *testing.B) {
	data, err := ioutil.ReadFile("../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json")
	require.NoError(b, err)

	s := map[string]interface{}{}
	require.NoError(b, json.Unmarshal(data, &s))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&s)
	}
}

func BenchmarkSchema_MarshalJSON(b *testing.B) {
	data, err := ioutil.ReadFile("../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json")
	require.NoError(b, err)

	s := entities.Schema{}
	require.NoError(b, json.Unmarshal(data, &s))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&s)
	}
}
