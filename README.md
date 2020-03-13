<h1 align="center">Welcome to Go Message Broker 👋</h1>

> Example of implementing message brokers in golang

## Message Brokers

  - [x] RabbitMQ
  - [x] Apache Kafka
  - [ ] SQS AWS
  - [ ] Apache ActiveMQ

## Test implements

#### RabbitMQ

- Run
```sh
make up-rabbitmq
```

- Open [RabbitMQ](http://localhost:15672)
```sh
Username: guest
Password: guest
```

- Publish message
```sh
make publish-rabbitmq
```

- Consume message
```sh
make consumer-rabbitmq
```

### Kafka

- Run
```sh
make up-kafka
```

- Publish message
```sh
make publish-kafka
```

- Consume message
```sh
make consumer-kafka
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