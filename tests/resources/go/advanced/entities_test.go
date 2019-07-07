package entities

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalUnmarshal(t *testing.T) {
	entity := Properties{
		MessageID: &ShortStr{
			Value: "foo",
		},
		MapOfAnythingValues: map[string]interface{}{
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
	assert.NoError(t, err)
	assert.Equal(t,
		`{"message-id":{"value":"foo","type":"shortstr"},"x-whatever":"hello!","additional1":{"type":"longstr","value":"baaar"},"additional2":{"type":"short","value":"123"}}`,
		string(data))

	unmarshaled := Properties{}
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	anotherData, err := json.Marshal(unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, string(data), string(anotherData))
}
