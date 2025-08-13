package routes

import (
	"fmt"
	"go_api/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func SumHandler (r *chi.Mux){
	r.Use(chimiddle.StripSlashes)
	
	r.Route("/calculator", func (router chi.Router)  {
		
	})

	(r.ResponseWriter w *)

}