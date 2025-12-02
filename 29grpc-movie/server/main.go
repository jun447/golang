package main

import (
	"database/sql"
	"fmt"
	"net"
	"time"

	pb "grpc-movie/proto"

	_ "github.com/go-sql-driver/mysql"
	"context"
	grpc "google.golang.org/grpc"
)

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
type server struct{
	pb.UnimplementedMovieServiceServer
}

func (s *server) GetAllMovies(ctx context.Context,_ *pb.Empty) (*pb.MovieList, error) {
	fmt.Println("Getting all Movies")
	rows, err := db.Query("Select * from movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// var movies *pb.MovieList //panic and create error
	movies := &pb.MovieList{} // initialize pointer to struct

	for rows.Next(){
		var movie pb.Movie
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Watched)
		if err != nil {
			return nil, err
		}
		movies.Movies=append(movies.Movies,&movie)
	}
	return movies,nil

}
func (s *server) GetMovieById( ctx context.Context, req *pb.MovieId) (*pb.Movie,error){
	fmt.Println("Get Movie  by Id")

	row:=db.QueryRow("select * from movies where id=?", req.Id)

	 movie :=&pb.Movie{}
	err:=row.Scan(&movie.Id,&movie.Title,&movie.Watched)
	if err != nil {
		fmt.Println("Error in scanning data from mysql", err)
		return nil,err
	}
	return movie,nil
}
//deletion
func (s *server) DeleteMovieById(ctx context.Context, req *pb.MovieId) (*pb.Empty,error) {
	fmt.Println("Delete Movie by Id")
	_,err := db.Exec("Delete from movies where id=?",req.Id)
	if err != nil {
		return nil,err
	}
	return &pb.Empty{},nil
}

func (s *server) CreateMovie(ctx context.Context, req *pb.Movie) (*pb.Movie,error) { 
	fmt.Println("Inserting Movie")

	 movie :=&pb.Movie{
		Id: req.Id,
		Title: req.Title,
		Watched: req.Watched,
	 }
	db.Exec("INSERT INTO `movies` (`movie`, `watched`) VALUES (?, ?)", movie.Title, movie.Watched)
	return movie,nil
}
//update an row function
func (s *server) UpdateMovieById(ctx context.Context, req *pb.Movie) (*pb.Movie,error) { 
	fmt.Println("Update Movie by Id")
	//grabing id
    movie :=&pb.Movie{
		Id: req.Id,
		Title: req.Title,
		Watched: req.Watched,
	}

	_,err:=db.Exec("update movies set movie=?,watched=? where id=?",movie.Title,movie.Watched,req.Id)
	if err!=nil {
		return nil,err
	}
	return movie,nil
}
func main() {
    // 1. Create TCP listener on port 50051
    lis, err := net.Listen("tcp", ":50051")
	
	if err != nil {
		panic(err)
	}
    
    // 2. Create gRPC server
    grpcServer := grpc.NewServer()
    
    // 3. Register your service (look for RegisterMovieServiceServer in generated code)
    pb.RegisterMovieServiceServer(grpcServer, &server{})
    
    // 4. Start serving
    grpcServer.Serve(lis)

}