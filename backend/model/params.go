package model

//
//oId: oId.value,
//pId: pId.value,
//dId: dId.value,
//oRecord: reason.value,
//oDrug: drugTotalPrice.value,
//oCheck: checkTotalPrice.value,
//oTotalPrice: dataPackage().oTotalPrice,
//oCheckBuyData:checkBuyData,
//oDrugBuyData:drugBuyData.value

type DoctorData struct {
	DID           string `gorm:"primaryKey" json:"DId"`
	DGender       string // 例如：男、女等
	DPassword     string
	DName         string
	DPost         string
	DSection      string // 医生科室
	DPhone        int    // 电话号码
	DEmail        string
	DCard         string // 身份证号
	DPrice        string // 价格
	DIntroduction string // 医生简介，例如：很厉害的医生
	DState        int    `gorm:"default:1"`
	DNum          string `gorm:"default:30"`
	DNonWorkDays  string `gorm:"type:varchar(255);"`
}

type Order struct {
	OId           string        `json:"oId"`
	PId           string        `json:"pId"`
	DId           string        `json:"dId"`
	ORecord       string        `json:"oRecord"`
	ODrug         int           `json:"oDrug"`
	OCheck        int           `json:"oCheck"`
	OTotalPrice   int           `json:"oTotalPrice"`
	OCheckBuyData []CheckData   `json:"oCheckBuyData"`
	ODrugBuyData  []DrugBuyData `json:"oDrugBuyData"`
}

type CheckData struct {
	ChId    string `json:"chId"`
	ChPrice int    `json:"chPrice"`
	ChName  string `json:"chName"`
}

type DrugBuyData struct {
	DrId    string `json:"drId"`
	DrPrice int    `json:"drPrice"`
	DrName  string `json:"drName"`
	DrNum   int    `json:"drNum"`
	DrSum   int    `json:"drSum"`
}
