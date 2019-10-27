package entities

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := Info{
		Version: "v1",
		MapOfAnything: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := Info{
		Version: "v1",
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"version":"v1"}`, string(res))
}

func TestAsyncAPI_MarshalJSON(t *testing.T) {
	data, err := ioutil.ReadFile("streetlights.json")
	require.NoError(t, err)

	var a AsyncAPI

	require.NoError(t, json.Unmarshal(data, &a))

	marshaled, err := json.Marshal(a)
	require.NoError(t, err)
	assertjson.Equal(t, data, marshaled)
}
