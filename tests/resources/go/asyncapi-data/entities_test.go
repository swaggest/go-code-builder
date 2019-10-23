package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
)

func Test_MarshalUnmarshal(t *testing.T) {
	entity := MessagingReaderReads{
		Reads: []Book{
			{
				Amount:   100,
				EntityID: "ISBN123",
				Strategy: PlotStrategyFast,
			},
		},
		Country:        "US",
		ReaderID:       123,
		Week:           "2008-W05",
		SubscriptionID: 456,
	}

	encodedJSON, err := json.Marshal(entity)
	require.NoError(t, err)

	expectedJSON := []byte(`{
  "reads": [
    {
      "amount": 100,
      "entity_id": "ISBN123",
      "strategy": "fast",
      "entity_type": "book",
      "reason": "premium"
    }
  ],
  "country": "US",
  "reader_id": 123,
  "week": "2008-W05",
  "subscription_id": 456
}`)

	assertjson.Equal(t, expectedJSON, encodedJSON)
	decodedEntity := MessagingReaderReads{}
	err = json.Unmarshal(encodedJSON, &decodedEntity)
	require.NoError(t, err)
	assert.Equal(t, entity.Reads[0], decodedEntity.Reads[0])
	// assert.Equal(t, entity, decodedEntity)
	encodedJSON, err = json.Marshal(decodedEntity)
	require.NoError(t, err)
	assertjson.Equal(t, expectedJSON, encodedJSON)
}
