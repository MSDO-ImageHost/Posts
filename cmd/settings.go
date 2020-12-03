package main

import broker "github.com/MSDO-ImageHost/Posts/internal/broker"

var (
	_LOG_TAG string = "App:\t"

	createSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "create-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"created-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	readSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "get-single-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"single-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	readPostHistoryConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "get-history-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"history-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	readUserPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "get-user-posts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"user-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	readManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "get-many-posts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"many-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	updateOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "update-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"updated-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	deleteOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "delete-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"deleted-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}

	deleteManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf: broker.QueueConfig{Name: "delete-many-post", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf: broker.QueueConfig{Name: "test-queue" /*"deleted-many-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf: broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
	}
)
