package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestRequiredNullable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"a":{"b":"aeff"},"an":{"b":"effefd"},"axn":{"b":"dd"},"b":"cfc"}`)
		v RequiredNullable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestDef_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"b":"cbbdbd"}`)
		v Def
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestDefNullable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"b":"efcbfc"}`)
		v DefNullable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}

func TestDefXNullable_MarshalJSON_roundtrip(t *testing.T) {
	var (
		jsonValue = []byte(`{"b":"cdff"}`)
		v DefXNullable
	)

	require.NoError(t, json.Unmarshal(jsonValue, &v))

	marshaled, err := json.Marshal(v)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(marshaled, &v))
	assertjson.Equal(t, jsonValue, marshaled)
}
