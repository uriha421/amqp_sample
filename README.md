# What is this?
This is an introduction to amqp.

# Prerequisites
 - golang v1.10.3
 - docker v18.09.0

# How to use?
First, start amqp.
```
$ docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Then, create an exchange.
```
$ git clone https://github.com/uriha421/amqp_sample.git
$ cd publisher/
$ go run main.go
```

After that, prepare a queue for receiving messages.
```
$ cd ../queue/
$ go run main.go
```

Finally, open a new tab and publish a message.
```
$ cd ../publisher/
$ go run main.go
```

You should see the following message.
```
message received: Hello Wolrd
```

# References
- https://github.com/PacktPublishing/Cloud-Native-programming-with-Golang

Thank you.
