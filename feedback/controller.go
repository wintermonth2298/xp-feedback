package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wintermonth2298/xp-feedback/utils"
)

type Controller struct {
	service *Service
}

func InitRoutes(s *Service, api *gin.RouterGroup) {
	c := Controller{service: s}

	feedbacks := api.Group("/feedback")
	{
		feedbacks.POST("", c.create)
		feedbacks.GET("/csv", c.csv)
	}
}

func (c *Controller) create(ctx *gin.Context) {
	inp := new(Feedback)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := c.service.Create(ctx.Request.Context(), inp); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrResponse{Err: err.Error()})
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) csv(ctx *gin.Context) {
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename=feedbacks.csv")

	csv, err := c.service.CSV(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrResponse{Err: err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "text/csv", csv)
}
