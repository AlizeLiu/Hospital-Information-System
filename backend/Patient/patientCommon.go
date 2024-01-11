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

func AddOrder(ctx *gin.Context) {
	DB := common.GetDB()

	pId := ctx.Query("pId")
	oStart := ctx.Query("oStart")
	dId := ctx.Query("dId")

	order := &model.Registration{
		Account:         pId,
		DID:             dId,
		OTime:           oStart,
		MedicineRecord:  "",
		DiagnosisRecord: "",
		OCheckBuyData:   "1",
		ODrugBuyData:    "1",
	}

	if err := DB.Create(order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "挂号成功"})
	model.OrderNum++
}

func BadPeople(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": int(0),
	})
}
