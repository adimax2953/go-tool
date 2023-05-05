package encrypt

// import (
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	internal "gitlab.baifu-tech.net/dsg-game/backend-server/internal/common"
// )

// var CommonTokenSecret = "game"

// var jwtSecret = []byte(CommonTokenSecret)

// type Claims struct {
// 	Uid     int32  `json:"uid"`
// 	Account string `json:"account"`
// 	Name    string `json:"name"`
// 	jwt.StandardClaims
// }

// func GenerateToken(uid int32, account, name string, token string) (string, error) {
// 	nowTime := time.Now()
// 	expireTime := nowTime.Add(internal.ExpiredTime)

// 	claims := Claims{
// 		uid,
// 		account,
// 		name,
// 		jwt.StandardClaims{
// 			IssuedAt:  nowTime.Unix(),
// 			ExpiresAt: expireTime.Unix(),
// 			Issuer:    token,
// 		},
// 	}

// 	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

// 	jwtToken, err := tokenClaims.SignedString(jwtSecret)

// 	return jwtToken, err
// }

// func ParseToken(token string) (*Claims, error) {
// 	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
// 			return claims, nil
// 		}
// 	}
// 	return nil, err
// }

// type CustomClaims struct {
// 	VendorId int32  `json:"vendor_id"`
// 	VendorIp string `json:"vendor_ip"`
// 	BetNo    string `json:"bet_no"`
// 	jwt.StandardClaims
// }

// func GenerateCustomToken(claims CustomClaims) (string, error) {
// 	nowTime := time.Now()
// 	expireTime := nowTime.Add(10 * time.Minute)

// 	claims.StandardClaims = jwt.StandardClaims{
// 		IssuedAt:  nowTime.Unix(),
// 		ExpiresAt: expireTime.Unix(),
// 		Issuer:    claims.VendorIp,
// 	}

// 	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

// 	jwtToken, err := tokenClaims.SignedString(jwtSecret)
// 	return jwtToken, err
// }

// func ParseCustomToken(token string) (*CustomClaims, error) {
// 	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
// 			return claims, nil
// 		}
// 	}
// 	return nil, err
// }
