package controller

// import (
// 	"context"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	// "go.mongodb.org/mongo-driver/v2/mongo"
// )

// const connectionString = "mongodb+srv://devjun:devjun123@cluster0.009hsx9.mongodb.net/?appName=Cluster0"
// const dbName = "netflix"
// const colName = "watchlist"

// //Most Important

// var collection *mongo.Collection

// func init(){
// 	//client option
// 	clientOption := options.Client().ApplyURI(connectionString)

// 	//conect to mongodb

//     client,err :=	mongo.Connect(context.TODO(),clientOption)
// 	handleNilErr(err)

// 	collection=client.Database(dbName).Collection(colName)
// 	log.Println("MongoDB connection successfull")

// }

// func handleNilErr(err error){
// 	if err!=nil{
// 		log.Fatal(err)
// 		panic(err)
// 	}
// }

// will ue mysql onwards'

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jun447/mongoApi/model"
)

// ...
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "go_user:abcd1234@tcp(127.0.0.1:3306)/netflix")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("DB connected successfully Bro")
}

// GetDB returns the package-level database connection (useful for tests or other packages that import this package).
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection if it is open.
func CloseDB() error {
	if db == nil {
		return nil
	}
	return db.Close()
}

// mysql helper functions can be added here

func insertOneMovie(movie model.Netflix){
	db.Exec("INSERT INTO `movies` (`movie`, `watched`) VALUES (?, ?)", movie.Movie, movie.Watched)
}

func InsertMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("Inserting Movie")
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	//decoder
	// body is stream of data
	// we are decoding that stream of data into movie variable
	// &movie means we are passing address of movie variable
	// so that decoder can fill data at that address
	// similar to dereferencing pointer
	// in previous lesson

	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


//retrieval
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Movie  by Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)
	row:=db.QueryRow("select * from movies where id=?", params["id"])

	var movie model.Netflix
	err:=row.Scan(&movie.ID,&movie.Movie,&movie.Watched)
	if err != nil {
		fmt.Println("Error in scanning data from mysql", err)
		panic(err)
	}
	json.NewEncoder(w).Encode(movie)
}
// get all movies from mysql
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Movies")
	w.Header().Set("Content-Type", "application/json")
	
	rows,err := db.Query("Select * from movies")
	if (err!=nil){
		fmt.Println("Error in fetching data from mysql", err)
		panic(err)
	}

	var movies []model.Netflix
	for rows.Next(){
		var movie model.Netflix
		err=rows.Scan(&movie.ID,&movie.Movie,&movie.Watched)
		if(err!=nil){
			fmt.Println("Error in scanning data from mysql", err)
			panic(err)
		}
		movies=append(movies, movie)
		// fmt.Println("this simple row print ",rows)
		// fmt.Println("This is row.next()",rows.Next())
		// fmt.Println("this is row columns",rows.Columns())
		fmt.Println(reflect.TypeOf(movie).Kind())

	}
	fmt.Println("movies fetched from mysql ", movies)
	json.NewEncoder((w)).Encode(movies)

	defer rows.Close()

}

//deletion
func DeleteMovieByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Movie by Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)

	_,err := db.Exec("Delete from movies where id=?",params["id"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("Movie Deleted Successfully")
}
//update an row function
func UpdateMovieByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update Movie by Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)
    var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	_,err:=db.Exec("update movies set movie=?,watched=? where id=?",movie.Movie,movie.Watched,params["id"])
	if err!=nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(movie)

}