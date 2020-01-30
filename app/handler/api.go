package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
	"github.com/taka011002/go_sample_api_server/app/infra"
	"github.com/taka011002/go_sample_api_server/app/infra/persistence"
	"net/http"
	"os"
	"time"
)

func ApiHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := entity.ApiRequestLog{}
		log.Method = r.Method
		log.Path = r.URL.Path
		if loginUser, _ := GetLoginUser(r); loginUser != nil {
			log.UserId = loginUser.Id
		}

		pre := persistence.NewApiRequestLogPersistence(infra.GetDB())
		logService := service.NewApiRequestLogService(pre)
		_ = logService.Create(log)

		handler.ServeHTTP(w, r)
	}
}

func AuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		JwtMiddleware(handler).ServeHTTP(w, r)
	}
}

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("SECRET_KEY"))
			return b, nil
		})

		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if token.Valid {
			// レスポンスを返す
			next.ServeHTTP(w, r)
		} else {
			respondError(w, http.StatusUnauthorized, "token is invalid")
			return
		}
	}
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GetLoginUser(r *http.Request) (*entity.User, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SECRET_KEY"))
		return b, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userPersistence := persistence.NewUserPersistence(infra.GetDB())
		userService := service.NewUserService(userPersistence)
		return userService.GetByUsername(fmt.Sprintf("%s",claims["user"]))
	} else {
		return nil, nil
	}
}