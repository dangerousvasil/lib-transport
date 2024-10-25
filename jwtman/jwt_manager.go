package jwtman

import (
	"fmt"
	"github.com/gofrs/uuid"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User mapped from table <users>
type User struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OtpSetup  *bool      `gorm:"column:otp_setup" json:"otp_setup"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Roles     []string   `gorm:"column:roles;type:text[]" json:"roles"`
	Username  *string    `gorm:"column:username" json:"username"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// JWTManager is a JSON web token manager
type JWTManager struct {
	secretKey string
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	ID       int64    `json:"id"`
	// scanner UUID для привязки бота к сканнеру источника
	SUUID uuid.UUID `json:"suuid"`
	BotID int64     `json:"botid"`
	// отпи был установлен / false нужно проказать кр код
	SetupOTP bool `json:"sotp"`
	// otp код невведен и нужно его запросить если true значит пользователь уже успешно залогинелся
	OTP bool `json:"otp"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey}
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(
	user *User,
	tokenDuration time.Duration,
	otp bool,
) (string, error) {
	claims := generateUserClaimsForUser(user, otp, tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

//
//func (manager *JWTManager) GenerateForBot(bot *models.Robohammster) (string, error) {
//	var claims UserClaims
//	claims = generateUserClaimsBot(bot)
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(manager.secretKey))
//}
//
//func (manager *JWTManager) GenerateForStream(
//	bot *models.Robohammster,
//	scunnerUUID uuid.UUID,
//) (string, error) {
//	var claims UserClaims
//	claims = generateUserClaimsStream(bot, scunnerUUID)
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(manager.secretKey))
//}

func generateUserClaimsForUser(
	user *User,
	otp bool,
	tokenDuration time.Duration,
) UserClaims {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
		},
		Username: *user.Username,
		Roles:    user.Roles,
		ID:       user.ID,
		OTP:      otp,
	}

	if user.OtpSetup != nil {
		claims.SetupOTP = *user.OtpSetup
	}
	return claims
}

//
//func generateUserClaimsBot(bot *models.Robohammster) UserClaims {
//	return UserClaims{
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: bot.OauthTokenExpiration.Unix(),
//		},
//		Username: bot.Name,
//		Roles:    bot.Roles,
//		ID:       bot.UserID,
//		BotID:    bot.ID,
//	}
//}
//
//// generateUserClaimsStream генерирует UserClaims для бота и сканера
//func generateUserClaimsStream(bot *models.Robohammster, scunnerUUID uuid.UUID) UserClaims {
//	return UserClaims{
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: bot.OauthTokenExpiration.Unix(),
//		},
//		Username: bot.Name,
//		Roles:    bot.Roles,
//		ID:       bot.UserID,
//		BotID:    bot.ID,
//		SUUID:    scunnerUUID,
//	}
//}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
