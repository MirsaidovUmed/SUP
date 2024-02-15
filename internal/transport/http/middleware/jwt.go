// middleware/middleware.go
package middleware

import (
	"SUP/pkg/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
	"net/http"
	"strings"
	"time"
)

func (mw *Middleware) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp response.Response

		token := r.Header.Get("Authorization")

		tokenList := strings.Split(token, " ")

		if len(tokenList) != 2 {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}
		token = tokenList[1]

		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(mw.config.JwtSecretKey), nil
		})

		if err != nil {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		if !tokenObj.Valid {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		tokenMap, ok := tokenObj.Claims.(jwt.MapClaims)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		expireDate, ok := tokenMap["exp"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		if time.Now().Unix() > int64(expireDate) {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		userID, ok := tokenMap["user_id"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		roleID, ok := tokenMap["role_id"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		// Проверка роли (в данном случае, что не является админом)
		if strings.Contains(r.URL.String(), "/product/") && roleID != 1 {
			resp = response.Forbidden
			resp.WriteJSON(w)
			return
		}

		if strings.Contains(r.URL.String(), "/project/") && roleID != 2 {
			resp = response.Forbidden
			resp.WriteJSON(w)
			return
		}

		err = mw.service.CheckUserById(int(userID))

		if err != nil {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		context.Set(r, "role_id", int64(roleID))
		context.Set(r, "user_id", int64(userID))

		next.ServeHTTP(w, r)
	})
}
