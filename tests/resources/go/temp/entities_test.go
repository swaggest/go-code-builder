package entities

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalUnmarshal(t *testing.T) {
	entity := AsyncAPI{
		Asyncapi: Asyncapi120,
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
		MapOfAnythingValues: map[string]interface{}{
			"x-whatever": "hello!",
		},
	}

	data, err := json.Marshal(entity)
	assert.NoError(t, err)
	assert.Equal(t,
		`{"asyncapi":"1.2.0","components":{"securitySchemes":{"foo":{"in":"user","type":"apiKey"}}},"x-whatever":"hello!"}`,
		string(data))

	unmarshaled := AsyncAPI{}
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	anotherData, err := json.Marshal(unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, string(data), string(anotherData))
}

func Benchmark_Marshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entity := AsyncAPI{
			Asyncapi: Asyncapi120,
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
			MapOfAnythingValues: map[string]interface{}{
				"x-whatever": "hello!",
			},
		}

		_, _ = json.Marshal(entity)
	}
}

func Benchmark_Unmarshal(b *testing.B) {
	data := []byte(`{"asyncapi":"1.2.0","components":{"securitySchemes":{"foo":{"in":"user","type":"apiKey"}}},"x-whatever":"hello!"}`)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unmarshaled := AsyncAPI{}
		_ = json.Unmarshal(data, &unmarshaled)
	}
}
