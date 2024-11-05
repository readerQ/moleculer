package serializer

import (
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/payload"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"go.mongodb.org/mongo-driver/bson"
)

type JSONPayload struct {
	result gjson.Result
	logger *log.Entry
}

func (jpayload JSONPayload) Remove(fields ...string) moleculer.Payload {
	var err error
	json := jpayload.result.Raw
	for _, item := range fields {
		json, err = sjson.Delete(json, item)
		if err != nil {
			return payload.Error("Error serializng value into JSON. error: ", err.Error())
		}
	}
	return JSONPayload{gjson.Parse(json), jpayload.logger}
}

func (jpayload JSONPayload) AddItem(value interface{}) moleculer.Payload {
	if !jpayload.IsArray() {
		return payload.Error("payload.AddItem can only deal with lists/arrays.")
	}
	arr := jpayload.Array()
	arr = append(arr, payload.New(value))
	return payload.New(arr)
}

func (jpayload JSONPayload) Add(field string, value interface{}) moleculer.Payload {
	if !jpayload.IsMap() {
		return payload.Error("payload.Add can only deal with map payloads.")
	}
	var err error
	json := jpayload.result.Raw
	json, err = sjson.Set(json, field, value)
	if err != nil {
		return payload.Error("Error serializng value into JSON. error: ", err.Error())
	}
	return JSONPayload{gjson.Parse(json), jpayload.logger}
}

func (jpayload JSONPayload) AddMany(toAdd map[string]interface{}) moleculer.Payload {
	if !jpayload.IsMap() {
		return payload.Error("payload.Add can only deal with map payloads.")
	}
	var err error
	json := jpayload.result.Raw
	for key, value := range toAdd {
		json, err = sjson.Set(json, key, value)
		if err != nil {
			return payload.Error("Error serializng value into JSON. error: ", err.Error())
		}
	}
	return JSONPayload{gjson.Parse(json), jpayload.logger}
}

func (payload JSONPayload) MapArray() []map[string]interface{} {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]map[string]interface{}, len(source))
		for index, item := range source {
			array[index] = resultToMap(item, true)
		}
		return array
	}
	return nil
}

func (payload JSONPayload) ValueArray() []interface{} {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]interface{}, len(source))
		for index, item := range source {
			array[index] = item.Value()
		}
		return array
	}
	return nil
}

func (payload JSONPayload) IntArray() []int {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]int, len(source))
		for index, item := range source {
			array[index] = int(item.Int())
		}
		return array
	}
	return nil
}

func (payload JSONPayload) Int64Array() []int64 {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]int64, len(source))
		for index, item := range source {
			array[index] = item.Int()
		}
		return array
	}
	return nil
}

func (payload JSONPayload) UintArray() []uint64 {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]uint64, len(source))
		for index, item := range source {
			array[index] = item.Uint()
		}
		return array
	}
	return nil
}

func (payload JSONPayload) Float32Array() []float32 {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]float32, len(source))
		for index, item := range source {
			array[index] = float32(item.Float())
		}
		return array
	}
	return nil
}

func (payload JSONPayload) FloatArray() []float64 {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]float64, len(source))
		for index, item := range source {
			array[index] = item.Float()
		}
		return array
	}
	return nil
}

func (jp JSONPayload) BsonArray() bson.A {
	if jp.IsArray() {
		ba := make(bson.A, jp.Len())
		for index, value := range jp.Array() {
			if value.IsMap() {
				ba[index] = value.Bson()
			} else if value.IsArray() {
				ba[index] = value.BsonArray()
			} else {
				ba[index] = value.Value()
			}
		}
		return ba
	}
	return nil
}

func (jp JSONPayload) Bson() bson.M {
	if jp.IsMap() {
		bm := bson.M{}
		for key, value := range jp.Map() {
			if value.IsMap() {
				bm[key] = value.Bson()
			} else if value.IsArray() {
				bm[key] = value.BsonArray()
			} else {
				bm[key] = value.Value()
			}
		}
		return bm
	}
	return nil
}

func (payload JSONPayload) BoolArray() []bool {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]bool, len(source))
		for index, item := range source {
			array[index] = item.Bool()
		}
		return array
	}
	return nil
}

func (payload JSONPayload) ByteArray() []byte {
	return []byte(payload.result.Raw)
}

func (payload JSONPayload) TimeArray() []time.Time {
	if source := payload.result.Array(); source != nil {
		array := make([]time.Time, len(source))
		for index, item := range source {
			array[index] = item.Time()
		}
		return array
	}
	return nil
}

func (payload JSONPayload) At(index int) moleculer.Payload {
	if payload.IsArray() {
		source := payload.result.Array()
		if index >= 0 && index < len(source) {
			item := source[index]
			return JSONPayload{item, payload.logger}
		}
	}
	return nil
}

func (payload JSONPayload) Array() []moleculer.Payload {
	if payload.IsArray() {
		source := payload.result.Array()
		array := make([]moleculer.Payload, len(source))
		for index, item := range source {
			array[index] = JSONPayload{item, payload.logger}
		}
		return array
	}
	return nil
}

func (p JSONPayload) Sort(field string) moleculer.Payload {
	if !p.IsArray() {
		return p
	}
	ps := &payload.Sortable{field, p.Array()}
	sort.Sort(ps)
	return ps.Payload()
}

func (payload JSONPayload) IsArray() bool {
	return payload.result.IsArray()
}

func (payload JSONPayload) IsMap() bool {
	return payload.result.IsObject()
}

func (payload JSONPayload) ForEach(iterator func(key interface{}, value moleculer.Payload) bool) {
	payload.result.ForEach(func(key, value gjson.Result) bool {
		return iterator(key.Value(), &JSONPayload{value, payload.logger})
	})
}

func (p JSONPayload) MapOver(transform func(in moleculer.Payload) moleculer.Payload) moleculer.Payload {
	if p.IsArray() {
		list := []moleculer.Payload{}
		for _, value := range p.Array() {
			list = append(list, transform(value))
		}
		return payload.New(list)
	} else {
		return payload.Error("payload.MapOver can only deal with array payloads.")
	}
}

func (payload JSONPayload) Bool() bool {
	return payload.result.Bool()
}

func (payload JSONPayload) Float() float64 {
	return payload.result.Float()
}

func (payload JSONPayload) Float32() float32 {
	return float32(payload.result.Float())
}

func (payload JSONPayload) IsError() bool {
	return payload.IsMap() && payload.Get("error").Exists()
}

func (payload JSONPayload) Error() error {
	if payload.IsError() {
		return errors.New(payload.Get("error").String())
	}
	return nil
}

func (p JSONPayload) ErrorPayload() moleculer.Payload {
	if p.IsError() {
		return p
	}
	return nil
}

func orderedKeys(m map[string]moleculer.Payload) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

func (jp JSONPayload) StringIdented(ident string) string {
	return jp.String()
}

func (jp JSONPayload) String() string {
	if jp.IsMap() {
		ident := "  "
		m := jp.Map()

		out := "(len=" + strconv.Itoa(len(m)) + ") {\n"
		for _, key := range orderedKeys(m) {
			out = out + ident + `"` + key + `": ` + m[key].String() + "," + "\n"
		}
		if len(m) == 0 {
			out = out + "\n"
		}
		out = out + "}"
		return out
	}
	return jp.result.String()
}

func (payload JSONPayload) RawMap() map[string]interface{} {
	mapValue, ok := payload.result.Value().(map[string]interface{})
	if !ok {
		payload.logger.Warn("RawMap() Could not convert result.Value() into a map[string]interface{} - result: ", payload.result)
		return nil
	}
	return mapValue
}

func (payload JSONPayload) Map() map[string]moleculer.Payload {
	if source := payload.result.Map(); source != nil {
		newMap := make(map[string]moleculer.Payload, len(source))
		for key, item := range source {
			newMap[key] = &JSONPayload{item, payload.logger}
		}
		return newMap
	}
	return nil
}
