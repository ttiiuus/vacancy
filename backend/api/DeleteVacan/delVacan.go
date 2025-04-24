package DelVac

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"

	vac "backend/packages"
)

func DelVacancy(c echo.Context) error {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Удаление вакансии")
	// Удаление вакансии
	var a vac.VacancyStruct
	err := c.Bind(&a)
	if err != nil {
		logger.Logger.Error("Ошибка при удалении вакансии", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Неверный формат данных",
		})
	}
	cfg := config.LoadConfig()
	collection, err := db.СonnectCollection(cfg)
	if err != nil {
		logger.Logger.Error("Ошибка при подключении к коллекции: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось подключиться к базе данных",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	filter := bson.M{"id": a.ID}
	inf, err := collection.DeleteOne(ctx, filter)
	logger.Logger.Info("Удаление вакансии", zap.Any("id", a.ID))
	if err != nil {
		logger.Logger.Error("Ошибка при удалении обьекта: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось удалить обьект",
		})
	}
	logger.Logger.Info("Обьект успешно удален", zap.Any("id", a.ID))
	return c.JSON(http.StatusOK, inf)
}
