package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestObjectOrNull_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"a":"aeff"}`)
		v ObjectOrNull
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}
