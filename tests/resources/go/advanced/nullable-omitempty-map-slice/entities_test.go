package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestRequiredNullable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"slice":["aeff"],"map":{"effefd":"dd"}}`)
		v RequiredNullable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}
