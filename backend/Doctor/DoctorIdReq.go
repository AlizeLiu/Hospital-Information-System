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
		transformedReg = append(transformedReg, gin.H{
			"oId":           registration.ID,
			"dId":           registration.DID,
			"pId":           registration.Account,
			"pName":         user.Username,
			"oStart":        registration.OTime,
			"oEnd":          registration.UpdatedAt,
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

	// 检查UpdatedAt字段是否为空
	var status string
	var time string
	if registration.UpdatedAt.IsZero() {
		status = "未就诊"
		time = "1970-01-01 00:00:00" // 默认时间
	} else {
		status = "已就诊"
		time = registration.UpdatedAt.Format("2006-01-02 15:04:05") // 格式化时间
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"time":   time,
	})
}
