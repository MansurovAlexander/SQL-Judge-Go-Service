package handler

import (
	middleware "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Middleware"
	service "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(server_ip_4, server_ip_6 string) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	whitelist := make(map[string]bool)
	whitelist[server_ip_4] = true
	whitelist[server_ip_6] = true

	router.Use(middleware.IPWhiteList(whitelist))
	api := router.Group("/api")
	{
		submissions := api.Group("/submissions")
		{
			submissions.GET("/", h.getAllSubmissions)
			submissions.GET("/:student_id/:assign_id", h.getSubmissionByID)
			check := submissions.Group("/check")
			{
				check.POST("/", h.checkSubmission)
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
			databases.DELETE("/:id", h.deleteDataBaseByID)
		}
		assignes := api.Group("/assignes")
		{
			assignes.POST("/", h.createAssign)
			assignes.POST("/:id", h.updateAssign)
			assignes.DELETE("/:id", h.deleteAssign)
			assignes.GET("/", h.getAllAssignes)
			assignes.GET("/:id", h.getAssignByID)
		}
		statuses := api.Group("/statuses")
		{
			statuses.GET("/", h.getAllStatuses)
		}
	}
	return router
}
