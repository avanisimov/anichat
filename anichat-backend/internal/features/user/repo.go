package user

import (
	"context"
	"errors"

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

type UsersMeDTO struct {
	ID             int64
	Email          string
	ProfileCreated bool

	FirstName *string
	LastName  *string
	AvatarURL *string
}

func (r *Repository) GetProfileByUserID(
	ctx context.Context,
	userID string,
) (*UsersMeDTO, error) {

	query := `
		SELECT
			u.id,
			u.email,
			p.first_name,
			p.last_name,
			p.avatar_url
		FROM users u
		LEFT JOIN user_profiles p
			ON p.user_id = u.id
		WHERE u.id = $1
	`

	var dto UsersMeDTO

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&dto.ID,
		&dto.Email,
		&dto.FirstName,
		&dto.LastName,
		&dto.AvatarURL,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	dto.ProfileCreated =
		dto.FirstName != nil &&
			dto.LastName != nil

	return &dto, nil
}

func (r *Repository) UpsertProfile(
	ctx context.Context,
	userID string,
	firstName string,
	lastName string,
) error {

	query := `
		INSERT INTO user_profiles (
			user_id,
			first_name,
			last_name
		)
		VALUES ($1,$2,$3)
		ON CONFLICT (user_id)
		DO UPDATE SET
			first_name = EXCLUDED.first_name,
			last_name = EXCLUDED.last_name,
			updated_at = NOW()
	`

	_, err := r.db.Exec(
		ctx,
		query,
		userID,
		firstName,
		lastName,
	)

	return err
}

func (r *Repository) UpdateAvatar(
	ctx context.Context,
	userID string,
	avatarURL string,
) error {

	query := `
		UPDATE user_profiles
		SET
			avatar_url = $2,
			updated_at = NOW()
		WHERE user_id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		userID,
		avatarURL,
	)

	return err
}
