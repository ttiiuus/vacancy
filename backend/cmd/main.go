package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	adds "backend/api/AddVacan"
	dels "backend/api/DeleteVacan"
	gets "backend/api/GetVacan"
	logger "backend/internal/logger"

	log "backend/api/Login"
	reg "backend/api/Register"
)
//реализовать подключение к БД и поднятие http сервера
func main() {
	// Echo instance
	e := echo.New()

	//init logger
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Запуск приложения")
	// Middleware
	m := logger.ZapLogger(logger.Logger)
	e.Use(m)
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	e.GET("/api", api)

	e.POST("/vacancyadd", adds.AddVacancy)

	e.GET("/api/vacancy", gets.GetVacancys)

	e.DELETE("/api/vacancy", dels.DelVacancy)

	e.POST("/login", log.LoginSite)

	e.POST("/register", reg.RegisterSite)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Fatal("shutting down the server", zap.Error(err))
	}
}

// Handler
func hello(c echo.Context) error {
	content, err := os.ReadFile("C:\\golanglearn\\mysite\\backend\\cmd\\index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read HTML file")
	}
	return c.HTML(http.StatusOK, string(content))
}

// функция проверки работы api
func api(c echo.Context) error {
	return c.String(http.StatusOK, "this is api site")
}
