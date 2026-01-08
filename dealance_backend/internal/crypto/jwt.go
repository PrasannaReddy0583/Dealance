package crypto

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	issuer     string
	audience   string
	ttl        time.Duration
}

type Claims struct {
	UserID   string `json:"sub"`
	Role     string `json:"role"`
	KYCLevel string `json:"kyc_level"`
	jwt.RegisteredClaims
}

func NewJWTManager(
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
	issuer string,
	audience string,
	ttl time.Duration,
) *JWTManager {
	return &JWTManager{
		privateKey: privateKey,
		publicKey:  publicKey,
		issuer:     issuer,
		audience:   audience,
		ttl:        ttl,
	}
}

func (m *JWTManager) GenerateToken(
	userID string,
	role string,
	kycLevel string,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID:   userID,
		Role:     role,
		KYCLevel: kycLevel,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Audience:  jwt.ClaimStrings{m.audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        uuid.NewString(), // jti
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.privateKey)
}

func (m *JWTManager) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodRS256 {
				return nil, errors.New("unexpected signing method")
			}
			return m.publicKey, nil
		},
		jwt.WithAudience(m.audience),
		jwt.WithIssuer(m.issuer),
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
