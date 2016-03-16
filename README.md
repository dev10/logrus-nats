[![Build Status](https://semaphoreci.com/api/v1/ehoward/logrus-nats/branches/master/shields_badge.svg)](https://semaphoreci.com/ehoward/logrus-nats)

# logrus-nats

A [Logrus]( https://github.com/Sirupsen/logrus) hook for sending data to [NATS]( http://nats.io/)

## Installation
	go get github.com/trillium-labs/logrus-nats

## Usage

```go
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/trillium-labs/logrus-nats"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stderr)

	servers := []string{"nats://localhost:4222"}
	hook := logrusnats.NewNATSHook(servers, "my-subject")
	logrus.AddHook(hook)

	logrus.Warn("warning")
	logrus.Info("info")
	logrus.Debug("debug")
}
```

## Credits
Based on AMQP handler by [vladoatanasov](https://github.com/vladoatanasov/logrus_amqp).

## License
Released under the [MIT License](http://www.opensource.org/licenses/MIT).
