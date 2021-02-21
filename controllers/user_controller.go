package controllers

import (
	"github.com/whary11/restapi/database"

	"github.com/gin-gonic/gin"
)

type Person struct {
	id      int
	name    string
	surname string
}

func HomeHandler(c *gin.Context) {

	id_init := c.Param("id")
	db := database.GetConnection()
	var person Person

	row := db.Raw("select id, primer_nombre name, primer_apellido surname from users where id =?", id_init).Row()
	row.Scan(&person.id, &person.name, &person.surname)

	c.JSON(200, gin.H{"id": person.id, "name": person.name, "surname": person.surname})

}
