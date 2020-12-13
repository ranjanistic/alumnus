package token

import(
	// "github.com/joho/godotenv"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"os"
)

//Claims is
type Claims struct {
	Value string `json:"value"`
	jwt.StandardClaims
}

//IsDev is
var IsDev bool = os.Getenv("APP_ENV") != "prod"
var ssh = []byte(os.Getenv("SSH"))

//Sign is
func Sign(value string) string{
	claims := &Claims{
		Value: value,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(ssh)
	if(err!= nil){
		return ""
	}
	return tokenString
}

//Verify is
func Verify(tokenString string) jwt.MapClaims{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return ssh, nil
	})
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if(err!=nil||ok){ return claims}
	fmt.Println(claims)
	return claims
}

