package kitty

import (
	"net/http"
	"rest-api/app/common"
	"github.com/go-chi/chi"
)

type IndexController struct {
	common.Controller
}

// Index func return all kitties in database
func (c *IndexController) Index (w http.ResponseWriter, r *http.Request)  {
	c.SendJSON(
				w, 
				r, 
				[] *Kitty{
					NewKitty("Gaspart", "British", "2016-07-05"),
					NewKitty("Marcel", "European", "2014-05-02")},
				http.StatusOK,
	)
}


func (c *IndexController) FindByName (w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	kitties := Kitty{
		Name:  name,
		Breed: "Postman",
		BirthDate:  "today",
	}
	c.SendJSON(
		w, 
		r, 
		kitties,
		http.StatusOK,
	)
}