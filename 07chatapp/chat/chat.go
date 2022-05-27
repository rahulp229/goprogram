package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("received message body from client: %s", message.Body)
	return &Message{Body: "Hello from server"}, nil
}

func (s *Server) BroadCastMessage(stream ChatService_BroadCastMessageServer) error {
	m, err := stream.Recv()
	if err != nil {
		log.Fatalf("error from client : %v", err.Error())
	}

	log.Printf("broadcast received message body from client: %s", m)

	//	m1 := Message{Body: "bcmHello from server"}
	// b, err := proto.Marshal(&m1)
	// if err != nil {
	// 	log.Fatalf("error while marsha msg of client : %v", err.Error())
	// }
	// err = stream.SendMsg(b)
	// if err != nil {
	// 	log.Fatalf("error while send msg to client : %v", err.Error())
	// }
	return nil
}
