package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func main() {
	r := gin.Default()

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users", deleteUser)

	if err := r.Run(); err != nil {
		log.Fatalf("Unable to Start Server %v", err)
	}

}

func getUsers(g *gin.Context) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to db, %v", err)
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, address, dob, description, name FROM users")

	if err != nil {
		log.Fatalf("Unable to query to db, %v", err)
	}
	var Users []User
	for rows.Next() {
		var User User
		if err := rows.Scan(&User.Id, &User.Name, &User.Dob, &User.Address, &User.Description); err != nil {
			log.Fatalf("Unable to List users from db, %v", err)
		}

		Users = append(Users, User)
	}

	g.JSON(http.StatusOK, gin.H{"Users": Users})
}
func getUser(g *gin.Context) {
	id := g.Param("id")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to db, %v", err)
	}

	defer conn.Close(context.Background())

	rows := conn.QueryRow(context.Background(), "SELECT id, address, dob, description, name FROM users WHERE id = $1", id)
	var User User
	if err := rows.Scan(&User.Id, &User.Name, &User.Dob, &User.Address, &User.Description); err != nil {
		log.Fatalf("Unable to List users from db, %v", err)
	}

	g.JSON(http.StatusOK, gin.H{"User": "User"})

}
func createUser(g *gin.Context) {}
func updateUser(g *gin.Context) {}
func deleteUser(g *gin.Context) {}
