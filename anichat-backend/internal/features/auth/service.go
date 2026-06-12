package auth

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo        *Repository
	emailSender *EmailSender
	jwtManager  *JwtManager
}

var ErrInvalidOTP = errors.New("invalid otp code")
var NoUserError = errors.New("user not found")

func NewService(
	repo *Repository,
	emailSender *EmailSender,
	jwtManager *JwtManager,
) *Service {

	return &Service{
		repo:        repo,
		emailSender: emailSender,
		jwtManager:  jwtManager,
	}
}

func (s *Service) LoginByEmail(context context.Context, email string) (*uuid.UUID, error) {
	code, err := GenerateOTP()
	if err != nil {
		return nil, err
	}
	log.Printf("Generated OTP for %s: %s", email, code) // Логируем код для отладки
	ticketId, err := s.repo.SaveOTP(context, email, code, 10)
	if err != nil {
		log.Printf("Failed to save OTP for %s: %v", email, err) // Логируем ошибку сохранения OTP
		return nil, err
	}
	log.Printf("Saved OTP for %s in database with id: %s", email, ticketId) // Логируем успешное сохранение OTP
	// if err := s.emailSender.SendOtp(email, code); err != nil {
	// 	return err
	// }
	log.Printf("Sent OTP email to %s", email) // Логируем успешную отправку письма
	return &ticketId, nil
}

func (s *Service) VerifyLoginOtp(context context.Context, challengeId string, code string) (*string, *string, error) {
	result, err := s.repo.VerifyOTP(context, challengeId, code)
	log.Printf("Verifying OTP for %s: code=%s, error=%v", challengeId, code, err) // Логируем результат проверки OTP
	if err != nil {
		return nil, nil, err
	}

	if !result.Success {
		return nil, nil, ErrInvalidOTP
	}
	email := result.Email
	user, err := s.repo.FindUserByEmail(context, email)
	log.Printf("Finding user by email %s: user=%+v, error=%v", email, user, err) // Логируем результат поиска пользователя

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("User with email %s not found, creating new user", email) // Логируем создание нового пользователя
			if err := s.repo.SaveUser(context, email); err != nil {
				log.Printf("Failed to save new user with email %s: %v", email, err) // Логируем ошибку сохранения нового пользователя
				return nil, nil, err
			}
			user, err = s.repo.FindUserByEmail(context, email) // И снова пытаемся найти пользователя после создания
			if err != nil {
				log.Printf("Failed to find user after creation with email %s: %v", email, err) // Логируем ошибку поиска пользователя после создания
				return nil, nil, err
			}
		} else {
			log.Printf("Error finding user with email %s: %v", email, err) // Логируем ошибку поиска пользователя, если она не связана с отсутствием пользователя
			return nil, nil, err
		}
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user) // Генерируем токены для пользователя
	if err != nil {
		log.Printf("Failed to generate access token for user %s: %v", email, err)
		return nil, nil, err
	}
	refreshToken, err := s.jwtManager.GenerateRefreshToken()
	if err != nil {
		log.Printf("Failed to generate refresh token for user %s: %v", email, err)
		return nil, nil, err
	}
	log.Printf("Generated tokens for user %s: accessToken=%s, refreshToken=%s", email, accessToken, refreshToken) // Логируем успешную генерацию токенов

	return &accessToken, &refreshToken, nil
}
