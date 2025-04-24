package logger

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init инициализирует логгер
func Init() error {
	// Настройка конфигурации логгера
	config := zap.NewProductionConfig()

	// Настройка уровня логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// Создание логгера
	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	// Заменяем глобальный логгер
	zap.ReplaceGlobals(Logger)

	// Создаем директорию для логов, если её нет
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Sync синхронизирует буферы логгера
func Sync() {
	_ = Logger.Sync()
}

func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			logger.Info("Incoming request",
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.String("remote_addr", c.RealIP()),
				zap.String("path", c.Path()),
			)

			err := next(c)

			logger.Info("Outgoing response",
				zap.Int("status", c.Response().Status),
				zap.String("latency", time.Since(start).String()),
			)
			return err
		}
	}
}
