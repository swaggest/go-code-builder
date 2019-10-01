package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
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
	require.NoError(t, err)
	assertjson.Equal(t,
		[]byte(`{"asyncapi":"1.2.0","components":{"securitySchemes":{"foo":{"in":"user","type":"apiKey"}}},"x-whatever":"hello!"}`),
		data)

	unmarshaled := AsyncAPI{}
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	anotherData, err := json.Marshal(unmarshaled)
	require.NoError(t, err)
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

func TestAsyncAPI_MarshalJSON(t *testing.T) {
	data := `{
 "asyncapi": "1.2.0",
 "info": {
  "title": "Streetlights API",
  "version": "1.0.0",
  "description": "The Smartylighting Streetlights API allows you to remotely manage the city lights.\n\n### Check out its awesome features:\n\n* Turn a specific streetlight on/off 🌃\n* Dim a specific streetlight 😎\n* Receive real-time information about environmental lighting conditions 📈\n",
  "license": {
   "name": "Apache 2.0",
   "url": "https://www.apache.org/licenses/LICENSE-2.0"
  }
 },
 "baseTopic": "smartylighting.streetlights.1.0",
 "servers": [
  {
   "url": "api.streetlights.smartylighting.com:{port}",
   "description": "Test broker",
   "scheme": "mqtt",
   "variables": {
    "port": {
     "enum": [
      "1883",
      "8883"
     ],
     "default": "1883",
     "description": "Secure connection (TLS) is available through port 8883."
    }
   }
  }
 ],
 "topics": {
  "action.{streetlightId}.dim": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/dimLight"
   }
  },
  "action.{streetlightId}.turn.off": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/turnOnOff"
   }
  },
  "action.{streetlightId}.turn.on": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/turnOnOff"
   }
  },
  "event.{streetlightId}.lighting.measured": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "publish": {
    "$ref": "#/components/messages/lightMeasured"
   }
  }
 },
 "components": {
  "schemas": {
   "dimLightPayload": {
    "properties": {
     "percentage": {
      "description": "Percentage to which the light should be dimmed to.",
      "maximum": 100,
      "minimum": 0,
      "type": "integer"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   },
   "lightMeasuredPayload": {
    "properties": {
     "lumens": {
      "description": "Light intensity measured in lumens.",
      "minimum": 0,
      "type": "integer"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   },
   "sentAt": {
    "description": "Date and time when the message was sent.",
    "format": "date-time",
    "type": "string"
   },
   "turnOnOffPayload": {
    "properties": {
     "command": {
      "description": "Whether to turn on or off the light.",
      "enum": [
       true,
       false
      ],
      "type": "string"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   }
  },
  "messages": {
   "dimLight": {
    "payload": {
     "$ref": "#/components/schemas/dimLightPayload"
    },
    "summary": "Command a particular streetlight to dim the lights."
   },
   "lightMeasured": {
    "payload": {
     "$ref": "#/components/schemas/lightMeasuredPayload"
    },
    "summary": "Inform about environmental lighting conditions for a particular streetlight."
   },
   "turnOnOff": {
    "payload": {
     "$ref": "#/components/schemas/turnOnOffPayload"
    },
    "summary": "Command a particular streetlight to turn the lights on or off."
   }
  },
  "securitySchemes": {
   "apiKey": {
    "in": "user",
    "description": "Provide your API key as the user and leave the password empty.",
    "type": "apiKey"
   }
  },
  "parameters": {
   "streetlightId": {
    "description": "The ID of the streetlight.",
    "name": "streetlightId",
    "schema": {
     "type": "string"
    }
   }
  }
 },
 "security": [
  {
   "apiKey": []
  }
 ]
}`

	var a AsyncAPI
	require.NoError(t, json.Unmarshal([]byte(data), &a))

	marshaled, err := json.MarshalIndent(a, "", " ")
	require.NoError(t, err)

	assertjson.Equal(t, []byte(data), marshaled)
}
