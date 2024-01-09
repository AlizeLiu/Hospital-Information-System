package Patient

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindOrderByPid(ctx *gin.Context) {
	DB := common.GetDB()

	pId := ctx.Query("pId")

	var registrations []model.Registration

	if err := DB.Where("account = ?", pId).Find(&registrations).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询挂号信息失败"})
		return
	}

	var transformedRegs []gin.H
	for _, registration := range registrations {
		transformedRegs = append(transformedRegs, gin.H{
			"oId":             registration.ID,
			"dId":             registration.DID,
			"pId":             registration.Account,
			"pName":           registration.Patient.Username,
			"oRecord":         registration.ORecord,
			"oTime":           registration.OTime.Format("2006-01-02 15:04:05"), // 格式化时间
			"medicineRecord":  registration.MedicineRecord,
			"diagnosisRecord": registration.DiagnosisRecord,
			"oTotalPrice":     registration.OTotalPrice,
			"oCheckBuyData":   registration.OCheckBuyData,
			"oDrugBuyData":    registration.ODrugBuyData,
			"oState":          int(1),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"registrations": transformedRegs,
	})
}

func UpdatePrice(ctx *gin.Context) {
	DB := common.GetDB()

	oId := ctx.Query("oId")

	// 查询对应的挂号记录
	var registration model.Registration
	if err := DB.First(&registration, oId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询挂号记录失败"})
		return
	}

	// 更新OPriceState字段为1
	registration.OPriceState = 1
	if err := DB.Save(&registration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新价格状态失败"})
		return
	}

	// 查询医生信息
	var doctor model.Doctor
	if err := DB.Where("DID = ?", registration.DID).First(&doctor).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询医生信息失败"})
		return
	}

	sumPrice := registration.OTotalPrice

	ctx.JSON(http.StatusOK, gin.H{
		"dName":    doctor.DName,
		"sumPrice": sumPrice,
	})
}

func FindOrderInfo(ctx *gin.Context) {
	DB := common.GetDB()

	oId := ctx.Query("oId")
	var registration model.Registration
	if err := DB.Where("id = ?", oId).First(&registration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询挂号信息失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"dName":         registration.Doctor.DName,
		"oCheckBuyData": registration.OCheckBuyData,
		"oDrugBuyData":  registration.ODrugBuyData,
		"oRecord":       registration.ORecord,
	})
}
