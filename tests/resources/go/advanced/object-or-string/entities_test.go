package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func Test_MarshalUnmarshalRoundtrip(t *testing.T) {
	jsonPayload := []byte(`[
  "abc",
  {
    "id": "def",
    "val": 1
  }
]
`)

	var v []ElementConditional
	err := json.Unmarshal(jsonPayload, &v)
	require.NoError(t, err)

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	assertjson.Equal(t, jsonPayload, marshaled)
}
