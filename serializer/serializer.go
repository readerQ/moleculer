package serializer

import (
	"io"

	"github.com/moleculer-go/moleculer"
)

type Serializer interface {
	ReaderToPayload(io.Reader) moleculer.Payload
	BytesToPayload(*[]byte) moleculer.Payload
	PayloadToBytes(moleculer.Payload) []byte
	PayloadToString(moleculer.Payload) string
	MapToString(interface{}) string
	StringToMap(string) map[string]interface{}
	PayloadToContextMap(moleculer.Payload) map[string]interface{}
	MapToPayload(*map[string]interface{}) (moleculer.Payload, error)
}

func New(broker *moleculer.BrokerDelegates) Serializer {
	switch broker.Config.Serializer {
	case "CBOR":
		{
			return CreateCBORSerializer(broker.Logger("serializer", "cbor"))
		}
	default:
		{
			return CreateJSONSerializer(broker.Logger("serializer", "json"))
		}
	}
}
