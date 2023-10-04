package router

import (
	"backend-food/internal/pkg/handler"
	"backend-food/pkg/infrastucture/db"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/validators"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() {
	var err error
	r.Engine = gin.Default()
	r.DB, err = db.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	r.DB.MigrateDBWithGorm()
	validators.SetUpValidator()
	h := handler.NewHTTPHandler(r.DB)
	hClient := handler.NewHTTPClientHandler(r.DB)
	hAdmin := handler.NewHTTPAdminHandler(r.DB)
	webAPI := r.Engine.Group("/app")
	{
		webAPI.POST("/query", h.Handle)
		musicAPI := webAPI.Group("/account")
		{
			musicAPI.Use(middleware.AuthMiddleware(r.DB))
			{
				musicAPI.POST("/query", hClient.Handle)

				adminAPI := webAPI.Group("/admin")
				{
					adminAPI.Use(middleware.AuthAdminMiddleware(r.DB))
					{
						adminAPI.POST("/query", hAdmin.Handle)
					}
				}
			}
		}

	}
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
