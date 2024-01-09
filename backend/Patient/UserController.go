package Patient

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
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

	var user interface{} // 声明一个空接口以存储用户信息

	switch role {
	case "admin":
		var admin model.Admin
		if err := DB.Where("r_name = ?", account).First(&admin).Error; err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
		user = admin

	case "patient":
		var patient model.User
		if err := DB.Where("account = ?", account).First(&patient).Error; err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
		user = patient

	case "doctor":
		var doctor model.Doctor
		if err := DB.Where("d_id = ?", account).First(&doctor).Error; err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
		user = doctor

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "无效的角色"})
		return
	}

	// 根据角色进行密码验证
	switch role {
	case "admin":
		admin, ok := user.(model.Admin)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户类型错误"})
			return
		}
		if admin.RPassword != password {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
			return
		}

	case "patient":
		patient, ok := user.(model.User)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户类型错误"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(password)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
			return
		}

	case "doctor":
		doctor, ok := user.(model.Doctor)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户类型错误"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(doctor.DPassword), []byte(password)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
			return
		}

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "无效的角色"})
		return
	}

	// 生成令牌（示例）
	token := "patient123"

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
	})
}

func FindAllPatient(ctx *gin.Context) {
	DB := common.GetDB()
	pageNumberStr := ctx.Query("pageNumber")
	sizeStr := ctx.Query("size")
	query := ctx.Query("query")

	pageNumber, _ := strconv.Atoi(pageNumberStr)
	size, _ := strconv.Atoi(sizeStr)

	offset := (pageNumber - 1) * size

	if query == "" {
		query = "%"
	} else {
		query = "%" + query + "%"
	}

	var total int64
	var patients []model.User
	DB.Model(&model.User{}).Where("username LIKE ?", query).
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&patients)

	var transformedPatients []gin.H
	for _, patient := range patients {
		age, _ := CalculateAge(patient.Date_of_birth)
		transformedPatients = append(transformedPatients, gin.H{
			"pId":     patient.Account,
			"pName":   patient.Username,
			"pGender": patient.Gender,
			"pAge":    age,
			"pCard":   patient.Id_card_number,
			"pPhone":  patient.Phone_number,
			"pEmail":  patient.Email,
			"pState":  int(1),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"patients": transformedPatients,
		"total":    total,
	})

}

func DeletePatient(ctx *gin.Context) {
	DB := common.GetDB()
	account := ctx.Query("pId")
	if account == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID不能为空"})
		return
	}
	if err := DB.Where("account = ?", account).Delete(&model.User{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "删除失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "删除成功",
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

func CalculateAge(birthDateStr string) (int, error) {
	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, birthDateStr)
	if err != nil {
		return 0, err
	}

	currentDate := time.Now()
	years := currentDate.Year() - birthDate.Year()

	if currentDate.Month() < birthDate.Month() || (currentDate.Month() == birthDate.Month() && currentDate.Day() < birthDate.Day()) {
		years--
	}

	return years, nil
}
