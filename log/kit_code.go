package main

import (
	"os"
	"github.com/go-kit/kit/log"
)

func main() {
	//log.Logger.Log("transport", "HTTP", "addr", "localhost", "msg", "listening")
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	logger.Log("question", "what is the meaning of life?", "answer", 42)
}



