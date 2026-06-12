package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string
const UserIDKey contextKey = "user_id"

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
    jwt.RegisteredClaims
}

type JwtManager struct {
    secret string
}

func NewJwtManager(secret string) *JwtManager {
    return &JwtManager{secret: secret}
}

func (jm *JwtManager) GenerateAccessToken(userId int64, role string) (string, error) {
    claims := JWTClaims{
        UserID: strconv.FormatInt(userId, 10),
        Role:   role,  
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(jm.secret))
}

func HashToken(token string) string {
    hash := sha256.Sum256([]byte(token))
    return hex.EncodeToString(hash[:])
}

func (jm *JwtManager) GenerateRefreshToken() (string, error) {
    b := make([]byte, 32)

    if _, err := rand.Read(b); err != nil {
        return "", err
    }

    return base64.RawURLEncoding.EncodeToString(b), nil
}

func (jm *JwtManager) VerifyToken(tokenStr string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(jm.secret), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok || !token.Valid {
        return nil, jwt.ErrInvalidKey
    }

    return claims, nil
}
