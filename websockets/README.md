Websockets
This example will show how to work with websockets in Go. We will build a simple server which echoes back everything we send to it. For this we have to go get the popular gorilla/websocket library like so:
```
$ go get github.com/gorilla/websocket
```
From now on, every application we write will be able to make use of this library