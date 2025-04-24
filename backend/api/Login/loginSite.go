package login

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/logger"
	"context"

	hsh "backend/internal/auth"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func LoginSite(c echo.Context) error {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Logger.Info("Логин")
	var login LoginStruct
	if err := c.Bind(&login); err != nil {
		logger.Logger.Error("Ошибка при логине", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Неверный формат данных"})
	}

	cfg := config.LoadConfig()
	collection, err := db.СonnectCollectionSecond(cfg)
	if err != nil {
		logger.Logger.Error("Ошибка при подключении к коллекции: %v", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Ошибка при подключении к базе данных"})
	}

	// Используем правильный фильтр
	filter := bson.M{"login": login.Login}
	result := collection.FindOne(context.TODO(), filter)

	var user LoginStruct
	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Logger.Error("Пользователь был не найден в бд, ему нужно зарегаться", zap.Error(err))
			return c.JSON(404, map[string]string{
				"error": "Пользователь не найден, нужно зарегистрироваться",
			})
		} else {
			logger.Logger.Error("Ошибка при декодировании данных", zap.Error(err))
			return c.JSON(500, map[string]string{
				"error": "Ошибка при декодировании данных",
			})
		}
	}

	// Проверка пароля
	if err := hsh.CheckPassword(user.Password, login.Password); err != nil {
		logger.Logger.Error("Пользователь ввел неверный пароль", zap.Error(err))
		return c.JSON(401, map[string]string{
			"error": "Неверный пароль",
		})
	}
	logger.Logger.Info("Успешно залогинились")
	return c.JSON(200, map[string]string{
		"OK":       "Успешно залогинились, теперь можно работать",
		"Name":     login.Login,
		"Password": login.Password,
	})
}
