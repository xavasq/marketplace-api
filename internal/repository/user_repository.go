package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang-store/internal/domain"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *zap.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (email, password, role, balance)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query, user.Email, user.Password, user.Role, user.Balance).Scan(&user.ID)
	if err != nil {
		r.logger.Error("не удалось создать пользователя", zap.Error(err), zap.String("email", user.Email))
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT id, email, password, role, balance
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Balance)
	if err != nil {
		r.logger.Error("не удалось получить пользователя по ID", zap.Error(err), zap.Int64("id", id))
		return nil, fmt.Errorf("ошибка получения пользователя по ID: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password, role, balance
		FROM users
		WHERE email = $1
	`

	var user domain.User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Balance)
	if err != nil {
		r.logger.Error("не удалось получить пользователя по email", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("ошибка получения пользователя по email: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	query := `
		SELECT id, email, password, role, balance
		FROM users
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("не удалось получить список пользователей", zap.Error(err))
		return nil, fmt.Errorf("ошибка получения списка пользователей: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Balance); err != nil {
			r.logger.Error("ошибка при сканировании строки пользователя", zap.Error(err))
			return nil, fmt.Errorf("ошибка чтения данных пользователя: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("ошибка при чтении всех пользователей", zap.Error(err))
		return nil, fmt.Errorf("ошибка итерации по пользователям: %w", err)
	}
	return users, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("не удалось удалить пользователя", zap.Error(err), zap.Int64("id", id))
		return fmt.Errorf("ошибка при удалении пользователя: %w", err)
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int64, email, password, role string, balance int) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, role = $3, balance = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(ctx, query, email, password, role, balance, id)
	if err != nil {
		r.logger.Error("не удалось обновить пользователя", zap.Error(err), zap.Int64("id", id))
		return fmt.Errorf("ошибка обновления пользователя: %w", err)
	}
	return nil
}
