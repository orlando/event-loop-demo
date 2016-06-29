package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/googollee/go-socket.io"
)

var messages = []map[string]interface{}{
	map[string]interface{}{"color": "#00d2ff", "index": 1},
	map[string]interface{}{"color": "#3a7bd5", "index": 2},
	map[string]interface{}{"color": "#D3959B", "index": 3},
	map[string]interface{}{"color": "#BFE6BA", "index": 4},
	map[string]interface{}{"color": "#DAD299", "index": 5},
	map[string]interface{}{"color": "#B0DAB9", "index": 6},
	map[string]interface{}{"color": "#E6DADA", "index": 7},
	map[string]interface{}{"color": "#274046", "index": 8},
	map[string]interface{}{"color": "#5D4157", "index": 9},
	map[string]interface{}{"color": "#A8CABA", "index": 10},
	map[string]interface{}{"color": "#ddd6f3", "index": 11},
	map[string]interface{}{"color": "#faaca8", "index": 12},
	map[string]interface{}{"color": "#616161", "index": 13},
	map[string]interface{}{"color": "#9bc5c3", "index": 14},
	map[string]interface{}{"color": "#50C9C3", "index": 15},
	map[string]interface{}{"color": "#96DEDA", "index": 16},

	map[string]interface{}{"color": "#00d2ff", "index": 17},
	map[string]interface{}{"color": "#3a7bd5", "index": 18},
	map[string]interface{}{"color": "#D3959B", "index": 19},
	map[string]interface{}{"color": "#BFE6BA", "index": 20},
	map[string]interface{}{"color": "#DAD299", "index": 21},
	map[string]interface{}{"color": "#B0DAB9", "index": 22},
	map[string]interface{}{"color": "#E6DADA", "index": 23},
	map[string]interface{}{"color": "#274046", "index": 24},
	map[string]interface{}{"color": "#5D4157", "index": 25},
	map[string]interface{}{"color": "#A8CABA", "index": 26},
	map[string]interface{}{"color": "#ddd6f3", "index": 27},
	map[string]interface{}{"color": "#faaca8", "index": 28},
	map[string]interface{}{"color": "#616161", "index": 29},
	map[string]interface{}{"color": "#9bc5c3", "index": 30},
	map[string]interface{}{"color": "#50C9C3", "index": 31},
	map[string]interface{}{"color": "#96DEDA", "index": 32},

	map[string]interface{}{"color": "#00d2ff", "index": 33},
	map[string]interface{}{"color": "#3a7bd5", "index": 34},
	map[string]interface{}{"color": "#D3959B", "index": 35},
	map[string]interface{}{"color": "#BFE6BA", "index": 36},
	map[string]interface{}{"color": "#DAD299", "index": 37},
	map[string]interface{}{"color": "#B0DAB9", "index": 38},
	map[string]interface{}{"color": "#E6DADA", "index": 39},
	map[string]interface{}{"color": "#274046", "index": 40},
	map[string]interface{}{"color": "#5D4157", "index": 41},
	map[string]interface{}{"color": "#A8CABA", "index": 42},
	map[string]interface{}{"color": "#ddd6f3", "index": 43},
	map[string]interface{}{"color": "#faaca8", "index": 44},
	map[string]interface{}{"color": "#616161", "index": 45},
	map[string]interface{}{"color": "#9bc5c3", "index": 46},
	map[string]interface{}{"color": "#50C9C3", "index": 47},
	map[string]interface{}{"color": "#96DEDA", "index": 48},

	map[string]interface{}{"color": "#00d2ff", "index": 49},
	map[string]interface{}{"color": "#3a7bd5", "index": 50},
	map[string]interface{}{"color": "#D3959B", "index": 51},
	map[string]interface{}{"color": "#BFE6BA", "index": 52},
	map[string]interface{}{"color": "#DAD299", "index": 53},
	map[string]interface{}{"color": "#B0DAB9", "index": 54},
	map[string]interface{}{"color": "#E6DADA", "index": 55},
	map[string]interface{}{"color": "#274046", "index": 56},
	map[string]interface{}{"color": "#5D4157", "index": 57},
	map[string]interface{}{"color": "#A8CABA", "index": 58},
	map[string]interface{}{"color": "#ddd6f3", "index": 59},
	map[string]interface{}{"color": "#faaca8", "index": 60},
	map[string]interface{}{"color": "#616161", "index": 61},
	map[string]interface{}{"color": "#9bc5c3", "index": 62},
	map[string]interface{}{"color": "#50C9C3", "index": 63},
	map[string]interface{}{"color": "#96DEDA", "index": 64},
}

var ch = make(chan map[string]interface{}, 64)

func main() {
	for _, msg := range messages {
		ch <- msg
	}

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.On("ready", func(msg string) {
			for msg := range ch {
				sendTick(so, msg)
			}
		})
	})

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000/")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})

	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendTick(so socketio.Socket, msg map[string]interface{}) {
	log.Println("on tick")

	amt := time.Duration(rand.Intn(5))
	time.Sleep(time.Second * amt)

	log.Println("emit: ", msg, "sleep: ", time.Second*amt)
	so.Emit("paint", msg)
}
