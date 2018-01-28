package main

import (
	"log"
	"net/http"
	"time"

	"github.com/3cb/ssc"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static/")))

	config := ssc.PoolConfig{
		IsReadable:   true,
		IsWritable:   true,
		IsJSON:       false,
		PingInterval: time.Second * 30,
	}

	pool, err := ssc.NewSocketPool(config)
	if err != nil {
		log.Fatal("Unable to start socket pool. Shutting down.")
	}

	upgrader := &websocket.Upgrader{}
	r.Handle("/ws", wsHandler(upgrader, pool))

	log.Fatal(http.ListenAndServe(":3000", r))
}

func wsHandler(upgrader *websocket.Upgrader, pool *ssc.SocketPool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := pool.AddClientSocket(upgrader, w, r)
		if err != nil {
			log.Printf("Unable to connect new websocket: %v\n", err)
		} else {
			log.Printf("New websocket client connected.")
		}
	})
}
