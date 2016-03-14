package logrusnats

import (
	"github.com/Sirupsen/logrus"
	"github.com/nats-io/nats"
	"time"
)

type NATSHook struct {
	Servers       []string
	Subject       string
	MaxReconnect  int
	ReconnectWait int
	NoRandomize   bool
}

func NewNATSHook(servers []string, subject string) *NATSHook {
	hook := &NATSHook{Servers: servers,
		Subject: subject,
	}

	return hook
}

func (hook *NATSHook) Fire(entry *logrus.Entry) error {
	opts := nats.DefaultOptions
	opts.Servers = hook.Servers
	opts.MaxReconnect = hook.MaxReconnect
	opts.ReconnectWait = (time.Duration(hook.ReconnectWait) * time.Second)
	opts.NoRandomize = hook.NoRandomize

	client, err := opts.Connect()
	if err != nil {
		return err
	}
	defer client.Close()

	msg, err := entry.String()
	if err != nil {
		return err
	}

	client.Publish(hook.Subject, []byte(msg))
	client.Flush()

	return nil
}

func (hook *NATSHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
