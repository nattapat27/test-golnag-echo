package main

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	handler "github.com/nattapat27/test-golnag-echo/http"
	"github.com/nattapat27/test-golnag-echo/repository"
	"github.com/nattapat27/test-golnag-echo/useCase"
	"log"
	"net/http"
)

func main() {


	con := newConnection("postgres://postgres:postgres@10.10.9.12:5432/practice_benz?sslmode=disable")
	defer con.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return  c.String(http.StatusOK, "Hello World")
	})

	userRepo := repository.NewUserRepository(con)
	relationRepo := repository.NewRelationRepository(con)
	userUseCase := useCase.NewUserUseCase(userRepo)
	relationUseCase := useCase.NewRelationUseCase(relationRepo)
	handler.NewUserHandler(e, userUseCase, relationUseCase)
	handler.NewRelationHandler(e, relationUseCase)

	e.Logger.Fatal(e.Start(":1323"))
}

func newConnection(url string) *pg.DB {
	//url := "postgres://postgres:postgres@10.10.9.12:5432/practice_benz?ssl_mode=disable"
	option, err := pg.ParseURL(url)
	if err != nil{
		log.Panic(err)
	}
	db := pg.Connect(option)
	log.Println("connecting")
	return db
}
