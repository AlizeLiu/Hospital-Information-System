package Patient

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func PatientRegister(ctx *gin.Context) {
	DB := common.GetDB()

	// 从 URL 查询参数中获取数据
	username := ctx.Query("pName")
	account := ctx.Query("pId")
	password := ctx.Query("pPassword")
	gender := ctx.Query("pGender")
	email := ctx.Query("pEmail")
	phone_number := ctx.Query("pPhone")
	date_of_birth := ctx.Query("pBirthday")
	id_card := ctx.Query("pCard")

	role := "patient"

	if len(phone_number) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if IsTelephoneExist(DB, phone_number) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经被注册"})
		return
	}

	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
		return
	}

	// 创建新用户并保存到数据库
	newUser := model.User{
		Username:       username,
		Account:        account,
		Password:       string(hasedPassword),
		Role:           role,
		Gender:         gender,
		Email:          email,
		Phone_number:   phone_number,
		Id_card_number: id_card,
		Date_of_birth:  date_of_birth,
	}

	DB.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{"desc": "注册成功"})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 从 URL 查询参数中获取数据
	account := ctx.Query("username")
	password := ctx.Query("password")
	role := ctx.Query("role")

	// 检查参数是否存在并转换为字符串类型
	if account == "" || password == "" || role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "必要参数缺失"})
		return
	}

	// 查询用户
	var user model.User
	DB.Where("account= ?", account).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 检查密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 检查角色
	if user.Role != role {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户身份核验错误"})
		return
	}

	// 生成令牌（示例）
	token := "patient123"

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
	})
}
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("phone_number = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
