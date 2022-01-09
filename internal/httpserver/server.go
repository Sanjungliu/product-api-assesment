package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sanjungliu/product-api-assesment/internal/app"
	"github.com/Sanjungliu/product-api-assesment/internal/httpserver/controller"
	"github.com/go-chi/chi"
)

type Server struct {
	Controller *controller.Controller
}

func NewServer(app *app.App) *Server {
	productController := controller.NewController(&app.Product)

	return &Server{
		Controller: productController,
	}
}

func (s *Server) compileRouter() chi.Router {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Post("/products", s.Controller.AddProduct)
		r.Get("/products", s.Controller.GetListProduct)
	})

	return router
}

func (s *Server) Serve() {
	r := s.compileRouter()

	log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080")
	srv := http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
