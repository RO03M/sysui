package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("GET /resource-info", getResourceInfo)
	http.HandleFunc("GET /sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		fmt.Printf("teste", flusher)
		// Send events in a loop
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data: Message %d at %s\n\n", i, time.Now().Format(time.RFC3339))
			flusher.Flush()
			time.Sleep(1 * time.Second)
		}
	})

	fmt.Println("Running")
	http.ListenAndServe(":8000", nil)
}
