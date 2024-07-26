package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/PedroReves/wiredcraft-golang/db"
	"github.com/PedroReves/wiredcraft-golang/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(g *gin.Context) {
	conn := db.InitConn()

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, address, dob, description, name FROM users")

	if err != nil {
		log.Fatalf("Unable to query to db, %v", err)
	}
	var Users []model.User
	for rows.Next() {
		var User model.User
		if err := rows.Scan(&User.Id, &User.Name, &User.Dob, &User.Address, &User.Description); err != nil {
			log.Fatalf("Unable to List users from db, %v", err)
		}

		Users = append(Users, User)
	}

	g.JSON(http.StatusOK, gin.H{"Users": Users})
}
func GetUser(g *gin.Context) {
	id := g.Param("id")
	conn := db.InitConn()

	defer conn.Close(context.Background())

	rows := conn.QueryRow(context.Background(), "SELECT id, address, dob, description, name FROM users WHERE id = $1", id)
	var User model.User
	if err := rows.Scan(&User.Id, &User.Name, &User.Dob, &User.Address, &User.Description); err != nil {
		log.Fatalf("Unable to List users from db, %v", err)
	}

	g.JSON(http.StatusOK, gin.H{"User": User})

}
func CreateUser(g *gin.Context) {
	var user model.User

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Bad Request": "There is an error with the request, Try Again!"})
	}

	conn := db.InitConn()

	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "INSERT INTO users (name, dob, description, address) VALUES ($1, $2, $3, $4)", &user.Name, &user.Dob, &user.Address, &user.Description)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": "Unable to Finish Query!"})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"User": user})
}
func UpdateUser(g *gin.Context) {
	id := g.Param("id")

	var user model.User

	conn := db.InitConn()

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Bad Request": "There is an error with the request, Try Again!"})
	}

	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "UPDATE users SET name = $1, dob = $2, address = $3, description = $4 WHERE id = $5", &user.Name, &user.Dob, &user.Address, &user.Description, id)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": "Unable to Finish Query!"})
	}

	g.JSON(http.StatusCreated, gin.H{"Tool": user})

}
func DeleteUser(g *gin.Context) {
	id := g.Param("id")

	conn := db.InitConn()

	defer conn.Close(context.Background())

	query, err := conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)

	if err != nil {
		log.Fatalf("Unable to query to db, %v", err)
	}

	if query.RowsAffected() == 0 {
		g.JSON(http.StatusNotFound, gin.H{"error": "User not found\n"})
		return
	}

	g.JSON(http.StatusNoContent, nil)

}
