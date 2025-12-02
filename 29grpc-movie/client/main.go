package main

import (
	"context"
	"fmt"

	pb "grpc-movie/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    // 1. Connect to server
    conn, err := grpc.NewClient("localhost:50051", 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    // handle err
	if err!=nil{
		panic(err)
	}
    defer conn.Close()
    
    // 2. Create client from connection
    client := pb.NewMovieServiceClient(conn)
    
    // 3. Call methods!
    // Example: Get all movies
   

	movie,err :=client.CreateMovie(context.Background(),&pb.Movie{
		Title: "A Sepration",
		Watched: true,
	})
	if err!=nil{
		panic(err)
	}

	fmt.Println("this new movie created ",movie)

	movies, err := client.GetAllMovies(context.Background(), &pb.Empty{})
	if err!=nil{
		panic(err)
	}
	
	fmt.Println("This is All Movie",movies)

	movieUpdated,err:=client.UpdateMovieById(context.Background(),&pb.Movie{
		Id: 7,
		Title: "Patakha 2018",
		Watched: true,
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println("This is Updated Movie",movieUpdated)
}