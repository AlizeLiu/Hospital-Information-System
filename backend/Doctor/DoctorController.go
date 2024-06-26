package Doctor

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DoctorInit(ctx *gin.Context) {
	DB := common.GetDB()

	dId := ctx.Query("dId")
	dGender := ctx.Query("dGender")
	dPassword := ctx.Query("dPassword")
	dName := ctx.Query("dName")
	dPost := ctx.Query("dPost")
	dSection := ctx.Query("dSection")
	dPhoneStr := ctx.Query("dPhone")
	dEmail := ctx.Query("dEmail")
	dCardStr := ctx.Query("dCard")
	dPriceStr := ctx.Query("dPrice")
	dIntroduction := ctx.Query("dIntroduction")

	dCard, err := strconv.Atoi(dCardStr)
	dPrice, err := strconv.Atoi(dPriceStr)

	if len(dPhoneStr) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(dPassword) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if IsTelephoneExist(DB, dPhoneStr) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "手机号已经被注册"})
		return
	}

	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(dPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusInternalServerError, "msg": "加密错误"})
		return
	}
	dPhone, err := strconv.Atoi(dPhoneStr)

	newDoctor := model.Doctor{
		DID:           dId,
		DGender:       dGender,
		DPassword:     string(hasedPassword),
		DName:         dName,
		DPost:         dPost,
		DSection:      dSection,
		DPhone:        dPhone,
		DEmail:        dEmail,
		DCard:         dCard,
		DPrice:        dPrice,
		DIntroduction: dIntroduction,
	}
	DB.Create(&newDoctor)

	ctx.JSON(http.StatusOK, gin.H{"success": "增加医生：" + dName + "成功"})

}

func FindAllDoctor(ctx *gin.Context) {
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
	var doctors []model.Doctor
	DB.Model(&model.Doctor{}).Where("d_name LIKE ?", query).
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&doctors)

	var transformedDoctors []gin.H
	for _, doctor := range doctors {
		transformedDoctors = append(transformedDoctors, gin.H{
			"dEmail":        doctor.DEmail,
			"dPrice":        doctor.DPrice,
			"dName":         doctor.DName,
			"dState":        doctor.DState,
			"dPost":         doctor.DPost,
			"dCard":         doctor.DCard,
			"dId":           doctor.DID,
			"dAvgStar":      int(5), // 注意：需要为 Doctor 结构体添加 DAvgStar 字段
			"dSection":      doctor.DSection,
			"dGender":       doctor.DGender,
			"dPhone":        doctor.DPhone,
			"dIntroduction": doctor.DIntroduction,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doctors": transformedDoctors,
		"total":   total,
	})

}

func DeleteDoctor(ctx *gin.Context) {
	DB := common.GetDB()

	dId := ctx.Query("dId")
	if dId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "医生ID不能为空"})
		return
	}
	if err := DB.Where("d_id = ?", dId).Delete(&model.Doctor{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "删除成功",
	})
}

func GetFindDoctor(ctx *gin.Context) { //根据ID查询医生信息
	DB := common.GetDB()
	dId := ctx.Query("dId")
	if dId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "医生ID不能为空"})
		return
	}
	var doctor model.Doctor
	result := DB.Where("d_id=?", dId).First(&doctor)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "未查询到该医生记录"})
		return
	}
	data := struct {
		DEmail        string `json:"dEmail"`
		DPrice        int    `json:"dPrice"`
		DName         string `json:"dName"`
		DState        int    `json:"dState"`
		DPost         string `json:"dPost"`
		DCard         int    `json:"dCard"`
		DId           string `json:"dId"`
		DSection      string `json:"dSection"`
		DGender       string `json:"dGender"`
		DPhone        int    `json:"dPhone"`
		DIntroduction string `json:"dIntroduction"`
	}{
		DEmail:        doctor.DEmail,
		DPrice:        doctor.DPrice,
		DName:         doctor.DName,
		DState:        1,
		DPost:         doctor.DPost,
		DCard:         doctor.DCard,
		DId:           doctor.DID,
		DSection:      doctor.DSection,
		DGender:       doctor.DGender,
		DPhone:        doctor.DPhone,
		DIntroduction: doctor.DIntroduction,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var doctor model.Doctor
	db.Where("dPhone = ?", telephone).First(&doctor)
	if doctor.ID != 0 {
		return true
	}
	return false
}
