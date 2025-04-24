package register

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"context"

	hsh "backend/internal/auth"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func RegisterSite(c echo.Context) error {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Запуск приложения")
	var reg registerStruct
	if err := c.Bind(&reg); err != nil {
		return c.JSON(400, map[string]string{"error": "Неверный формат данных"})
	}
	// Генерация UUID
	reg.ID = uuid.New().String()
	logger.Logger.Info("Регистрация пользователя",
		zap.String("login", reg.Login),
		zap.String("password", reg.Password),
		zap.String("id", reg.ID),
	)
	reg.Password, _ = hsh.HashPassword(reg.Password)
	cfg := config.LoadConfig()
	collection, err := db.СonnectCollectionSecond(cfg)
	if err != nil {
		logger.Logger.Error("Ошибка при подключении к коллекции: %v", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Ошибка при подключении к базе данных"})
	}

	// Проверка, существует ли пользователь
	filter := bson.M{"login": reg.Login}
	result := collection.FindOne(context.TODO(), filter)
	if result == nil {
		logger.Logger.Error("Пользователь уже зарегистрирован", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Пользователь уже зарегистрирован"})

	}
	// Вставка нового пользователя
	_, err = collection.InsertOne(context.TODO(), reg)
	if err != nil {
		logger.Logger.Error("Ошибка при регистрации пользователя: %v", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Ошибка при регистрации"})
	}
	logger.Logger.Info("Пользователь успешно зарегистрирован",
		zap.String("login", reg.Login),
	)
	return c.JSON(200, map[string]interface{}{
		"OK":       "Успешно зарегистрировались, теперь можно работать",
		"Name":     reg.Login,
		"Password": reg.Password,
		"ID":       reg.ID,
	})
}
