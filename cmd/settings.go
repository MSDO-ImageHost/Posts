package main

import (
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	"github.com/beevik/guid"
)

var (
	_LOG_TAG   string = "App:\t"
	InstanceID        = guid.New()

	// Exchanges
	rapidExchange = broker.ExchangeConfig{Name: "rapid", Kind: "direct", Durable: false, AutoDelete: false, Internal: false, NoWait: false, Args: nil}

	// Request queues
	createOnePost = broker.QueueConfig{
		Intents: []broker.Intent{broker.CreateOnePost},
		Name:    "posts.create.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	createManyPosts = broker.QueueConfig{
		Intents: []broker.Intent{broker.CreateManyPosts},
		Name:    "posts.create.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	readOnePost = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadOnePost},
		Name:    "posts.read.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	readManyPosts = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadManyPosts},
		Name:    "posts.read.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	readOnePostHistory = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadOnePostHistory},
		Name:    "posts.read.one.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	readManyPostHistories = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadManyPostHistories},
		Name:    "posts.read.many.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	readUserPosts = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadUserPosts},
		Name:    "posts.read.userposts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	updateOne = broker.QueueConfig{
		Intents: []broker.Intent{broker.UpdateOnePost},
		Name:    "posts.update.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	deleteOne = broker.QueueConfig{
		Intents: []broker.Intent{broker.DeleteOnePost},
		Name:    "posts.delete.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	deleteMany = broker.QueueConfig{
		Intents: []broker.Intent{broker.DeleteManyPosts},
		Name:    "posts.delete.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}

	// Response queues
	commonReturnQueue = broker.QueueConfig{Intents: broker.AllIntents, Name: "Posts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	returnOneQueue    = broker.QueueConfig{
		Intents: []broker.Intent{broker.ConfirmOnePostCreation, broker.ConfirmOnePostUpdate, broker.ConfirmOnePostDeletion},
		Name:    "posts.return.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	returnMany = broker.QueueConfig{
		Intents: []broker.Intent{broker.ConfirmManyPostCreations, broker.ConfirmManyPostUpdates, broker.ConfirmManyPostDeletions},
		Name:    "posts.return.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}
	returnOneHistory = broker.QueueConfig{
		Intents: []broker.Intent{broker.ReadOnePostHistory},
		Name:    "posts.return.one.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false,
	}

	// Message consumer
	consumer = broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil}

	// Complete queue configurations
	createSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  createOnePost,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  readOnePost,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readPostHistoryConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  readOnePostHistory,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneHistory},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readUserPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  readUserPosts,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnMany},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  readManyPosts,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnMany},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	updateOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  updateOne,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	deleteOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  deleteOne,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	deleteManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  deleteMany,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnMany},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}
)
