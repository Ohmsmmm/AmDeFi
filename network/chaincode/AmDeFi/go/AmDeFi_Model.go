package main

import (
	"time"
)

/*
	=========================================================
	================== Structure Dictionary =================
	=========================================================
*/
const marketKey = "market|0x8yNXrOSmrMrnrZR3z9Y3DPaNbJZjRwrFN46osUVCsWB3xKeiEOP2My74vhaxNSNu"

type Market struct {
	LoanId []string `json:"loan_id"`
}

type DigitalWallet struct {
	Address       	string              `json:"address"`
	Balance       	float64             `json:"balance"`
	BorrowerAsset 	[]string            `json:"borrower_asset"`
	BorrowerToken 	[]string            `json:"borrower_token"`
	LenderLoan    	[]string            `json:"lender_loan"` //loan id
	LenderPromotion []PromotionList		  	`json:"lender_promotion"`
	PromotionOrder	[]string			`json:"promotion_order"`
	Transaction   	[]WalletTransaction `json:"transaction"`
	LoanDocument  	[]string            `json:"loan_document"`
}

type PromotionList struct {
	Address string `json:"address"`
	Asset Asset `json:"asset"`
	RiskRate	int `json:"risk_rate"`
	Status	string `json:"status"`
	IssueDate time.Time `json:"issue_date"`
}

type PromotionOrder struct {
	Address string `json:"address"`
	TransactionName string `json:"transaction_name"`
	RiskRate	int `json:"risk_rate"`
	Status	string `json:"status"`
	Interest float64 `json:"interest"`
	IssueDate time.Time `json:"issue_date"`
}

type WalletTransaction struct {
	Address         string    `json:"address"`
	IssueDate       time.Time `json:"issue_date"`
	TransactionName string    `json:"transaction_name"`
	TransactionType string    `json:"transaction_type"` //buy sell token
	Total           float64   `json:"total"`
	SnapshotBalance float64   `json:"snapshot_balance"`
}
type Asset struct {
	AssetId         string    `json:"asset_id"`
	Address         string    `json:"address"`
	AssetName       string    `json:"asset_name"`
	Value           float64   `json:"value"`
	LoanType        string    `json:"loan_type"` //secure or un secure
	LoanDuration    int       `json:"loan_duration"`
	LoanInterest    float64   `json:"loan_interest"`
	LoanMin         float64   `json:"loan_min"`
	LoanMax         float64   `json:"loan_max"`
	ValuateInterest float64   `json:"valuate_interest"`
	TokenAmount     int       `json:"token_amount"`
	TokenBalance	int		  `json:"token_balance"`
	Status          string    `json:"status"`
	IssueDate       time.Time `json:"issue_date"` //วันที่เอาสินทรัพเข้า
	LoanId          string    `json:"loan_id"`
}

type LoanDocument struct {
	LoanId          string    `json:"loan_id"`
	AssetId         string    `json:"asset_id"`
	Address         string    `json:"address"` //borrower
	Loan            float64   `json:"loan"`
	RemainDebt      float64   `json:"remain_debt"`
	MinDebtPerMonth float64   `json:"min_debt_per_month"`
	IssueDate         time.Time `json:"issue_date"` // วันที่หมดสัญญา
	Interest        float64   `json:"interest"`
	Status          string    `json:"status"`
	Token           []string  `json:"token"`
}

type Token struct {
	TokenId       string    `json:"token_id"`
	AssetId       string    `json:"asset_id"`
	LenderAddress string    `json:"lender_address"`
	Rate          float64   `json:"rate"`
	IsSell        bool      `json:"is_sell"`    // ขาย  ไม่ขาย
	IssueDate     time.Time `json:"issue_date"` //วันที่ออกเหรียญ
}

type BorrowerGetOwnerAssetList struct {
	AssetList []AssetList `json:"asset_list"`
}

type MarketAssetList struct {
	AssetList []AssetList `json:"asset_list"`
}

type AssetList struct {
	AssetInfo  Asset        `json:"asset_info"`
	LoanInfo   LoanDocument `json:"loan_info"`
	AssetToken []Token      `json:"asset_token"`
}
