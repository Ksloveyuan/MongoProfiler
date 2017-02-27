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
	"fmt"
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

		profiler := applicationContext.GetProfiler()

		if result, err = router.ProfileByGroupMethod(groupMethod,startDateString, profiler); err != nil{
			fmt.Println("An error happens: %+v", err)
			c.JSON(http.StatusBadRequest, err.Error())
		}else {
			c.JSON(http.StatusOK, gin.H{"groupMethod": groupMethod, "statrDate": startDateString, "result": result})
		}
	})


	r.Run() // listen and serve on 0.0.0.0:8080
}