package Admin

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRole(ctx *gin.Context) {
	DB := common.GetDB()
	rId := ctx.Query("mId")

	if rId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID不能为空"})
		return
	}
	if err := DB.Where("id = ?", rId).Delete(&model.Admin{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "删除成功",
	})
}

func AddRole(ctx *gin.Context) {
	DB := common.GetDB()

	roleName := ctx.Query("roleName")
	role := ctx.Query("role")
	statestr := ctx.Query("state")
	state, _ := strconv.Atoi(statestr)

	newAdmin := model.Admin{
		RName:  roleName,
		Role:   role,
		Rstate: state,
	}
	DB.Create(&newAdmin)
	ctx.JSON(http.StatusOK, gin.H{"success": "添加成功"})
}

func LoadRole(ctx *gin.Context) {
	DB := common.GetDB()

	var admins []model.Admin
	// 直接从admin表中获取所有记录
	if err := DB.Find(&admins).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load roles",
		})
		return
	}

	// 转换并映射数据
	var adminResponses []model.AdminResponse
	for _, admin := range admins {
		adminResponses = append(adminResponses, model.AdminResponse{
			ID:        admin.ID,
			Rolename:  "管理员",                                         // 假设数据库中字段名为RName
			Role:      admin.Role,                                    // 假设数据库中字段名为Role
			CreatedAt: admin.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间
			State:     admin.Rstate,                                  // 假设数据库中字段名为RState
		})
	}

	// 返回查询结果
	ctx.JSON(http.StatusOK, gin.H{
		"": adminResponses,
	})
}

func ConvertStringSliceToDrugSlice(data []model.FrontendDrug) []model.Drug {
	var drugs []model.Drug
	for _, item := range data {
		drug := model.Drug{
			DrId:     item.DrId,
			DrName:   item.DrName,
			DrNumber: item.DrNumber,
			DrPrice:  item.DrPrice,
		}
		drugs = append(drugs, drug)
	}
	return drugs
}

func ConvertStringSliceToCheckSlice(data []model.FrontendCheck) []model.Check {
	var checks []model.Check
	for _, item := range data {
		check := model.Check{
			ChId:    item.ChId,
			ChName:  item.ChName,
			ChPrice: item.ChPrice,
		}
		checks = append(checks, check)
	}
	return checks
}
