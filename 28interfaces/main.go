package main

import "fmt"

// type Animal interface{
// 	Speak() string
// 	Move() string
// }

// func MakeItSpeak(s Animal) {
// 	fmt.Println(s.Speak())
// 	fmt.Println(s.Move())
// }

// type Dog struct{
// 	Name string
// }
// type Cat struct{
// 	Name string
// }


// func (d *Dog) Speak() string {
// 	return "Woof! I am " + d.Name
// }
// func (d *Dog) Move() string {
// 	return "Dog is running"
// }

// func (c *Cat) Speak() string {
// 	return "Meow! I am " + c.Name
// }

// func (c *Cat) Move() string {
// 	return "Cat is running"
// }

// func main() {
//     fmt.Println("Interfaces learning")
// 	dog:=Dog{Name: "Buddy"}
// 	cat:=Cat{Name: "Whiskers"}

// 	// animals:=[]interface{}{&dog,&cat} //slice of interfaces
// 	// fmt.Println(dog.Speak())
// 	// fmt.Println(cat.Speak())
// 	MakeItSpeak(&dog)
// 	MakeItSpeak(&cat)

// }
type MovieStorage interface {
	Save(movie string) error
	GetAll() []string
}

type InMemoryMovieStorage struct {
	movies []string
}

func (s *InMemoryMovieStorage) Save(movie string) error{
	fmt.Println("Saving movie in memory:", movie)
	s.movies = append(s.movies, movie)
	return nil
}

func (s *InMemoryMovieStorage) GetAll() []string{
	return s.movies
}

type InDBMovieStorage struct {
	// dbConnection string
	movies []string
}

func (s *InDBMovieStorage) Save(movie string) error{
	fmt.Println("Saving movie in database:", movie)
	s.movies = append(s.movies, movie)
	return nil
}

func (s *InDBMovieStorage) GetAll() []string{
	return s.movies
}

//creating the service that uses the interface
type MovieStorageService struct{
	storage MovieStorage
}

func (s *MovieStorageService) AddMovie(movie string){
	s.storage.Save(movie)
}

func (s *MovieStorageService) ListMovies() {
	movies := s.storage.GetAll()
	fmt.Println("Movies:")
	for _, movie := range movies {
		fmt.Println("-", movie)
	}
}

func main() {

	// var storage MovieStorage
	
	// // Use in-memory storage
	// storage = &InMemoryMovieStorage{}
	// storage.Save("Inception")
	// storage.Save("The Matrix")
	// fmt.Println("Movies in memory:", storage.GetAll())

	// // Switch to database storage
	// storage = &InDBMovieStorage{}
	// storage.Save("Interstellar")
	// storage.Save("The Dark Knight")
	// fmt.Println("Movies in database:", storage.GetAll())

	// using memory
	service1 := MovieStorageService{
		storage: &InMemoryMovieStorage{},
	}
	service1.AddMovie("Inception")
	service1.AddMovie("The Matrix")
	service1.ListMovies()

		// using mockDB
	service2 := MovieStorageService{
		storage: &InDBMovieStorage{},
	}
	service2.AddMovie("Inception")
	service2.AddMovie("The Matrix")
	service2.ListMovies()
}