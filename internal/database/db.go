package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang-store/internal/config"
)

type DBPool struct {
	Pool   *pgxpool.Pool
	logger *zap.Logger
}

func ConnectDB(ctx context.Context, logger *zap.Logger) (*DBPool, error) {
	cfg, err := config.LoadEnv()
	if err != nil {
		logger.Error("проблема c загрузкой конфигурации", zap.Error(err))
		return nil, fmt.Errorf("ошибка в конфигурации: %w", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error("ошибка при подключении к базе данных", zap.Error(err), zap.String("dsn", dsn))
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		logger.Error("ошибка при проверке соединения с базой данных", zap.Error(err))
		return nil,
			fmt.Errorf("не удалось проверить соединение с БД: %w", err)
	}
	logger.Info("база данных подключена успешно")
	return &DBPool{Pool: pool, logger: logger}, nil
}

func (db *DBPool) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		db.logger.Info("соединение с базой данных успешно закрыто")
	}
}
