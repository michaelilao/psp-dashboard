package api

import (
	"log"
	"net/http"
	"psp-dashboard-be/service/transaction"
	"psp-dashboard-be/service/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
	db *mongo.Client
}

func NewAPIServer(addr string, db *mongo.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()

	userStore := user.NewStore(s.db)
	transactionStore := transaction.NewStore(s.db)


	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	transactionHandler := transaction.NewHandler(transactionStore, userStore)
	transactionHandler.RegisterRoutes(router)

	
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

	log.Println("server running on port"+s.addr)
	return server.ListenAndServe()
}