package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	// instantiate a new node of maelstrom
	n := maelstrom.NewNode()

	// register a handler cb function for "echo" messages
	n.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshall the json msg body as loosely-typed map

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// update the message type to return back
		body["type"] = "echo_ok"

		// return the modified message
		return n.Reply(msg, body)
	})

	// finally run the echo node
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
	
}

