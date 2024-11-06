package main

import (
	"fmt"
	"time"

	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"github.com/moleculer-go/moleculer/payload"
)

var mathService = moleculer.ServiceSchema{
	Name: "math",
	Actions: []moleculer.Action{
		{
			Name: "add",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				return params.Get("a").Int() + params.Get("b").Int()
			},
		},
	},
}

func main() {
	var bkr = broker.New(&moleculer.Config{
		Namespace:   "dev",
		LogLevel:    "info",
		Transporter: "nats://192.168.128.101:34222",
		Serializer:  "CBOR",
	})

	bkr.Publish(mathService)
	bkr.Start()
	fmt.Println(bkr.IsStarted())
	defer bkr.Stop()

	for i := 0; i < 100; i++ {

		t := <-time.After(time.Microsecond * 20)

		fmt.Println(bkr.IsStarted(), t)

		result := <-bkr.Call("vendors.get", payload.New("dalimo"))

		if result.Error() != nil {
			continue
		}
		slug := result.Get("slug").String()
		if slug != "" {
			break

		}

		fmt.Println(result)

	}

	fmt.Println(time.Now())

	for i := 0; i < 1000; i++ {
		result2 := <-bkr.Call("vendors.get", payload.New("dalimo"))
		if !result2.Exists() {
			panic("2")
		}
	}

	fmt.Println(time.Now())

	// for {
	// 	select {
	// 	case t := <-time.After(time.Second * 5):
	// 		{
	// 			fmt.Println(t)

	// 			result := <-bkr.Call("vendors.get", payload.New("dalimo"))

	// 			fmt.Println(result)

	// 		}
	// 	}
	// }

	kill := make(chan interface{}, 0)
	<-kill

}
