package Public

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginInf(ctx *gin.Context) {
	DB := common.GetDB()

	account := ctx.Query("userId")
	role := ctx.Query("role")

	if account == "" || role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "缺失必要参数"})
		return
	}

	var name string

	switch role {
	case "admin":
		var admin model.Admin
		if err := DB.Where("r_name = ?", account).First(&admin).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "用户不存在"})
			return
		}
		name = admin.RName

	case "doctor":
		var doctor model.Doctor
		if err := DB.Where("d_id = ?", account).First(&doctor).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "用户不存在"})
			return
		}
		name = doctor.DName

	case "patient":
		var patient model.User
		if err := DB.Where("account = ?", account).First(&patient).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "用户不存在"})
			return
		}
		name = patient.Username

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无效的角色"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username": name,
		"role":     role,
	})
}

func GetMenu(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "缺失必要参数"})
		return
	}
	title := "首页大屏"

	newMeta := model.Meta{
		Title:       title,
		RequireAuth: true,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"component": "/admin/AdminLayout.vue",
		"path":      "/adminLayout",
		"meta":      newMeta,
	})
}

func OrderPeople(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": model.OrderNum,
	})
}
