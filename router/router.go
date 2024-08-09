package router

import (
	"net/http"
)

// make a default router to initiate at the root

func BootstrapRouter() http.Handler {
	root := http.NewServeMux()
	root.HandleFunc("/", handleRoot)

	root.Handle("/api/", MakeAPIHandler("/api"))

	return root
}

func handleRoot(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello from root"))
}
