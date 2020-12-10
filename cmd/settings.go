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
	createOnePostReqQueue         = broker.QueueConfig{Name: "posts.create.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	createManyPostsReqQueue       = broker.QueueConfig{Name: "posts.create.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	readOnePostReqQueue           = broker.QueueConfig{Name: "posts.read.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	readManyPostsReqQueue         = broker.QueueConfig{Name: "posts.read.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	readOnePostHistoryReqQueue    = broker.QueueConfig{Name: "posts.read.one.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	readManyPostHistoriesReqQueue = broker.QueueConfig{Name: "posts.read.many.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	readUserPostsReqQueue         = broker.QueueConfig{Name: "posts.read.userposts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	updateOneReqQueue             = broker.QueueConfig{Name: "posts.update.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	deleteOneReqQueue             = broker.QueueConfig{Name: "posts.delete.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	deleteManyReqQueue            = broker.QueueConfig{Name: "posts.delete.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}

	// Response queues
	commonReturnQueue     = broker.QueueConfig{Name: "Posts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	returnOneQueue        = broker.QueueConfig{Name: "posts.return.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	returnManyQueue       = broker.QueueConfig{Name: "posts.return.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}
	returnOneHistoryQueue = broker.QueueConfig{Name: "posts.return.one.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false}

	// Message consumer
	consumer = broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil}

	// Complete queue configurations
	createSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.CreateOnePost,
		SubQueueConf:  createOnePostReqQueue,
		PubIntent:     broker.ConfirmOnePostCreation,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.RequestOnePost,
		SubQueueConf:  readOnePostReqQueue,
		PubIntent:     broker.ReturnOnePost,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readPostHistoryConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.RequestOnePostHistory,
		SubQueueConf:  readOnePostHistoryReqQueue,
		PubIntent:     broker.ReturnOnePostHistory,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneHistoryQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readUserPostsConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.RequestUserPosts,
		SubQueueConf:  readUserPostsReqQueue,
		PubIntent:     broker.ReturnUserPosts,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnManyQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	readManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.RequestManyPosts,
		SubQueueConf:  readManyPostsReqQueue,
		PubIntent:     broker.ReturnManyPosts,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnManyQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	updateOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.UpdateOnePost,
		SubQueueConf:  updateOneReqQueue,
		PubIntent:     broker.ConfirmOnePostUpdate,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	deleteOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.DeleteOnePost,
		SubQueueConf:  deleteOneReqQueue,
		PubIntent:     broker.ConfirmOnePostDeletion,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnOneQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}

	deleteManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubIntent:     broker.DeleteManyPosts,
		SubQueueConf:  deleteManyReqQueue,
		PubIntent:     broker.ConfirmManyPostDeletions,
		PubQueueConfs: []broker.QueueConfig{commonReturnQueue, returnManyQueue},
		ExchangeConf:  rapidExchange,
		ConsumerConf:  consumer,
	}
)
