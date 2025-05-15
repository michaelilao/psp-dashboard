package api

import (
	"log"
	"net/http"
	"psp-dashboard-be/service/transaction"
	"psp-dashboard-be/service/user"

	httpSwagger "github.com/swaggo/http-swagger"
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

	corsWrapper := CORSMiddleware(router)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return 
		}
		w.Write([]byte("welcome to psp-dashboard go be"))
	})


	router.Handle("/swagger/", httpSwagger.WrapHandler)
	server := http.Server{
		Addr: s.addr,
		Handler: corsWrapper,
	}

	log.Println("server running on port"+s.addr)
	return server.ListenAndServe()
}

func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins (adjust as needed for security)
        w.Header().Set("Access-Control-Allow-Origin", "*")
        // Allow specific HTTP methods
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        // Allow specific headers
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-reset-token")
        
        // Handle preflight requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
