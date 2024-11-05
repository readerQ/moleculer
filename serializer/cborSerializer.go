package serializer

import (
	"io"

	"github.com/fxamacker/cbor/v2"
	"github.com/moleculer-go/moleculer"
	log "github.com/sirupsen/logrus"
)

type CBORSerializer struct {
	logger *log.Entry
}

func CreateCBORSerializer(logger *log.Entry) CBORSerializer {
	return CBORSerializer{logger}
}

func (serializer CBORSerializer) ReaderToPayload(io.Reader) moleculer.Payload {
	// TODO implement
	panic("not implemented")
}

func (serializer CBORSerializer) BytesToPayload(data *[]byte) moleculer.Payload {

	//fmt.Println(string(*data))
	b := map[string]interface{}{}

	err := cbor.Unmarshal(*data, &b)
	if err != nil {
		serializer.logger.Error(err)
	}

	c := deepCopyMap(b)
	return CBORPayload{
		data:   c,
		logger: serializer.logger,
	}

	// TODO implement
	panic("not implemented")

}

func deepCopyMap(src interface{}) interface{} {

	if src == nil {
		return nil
	}

	switch src.(type) {
	case map[string]interface{}:
		{
			dst := map[string]interface{}{}

			for k, v := range src.(map[string]interface{}) {
				dst[k] = deepCopyMap(v)
			}

			return dst

		}
	case map[interface{}]interface{}:
		{
			dst := map[string]interface{}{}

			for ki, v := range src.(map[interface{}]interface{}) {
				k, ok := ki.(string)
				if !ok {
					//	fmt.Println("deep copy error")
					continue

				}
				dst[k] = deepCopyMap(v)
			}

			return dst

		}

	case []interface{}:
		{

			dst := []interface{}{}

			for _, v := range src.([]interface{}) {
				dst = append(dst, deepCopyMap(v))
			}

			return dst

		}
	case bool, int64, uint64, float64, string:
		{
			return src
		}
	default:
		{
			return src
		}

	}
}

func (serializer CBORSerializer) PayloadToBytes(payload moleculer.Payload) []byte {

	b, err := cbor.Marshal(payload.Value())
	if err != nil {
		serializer.logger.Error(err)
	}
	return b
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) PayloadToString(payload moleculer.Payload) string {
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) MapToString(interface{}) string {
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) StringToMap(string) map[string]interface{} {
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) PayloadToContextMap(moleculer.Payload) map[string]interface{} {
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) MapToPayload(mapValue *map[string]interface{}) (moleculer.Payload, error) {
	// TODO implement

	return CBORPayload{
		data: mapValue,
	}, nil

}
