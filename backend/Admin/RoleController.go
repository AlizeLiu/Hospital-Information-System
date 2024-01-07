package Admin

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRole(ctx *gin.Context) {
	DB := common.GetDB()

	roleName := ctx.Query("roleName")
	role := ctx.Query("role")
	state := ctx.Query("state")

	newAdmin := model.Admin{
		RName:  roleName,
		Role:   role,
		Rstate: state,
	}
	DB.Create(&newAdmin)
	ctx.JSON(http.StatusOK, gin.H{"success": "添加成功"})
}
