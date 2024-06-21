package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v4"
	// "encoding/json"
	// "fmt"
	// "log"
	// "log"
	// "net/http"
	// "github.com/MoCarterWoods/carterwoods"
	// "github.com/google/uuid"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	email := claims["email"].(string)
	fmt.Println(email)

	return c.Next()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Book One", Author: &Author{FirstName: "Mo", LastName: "Woods"}})
	books = append(books, Book{ID: 2, Title: "Book Two", Author: &Author{FirstName: "View", LastName: "Woods"}})


	app.Post("/login", login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("os.env.JWT_SECRET")),
	}))

	app.Use(checkMiddleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)

	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./upload/"+ file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}


func testHTML(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func getEnv(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email == "" || user.Password == "" {
		return c.SendString("Please enter email and password")
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Login Failed")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = memberUser.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"message": "Login Success", "token": t})
}



type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User {
    Email:    "mo@view.com",
	Password: "123456",
}



// -------------------------.net/http---------------------------------
// func main() {
	
// 	http.HandleFunc("/hello", SayHello)

// 	fmt.Println("Server listening on port 8080")
// 	if  err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
		
// 	}
// }

// func SayHello (w http.ResponseWriter, r *http.Request)  {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello!")
// }


// func SayHello(w http.ResponseWriter, r *http.Request)  {

// 	id := uuid.New()
// 	fmt.Println("Hello, World!")
// 	fmt.Println("UUID: %s", id)
// 	carterwoods.SayHelloCarter()
// 	carterwoods.SayHelloView()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Number:", i)
		
	// }
// }



// func getBooks(w http.ResponseWriter, r *http.Request) []Book {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// func main()  {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/books", getBooks).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8080", r))
// }