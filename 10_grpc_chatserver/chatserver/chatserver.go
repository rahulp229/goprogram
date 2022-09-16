package chatserver

import (
	"math/rand"
	"sync"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
}

type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

type ChatServer struct {
}

//define chat service
func (is *ChatServer) ChatService(csi Services_ChatServiceServer) error {
	clientUniqueCode := rand.Intn(1e6)

	errch := make(chan error)

	//receive message

	//send message

	return <-errch
}

func receiveFromStream(csi_ servicesChatServiceServer, clientUniqueCode_ int, errch_ chan error )  {
	//implementing loop

	for{
		mssg, err := csi_.Recv()
		if err != nil{
			log.Println
		}
	}
}