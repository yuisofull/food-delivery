package main

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub/localpubsub"
	"log"
	"time"
)

func main() {
	localPb := localpubsub.NewPubsub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPb.Subscribe(context.Background(), topic)
	sub2, _ := localPb.Subscribe(context.Background(), topic)

	if err := localPb.Publish(context.Background(), topic, pubsub.NewMessage(1)); err != nil {
		return
	}

	if err := localPb.Publish(context.Background(), topic, pubsub.NewMessage(2)); err != nil {
		return
	}

	go func() {
		for {
			log.Println("Sub 1: ", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sub 2: ", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()
	//close2()
	//
	if err := localPb.Publish(context.Background(), topic, pubsub.NewMessage(3)); err != nil {
		return
	}

	time.Sleep(time.Second * 2)

}
