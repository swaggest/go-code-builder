package entities

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/yudai/gojsondiff/formatter"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := Info{
		Title:   "Foo",
		Version: "v1",
		MapOfAnything: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"title":"Foo","version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := Info{
		Title:   "Foo",
		Version: "v1",
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"title":"Foo","version":"v1"}`, string(res))
}

func TestAsyncAPI_MarshalJSON(t *testing.T) {
	data, err := ioutil.ReadFile("streetlights.json")
	require.NoError(t, err)

	var a AsyncAPI

	require.NoError(t, json.Unmarshal(data, &a))

	marshaled, err := json.Marshal(a)
	require.NoError(t, err)

	aj := assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}

	aj.Equal(t, data, marshaled)
}

// BenchmarkAsyncAPI_MarshalJSON-4   	     495	   2523002 ns/op	  283862 B/op	    4447 allocs/op

// mustUnmarshal inlined
// BenchmarkAsyncAPI_MarshalJSON-4   	     391	   2859061 ns/op	  282034 B/op	    4355 allocs/op

// mayUnmarshal inlined, constValues removed from mayUnmarshal
// BenchmarkAsyncAPI_MarshalJSON-4   	     622	   2112988 ns/op	  223881 B/op	    3388 allocs/op

// extracting reusable ignoreKeys slices
// BenchmarkAsyncAPI_MarshalJSON-4   	     404	   2951737 ns/op	  217140 B/op	    3302 allocs/op

// inlining const and pattern properties
// BenchmarkAsyncAPI_MarshalJSON-4   	     374	   2824172 ns/op	  241703 B/op	    3853 allocs/op

// inlining additional properties, retiring unionmap
// BenchmarkAsyncAPI_MarshalJSON-4   	     500	   2718313 ns/op	  200211 B/op	    2982 allocs/op

func BenchmarkAsyncAPI_MarshalJSON(b *testing.B) {
	data, err := ioutil.ReadFile("streetlights.json")
	require.NoError(b, err)

	var a AsyncAPI

	require.NoError(b, json.Unmarshal(data, &a))

	marshaled, err := json.Marshal(a)
	require.NoError(b, err)

	aj := assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}

	aj.Equal(b, data, marshaled)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &a)
	}
}
