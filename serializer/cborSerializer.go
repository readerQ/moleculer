package serializer

import (
	"bytes"
	"fmt"
	"io"

	"github.com/fxamacker/cbor/v2"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/payload"
	log "github.com/sirupsen/logrus"
)

type CBORSerializer struct {
	logger *log.Entry
}

func CreateCBORSerializer(logger *log.Entry) CBORSerializer {
	return CBORSerializer{logger}
}

func (serializer CBORSerializer) ReaderToPayload(r io.Reader) moleculer.Payload {
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	data := buf.Bytes()

	b, err := bytesToMap(data)
	if err != nil {
		return payload.New(fmt.Errorf("cbor unmarshal error"))
	}

	payload := CBORPayload{data: b, logger: serializer.logger}
	return payload
}

func bytesToMap(data []byte) (map[string]interface{}, error) {

	b := map[string]interface{}{}

	err := cbor.Unmarshal(data, &b)
	return b, err
}

func (serializer CBORSerializer) BytesToPayload(data *[]byte) moleculer.Payload {

	b, err := bytesToMap(*data)
	if err != nil {
		return payload.New(fmt.Errorf("cbor unmarshal error"))
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

}

func (serializer CBORSerializer) PayloadToString(payload moleculer.Payload) string {

	return string(serializer.PayloadToBytes(payload))

}

func (serializer CBORSerializer) MapToString(interface{}) string {
	// TODO implement
	panic("not implemented")

}

func (serializer CBORSerializer) StringToMap(string) map[string]interface{} {
	// TODO implement
	panic("not implemented")

}
func (serializer CBORSerializer) cleanContextMap(values map[string]interface{}) map[string]interface{} {
	if values["level"] != nil {

		float, ok := values["level"].(float64)
		if ok {
			values["level"] = int(float)
		} else {
			values["level"] = int(values["level"].(uint64))
		}

	}
	if values["timeout"] != nil {
		float, ok := values["timeout"].(float64)
		if ok {
			values["timeout"] = int(float)
		} else {
			values["timeout"] = int(values["timeout"].(uint64))
		}

	}
	return values
}

func (serializer CBORSerializer) PayloadToContextMap(payload moleculer.Payload) map[string]interface{} {
	return serializer.cleanContextMap(payload.RawMap())

}

func (serializer CBORSerializer) MapToPayload(mapValue *map[string]interface{}) (moleculer.Payload, error) {
	// TODO implement

	return CBORPayload{
		data: mapValue,
	}, nil

}
