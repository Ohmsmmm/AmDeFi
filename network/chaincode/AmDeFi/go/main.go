package main

import (
	"encoding/json"

	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

/*
	=============================================================
	================== Handle Map Function Name =================
	=============================================================
*/
// Invoke function request form API server is SDK
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(function)
	if function == "query" {
		return t.Query(stub, args)
	} else if function == "init" {
		return t.Init(stub)
	} else if function == "GetMarketplace" {
		return t.GetMarketplace(stub)
	} else if function == "BorrowerGetOwnerAssetList" {
		return t.BorrowerGetOwnerAssetList(stub, args)
	} else if function == "IssueBorrow" {
		return t.Borrow(stub, args)
	} else if function == "LenderGetAssetLendingList" {
		return t.LenderGetAssetLendingList(stub, args)
	} else if function == "LenderGetPromotionOrder" {
		return t.LenderGetPromotionOrder(stub, args)
	} else if function == "IssuePromotionOrder" {
		return t.IssuePromotionOrder(stub, args)
	} else if function == "LenderBuyToken" {
		return t.LenderBuyToken(stub, args)
	} else if function == "LenderSellToken" {
		return t.LenderSellToken(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting " +
		"\"query\" " +
		"\"GetMarketplace\" " +
		"\"BorrowerGetOwnerAssetList\" " +
		"\"IssueBorrow\" " +
		"\"LenderGetAssetLendingList\" " +
		"\"LenderGetPromotionOrder\" " +
		"\"IssuePromotionOrder\" " +
		"\"LenderBuyToken\" " +
		"\"LenderSellToken\" ")
}

func main() {
	// // Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

// Init function for start network in step instantiate first block
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	functionName := "[Init]"
	println("=======================" + functionName + "=======================")
	println("=======================" + "Create Market" + "=======================")
	marketKey := marketKey
	market := Market{LoanId: []string{
		"0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJWEaase223DKwkoIK",
		"0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJffafASFasdaASDda",
	}}

	println(">> START parse market Model to ByteArray <<")
	marketAsBytes, err := json.Marshal(market)
	if err != nil {
		println("Marshal parser market as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser marketModel as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse marketModel Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START marketAsBytes PutState to state blockchain <<")
	err = stub.PutState(marketKey, marketAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END marketAsBytes PutState to state blockchain <<")

	println("=======================" + "Create Wallet" + "=======================")

	//create Digital Wallet > 2
	walletAKey := "0x3921h3921hg7wRipjUIG39sdkOgfkbJe3KWDindoHIsdMiiNSKw"
	walletA := DigitalWallet{
		Address: walletAKey,
		Balance: 1000000.99, //DAI
		BorrowerAsset: []string{
			"0xGDhsud34nsaJHDw59Jh2HDuhJDEJdsfpkKi09YhosII114Sccks",
			"0xOGhsud34nsaJHDw59Jh2AVtuJDEJdsfrkKi09YhosII114Sookr",
		},
		BorrowerToken: nil,
		LenderLoan: []string{
			"0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJWEaase223DKwkoIK",
			"0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJffafASFasdaASDda",
		},
		LenderPromotion: []PromotionList{
			{
				Asset: Asset{
					AssetName: "บ้ายายมี",
					Value:     1000000.00,
				},
				RiskRate:  5,
				Status:    "Accept",
				IssueDate: time.Now(),
			},
			{
				Asset: Asset{
					AssetName: "บ้าขายหอย",
					Value:     2000000.00,
				},
				RiskRate:  1,
				Status:    "Accept",
				IssueDate: time.Now(),
			},
		},
		PromotionOrder: []string{
			"0x3921h3921hg7wRipjUIG39sdkOgfkbJe3KWDindoHIsdMiiNSAA",
			"0x3921h3921hg7wRipjUIG39sdkOgfkbJe3KWDindoHIsdMiiNSBB",
		},
		Transaction: []WalletTransaction{
			{
				Address:         "0xwW9OADwy9QmhVsBmff7n6L8vpWs1MricqCoMME8KMuVtxRaiOJZ54ZrnjR6BEVCI",
				IssueDate:       time.Now().AddDate(0, 0, -1),
				TransactionName: "Buy Asset Token",
				TransactionType: "Buy",
				Total:           100.00,
				SnapshotBalance: 999900.99,
			},
			{
				IssueDate:       time.Now(),
				TransactionName: "loan by Asset Token",
				TransactionType: "Sell",
				Total:           100.00,
				SnapshotBalance: 1000000.99,
			},
		},
		LoanDocument: nil,
	}

	walletBKey := "0xSakdsjakIhdsklmLmJDpjIOhodsol93GHudg2kDNhKHodsb9dSk"
	walletB := DigitalWallet{
		Address:       walletBKey,
		Balance:       1000000.99, //DAI
		BorrowerAsset: []string{"0xGDhsud34nsaJHDw59Jh2HDuhJDEJdsfpkKi09YhosII112Sucks"},
		BorrowerToken: nil,
		LenderLoan:    nil,
		Transaction:   nil,
		LoanDocument:  nil,
	}
	println(">> START parse Promo Order Model to ByteArray <<")
	PromoOrderKey1 := "0x3921h3921hg7wRipjUIG39sdkOgfkbJe3KWDindoHIsdMiiNSAA"
	PromoOrder1 := PromotionOrder{
		Address:         PromoOrderKey1,
		TransactionName: "Order1",
		RiskRate:        3,
		Status:          "Pending",
		IssueDate:       time.Now(),
	}
	PromoOrderKey2 := "0x3921h3921hg7wRipjUIG39sdkOgfkbJe3KWDindoHIsdMiiNSBB"
	PromoOrder2 := PromotionOrder{
		Address:         PromoOrderKey2,
		TransactionName: "Order2",
		RiskRate:        8,
		Status:          "Pending",
		IssueDate:       time.Now(),
	}
	PromoOrder1AsBytes, err := json.Marshal(PromoOrder1)
	if err != nil {
		println("Marshal parser wallet as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser walletAsBytesModel as Model to ByteArray is error" + err.Error())
	}
	PromoOrder2AsBytes, err := json.Marshal(PromoOrder2)
	if err != nil {
		println("Marshal parser wallet as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser walletAsBytesModel as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse Promo Order Model to ByteArray <<")
	//byteArray put to state blockchain
	println(">> START PromoOrderAsBytes PutState to state blockchain <<")
	err = stub.PutState(PromoOrderKey1, PromoOrder1AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(PromoOrderKey2, PromoOrder2AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END PromoOrderAsBytes PutState to state blockchain <<")

	println(">> START parse wallet Model to ByteArray <<")
	walletAAsBytes, err := json.Marshal(walletA)
	if err != nil {
		println("Marshal parser wallet as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser walletAsBytesModel as Model to ByteArray is error" + err.Error())
	}
	walletBAsBytes, err := json.Marshal(walletB)
	if err != nil {
		println("Marshal parser wallet as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser walletAsBytesModel as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse walletModel Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START walletAsBytes PutState to state blockchain <<")
	err = stub.PutState(walletAKey, walletAAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(walletBKey, walletBAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END walletAsBytes PutState to state blockchain <<")

	println("=======================" + "Create Asset" + "=======================")
	//create Asset
	assetKey := "0xGDhsud34nsaJHDw59Jh2HDuhJDEJdsfpkKi09YhosII114Sccks"
	asset := Asset{
		LoanId:          "0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJWEaase223DKwkoIK",
		AssetId:         assetKey,
		Address:         walletAKey,
		AssetName:       "บ้าน",
		Value:           10000, //DAI
		LoanType:        "secure loan",
		LoanDuration:    76,   //M = 3 Year
		LoanInterest:    4,    //Percent
		LoanMin:         1000, //10%
		LoanMax:         8000, //80%
		ValuateInterest: 4,    //Percent
		TokenAmount:     10,
		TokenBalance:    10,
		Status:          "approved",
		IssueDate:       time.Now(),
	}

	assetA2Key := "0xOGhsud34nsaJHDw59Jh2AVtuJDEJdsfrkKi09YhosII114Sookr"
	assetA2 := Asset{
		LoanId:          "0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJffafASFasdaASDda",
		AssetId:         assetA2Key,
		Address:         walletAKey,
		AssetName:       "บ้าน",
		Value:           20000, //DAI
		LoanType:        "secure loan",
		LoanDuration:    76,   //M = 3 Year
		LoanInterest:    3,    //Percent
		LoanMin:         1000, //10%
		LoanMax:         8000, //80%
		ValuateInterest: 4,    //Percent
		TokenAmount:     10,
		TokenBalance:    10,
		Status:          "approved",
		IssueDate:       time.Now(),
	}

	assetKey2 := "0xGDhsud34nsaJHDw59Jh2HDuhJDEJdsfpkKi09YhosII112Sucks"
	assetB := Asset{
		LoanId:          "0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJWEaase223Dkigluay",
		AssetId:         assetKey2,
		Address:         walletBKey,
		AssetName:       "บ้าน",
		Value:           10, //DAI
		LoanType:        "secure loan",
		LoanDuration:    76,   //M = 3 Year
		LoanInterest:    4,    //Percent
		LoanMin:         1000, //10%
		LoanMax:         8000, //80%
		ValuateInterest: 4,    //Percent
		TokenAmount:     10,
		TokenBalance:    10,
		Status:          "approved",
		IssueDate:       time.Now(),
	}
	println(">> START parse asset Model to ByteArray <<")
	assetBAsBytes, err := json.Marshal(assetB)
	if err != nil {
		println("Marshal parser asset as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser asset as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse assetModel Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START assetAsBytes PutState to state blockchain <<")
	err = stub.PutState(assetKey2, assetBAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	println(">> START parse asset Model to ByteArray <<")
	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		println("Marshal parser asset as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser asset as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse assetModel Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START assetAsBytes PutState to state blockchain <<")
	err = stub.PutState(assetKey, assetAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END assetAsBytes PutState to state blockchain <<")

	println(">> START parse asset Model to ByteArray <<")
	assetA2AsBytes, err := json.Marshal(assetA2)
	if err != nil {
		println("Marshal parser asset as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser asset as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse assetModel Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START assetAsBytes PutState to state blockchain <<")
	err = stub.PutState(assetA2Key, assetA2AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	println("=======================" + "Create loanDocument" + "=======================")

	loanDocumentKey := "0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJWEaase223DKwkoIK"
	loanDocument := LoanDocument{
		AssetId:         assetKey,
		Address:         walletAKey,
		LoanId:          loanDocumentKey,
		Loan:            8000,
		RemainDebt:      8000,
		Interest:        asset.LoanInterest,
		MinDebtPerMonth: (8000 / float64(asset.LoanDuration)) * (float64(asset.LoanInterest) / float64(100)),
		IssueDate:       time.Now(),
		Status:          "callateral",
		Token: []string{
			"0x534adsDAdsfkewlvkkjsfweiQQdjefn398fps025hc8NEEIjfoe",
			"0xIJdsdnauGdeworj39NJndop0389djDjfes5iht4ncKJsweovjKa",
			"0x0347idsjfsjjfrmskgnoierl38chfs7jdJJFmekkedsojwWWfkv",
			"0x95LJUObmMKIDSt1hZoChoXz1rUH3QPTaRcNiGVH51foBAWDj90nT1z8824MVVzyT",
			"0xVx1my9LtDpa3VxVasaxzu3j2YApdKeE2GQDcMARY2zsUi5TzLPC3wyNa2hLTpTyf",
			"0xLxU81ZcCKCTOvmFZQL9UUZ4aB7IiJEu01iEtFEuCUoirBGz4AAK8P81mGUMM1TBa",
			"0xRxcbOOUjJmcOoj20ZSuPOT2Sv1JB8aaj9SWEQEi09mQbHC2U5TxDIV80zU9tGRNL",
			"0xmKVd9OnvjRt7LoGoowwRB3tiYZzM3r4Zc61hdnnrJ8H6Ch5r331vjJQsL8xCPZH4",
			"0xtLPCPT5y2qcN7s3MHxyySTn6hB865CL1n0YN8WVTvhcXP2rwZ0QnMTMPZBTB1A2y",
			"0xCBvm648FvE3atAIQ72m1auNxAIE1uqqifZti0vZju0PY6S51YteEuaJtLV6YD2hE"},
	}

	loanDocumentKey2 := "0xfdsp3ofjsi32jr9hcsm2FkewpKwoOW00fsdJffafASFasdaASDda"
	loanDocument2 := LoanDocument{
		AssetId:         assetA2Key,
		Address:         walletAKey,
		LoanId:          loanDocumentKey2,
		Loan:            5000,
		RemainDebt:      1000,
		Interest:        assetA2.LoanInterest,
		MinDebtPerMonth: (5000 / float64(assetA2.LoanDuration)) * (float64(assetA2.LoanInterest) / float64(100)),
		IssueDate:       time.Now(),
		Status:          "callateral",
		Token: []string{
			"0x534adsDAdsfkewlvkkjsfweiQQdjefn398fps025hc8NEEIjf11",
			"0xIJdsdnauGdeworj39NJndop0389djDjfes5iht4ncKJsweovj22",
			},
	}
	println(">> START parse loan Model to ByteArray <<")
	loanAsBytes, err := json.Marshal(loanDocument)
	if err != nil {
		println("Marshal parser loan as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser loan as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse loan Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START loanAsBytes PutState to state blockchain <<")
	err = stub.PutState(loanDocumentKey, loanAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END loanAsBytes PutState to state blockchain <<")
	println(">> START parse loan Model to ByteArray <<")
	loan2AsBytes, err := json.Marshal(loanDocument2)
	if err != nil {
		println("Marshal parser loan as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser loan as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse loan Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START loanAsBytes PutState to state blockchain <<")
	err = stub.PutState(loanDocumentKey2, loan2AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END loanAsBytes PutState to state blockchain <<")

	println("=======================" + "Create token" + "=======================")

	token1Key := "0x534adsDAdsfkewlvkkjsfweiQQdjefn398fps025hc8NEEIjfoe"
	token1 := Token{
		TokenId:       token1Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        false,
		IssueDate:     time.Now(),
	}
	token2Key := "0xIJdsdnauGdeworj39NJndop0389djDjfes5iht4ncKJsweovjKa"
	token2 := Token{
		TokenId:       token2Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        false,
		IssueDate:     time.Now(),
	}
	token3Key := "0x0347idsjfsjjfrmskgnoierl38chfs7jdJJFmekkedsojwWWfkv"
	token3 := Token{
		TokenId:       token3Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        false,
		IssueDate:     time.Now(),
	}
	token4Key := "0x95LJUObmMKIDSt1hZoChoXz1rUH3QPTaRcNiGVH51foBAWDj90nT1z8824MVVzyT"
	token4 := Token{
		TokenId:       token4Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        false,
		IssueDate:     time.Now(),
	}

	token5Key := "0xVx1my9LtDpa3VxVasaxzu3j2YApdKeE2GQDcMARY2zsUi5TzLPC3wyNa2hLTpTyf"
	token5 := Token{
		TokenId:       token5Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        false,
		IssueDate:     time.Now(),
	}

	token6Key := "0xLxU81ZcCKCTOvmFZQL9UUZ4aB7IiJEu01iEtFEuCUoirBGz4AAK8P81mGUMM1TBa"
	token6 := Token{
		TokenId:       token6Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token7Key := "0xRxcbOOUjJmcOoj20ZSuPOT2Sv1JB8aaj9SWEQEi09mQbHC2U5TxDIV80zU9tGRNL"
	token7 := Token{
		TokenId:       token7Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token8Key := "0xmKVd9OnvjRt7LoGoowwRB3tiYZzM3r4Zc61hdnnrJ8H6Ch5r331vjJQsL8xCPZH4"
	token8 := Token{
		TokenId:       token8Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token9Key := "0xtLPCPT5y2qcN7s3MHxyySTn6hB865CL1n0YN8WVTvhcXP2rwZ0QnMTMPZBTB1A2y"
	token9 := Token{
		TokenId:       token9Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token10Key := "0xCBvm648FvE3atAIQ72m1auNxAIE1uqqifZti0vZju0PY6S51YteEuaJtLV6YD2hE"
	token10 := Token{
		TokenId:       token10Key,
		AssetId:       assetKey,
		LenderAddress: "",
		Rate:          loanDocument.Loan / float64(asset.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token11Key := "0x534adsDAdsfkewlvkkjsfweiQQdjefn398fps025hc8NEEIjf11"
	token11 := Token{
		TokenId:       token11Key,
		AssetId:       assetA2Key,
		LenderAddress: "",
		Rate:          loanDocument2.Loan / float64(assetA2.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	token12Key := "0xIJdsdnauGdeworj39NJndop0389djDjfes5iht4ncKJsweovj22"
	token12 := Token{
		TokenId:       token12Key,
		AssetId:       assetA2Key,
		LenderAddress: "",
		Rate:          loanDocument2.Loan / float64(assetA2.TokenAmount),
		IsSell:        true,
		IssueDate:     time.Now(),
	}

	println(">> START parse token Model to ByteArray <<")
	token1AsBytes, err := json.Marshal(token1)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token2AsBytes, err := json.Marshal(token2)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token3AsBytes, err := json.Marshal(token3)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token4AsBytes, err := json.Marshal(token4)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token5AsBytes, err := json.Marshal(token5)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token6AsBytes, err := json.Marshal(token6)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token7AsBytes, err := json.Marshal(token7)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token8AsBytes, err := json.Marshal(token8)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token9AsBytes, err := json.Marshal(token9)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token10AsBytes, err := json.Marshal(token10)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token11AsBytes, err := json.Marshal(token11)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	token12AsBytes, err := json.Marshal(token12)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START tokenAsBytes PutState to state blockchain <<")
	err = stub.PutState(token1Key, token1AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token2Key, token2AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token3Key, token3AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token4Key, token4AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token5Key, token5AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token6Key, token6AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token7Key, token7AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token8Key, token8AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token9Key, token9AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token10Key, token10AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token11Key, token11AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	err = stub.PutState(token12Key, token12AsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END tokenAsBytes PutState to state blockchain <<")

	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}
