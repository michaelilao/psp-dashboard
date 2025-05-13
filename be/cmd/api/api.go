package api

import "net/http"

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return 
		}
		w.Write([]byte("welcome to psp-dashboard go be"))
	})

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}