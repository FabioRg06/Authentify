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
		is_active BOOLEAN NOT NULL
	);`)
	if err != nil {
		log.Fatalf("❌ Failed to create users table: %v", err)
	}
}
func (r *PostgresUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (username, email, password, role, is_active)
		VALUES ($1, $2, $3, $4, $5)`,
		user.Username, user.Email, user.Password, user.Role, true)
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
