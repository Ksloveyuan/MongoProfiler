package main

import (
	"db"
	"github.com/gin-gonic/gin"
	"net/http"
	"model"
	"middlewares"
	"gopkg.in/mgo.v2"
	"strings"
	"timeutil"
	"router"
)


func main() {

	db.Connect()

	r := gin.Default()

	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	r.GET("/profile/:groupMethod", func(c *gin.Context) {
		groupMethod := strings.ToLower(c.Param("groupMethod"))
		startDateString := c.DefaultQuery("startDate", timeutil.ToString(timeutil.LastYearOfToday()))
		db := c.MustGet("db").(*mgo.Database)

		profileRouter := router.ProfileRouter{
			DB:db,
			GroupMethod:groupMethod,
			StartDateString:startDateString,
		}

		var result []model.ProfileSummary
		var err error

		if result, err = profileRouter.ProfileByGroupMethod(); err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"groupMethod": groupMethod, "statrDate": startDateString, "result": result})
	})


	r.Run() // listen and serve on 0.0.0.0:8080
}