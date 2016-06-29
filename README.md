# Concurrencia, Paralelismo y el Event-loop

Este es un demo que hice para una platica. Los slides de la platica
estan en [SpeakerDeck](https://speakerdeck.com/orlando/concurrencia-paralelismo-y-el-event-loop)

## Como correrlo?

### Node.js

1. `npm install`
2. `node server.js`
3. `python -m SimpleHTTPServer`
4. `open http://localhost:8000/`

En el caso de Node.js, como es single threaded y como funciona
socket.io, no se renderea en el browser hasta que sucede el siguiente
tick. Por eso, el array de mensajes que usamos es mas corto.

### Go

1. `go get github.com/googollee/go-socket.io`
2. `go run main.go` sin goroutines o `go run main_goroutine.go` con
   goroutines
3. `python -m SimpleHTTPServer`
4. `open http://localhost:8000/`
