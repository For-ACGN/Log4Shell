package log4j2

import (
	"fmt"

	"golang.org/x/crypto/acme/autocert"
)

func testAutoCert() {

	listener := autocert.NewListener("")

	mgr := autocert.Manager{}
	mgr.TLSConfig()

	conn, err := listener.Accept()
	fmt.Println(err)

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	fmt.Println("asdasdads", err)
	fmt.Println(string(buf[:n]))

	fmt.Println(conn.RemoteAddr())

	// m:= autocert.Manager{}
}
