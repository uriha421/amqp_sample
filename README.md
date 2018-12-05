# What is this?
This is an introduction to amqp.

# Prerequisites
 - golang
 - docker

# How to use?
First, start amqp.
```
$ docker run --detech --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Then, prepare queue for receiving messages.
```
$ git clone https://github.com/uriha421/amqp_sample.git
$ cd queue/
$ go run main.go
```

Finally, open a new tab and publish a message.
```
$ cd ../publisher/
$ go run main.go
```

you should see the following message.
```
message received: Hello Wolrd
```

thank you.
