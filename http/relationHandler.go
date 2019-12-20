package http

import (
	"github.com/labstack/echo"
	"github.com/nattapat27/test-golnag-echo/useCase"
	"github.com/nattapat27/test-golnag-echo/helper/respone"
	"net/http"
	"strconv"
)

type relationHandler struct {
	useCase useCase.RelationUseCaseInf
}

func NewRelationHandler(e *echo.Echo, useCase useCase.RelationUseCaseInf){
	handler := &relationHandler{
		useCase:useCase,
	}
	e.GET("/relation/:id", handler.fetchByUserId)
}

func (h *relationHandler) fetchByUserId(context echo.Context) error {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong Input")
	}
	relation, er := h.useCase.FetchByUserId(id)
	if er != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't find data")
	}
	return context.JSON(http.StatusOK, respone.ResponseData("relation", relation))
}