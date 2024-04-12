package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/medali/test/pkg/models"
	"github.com/medali/test/pkg/utils"
)
var NewBook models.Book




func GetBook(res http.ResponseWriter, req *http.Request){
	newBooks := models.GetAllBooks()
	results, _ := json.Marshal(newBooks)
	res.Header().Set( "Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(results)
}


func GetSingleBook (res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0 ,0);
	if err != nil {
		log.Fatal(err)
	}
	bookDetails, _ := models.GetBookById(ID)
	results, _ := json.Marshal(bookDetails)
	res.Header() .Add("Content-type","Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(results)
}

// func createMovie(res http.ResponseWriter, req *http.Request){
// 		res.Header().Set("Content-Type", "application/json")
// 		var movie Movie 
// 		_ = json.NewDecoder(req.Body).Decode(&movie)
// 		movie.Id = strconv.Itoa(rand.Intn(1000000) + 3333) 
// 		movies = append(movies, movie)
// 		json.NewEncoder(res).Encode(movie)
// 	}




func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)

	b := CreateBook.CreateBook()
	results, _ := json.Marshal(b)

	res.WriteHeader(http.StatusOK)
	res.Write(results)

		
}	



func DeleteSingleBook (res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0 ,0);
	if err != nil {
		log.Fatal(err)
	}
	bookDetails := models.DeleteBook(ID)
	results, _ := json.Marshal(bookDetails)
	res.Header() .Add("Content-type","Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(results)
}

func UpdateSingleBook (res http.ResponseWriter, req *http.Request){
	updateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)

	vars := mux.Vars(req)
	bookId := vars["id"]


	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		log.Fatal(err)
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
results, _ := json.Marshal(bookDetails)
	res.Header() .Add("Content-type","Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(results)
	

	
	

}