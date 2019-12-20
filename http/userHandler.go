package http

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/nattapat27/test-golnag-echo/helper/respone"
	"github.com/nattapat27/test-golnag-echo/model"
	"github.com/nattapat27/test-golnag-echo/useCase"
	"log"
	"net/http"
	"strconv"
)
type userHandler struct {
	userUseCase useCase.UserUseCaseInf
	relation useCase.RelationUseCaseInf
}

func NewUserHandler(e *echo.Echo, useCase useCase.UserUseCaseInf, relationUseCase useCase.RelationUseCaseInf)  {
	handler := &userHandler{
		userUseCase:useCase,
		relation:relationUseCase,
	}
	e.POST("/user/add", handler.addUser)
	e.GET("/user/:id", handler.findUserById)
	e.GET("/user", handler.findAll)
}

func (u *userHandler)addUser(context echo.Context) error {
	url, _ := context.FormParams()
	jsonData := url.Get("user")
	var userData map[string]interface{}
	log.Println(userData)
	err := json.Unmarshal([]byte(jsonData), &userData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user :=model.NewUserWithParam(userData, nil)
	err = u.userUseCase.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i:= range user.Relation{
		user.Relation[i].UserId = user.Id
		log.Println(user.Relation[i].UserId)
		err = u.relation.Create(user.Relation[i])
		if err != nil{
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return context.JSON(http.StatusOK, respone.ResponseData("user", user))
}

func (u *userHandler)findUserById(context echo.Context) error{
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong request")
	}
	user, err := u.userUseCase.FetchOne(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Relation, err = u.relation.FetchByUserId(user.Id)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, respone.ResponseData("user", user))
}

func (u *userHandler) findAll(context echo.Context) error {
	users, err := u.userUseCase.FetchAll()
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	for i := range users{
		users[i].Relation, err = u.relation.FetchByUserId(users[i].Id)
		if err != nil{
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return context.JSON(http.StatusOK, respone.ResponseData("user", users))
}


