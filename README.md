# streadway/amqp Conneciton/Channel auto reconnect wrap
streadway/amqp Connection/Channel does not reconnect if rabbitmq server restart/down.

To simply developers, here is auto reconnect wrap with detail comments.

## Example
### Close by developer
> go run example/close/demo.go

### Auto reconnect
> go run example/reconnect/demo.go

