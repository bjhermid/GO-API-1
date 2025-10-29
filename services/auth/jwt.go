package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bjhermid/go-api-1/config"
	"github.com/bjhermid/go-api-1/types"
	"github.com/bjhermid/go-api-1/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExperetionInSecond)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	stringToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return stringToken, nil

}

// middleware ?
func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the token from the user request
		tokenString := getTokenFromRequest(r)

		// Validate the JWT
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token : %v", err)
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		//if there is, we need to fect userID from the database (id from the token)

		claims := token.Claims.(jwt.MapClaims)
		userID, err := extractUserID(claims)
		if err!= nil {
			log.Printf("bad userID %w", err)
		}


		u, err := store.GetUserById(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return 
		}
		//set the context "userID" to the user id
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func GetUserIDFromContext(ctx context.Context) int{
	uID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return uID
}

// helper func middleware
func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unxepected signing method : %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func extractUserID(claims jwt.MapClaims) (int, error){
	if s,ok := claims["userID"].(string); ok{
		id, err := strconv.Atoi(s)
		if err != nil {
			return 0 , fmt.Errorf("invalid userID string: %v", err)
		}
		return id, nil
	}

	if f, ok := claims["userID"].(float64) ; ok {
		return int(f), nil
	}
	return 0, fmt.Errorf("userID claims not found or invalid tyoe")
}