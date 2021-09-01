package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router{
router := httprouter.New()

router.HandlerFunc(http.MethodGet,"/status",app.statusHandler)

router.HandlerFunc(http.MethodGet,"/v1/book/:id",app.getOneBook)
router.HandlerFunc(http.MethodGet,"/v1/books",app.getAllBooks)
router.HandlerFunc(http.MethodGet,"/v1/genres",app.getAllGenres)

return router
}