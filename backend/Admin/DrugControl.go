package Admin

import (
	"backend/common"
	"backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateOrder(ctx *gin.Context) {
	DB := common.GetDB()

	oId := ctx.Query("oId")
	pId := ctx.Query("pId")
	dId := ctx.Query("dId")
	oRecord := ctx.Query("oRecord")
	oDrug := ctx.Query("oDrug")
	oCheck := ctx.Query("oCheck")
	oTotalPrice := ctx.Query("oTotalPrice")

	var oDrugBuyData []model.FrontendDrug
	if err := ctx.BindJSON(&oDrugBuyData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var oCheckBuyData []model.FrontendCheck
	if err := ctx.BindJSON(&oCheckBuyData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// 使用转换函数将前端数据转换为数据模型
	drugs := ConvertStringSliceToDrugSlice(oDrugBuyData)
	checks := ConvertStringSliceToCheckSlice(oCheckBuyData)

	// 现在，你可以将这些数据保存到数据库或进行其他操作

	var existingOrder model.Registration
	result := DB.Where("id = ?", oId).First(&existingOrder)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	// 更新挂号记录
	existingOrder.Account = pId
	existingOrder.DID = dId
	existingOrder.ORecord = oRecord
	existingOrder.ODrugs = oDrug
	existingOrder.OChecks = oCheck
	existingOrder.OTotalPrice = oTotalPrice
	existingOrder.OCheckBuyData = checks
	existingOrder.ODrugBuyData = drugs

	DB.Save(&existingOrder)

	// 返回成功响应
	ctx.JSON(200, gin.H{
		"message": "Order successfully updated",
		"oId":     existingOrder.ID,
	})
}

func ReduceDrugs(ctx *gin.Context) {
	DB := common.GetDB()

	drId := ctx.Query("drId")
	userNumberStr := ctx.Query("usedNumber") //药物数量

	userNumber, err := strconv.Atoi(userNumberStr)
	if err != nil {
		// 如果转换失败，返回错误响应
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userNumber",
		})
		return
	}
	var drug model.Drug
	if err := DB.Where("dr_id = ?", drId).First(&drug).Error; err != nil {
		// 如果找不到匹配的药物，返回错误响应
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Drug not found",
		})
		return
	}

	if drug.DrNumber < userNumber {
		// 如果库存不足，返回错误响应
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient stock",
		})
		return
	}
	// 扣减库存
	drug.DrNumber -= userNumber
	if err := DB.Save(&drug).Error; err != nil {
		// 如果更新数据库出错，返回错误响应
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update stock",
		})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"success": "success reduce",
	})
}

func AddChecks(ctx *gin.Context) {
	DB := common.GetDB()
	chId := ctx.Query("chId")

	var checks model.Check

	// 在数据库中查找与给定 drId 匹配的药物
	if err := DB.Where("ch_id = ?", chId).First(&checks).Error; err != nil {
		// 如果查询出错或没有找到匹配的药物，则返回错误响应
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Drug not found",
		})
		return
	}

	// 构造响应数据
	response := gin.H{
		"chId":    checks.ChId,
		"chName":  checks.ChName,
		"chPrice": checks.ChPrice,
	}

	// 发送响应到前端
	ctx.JSON(http.StatusOK, response)
}
func AddDrugs(ctx *gin.Context) {
	DB := common.GetDB()
	drId := ctx.Query("drId")

	var drug model.Drug

	// 在数据库中查找与给定 drId 匹配的药物
	if err := DB.Where("dr_id = ?", drId).First(&drug).Error; err != nil {
		// 如果查询出错或没有找到匹配的药物，则返回错误响应
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Drug not found",
		})
		return
	}

	// 构造响应数据
	response := gin.H{
		"drId":     drug.DrId,
		"drName":   drug.DrName,
		"drNumber": drug.DrNumber,
		"drUnit":   drug.DrUnit,
		"drPrice":  drug.DrPrice,
	}

	// 发送响应到前端
	ctx.JSON(http.StatusOK, response)
}

func FindAllDrug(ctx *gin.Context) {
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
	var drugs []model.Drug
	DB.Model(&model.Drug{}).Where("dr_name LIKE ?", query).
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&drugs)

	var transformedDrugs []gin.H
	for _, drug := range drugs {
		transformedDrugs = append(transformedDrugs, gin.H{
			"drId":     drug.DrId,
			"drName":   drug.DrName,
			"drNumber": drug.DrNumber,
			"drUnit":   drug.DrUnit,
			"drPrice":  drug.DrPrice,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"drugs": transformedDrugs,
		"total": total,
	})
}

func FindAllCheck(ctx *gin.Context) {
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
	var checks []model.Check
	DB.Model(&model.Check{}).Where("ch_name LIKE ?", query).
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&checks)
	if DB.Error != nil {
		fmt.Println("Database error:", DB.Error)
	}
	var transformedChecks []gin.H
	for _, check := range checks {
		transformedChecks = append(transformedChecks, gin.H{
			"chId":    check.ChId,
			"chName":  check.ChName,
			"chPrice": check.ChPrice,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"checks": transformedChecks,
		"total":  total,
	})
}
