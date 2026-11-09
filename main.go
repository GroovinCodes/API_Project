package main

// The main thing we'll do with this API is perform CRUD operations like retrieval, edit, and more.
import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// we'll be using a type struct to store our books
// Ensure the vars have uppercase starts to allow outside packages to gain access and edit it (Makes them exported and accessible to outside packages/public)
// After adding things to the struct, we will add additional things to serialize the struct and convert it to JSON for our API
type Book struct {
	Id       string `json:"id"` // The json statement here acts as an identifier for when we're serializing between our api and program.
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
} // not making our struct vars public will return an empty json

// Create a slice of the struct to act as our database
var books = []Book{
	{Id: "1", Title: "Diary of a Wimpy Kid", Author: "Jeff Kingsley", Quantity: 10},
	{Id: "2", Title: "Big Nate", Author: "Michael Jordan", Quantity: 4},
	{Id: "3", Title: "Caramel", Author: "Tommy Felps", Quantity: 6},
}

// This func will return the json ver. of all our books (get requests)
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// In main, we'll set up our router which will handle the different routes and end points for our API
// GET == Getting something, POST == Adding info/creating something new, PATCH == Updating something
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById) // defaults to a string type
	router.POST("/books", createBooks)
	router.PATCH("/checkout", checkoutBook) // updates something that already exists
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}

// Once you've finished the router and get request implementation, you can test the API. Run your go file and use the "curl" command in another terminal
// Once you've run the command: curl localhost:8080/books, you will have your books be printed out in their json format.
// In the other terminal, it shows us the status code 200 and the successful getting of our BOOK struct values
// We use ctrl-C to quit the api scene.

// Now we'll add an endpoint to add books to the slice, or a post request
func createBooks(c *gin.Context) { // This var here stores everything related to a specific request
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil { // Binds the json of the request to the book obj
		return
	} // in case err != nil, we get an error. BindJSON will send an error response once we make it to the return
	// in the case the error hadn't occurred
	books = append(books, newBook) // adds to the books slice
	c.IndentedJSON(http.StatusCreated, newBook)

}

// after implementation, we'll create a json file with the new book to be passed. You can automate this with python, possibly with random values.
// After finishing your json file with your new book data, we'll write a new command in the cmd to append the new book into our slice
// write: curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
// Then provide the original localhost statement to retrieve the new list (curl localhost:8080/books)

// Helper func to retrieve books by ID
func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil // &books allows us to modify the fields of the struct from this function
		}
	}
	return nil, errors.New("Book not found!")
}

// Func to provide book to getBookById
func bookById(c *gin.Context) {
	id := c.Param("id") // Path parameter
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) // err message. Allows for custom json
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// Now return to main and set up the route router.GET(/books/:id)
// Now we'll complete our final two operations to add and subtract the quantity from the books

// if you notice, these are basically early if statements to exit early
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // creating a query parameter
	if ok == false {           // tells us if the key exists via boolean
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) // err message. Allows for custom json
		return
	}

	if book.Quantity <= 0 { // tells us if the key exists via boolean
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No more books to check out"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// now in main, add a patch request
// Then add the following command: curl localhost:8080/checkout?id=2 --request "PATCH"
// Your count will update in real time, play around with it

// now to return books
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // creating a query parameter
	if ok == false {           // tells us if the key exists via boolean
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) // err message. Allows for custom json
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// Now add a patch request within main
// Then you are done

// Ways to challenge yourself
// 1. Implement a real database here and try to connect it to work with golang
// 2. Try to implement a python script to see if you can have a body json randomly be generated
// 3. Try adding an if to prevent your ability to overreturn the original amount of books that you possess
// 4. Try doing as much as you can from memory and not by the video
// Remember to read the documentation!
