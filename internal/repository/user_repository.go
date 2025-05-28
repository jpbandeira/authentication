package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jp/authentication/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Printf("[%v] erro ao criar usuário: %v", ctx.Value("req_id"), err)
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
