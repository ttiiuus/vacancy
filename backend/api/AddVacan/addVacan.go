package AddVacan

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"context"
	"net/http"

	vac "backend/packages"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func AddVacancy(c echo.Context) error {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Добваление вакансии")
	var a vac.VacancyStruct
	err := c.Bind(&a)
	a.ID = uuid.New().String()
	if err != nil {
		logger.Logger.Error("Ошибка при генерации uuid", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	cfg := config.LoadConfig()
	collection, err := db.СonnectCollection(cfg)
	if err != nil {
		logger.Logger.Error("Ошибка при подключении к коллекции: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось подключиться к базе данных",
		})
	}

	result, err := collection.InsertOne(context.Background(), a)
	if err != nil {
		logger.Logger.Error("Ошибка при вставке данных: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось сохранить данные",
		})
	}
	logger.Logger.Info("Данные успешно сохранены", zap.Any("id", result.InsertedID))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Данные успешно сохранены",
		"id":      result.InsertedID,
	})
}
