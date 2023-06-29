package wsManage

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var hub = Hub{
	register:      make(chan *Conn),
	unregister:    make(chan *Conn),
	connections:   make(map[*Conn]bool),
	SubscribeList: make(map[string]map[*Conn]bool),
}

func WebsocketInit(url string, logger *log.Logger) {
	hub.Logger = logger
	// hub.SetNewGrpcServer()

	go hub.run()
	// go hub.PrintConnectionsCount()
	go hub.ListenMessageStream()
	http.HandleFunc("/ws", serveWs)

	// http.Handle("/ws", Middleware(
	// 	http.HandlerFunc(serveWs),
	// 	authMiddleware,
	// ))
	HttpListenAndServe(url)
}

func HttpListenAndServe(Wsurl string) {
	inc := func(Wsurl string) {
		fmt.Println("Start Listen And Serve WS")
		err := http.ListenAndServe(Wsurl, nil)
		var c *websocket.Conn
		if err != nil {
			fmt.Println("IT`S RESERVE WEBSOCKET")
			URL := url.URL{Scheme: "ws", Host: Wsurl, Path: "ws"}
			c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
			if err == nil {
				for {
					_, _, err := c.ReadMessage()
					if err != nil {
						break
					}
				}
			}
		}
		if c != nil {
			c.Close()
		}
	}
	for {
		inc(Wsurl)
	}
}

// func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
// 	for _, mw := range middleware {
// 		h = mw(h)
// 	}
// 	return h
// }

// func authMiddleware(next http.Handler) http.Handler {
// 	TestApiKey := "test_api_key"
// 	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		var apiKey string
// 		if apiKey = req.Header.Get("Token"); apiKey != TestApiKey {
// 			log.Printf("bad auth api key: %s", apiKey)
// 			rw.WriteHeader(http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(rw, req)
// 	})
// }
