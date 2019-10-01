package entities

import (
	"encoding/json"
	"reflect"
	"testing"

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

	j, err := json.Marshal(entity)
	require.NoError(t, err)
	assertjson.Equal(
		t,
		[]byte(`{"reads":[{"amount":100,"entity_id":"ISBN123","strategy":"fast","entity_type":"book","reason":"premium"}],"country":"US","reader_id":123,"week":"2008-W05","subscription_id":456}`),
		j)
	decodedEntity := MessagingReaderReads{}
	err = json.Unmarshal(j, &decodedEntity)
	require.NoError(t, err)
	if !reflect.DeepEqual(entity, decodedEntity) {
		t.Fatalf("not equal: %+v %+v", entity, decodedEntity)
	}
}
