package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

func ServeWebContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// read html file
	f, err := os.Open("web/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
}
