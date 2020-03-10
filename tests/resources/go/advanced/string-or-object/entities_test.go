package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestElementObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"id":"fffef","val":8556,"fdbdd":"cfc"}`)
		v ElementObject
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestElement_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"id":"cbbdbd","val":8781,"fcbfc":"cdff"}`)
		v Element
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}
