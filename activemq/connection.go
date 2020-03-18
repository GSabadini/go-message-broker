package activemq

import "github.com/go-stomp/stomp"

func Connect() (*stomp.Conn, error) {
	addr := "localhost:61616"

	return stomp.Dial("tcp", addr, stomp.ConnOpt.Login("user-amq", "password-amq"))
}
