package main

import (
	"net/http"
	"os"

	"github.com/agusbasari29/simple-xjx-tasking.git/database"
	"github.com/agusbasari29/simple-xjx-tasking.git/database/seeder"
	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
	"github.com/agusbasari29/simple-xjx-tasking.git/route"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	db = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.Tasks{})
	xjx := gin.Default()
	route.DefineApiRoute(xjx)
	xjx.Use(static.Serve("/", static.LocalFile("static", true)))
	xjx.GET("/db_seeds", func(c *gin.Context) {
		seeder.LetsSeed()
		c.JSON(http.StatusOK, gin.H{"message": "Succesfully seed database"})
	})

	xjx.Run(":" + os.Getenv("PORT"))
}
