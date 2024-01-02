package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/domain"
)

type movieHandler struct {
	g            *gin.RouterGroup
	movieUseCase domain.MovieUseCase
}

func New(g *gin.RouterGroup, movieUseCase domain.MovieUseCase) {
	handler := movieHandler{
		g:            g,
		movieUseCase: movieUseCase,
	}

	route := g.Group("/movies")
	route.GET("/:id", handler.GetDetailMovie)
	route.PATCH("/:id", handler.PatchUpdateMovie)
	route.DELETE("/:id", handler.DeleteMovie)
	route.POST("/", handler.PostCreateMovie)
	route.GET("/", handler.GetListMovie)

}

func (h *movieHandler) PostCreateMovie(ctx *gin.Context) {
	// TODO: function create movie
	// TODO: validate payload
	// TODO: call useCase createMovie
}

func (h *movieHandler) GetDetailMovie(ctx *gin.Context) {
	// TODO: function delete movie
	// TODO: validate param id
	// TODO: call useCase detailMovie
}

func (h *movieHandler) GetListMovie(ctx *gin.Context) {
	// TODO: function delete movie
	// TODO: call useCase listMovie
}

func (h *movieHandler) PatchUpdateMovie(ctx *gin.Context) {
	// TODO: function update movie
	// TODO: validate payload
	// TODO: call useCase updateMovie
}

func (h *movieHandler) DeleteMovie(ctx *gin.Context) {
	// TODO: function delete movie
	// TODO: validate param id
	// TODO: call useCase deleteMovie
}
