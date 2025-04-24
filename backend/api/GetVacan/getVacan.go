package GetVacan

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"context"
	"net/http"
	"time"

	vac "backend/packages"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func GetVacancys(c echo.Context) error {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Удаление вакансии")
	cfg := config.LoadConfig()
	collection, err := db.СonnectCollection(cfg)
	if err != nil {
		logger.Logger.Error("Ошибка при подключении к коллекции: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось подключиться к базе данных",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Logger.Error("Ошибка при получении данных: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось получить данные",
		})
	}
	defer cursor.Close(ctx)

	var vacancies []vac.VacancyStruct
	if err := cursor.All(ctx, &vacancies); err != nil {
		logger.Logger.Error("Ошибка при обработке данных: %v", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Не удалось обработать данные",
		})
	}

	// Возвращаем данные в формате JSON
	logger.Logger.Info("Данные успешно получены")
	return c.JSON(http.StatusOK, vacancies)

}
