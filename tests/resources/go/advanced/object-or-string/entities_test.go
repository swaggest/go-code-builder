package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestElementObject_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"id":"aeff","val":8686,"ffefd":"dd"}`)
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
		jsonValue = []byte(`{"id":"cfc","val":8463,"bbdb":"fefc"}`)
		v Element
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}
