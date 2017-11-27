package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ty-edelweiss/crawler/test"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Db	*sql.DB
}

func NewServer(text string) Server {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/test")
	if err != nil {
		log.Println(err)
	}

	log.Println(text)
	return Server{db}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	text := LogRequest(r)
	log.Println(text)
	w.WriteHeader(200)
	w.Write([]byte("Hello Go Server!"))
}
