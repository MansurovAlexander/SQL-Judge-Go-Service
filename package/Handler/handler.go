package handler

import (
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
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
		dbms := api.Group("/dbms")
		{
			dbms.POST("/", h.createDbms)
			dbms.GET("/", h.getAllDbms)
			dbms.GET("/:id", h.getDbmsByID)
		}
		databases := api.Group("/databases")
		{
			databases.POST("/", h.createDataBase)
			databases.GET("/", h.getAllDataBases)
			databases.GET("/:id", h.getDataBaseByID)
		}
		bannedWords := api.Group("/bannedwords")
		{
			bannedWords.POST("/", h.createBannedWord)
			bannedWords.GET("/", h.getAllBannedWords)
		}
		assignes := api.Group("/assignes")
		{
			assignes.POST("/", h.createAssign)
			assignes.GET("/", h.getAllAssignes)
			assignes.GET("/:id", h.getAssignByID)
		}
	}
	return router
}
