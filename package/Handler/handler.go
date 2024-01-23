package handler

import (
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		submissions := api.Group("/submissions")
		{
			submissions.GET("/", h.getAllSubmissions)
			submissions.GET("/:id", h.getSubmissionByID)
			check := submissions.Group("/check")
			{
				check.POST("/:id", h.checkSubmission)
			}
		}
		databases := api.Group("/databases")
		{
			databases.GET("/", h.getAllDataBases)
			databases.GET("/:id", h.getDataBaseByID)
		}
	}
	return router
}
