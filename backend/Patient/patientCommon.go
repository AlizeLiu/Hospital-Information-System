package Patient

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindPatientById(ctx *gin.Context) {
	DB := common.GetDB()
	account := ctx.Query("pId")

	if account == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "缺失必要参数"})
	}
	var user model.User
	result := DB.Where("account=?", account).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "未查询到该用户记录"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pName":   user.Username,
		"pGender": user.Gender,
		"pPhone":  user.Phone_number,
	})
}
