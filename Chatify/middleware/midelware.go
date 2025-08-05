package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func WebSocketAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(jwtSecret) == 0 {
			InitAuthMiddleware()
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			authHeader = c.Query("Authorization")
			if authHeader != "" {
				log.Println("⚠️ [WS-Middleware] Using Authorization from query params")
			}
		}
		if authHeader == "" {
			authHeader = c.Query("token")
			if authHeader != "" {
				log.Println("⚠️ [WS-Middleware] Using token from query params")
			}
		}
		if authHeader == "" {
			log.Println("❌ [WS-Middleware] Missing Authorization header and query param")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Println("🔐 [WS-Middleware] Raw Authorization:", authHeader)
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		tokenStr = strings.TrimSpace(tokenStr)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("❌ [WS-Middleware] Unexpected signing method: %v", t.Header["alg"])
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil {
			log.Printf("❌ [WS-Middleware] Token parsing error: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("❌ [WS-Middleware] Token is not valid")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("❌ [WS-Middleware] Cannot parse token claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, exists := claims["user_id"]
		if !exists || userID == nil {
			log.Println("❌ [WS-Middleware] user_id missing in claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userIDStr, ok := userID.(string)
		if !ok {
			log.Printf("❌ [WS-Middleware] user_id is not a string: %v (type: %T)", userID, userID)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Printf("✅ [WS-Middleware] Auth success. user_id: %s", userIDStr)
		c.Set("user_id", userIDStr)
		c.Next()
	}
}