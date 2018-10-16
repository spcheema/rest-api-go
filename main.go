package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/middleware"
	"fmt"
	"time"
	"rest-api/app/kitty"
	"github.com/gobuffalo/envy"
	
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "First route!")
}

func Routes() *chi.Mux{
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	  // Set a timeout value on the request context (ctx), that will signal
  	// through ctx.Done() that the request has timed out and further
  	// processing should be stopped.
  	router.Use(middleware.Timeout(60 * time.Second))
	  router.Route("/v1", func(r chi.Router){
		r.Get("/hello", sayHello);

		r.Mount("/kitties", kitty.Routes())
	})

	return router
}


func main()  {

	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", envy.Get("APP_PORT", "8001")), router))

}