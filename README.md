<h1 align="center">Welcome to Go Message Broker 👋</h1>

> Example of implementing message brokers in golang

## Message Brokers

  - [x] RabbitMQ
  - [x] Apache Kafka
  - [x] Apache ActiveMQ
  - [ ] Redis

## Test implements

#### RabbitMQ

- Run
```sh
make up-rabbitmq
```

- Open interface web [RabbitMQ](http://localhost:15672)
```sh
Username: guest
Password: guest
```

- Consume message
```sh
make consumer-rabbitmq
```

- Publish message
```sh
make publish-rabbitmq
```

### Kafka

- Run
```sh
make up-kafka
```

- Consume message
```sh
make consumer-kafka
```

- Publish message
```sh
make publish-kafka
```

### ActiveMQ

- Run
```sh
make up-activemq
```

- Open interface web [ActiveMQ](http://localhost:8161/console/login)
```sh
Username: guest
Password: guest
```

- Consume message
```sh
make consumer-activemq
```

- Publish message
```sh
make publish-activemq
```

## Author

👤 **Gabriel S. Facina**

* Website: gsabadini.github.io
* Github: [@GSabadini](https://github.com/GSabadini)
* LinkedIn: [@gabriel-sabadini-facina](https://linkedin.com/in/gabriel-sabadini-facina)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_