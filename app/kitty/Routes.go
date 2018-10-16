package kitty

import "github.com/go-chi/chi"


func Routes()  *chi.Mux{
	kc := IndexController{}

	router := chi.NewRouter()
	router.Get("/", kc.Index)
	router.Get("/{name}", kc.FindByName)
	return router
}