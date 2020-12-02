package main

import broker "github.com/MSDO-ImageHost/Posts/internal/broker"

var _LOG_TAG string = "App:\t"

var postCreationHandleConf broker.HandleConfig = broker.HandleConfig{
	SubQueueConf: broker.QueueConfig{Name: "create-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
	PubQueueConf: broker.QueueConfig{Name: "confirm-post-creation", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
	ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
}
