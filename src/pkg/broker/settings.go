package broker

const (
	Consumer = ""
)

var (
	EConfig = ExchangeConfig{
		Name:       "imagehost.posts",
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
		Consumer:  Consumer,
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	}
)
