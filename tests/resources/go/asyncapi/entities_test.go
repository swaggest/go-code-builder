package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := Info{
		Version: "v1",
		MapOfAnythingValues: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := Info{
		Version: "v1",
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"v1"}`, string(res))
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
    "type": "apiKey",
    "in": "user",
    "description": "Provide your API key as the user and leave the password empty."
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
	assert.NoError(t, json.Unmarshal([]byte(data), &a))

	marshaled, err := json.MarshalIndent(a, "", " ")
	assert.NoError(t, err)
	assert.Equal(t, data, string(marshaled))
}
