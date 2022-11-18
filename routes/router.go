package routes

import (
	"dans/controller"
	"dans/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Routes struct {
	DB *gorm.DB
}

func (r *Routes) Setup(port string) {

	app := gin.Default()
	user := app.Group("user")
	{
		usersCtrl := controller.UsersController{DB: r.DB}
		user.POST("/register", usersCtrl.Register)
		user.POST("/login", usersCtrl.Login)
	}
	//list jobs
	job := app.Group("jobs")
	job.Use(utils.Middleware)
	{
		jobCtrl := controller.JobsController{DB: r.DB}
		job.GET("/list", jobCtrl.GetJobs)
		job.GET("/:id", jobCtrl.GetJobsDetail)
	}

	runningOnPort := fmt.Sprintf(":%s", port)
	app.Run(runningOnPort)
}