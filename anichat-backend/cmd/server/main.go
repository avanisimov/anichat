package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/alanis/anichat-backend/internal/config"
	"github.com/alanis/anichat-backend/internal/db"
	"github.com/alanis/anichat-backend/internal/features/auth"
	"github.com/alanis/anichat-backend/internal/features/user"
	localhttp "github.com/alanis/anichat-backend/internal/http"
	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()

	_ = godotenv.Load(".env") // или .env.example для dev

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Default().Printf("config loaded: %+v", cfg.DatabaseURL)
	dbConn, err := db.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	jwtManager := auth.NewJwtManager(cfg.JWTSecret)
	authRepo := auth.NewRepository(dbConn.Pool)
	emailSender := auth.NewEmailSender(cfg.ResendAPIKey)
	authService := auth.NewService(authRepo, emailSender, jwtManager)
	authHandler := auth.NewHandler(authService, jwtManager)

	userRepo := user.NewRepository(dbConn.Pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	r := localhttp.NewRouter(authHandler, userHandler)
	log.Printf("starting backend on port %s", cfg.Port)
	httpErr := http.ListenAndServe(":8080", r)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
