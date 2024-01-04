package Doctor

import (
	"backend/common"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DoctorInfo(ctx *gin.Context) {
	DB := common.GetDB()

	// 从请求体中获取医生ID
	dId := ctx.PostForm("dId")

	// 使用 DID 查询医生信息
	var doctor model.Doctor
	if err := DB.First(&doctor, "DID = ?", dId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	doctor.DGender = ctx.PostForm("dGender")
	doctor.DName = ctx.PostForm("dName")
	doctor.DPost = ctx.PostForm("dPost")
	doctor.DSection = ctx.PostForm("dSection")
	doctor.DPhone, _ = strconv.Atoi(ctx.PostForm("dPhone"))
	doctor.DEmail = ctx.PostForm("dEmail")
	dCard, _ := strconv.Atoi(ctx.PostForm("dCard"))
	dPrice, _ := strconv.Atoi(ctx.PostForm("dPrice"))
	doctor.DCard = dCard
	doctor.DPrice = dPrice
	doctor.DIntroduction = ctx.PostForm("dIntroduction")
	dState, _ := strconv.Atoi(ctx.PostForm("dState"))
	doctor.DState = dState

	if err := DB.Save(&doctor).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update doctor"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor information updated successfully", "doctor": doctor})
}
