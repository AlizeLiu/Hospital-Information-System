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
