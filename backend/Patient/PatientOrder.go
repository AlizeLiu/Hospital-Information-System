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
		// 查询对应的patient和doctor信息
		var patient model.User
		if err := DB.Where("account = ?", registration.Account).First(&patient).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询病人信息失败"})
			return
		}

		var doctor model.Doctor
		if err := DB.Where("d_id = ?", registration.DID).First(&doctor).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询医生信息失败"})
			return
		}

		oStart := registration.CreatedAt.Format("2006-01-02 15:04")
		var oEnd string
		if registration.OTotalPrice == "" || registration.OTotalPrice == "0" {
			oEnd = "还未就诊" // 如果OTotalPrice为空或为零，则设置oEnd为空字符串
		} else {
			oEnd = registration.UpdatedAt.Format("2006-01-02 15:04")
		}

		transformedRegs = append(transformedRegs, gin.H{
			"oId":         registration.ID,
			"pName":       patient.Username,
			"dName":       doctor.DName,
			"oEnd":        oEnd,
			"oStart":      oStart,
			"oTotalPrice": registration.OTotalPrice,
			"oPrice":      registration.OTotalPrice,
			"oStatePrice": doctor.DPrice,
			"oState":      string(1),
		})
	}

	ctx.JSON(http.StatusOK, transformedRegs)
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
	var doctor model.Doctor
	if err := DB.Where("d_id = ?", registration.DID).First(&doctor).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询医生信息失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"dName":         doctor.DName,
		"oCheckBuyData": registration.OCheckBuyData,
		"oDrugBuyData":  registration.ODrugBuyData,
		"oRecord":       registration.ORecord,
	})
}
