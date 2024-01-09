package Doctor

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOrderByDid(ctx *gin.Context) {
	DB := common.GetDB()

	dId := ctx.Query("dId")
	pId := ctx.Query("pId")
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
	var registrations []model.Registration
	queryBuilder := DB.Model(&model.Registration{}).Where("d_id = ?", dId)

	// 根据 pId 过滤
	if pId != "" {
		queryBuilder = queryBuilder.Where("account = ?", pId)
	} else {
		queryBuilder = queryBuilder.Where("account LIKE ?", query)
	}

	queryBuilder.Count(&total).
		Offset(offset).
		Limit(size).
		Find(&registrations)

	var transformedReg []gin.H
	for _, registration := range registrations {
		// 获取患者名字
		var user model.User
		DB.Model(&model.User{}).Where("account = ?", registration.Account).First(&user)
		oStart := registration.CreatedAt.Format("2006-01-02 15:04")
		var oEnd string
		if registration.OTotalPrice == "" || registration.OTotalPrice == "0" {
			oEnd = "还未就诊" // 如果OTotalPrice为空或为零，则设置oEnd为空字符串
		} else {
			oEnd = registration.UpdatedAt.Format("2006-01-02 15:04")
		}
		transformedReg = append(transformedReg, gin.H{
			"oId":           registration.ID,
			"dId":           registration.DID,
			"pId":           registration.Account,
			"pName":         user.Username,
			"oStart":        oStart,
			"oEnd":          oEnd,
			"oRecored":      registration.ORecord,
			"oDrugBuyData":  registration.ODrugBuyData,
			"oCheckBuyData": registration.OCheckBuyData,
			"oTotalPrice":   registration.OTotalPrice,
			"oPriceState":   registration.OPriceState,
			"oState":        int(1),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": transformedReg,
		"total":  total,
	})
}

func FindByDid(ctx *gin.Context) {
	DB := common.GetDB()

	oId := ctx.Query("oId")
	pId := ctx.Query("pId")

	var registration model.Registration
	if err := DB.Model(&model.Registration{}).Where("id = ? AND account = ?", oId, pId).First(&registration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch registration",
		})
		return
	}

	// 初始化一个数组用于存储响应
	var responseData []gin.H

	// 检查UpdatedAt字段是否为空
	var status string
	var time string
	if registration.OTotalPrice == "" || registration.OTotalPrice == "0" {
		status = "未就诊"
		time = registration.CreatedAt.Format("2006-01-02 15:04:05") // 默认时间
		responseData = append(responseData, gin.H{
			"status": status,
			"time":   time,
		})
	} else {
		status = "已就诊"
		time = registration.UpdatedAt.Format("2006-01-02 15:04:05") // 格式化时间
		responseData = append(responseData, gin.H{
			"status": status,
			"time":   time,
		})
	}

	ctx.JSON(http.StatusOK, responseData)
}
