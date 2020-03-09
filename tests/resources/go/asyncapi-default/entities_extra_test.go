package entities

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func Test_MarshalUnmarshal(t *testing.T) {
	entity := AsyncAPI{
		Asyncapi: AsyncAPIAsyncapi120,
		Info: Info{
			Title:   "Foo",
			Version: "v1",
		},
		Components: &Components{
			SecuritySchemes: &ComponentsSecuritySchemes{
				MapOfComponentsSecuritySchemesAZAZ09Values: map[string]ComponentsSecuritySchemesAZAZ09{
					"foo": {
						SecurityScheme: &SecurityScheme{
							APIKey: &APIKey{
								In: APIKeyInUser,
							},
						},
					},
				},
			},
		},
		MapOfAnything: map[string]interface{}{
			"x-whatever": "hello!",
		},
	}

	data, err := json.Marshal(entity)
	require.NoError(t, err)
	assertjson.Equal(t,
		[]byte(`{
  "asyncapi": "1.2.0",
  "info": {
    "title": "Foo",
    "version": "v1"
  },
  "components": {
    "securitySchemes": {
      "foo": {
        "in": "user",
        "type": "apiKey"
      }
    }
  },
  "x-whatever": "hello!"
}`),
		data)

	unmarshaled := AsyncAPI{}
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	anotherData, err := json.Marshal(unmarshaled)
	require.NoError(t, err)
	assertjson.Equal(t, data, anotherData)
}

func Benchmark_Marshal(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entity := AsyncAPI{
			Asyncapi: AsyncAPIAsyncapi120,
			Components: &Components{
				SecuritySchemes: &ComponentsSecuritySchemes{
					MapOfComponentsSecuritySchemesAZAZ09Values: map[string]ComponentsSecuritySchemesAZAZ09{
						"foo": {
							SecurityScheme: &SecurityScheme{
								APIKey: &APIKey{
									In: APIKeyInUser,
								},
							},
						},
					},
				},
			},
			MapOfAnything: map[string]interface{}{
				"x-whatever": "hello!",
			},
		}

		_, _ = json.Marshal(entity)
	}
}

func Benchmark_Unmarshal(b *testing.B) {
	data := []byte(`{
  "asyncapi": "1.2.0",
  "components": {
    "securitySchemes": {
      "foo": {
        "in": "user",
        "type": "apiKey"
      }
    }
  },
  "x-whatever": "hello!"
}`)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		unmarshaled := AsyncAPI{}
		_ = json.Unmarshal(data, &unmarshaled)
	}
}

func TestAsyncAPI_MarshalJSON(t *testing.T) {
	data, err := ioutil.ReadFile("../asyncapi/streetlights.json")
	require.NoError(t, err)

	var a AsyncAPI

	require.NoError(t, json.Unmarshal(data, &a))

	marshaled, err := json.MarshalIndent(a, "", " ")
	require.NoError(t, err)

	assertjson.Equal(t, data, marshaled)
}
