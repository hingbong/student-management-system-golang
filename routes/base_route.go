package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func SetupRouter(s *http.Server) (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("./static"))

	e.POST("/student_op", addStudentPost)
	e.GET("/student_op", allStudentsGet)
	e.GET("/student_op/:id", oneStudentsGet)
	e.PUT("/student_op", modifyProfessionPut)
	e.DELETE("/student_op/:id", deleteStudent)
	e.Logger.Fatal(e.StartServer(s))
	return
}
