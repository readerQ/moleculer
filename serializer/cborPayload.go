package serializer

import (
	"errors"
	"fmt"
	"time"

	"github.com/moleculer-go/moleculer"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type CBORPayload struct {
	data   interface{}
	logger *log.Entry
}

func (pl CBORPayload) First() moleculer.Payload {
	// TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Sort(field string) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Remove(fields ...string) moleculer.Payload {

	tmp := pl.data
	mapa, ok := tmp.(*map[string]interface{})
	if !ok {
		return CBORPayload{data: fmt.Errorf("remove error")}
	}
	for _, key := range fields {
		delete(*mapa, key)
	}
	return CBORPayload{data: mapa, logger: pl.logger}
}

func (pl CBORPayload) AddItem(value interface{}) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Add(field string, value interface{}) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) AddMany(map[string]interface{}) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func i2m(src map[interface{}]interface{}) map[string]interface{} {
	smap := map[string]interface{}{}

	for k, v := range src {
		ks, ok3 := k.(string)
		if !ok3 {
			return map[string]interface{}{}
		}
		smap[ks] = v
	}

	return smap
}

func (pl CBORPayload) MapArray() []map[string]interface{} {

	data1, ok1 := pl.data.([]map[string]interface{})
	if ok1 {
		return data1
	}

	data, ok := pl.data.([]interface{})
	if !ok {
		return nil
	}

	result := []map[string]interface{}{}

	for _, val := range data {

		item0, ok0 := val.(map[string]interface{})

		if ok0 {

			result = append(result, item0)
			continue
		}

		item, ok1 := val.(map[interface{}]interface{})

		if ok1 {
			smap := i2m(item)
			result = append(result, smap)
		}
	}

	return result
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) RawMap() map[string]interface{} {
	//TODO implement

	raw, ok := pl.data.(map[string]interface{})

	if ok {
		return raw
	}

	return nil

}

func (pl CBORPayload) Map() map[string]moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Exists() bool {
	// TODO check logic
	return pl.data != nil
}

func (payload CBORPayload) IsError() bool {
	return payload.IsMap() && payload.Get("error").Exists()
}

func (payload CBORPayload) Error() error {
	if payload.IsError() {
		return errors.New(payload.Get("error").String())
	}
	return nil
}

func (pl CBORPayload) ErrorPayload() moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Value() interface{} {

	return pl.data

}

func (pl CBORPayload) ValueArray() []interface{} {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Int() int {

	switch pl.data.(type) {
	case int:
		{
			return int(pl.data.(int))
		}
	case float64:
		{
			return int(pl.data.(float64))
		}
	}

	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) IntArray() []int {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Int64() int64 {

	u, o := pl.data.(uint64)
	if o {
		return int64(u)
	}
	return pl.data.(int64)
}

func (pl CBORPayload) Int64Array() []int64 {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Uint() uint64 {
	s, o := pl.data.(int64)
	if o {
		return uint64(s)
	}
	return pl.data.(uint64)

}

func (pl CBORPayload) UintArray() []uint64 {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Float32() float32 {
	return pl.data.(float32)
}

func (pl CBORPayload) Float32Array() []float32 {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Float() float64 {
	return pl.data.(float64)
}

func (pl CBORPayload) FloatArray() []float64 {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) String() string {

	r, _ := pl.data.(string)
	return r

}

func (pl CBORPayload) StringArray() []string {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Bool() bool {

	b, ok := pl.data.(bool)
	return ok && b

}

func (pl CBORPayload) BoolArray() []bool {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) ByteArray() []byte {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Time() time.Time {
	return pl.data.(time.Time)
}

func (pl CBORPayload) TimeArray() []time.Time {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Array() []moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) At(index int) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Len() int {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) Get(path string, defaultValue ...interface{}) moleculer.Payload {

	switch pl.data.(type) {
	case nil:
	case map[string]interface{}:
		{
			m := pl.data.(map[string]interface{})
			v := m[path]

			switch t := v.(type) {
			case string:
				{
					break
				}
			case map[interface{}]interface{}:
				{
					smap := i2m(v.(map[interface{}]interface{}))
					return CBORPayload{
						data:   smap,
						logger: pl.logger,
					}

				}
			default:
				{
					pl.logger.Debug("Get", t, path)
				}
			}

			return CBORPayload{
				data:   v,
				logger: pl.logger,
			}
		}

	}

	//TODO implement
	panic("not implemented")
}

// Only return a payload containing only the field specified
func (pl CBORPayload) Only(path string) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) IsArray() bool {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) IsMap() bool {
	_, ok := pl.data.(map[string]interface{})
	return ok
}

func (pl CBORPayload) ForEach(iterator func(key interface{}, value moleculer.Payload) bool) {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) MapOver(tranform func(in moleculer.Payload) moleculer.Payload) moleculer.Payload {
	//TODO implement
	panic("not implemented")
}

func (pl CBORPayload) BsonArray() bson.A {
	if pl.IsArray() {
		ba := make(bson.A, pl.Len())
		for index, value := range pl.Array() {
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

func (pl CBORPayload) Bson() bson.M {
	if pl.IsMap() {
		bm := bson.M{}
		for key, value := range pl.Map() {
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
