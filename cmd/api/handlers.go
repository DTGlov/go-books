package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneBook(w http.ResponseWriter, r *http.Request){
	//This extracts the id passed in the  url
params := httprouter.ParamsFromContext(r.Context())

//the id comes in as a string so we convert it into an integer
id,err := strconv.Atoi(params.ByName("id"))
if err !=nil {
	app.errorLog.Print(errors.New("invalid id parameter"))
	app.errorJSON(w,err)
	return
}

//call the Get method from the model-->books.go file to get a single book
book,err := app.books.Get(id)
if err != nil{
	app.errorLog.Fatal(err)
}

//Getting a book to show up as json
err = app.writeJSON(w,http.StatusOK,book,"book")
	if err !=nil {
		app.errorJSON(w,err)
		return
	}
}

func (app *application) getAllBooks(w http.ResponseWriter, r *http.Request){
books,err := app.books.GetAll()
if err !=nil{
	app.errorLog.Fatal(err)
}
//getting all books as json
err = app.writeJSON(w,http.StatusOK,books,"books")
if err!=nil{
	app.errorJSON(w,err)
	return
}

}

func (app *application) getAllGenres(w http.ResponseWriter,r * http.Request){
	genre,err := app.books.GetGenres()
	if err !=nil{
		app.errorLog.Fatal(err)
	}
	//getting all the genres as json
	err = app.writeJSON(w,http.StatusOK,genre,"genre")
	if err !=nil{
		app.errorJSON(w,err)
		return
	}
}

