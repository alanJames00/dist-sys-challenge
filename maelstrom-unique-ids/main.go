package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	n *maelstrom.Node
}

func NewServer(n *maelstrom.Node) *server {
	return &server{
		n: n,
	}
}

func (s *server) uniqueIdHandler(msg maelstrom.Message) error {

	// generate a unique id
	var randomNum int64
	if err := binary.Read(rand.Reader, binary.BigEndian, &randomNum); err != nil {
		return err
	}
	

	response := map[string]any{
		"type": "generate_ok",
		"id": fmt.Sprintf("%v%v", time.Now().UnixNano(), randomNum),
	}

	return s.n.Reply(msg, response)
}

func main() {
	n := maelstrom.NewNode()

	s := NewServer(n)

	s.n.Handle("generate", s.uniqueIdHandler)

	if err := s.n.Run(); err != nil {
		log.Fatal(err)
	}
}
