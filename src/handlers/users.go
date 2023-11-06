package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"

	"github.com/rchargel/sabida/dao"
	"github.com/rchargel/sabida/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Conn      *dao.Queries
	secretKey []byte
}

func NewUserHandler(conn *dao.Queries) *UserHandler {
	return &UserHandler{
		conn,
		[]byte(os.Getenv("SECRET_KEY")),
	}
}

func (users *UserHandler) hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return string(hashed)
	} else {
		return ""
	}
}

func (users *UserHandler) validateUsernameExists(username string, ctx context.Context) bool {
	_, err := users.Conn.GetUserByUsername(ctx, username)
	return err == nil
}

func (users *UserHandler) validateEmailExists(email string, ctx context.Context) bool {
	_, err := users.Conn.GetUsersByEmail(ctx, email)
	return err == nil
}

func (users *UserHandler) DeleteUser(c *gin.Context) {
	ctx := context.Background()
	id := c.PostForm("uid")
	userId := uuid.Must(uuid.Parse(id))

	if err := users.Conn.DeleteUser(ctx, userId); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, models.Message{"User deleted"})
	}
}

func (users *UserHandler) UpdatePassword(c *gin.Context) {
	ctx := context.Background()
	id := c.PostForm("uid")
	userId := uuid.Must(uuid.Parse(id))

	hashedPassword := users.hashPassword(c.PostForm("password"))
	if err := users.Conn.UpdatePassword(ctx, dao.UpdatePasswordParams{userId, ToNullString(hashedPassword)}); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, models.Message{"Password updated"})
	}
}

func (users *UserHandler) CreateUser(c *gin.Context) {
	ctx := context.Background()

	createUserParams := dao.CreateUserParams{
		c.PostForm("username"),
		c.PostForm("email"),
	}

	if users.validateUsernameExists(createUserParams.Lower, ctx) || users.validateEmailExists(createUserParams.Lower_2, ctx) {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			"A user with that username or email address already exists",
			http.StatusBadRequest,
		})
	} else {
		if user, err := users.Conn.CreateUser(ctx, createUserParams); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			hashedPassword := users.hashPassword(c.PostForm("password"))

			users.Conn.UpdatePassword(ctx, dao.UpdatePasswordParams{
				user.ID,
				ToNullString(hashedPassword),
			})

			c.JSON(http.StatusOK, models.Message{"User created"})
		}
	}
}
