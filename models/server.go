package models

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	From string

	Payload []byte
}
type Server struct {
	ListnAddr string
	Ln        net.Listener
	Quitch    chan struct{}
	Msgch     chan Message
}

func Newserver(listneAddress string) *Server {

	return &Server{
		ListnAddr: listneAddress,
		Quitch:    make(chan struct{}),
		Msgch:     make(chan Message),
	}

}
func (s *Server) Start() error {

	ln, err := net.Listen("tcp", s.ListnAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Ln = ln
	go s.Accept()
	<-s.Quitch
	close(s.Msgch)
	return nil
}

func (s *Server) Accept() {

	for {
		conn, err := s.Ln.Accept()
		if err != nil {

			log.Printf("accept error %v", err)
			continue
		}
		fmt.Println("new conncetion to the server", conn.RemoteAddr())
		go s.readLoop(conn)
	}

}

func (s *Server) readLoop(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {

			log.Println("error reading the conncecion data")
		}

		s.Msgch <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}
		conn.Write([]byte("thank u for the message"))
		//fmt.Println(string(<-s.Msgch))
	}

}
