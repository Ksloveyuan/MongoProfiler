package router

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"model"
)

type ApplicationContext struct {
	ginContext *gin.Context
}

func NewApplicatonContext(ginContext *gin.Context) ApplicationContext {
	return ApplicationContext{ginContext}
}

func (context ApplicationContext) GetMongoDataSource() (*mgo.Database){
	return context.ginContext.MustGet("db").(*mgo.Database)
}

func (context ApplicationContext) GetProfiler() (model.Profiler)  {
	mongoDataSource := model.NewMongoDataSource(context.GetMongoDataSource())
	return mongoDataSource
}