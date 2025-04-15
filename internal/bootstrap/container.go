package bootstrap

import (
	"database/sql"

	"github.com/FabioRg06/Authentify/internal/app/user/adapter/http"
	"github.com/FabioRg06/Authentify/internal/app/user/app"
	"github.com/FabioRg06/Authentify/internal/infrastructure/persistence"
)

type Container struct {
	DB          *sql.DB
	UserHandler *http.UserHandler
}

func NewContainer() (*Container, error) {
	connector := persistence.GetConnector()
	db, err := connector.Connect()
	if err != nil {
		return nil, err
	}

	userRepo := persistence.NewPostgresUserRepository(db)
	userService := app.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	return &Container{
		DB:          db,
		UserHandler: userHandler,
	}, nil
}
