package persistence

import (
	"database/sql"
	"log"

	"github.com/FabioRg06/Authentify/internal/app/user/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	repo := &PostgresUserRepository{db: db}
	repo.initSchema()
	return repo
}
func (r *PostgresUserRepository) initSchema() {
	_, err := r.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`)
	if err != nil {
		log.Fatalf("❌ Failed to create users table: %v", err)
	}
}
func (r *PostgresUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (username, email, password, role)
		VALUES ($1, $2, $3, $4)`,
		user.Username, user.Email, user.Password, user.Role)
	return err
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, username, email, password, role, is_active FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *PostgresUserRepository) Get() ([]*domain.User, error) {
	rows, err := r.db.Query(`
		SELECT id, username, email, password, role, is_active, created_at, updated_at
		FROM users ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
