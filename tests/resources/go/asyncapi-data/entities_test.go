package entities

import (
	"encoding/json"
	"reflect"
	"testing"
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
	if err != nil {
		t.Fatal("unexpected error: " + err.Error())
	}
	if string(j) != `{"reads":[{"amount":100,"entity_id":"ISBN123","strategy":"fast","entity_type":"book","reason":"premium"}],"country":"US","reader_id":123,"week":"2008-W05","subscription_id":456}` {
		t.Fatal("unexpected marshal" + string(j))
	}
	decodedEntity := MessagingReaderReads{}
	err = json.Unmarshal(j, &decodedEntity)
	if err != nil {
		t.Fatal("unexpected error: " + err.Error())
	}
	if !reflect.DeepEqual(entity, decodedEntity) {
		t.Fatalf("not equal: %+v %+v", entity, decodedEntity)
	}
}
