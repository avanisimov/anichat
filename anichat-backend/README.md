# Anichat Backend

This folder contains the backend service for the messenger product.

## Features included

- Go HTTP server
- PostgreSQL connection
- Google OAuth 2.0 authentication
- JWT-based session tokens
- Docker support with `Dockerfile` and `docker-compose.yml`

## Run locally

1. Copy environment variables:
   ```sh
   cp .env.example .env
   ```
2. Set `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, and `JWT_SECRET` in `.env`.
3. Start services:
   ```sh
   docker compose up --build
   ```
4. Open the auth login endpoint:
   ```
   http://localhost:8080/auth/google/login
   ```

## Notes

The basic backend layout is intended as a starting point. Add application routes, message persistence, and mobile/web APIs in the `internal/` packages.
