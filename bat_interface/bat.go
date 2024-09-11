package bat_interface

import "net/http"

type Ws interface {
	SendMessage()
	ReadMessage(w http.ResponseWriter, r *http.Request)
	Start()
}
