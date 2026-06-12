package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveOTP(ctx context.Context, email, code string, ttlMinutes int) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRow(ctx,
		`
		INSERT INTO otp_codes(email, otp_hash, expires_at)
		VALUES ($1, $2, NOW() + ($3 || ' minutes')::interval)
		RETURNING id
		`,
		email,
		hashOTP(code),
		strconv.Itoa(ttlMinutes),
	).Scan(&id)

	return id, err
}

type OTPVerifyResult struct {
	Success  bool
	Attempts int
	Verified bool
	Email    string
}

func (r *Repository) VerifyOTP(ctx context.Context, challengeId, code string) (*OTPVerifyResult, error) {

	var result OTPVerifyResult

	err := r.db.QueryRow(
		ctx,
		`
        UPDATE otp_codes
        SET
            attempts = CASE
                WHEN otp_hash <> $2 THEN attempts + 1
                ELSE attempts
            END,
            verified = CASE
                WHEN otp_hash = $2 THEN TRUE
                ELSE verified
            END
        WHERE
            id = $1
            AND verified = FALSE
            AND expires_at > NOW()
        RETURNING
            otp_hash = $2,
            attempts,
            verified,
			email
        `,
		challengeId,
		hashOTP(code),
	).Scan(
		&result.Success,
		&result.Attempts,
		&result.Verified,
		&result.Email,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("OTP not found or already verified")
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func hashOTP(code string) string {
	hash := sha256.Sum256([]byte(code))
	return hex.EncodeToString(hash[:])
}

func (r *Repository) FindUserByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			role,
			created_at
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) SaveUser(ctx context.Context, email string) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users(email, role, created_at)
		VALUES ($1, 'user', NOW())
	`, email)

	return err
}

func (r *Repository) DeleteExpiredOTPs(ctx context.Context) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM otp_codes
		WHERE expires_at < NOW()
	`)

	return err
}

func (r *Repository) CreateUserSession(ctx context.Context, userID int64, refreshToken string, expiresAt time.Time) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO user_sessions(user_id, refresh_token, expires_at)
		VALUES ($1, $2, $3)
	`, userID, refreshToken, expiresAt)

	return err
}
