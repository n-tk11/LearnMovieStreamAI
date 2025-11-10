package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/n-tk11/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	controller "github.com/n-tk11/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/n-tk11/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/addmovie", controller.AddMovie(client))
	router.GET("/recommendedmovies", controller.GetRecommededMovies(client))
	router.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate(client))
}
