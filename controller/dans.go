package controller

import (
	"dans/host/dans"
	"dans/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)
type JobsController struct {
	DB *gorm.DB
}

func (controller *JobsController) GetJobs(ctx *gin.Context)  {
	page := ctx.Query(`page`)
	desc := ctx.Query("description")
	fullTime := ctx.Query("full_time")
	qParams := dans.QueryListJob{
		Page:        page,
		Description: desc,
		FullTime:    fullTime,
	}

	jobs, err := dans.GetListJob(qParams)
	if err != nil{
		res := utils.Response(http.StatusInternalServerError, "service dans not available", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := utils.Response(http.StatusOK, "Success", jobs)
	ctx.JSON(http.StatusOK, res)
	return
}

func (controller *JobsController) GetJobsDetail(ctx *gin.Context)  {
	Id := ctx.Params.ByName("id")

	job, err := dans.GetListJobDetail(Id)
	if err != nil{
		res := utils.Response(http.StatusInternalServerError, "service dans not available", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := utils.Response(http.StatusOK, "Success", job)
	ctx.JSON(http.StatusOK, res)
	return
}
