package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/rchargel/sabida/dao"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	Conn      *dao.Queries
	secretKey []byte
}

func NewLoginHandler(conn *dao.Queries) *LoginHandler {
	return &LoginHandler{
		conn,
		[]byte(os.Getenv("SECRET_KEY")),
	}
}

func (login *LoginHandler) ProcessLoginForm(c *gin.Context) {
	ctx := context.Background()
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := login.Conn.GetUsersByEmail(ctx, email)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if login.validatePassword(password, user.Password) {
		token, tkerr := login.generateJWT(user, ctx)

		if tkerr != nil {
			c.AbortWithError(http.StatusInternalServerError, tkerr)
		} else {
			result := make(map[string]string)
			result["token"] = token
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func (login *LoginHandler) generateJWT(user dao.User, ctx context.Context) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(30 * time.Minute)
	claims["authorized"] = true
	claims["uid"] = user.ID.String()
	claims["user"] = user.Username
	claims["email"] = user.Email
	orgs, err := login.Conn.GetOrganizationsByUser(ctx, user.ID)
	if err == nil {
		orgList := make([]string, len(orgs))
		for i, org := range orgs {
			orgList[i] = org.Name
		}
		claims["orgs"] = orgList
	}
	tokenString, tserr := token.SignedString(login.secretKey)
	if tserr != nil {
		return "", nil
	}
	return tokenString, nil
}

func (login *LoginHandler) validatePassword(supplied string, actual sql.NullString) bool {
	if actual.Valid {
		hashed := actual.String
		return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(supplied)) == nil
	} else {
		return false
	}
}
