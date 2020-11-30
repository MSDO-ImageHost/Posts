package broker

var (
	EConfig = ExchangeConfig{
		Name:       "posts",
		Kind:       "direct",
		Durable:    false,
		AutoDelete: true,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}

	QConfig = QueueConfig{
		Name:       "posts.default",
		Durable:    true,
		AutoDelete: true,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}

	CConfig = ConsumeConfig{
		Queue:     QConfig.Name,
		Consumer:  "",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	}

	QBConfig = QueueBindingConfig{
		Name:     QConfig.Name,
		Key:      "",
		Exchange: EConfig.Name,
		NoWait:   false,
		Args:     nil,
	}
)
