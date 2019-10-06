package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func Test_MarshalUnmarshal(t *testing.T) {
	entity := Properties{
		MessageID: &ShortStr{
			Value: "foo",
		},
		MapOfAnything: map[string]interface{}{
			"x-whatever": "hello!",
		},
		AdditionalProperties: map[string]Property{
			"additional1": {
				Scalar: &Scalar{
					Type:  StringedTypeLongstr,
					Value: "baaar",
				},
			},
			"additional2": {
				Scalar: &Scalar{
					Type:  StringedTypeShort,
					Value: "123",
				},
			},
		},
	}

	data, err := json.Marshal(entity)
	require.NoError(t, err)
	assertjson.Equal(t,
		[]byte(`{"message-id":{"value":"foo","type":"shortstr"},"x-whatever":"hello!","additional1":{"type":"longstr","value":"baaar"},"additional2":{"type":"short","value":"123"}}`),
		data)

	unmarshaled := Properties{}
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	anotherData, err := json.Marshal(unmarshaled)
	require.NoError(t, err)
	assert.Equal(t, string(data), string(anotherData))
}
