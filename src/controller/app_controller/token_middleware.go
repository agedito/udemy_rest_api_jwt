package app_controller

import (
	"agedito/udemy/rest_api_jwt/service/token"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var NoValidTokenError = errors.New("no valid token")

func (c *AppController) TokenMiddleware(callback http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestToken, tokenErr := c.getTokenFromRequest(r)
		if utils.AssertError(tokenErr) {
			c.responseError(w, http.StatusUnauthorized, NoValidTokenError)
			return
		}
		fmt.Println("TOKEN", requestToken)
		callback.ServeHTTP(w, r)
	}
}

func (c *AppController) getTokenFromRequest(r *http.Request) (token.Token, error) {
	authHeader := r.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer") {
		return token.Token{}, NoValidTokenError
	}
	bearerToken := strings.Split(authHeader, "Bearer ")[1]
	finalToken, tokenErr := token.NewFromId(bearerToken)

	if utils.AssertError(tokenErr) {
		return token.Token{}, NoValidTokenError
	}
	return finalToken, nil

}
