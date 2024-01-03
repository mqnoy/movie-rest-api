package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/domain"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/handler"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
	"github.com/mqnoy/movie-rest-api/pkg/cvalidator"
	"github.com/mqnoy/movie-rest-api/util"
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
	var payload dto.MovieCreatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	if err := cvalidator.Validator.Struct(&payload); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	param := dto.MovieCreateParam{
		Payload: payload,
	}

	result, err := h.movieUseCase.CreateMovie(ctx, param)
	if err != nil {
		handler.ParseToErrorMsg(ctx, err.StatusCode, err.Err)
		return
	}

	handler.ParseResponse(ctx, "Create movie successfully", result, nil)
}

func (h *movieHandler) GetDetailMovie(ctx *gin.Context) {
	var req dto.MovieDetailParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, cerror.ErrRequiredId))
		return
	}

	param := dto.MovieDetailParam{
		ID: req.ID,
	}

	result, err := h.movieUseCase.DetailMovie(ctx, param)
	if err != nil {
		handler.ParseToErrorMsg(ctx, err.StatusCode, err.Err)
		return
	}

	handler.ParseResponse(ctx, "Get detail movie Successfully", result, nil)
}

func (h *movieHandler) GetListMovie(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	title := util.ParseQueryToString(ctx.Query("title"))
	orders := util.ParseQueryToString(ctx.DefaultQuery("orders", "title asc"))

	params := dto.ListParam[dto.FilterMovieParams]{
		Filters: dto.FilterMovieParams{
			Title: title,
		},
		Orders: orders,
		Pagination: dto.Pagination{
			Page:   page,
			Limit:  limit,
			Offset: offset,
		},
	}

	result, err := h.movieUseCase.ListMovies(ctx, params)
	if err != nil {
		handler.ParseToErrorMsg(ctx, err.StatusCode, err.Err)
		return
	}

	handler.ParseResponse(ctx, "Successfully", result, nil)
}

func (h *movieHandler) PatchUpdateMovie(ctx *gin.Context) {
	var req dto.MovieDetailParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, cerror.ErrRequiredId))
		return
	}

	var payload dto.MovieUpdatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	if err := cvalidator.Validator.Struct(&payload); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	payload.ID = req.ID
	param := dto.MovieUpdateParam{
		Payload: payload,
	}

	result, err := h.movieUseCase.UpdateMovie(ctx, param)
	if err != nil {
		handler.ParseToErrorMsg(ctx, err.StatusCode, err.Err)
		return
	}

	handler.ParseResponse(ctx, "Update movie successfully", result, nil)
}

func (h *movieHandler) DeleteMovie(ctx *gin.Context) {
	var req dto.MovieDetailParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		handler.ParseResponse(ctx, "", nil, cerror.WrapError(http.StatusBadRequest, cerror.ErrRequiredId))
		return
	}

	param := dto.MovieDetailParam{
		ID: req.ID,
	}

	if err := h.movieUseCase.RemoveMovie(ctx, param); err != nil {
		handler.ParseToErrorMsg(ctx, err.StatusCode, err.Err)
		return
	}

	handler.ParseResponse(ctx, "Successfully", nil, nil)
}
