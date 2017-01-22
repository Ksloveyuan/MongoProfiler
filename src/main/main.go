package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"model"
	"middlewares"
	"gopkg.in/mgo.v2"
	"strings"
	"timeutil"
)


func main() {
	r := gin.Default()

	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	r.GET("/profile/:groupMethod", func(c *gin.Context) {
		groupMethod := strings.ToLower(c.Param("groupMethod"))
		groupID, ok := model.GetGroupID(groupMethod)

		if !ok {
			c.JSON(http.StatusBadRequest, "The group method is not supported")
			return
		}

		startDate, err := timeutil.ParseDate(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid time")
			return
		}

		db := c.MustGet("db").(*mgo.Database)

		result,err := model.Profile(db, startDate, groupID)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}else {
			c.JSON(http.StatusOK, gin.H{"groupMethod": groupMethod, "statrDate": timeutil.ToString(startDate), "result": result})
		}
	})


	r.Run() // listen and serve on 0.0.0.0:8080
}