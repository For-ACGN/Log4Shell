package log4shell

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/acme/autocert"
)

func TestNewListener(t *testing.T) {
	const testDomain = "test"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})
	server := http.Server{}
	server.Handler = mux
	go func() {

		http.DefaultClient.Transport = &http.Transport{}

		listener := autocert.NewListener(testDomain)
		conn, err := listener.Accept()
		require.NoError(t, err)

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		fmt.Println("err:", err)
		fmt.Println(string(buf[:n]))

		fmt.Println(conn.RemoteAddr())

		// log.Fatal(server.Serve(autocert.NewListener("example.com")))
	}()

	cfg := tls.Config{
		ServerName: testDomain,
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &cfg,
		},
	}

	req, err := http.NewRequest(http.MethodGet, "https://127.0.0.1:443/", nil)
	require.NoError(t, err)
	req.Host = testDomain

	resp, err := client.Do(req)
	require.NoError(t, err)

	fmt.Println(resp.StatusCode)

	// conn, err := tls.Dial("tcp", "127.0.0.1:443", &cfg)
	// require.NoError(t, err)
	//
	// _, err = conn.Write([]byte{1, 2, 3, 4})
	// require.NoError(t, err)
}
