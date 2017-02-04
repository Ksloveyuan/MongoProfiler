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
	"time"
	"gopkg.in/mgo.v2/bson"
)


func main() {

	db.Connect()

	r := gin.Default()

	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	r.GET("/profile/:groupMethod", func(c *gin.Context) {

		var groupID bson.M
		var startDate time.Time
		var result []model.ProfileSummary
		var err error

		groupMethod := strings.ToLower(c.Param("groupMethod"))

		if groupID, err = model.GetGroupID(groupMethod); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if startDate, err = timeutil.ParseDate(c); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		db := c.MustGet("db").(*mgo.Database)

		result,err = model.Profile(db, startDate, groupID)

		if result,err = model.Profile(db, startDate, groupID); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"groupMethod": groupMethod, "statrDate": timeutil.ToString(startDate), "result": result})
	})


	r.Run() // listen and serve on 0.0.0.0:8080
}