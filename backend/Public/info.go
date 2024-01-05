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

	var user model.User
	if err := DB.Where("account = ?", account).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "用户不存在"})
		return
	}
	var roleName string
	switch role {
	case "patient":
		roleName = "患者"
	case "doctor":
		roleName = "医生"
	case "admin":
		roleName = "管理员"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"role":     roleName,
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
