package main

import (
	"db"
	"github.com/gin-gonic/gin"
	"net/http"
	"model"
	"middlewares"
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

		applicationContext := router.NewApplicatonContext(c)

		var result []model.ProfileSummary
		var err error

		if result, err = router.ProfileByGroupMethod(groupMethod,startDateString, applicationContext.GetProfiler()); err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
		}else {
			c.JSON(http.StatusOK, gin.H{"groupMethod": groupMethod, "statrDate": startDateString, "result": result})
		}
	})


	r.Run() // listen and serve on 0.0.0.0:8080
}