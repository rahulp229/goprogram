package main

import (
	"07chatapp/chat"
	"context"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not establish connection : %v", err)
	}

	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{Body: "Hello from client"}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error while calling say hello: %s", err)
	}
	log.Printf("Response from server: %s", response)

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		j := strconv.Itoa(i)
		m := chat.Message{Body: string(j)}
		//	es := grpc.EmptyCallOption{}

		client, err := c.BroadCastMessage(context.Background())
		if err != nil {
			log.Fatalln("bcm : error while received response ", err)
		}
		err = client.Send(&m)
		if err != nil {
			log.Fatalln("bcm : error while send request ", err)
		}
		m2 := chat.Message{}
		err = client.RecvMsg(m2)
		if err != nil {
			log.Fatalln("bcm : error receive from server ", err)
		}
		log.Printf("bcm : Response from server: %v", m2.Body)

	}

}
