package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PaoloSini/OpenMetro/models"
	"github.com/gorilla/websocket"
)

func viewMap(mm *models.MetroMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Client Connected")

		err = ws.WriteMessage(websocket.TextMessage, mm.ToJSON())
		if err != nil {
			log.Println(err)
		}
		reader(ws)
	}
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))
	}
}

func SetupRoutes(mm *models.MetroMap) {
	http.HandleFunc("/ws", viewMap(mm))
}
