package natsutils

import (
	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func ConnectToNATS(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		logger.Log.Error("Failed to connect to NATS", zap.Error(err))
		return nil, err
	}
	logger.Log.Info("Connected to NATS server")
	return nc, nil
}

func SubscribeToTopic(nc *nats.Conn, topic string, handler func(*nats.Msg)) (*nats.Subscription, error) {
	sub, err := nc.Subscribe(topic, handler)
	if err != nil {
		logger.Log.Error("Failed to subscribe to topic",
			zap.String("topic", topic),
			zap.Error(err))
		return nil, err
	}
	logger.Log.Info("Subscribed to topic", zap.String("topic", topic))
	return sub, nil
}

func PublishMessage(nc *nats.Conn, topic string, data []byte) error {
	if err := nc.Publish(topic, data); err != nil {
		logger.Log.Error("Failed to publish message",
			zap.String("topic", topic),
			zap.Error(err))
		return err
	}
	logger.Log.Info("Message published", zap.String("topic", topic))
	return nil
}
