package main

import broker "github.com/MSDO-ImageHost/Posts/internal/broker"

var (
	_LOG_TAG string = "App:\t"

	rapidExchange = broker.ExchangeConfig{Name: "rapid", Kind: "direct", Durable: false, AutoDelete: false, Internal: false, NoWait: false, Args: nil}

	createSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.create.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.one" /*"created-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "CreateOnePost", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	readSinglePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.read.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.one" /*"single-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "RequestOnePost", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	readPostHistoryConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.read.history", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.one-history" /*"history-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "RequestOnePostHistory", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	readUserPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.read.userposts", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.many" /*"user-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "ReturnUserPosts", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	readManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.read.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.many" /*"many-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "ReturnManyPosts", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	updateOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.update.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.one-id" /*"updated-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "UpdateOnePost", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	deleteOnePostConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.delete.one", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.one-id" /*"deleted-post"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "DeleteOnePost", Exchange: rapidExchange, NoWait: false, Args: nil},
	}

	deleteManyPostsConf broker.HandleConfig = broker.HandleConfig{
		SubQueueConf:  broker.QueueConfig{Name: "posts.delete.many", Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		PubQueueConf:  broker.QueueConfig{Name: "posts.return.many-ids" /*"deleted-many-posts"*/, Durable: true, AutoDelete: false, Exclusive: false, NoWait: false},
		ConsumerConf:  broker.ConsumerConfig{AutoAck: false, Exclusive: false, NoLocal: false, NoWait: false, Args: nil},
		QueueBindConf: broker.QueueBindConfig{Name: "", Key: "DeleteManyPosts", Exchange: rapidExchange, NoWait: false, Args: nil},
	}
)
