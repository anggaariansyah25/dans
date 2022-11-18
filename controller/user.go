package controller

import (
	"dans/model"
	"dans/utils"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type UsersController struct {
	DB *gorm.DB
}
func (controller *UsersController) Login(ctx *gin.Context)  {
	var (
		req model.User
		users model.User
		empty model.Empty
		userView model.UserView
	)
	if err := ctx.BindJSON(&req); err != nil {
		res := utils.Response(http.StatusBadRequest, "can't bind struct, err:"+err.Error(), empty)
		ctx.JSON(http.StatusOK, res)
		return
	}

	if err := controller.DB.Set("gorm:auto_preload", true).Where("username LIKE ?", req.Username).Debug().Find(&users).Error; gorm.IsRecordNotFoundError(err) {
		res := utils.Response(http.StatusBadRequest, "username is incorrect", empty)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// create expired time for auth token
	expiresAt := time.Now().Add(time.Minute * 1000).Unix()
	// check if password not same
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(req.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		res := utils.Response(http.StatusBadRequest, "password is incorrect", empty)
		ctx.JSON(http.StatusOK, res)
		return
	}

	tk := &model.Token{
		ID: req.ID,
		Phone:  req.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	// generate jwt auth
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "can't signed token"})
		return
	}
	userView.Username = req.Username
	userView.AuthToken = tokenString
	res := utils.Response(http.StatusOK, "Success Login", userView)
	ctx.JSON(http.StatusOK, res)
	return
}

func (controller *UsersController) Register(ctx *gin.Context) {
	var (
		req model.User
		RegisView model.UserRegisView
		empty model.Empty
	)
	// get data from request parameter
	if err := ctx.BindJSON(&req); err != nil {
		res := utils.Response(http.StatusBadRequest, "can't get data, err :", empty)
		ctx.JSON(http.StatusOK, res)
		return
	}
	// validation request
	// request validation
	validate, trans := utils.InitValidate()
	if err := validate.Struct(req); err != nil {
		var errStrings []string
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			errStrings = append(errStrings, e.Translate(trans))
		}
		res := utils.Response(http.StatusBadRequest, strings.Join(errStrings, ", "), empty)
		ctx.JSON(http.StatusOK, res)
		return
	}
	// validation for password format
	if err := utils.VerifyPassword(req.Password); err != nil {
		res := utils.Response(http.StatusBadRequest, err.Error(), empty)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// check username is exist
	if err := controller.DB.Where("username = ?", req.Username).First(&req).Error; !(gorm.IsRecordNotFoundError(err)) {
		res := utils.Response(http.StatusBadRequest, "Username is already taken ", empty)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// assign hash password to user and create user
	pass, _ := utils.HashAndSalt([]byte(req.Password))
	req.Password = pass
	req.Username = req.Username
	if err := controller.DB.Create(&req).Error; err != nil {
		res := utils.Response(http.StatusInternalServerError, "can't create user, err: "+err.Error(), empty)
		ctx.JSON(http.StatusOK, res)
		return
	}
	RegisView.Username = req.Username
	res := utils.Response(http.StatusOK, "Success Register",RegisView)
	ctx.JSON(http.StatusOK, res)
}