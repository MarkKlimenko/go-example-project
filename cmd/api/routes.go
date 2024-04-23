package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// curl -i localhost:4000/v1/healthcheck
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	/*
			 curl -i --location 'localhost:4000/v1/movies' \
		     --header 'Content-Type: application/json' \
			 --data '{
			    "title": "Moana",
			    "year": 2016,
			    "runtime": "107 mins",
			    "runtimeTT": 107,
			    "genres": [
			        "animation",
			        "adventure"
			    ]
			 }'
	*/
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

	/*
				 curl -i --location --request PATCH 'localhost:4000/v1/movies/1' \
			     --header 'Content-Type: application/json' \
				 --data '{
				    "title": "Moana Edited",
				    "year": 2019,
				    "runtime": "110 mins",
				    "genres": [
				        "animation"
				    ]
				 }'

		OR MULTIPLE

		xargs -I % -P8 curl -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/4" < <(printf '%s\n' {1..8})

	*/
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)

	// curl -i localhost:4000/v1/movies/1
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	return router
}
