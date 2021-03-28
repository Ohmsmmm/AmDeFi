package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

/*
	=========================================================
	===================== Smart Contract  ===================
	=========================================================
*/

type SmartContract struct {
}

func (t *SmartContract) Query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[Query]"
	print("==========================================" + functionName + "====================================================")
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
	documentKey := args[0]
	// Get the state from the ledger
	documentAsBytes, err := stub.GetState(documentKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + documentKey + "\"}"
		return shim.Error(jsonResp)
	}
	if documentAsBytes == nil {
		jsonResp := documentKey
		return shim.Error(jsonResp)
	}
	jsonResp := string(documentAsBytes)
	fmt.Printf("%s Response:%s\n", functionName, jsonResp)
	return shim.Success(documentAsBytes)
}
func (t *SmartContract) LenderSellToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[LenderSellToken]"
	println("==========================================" + functionName + "====================================================")
	println("Parse args as string to array")
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")
	println("Parse args as string to array successfully")
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
	var err error
	walletKey := args[0]
	loanId := args[1]
	tokenAmount, err:= strconv.ParseInt(args[2], 10, 64)
	// Get the state from the ledger

	walletAsBytes, err := stub.GetState(walletKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}
	if walletAsBytes == nil {
		jsonResp := walletKey
		return shim.Error(jsonResp)
	}
	walletModel := DigitalWallet{}
	errWalletUnmarshal := json.Unmarshal(walletAsBytes, &walletModel)
	if errWalletUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}
	marketKey := marketKey
	//Get Market
	documentMarket, err := stub.GetState(marketKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + marketKey + "\"}"
		return shim.Error(jsonResp)
	}
	if documentMarket == nil {
		jsonResp := marketKey
		return shim.Error(jsonResp)
	}
	var market Market
	UnmarlshalMarket := json.Unmarshal(documentMarket, &market)
	if UnmarlshalMarket != nil {
		println(" Error " + functionName + " unmarshaling MarketModel : " + UnmarlshalMarket.Error())
		return shim.Error(" Error " + functionName + "  unmarshaling MarketModel : " + UnmarlshalMarket.Error())
	}
	var LenderLoan LoanDocument
	var token Token
	var soldToken []string

	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	if len(walletModel.LenderLoan) != 0 {
		for _, value := range walletModel.LenderLoan {
			println(value)
			if loanId == value {
				println("check loanId")
				LoanAsByte, err := stub.GetState(value)
				if err != nil {
					jsonResp := "{\"Error\":\"Failed to get state for " + value + "\"}"
					return shim.Error(jsonResp)
				}
				if LoanAsByte == nil {
					jsonResp := value
					return shim.Error(jsonResp)
				}
				errLoanUnmarshal := json.Unmarshal(LoanAsByte, &LenderLoan)
				if errLoanUnmarshal != nil {
					jsonResp := "{\"Error\":\"Failed to get unmarshall for " + value + "\"}"
					return shim.Error(jsonResp)
				}
				if (int(tokenAmount) <= len(LenderLoan.Token)) {
					println("check Token")
					for i := 0; i < int(tokenAmount); i++ {
						TokenAsByte, err := stub.GetState(LenderLoan.Token[i])
						if err != nil {
							jsonResp := "{\"Error\":\"Failed to get state for " + value + "\"}"
							return shim.Error(jsonResp)
						}
						if TokenAsByte == nil {
							jsonResp := LenderLoan.Token[i]
							return shim.Error(jsonResp)
						}
						errTokenUnmarshal := json.Unmarshal(TokenAsByte, &token)
						if errTokenUnmarshal != nil {
							jsonResp := "{\"Error\":\"Failed to get unmarshall for " + value + "\"}"
							return shim.Error(jsonResp)
						}
						if (token.IsSell == false) {
							println("check isSell")
							token.IsSell = true
							soldToken = append(soldToken, LenderLoan.Token[i])
						}

						println(">> START parse Token Model to ByteArray <<")
						tokenAsByte, err := json.Marshal(token)
						if err != nil {
							println("Marshal parser token as Model to ByteArray is error" + err.Error())
							return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
						}
						println(">> END parse token Model to ByteArray <<")

						//byteArray put tokenAsByte to state blockchain
						println(">> START TokenAsByte PutState to state blockchain <<")
						err = stub.PutState(LenderLoan.Token[i], tokenAsByte)
						if err != nil {
							println("PutState is error" + err.Error())
							return shim.Error("PutState is error" + err.Error())
						}
						println(">> END TokenAsByte PutState to state blockchain <<")

					}
					market.LoanId = append(market.LoanId, value)
				}
			}
		}
	}
	println("check loanId end")
	marketAsByte, err := json.Marshal(market)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")

	//byteArray put tokenAsByte to state blockchain
	println(">> START marketAsByte PutState to state blockchain <<")
	err = stub.PutState(marketKey, marketAsByte)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END marketAsByte PutState to state blockchain <<")


	WalletAsByte, err := json.Marshal(walletModel)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")

	//byteArray put tokenAsByte to state blockchain
	println(">> START WalletAsByte PutState to state blockchain <<")
	err = stub.PutState(walletKey, WalletAsByte)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END WalletAsByte PutState to state blockchain <<")
	sellTokenAsByte, err := json.Marshal(soldToken)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")
	return shim.Success(sellTokenAsByte)
}
func (t *SmartContract) IssuePromotionOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssuePromotionOrder]"
	print("==========================================" + functionName + "====================================================")
	println("Parse args as string to array")
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")
	println("Parse args as string to array successfully")
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
	var err error
	walletKey := args[0]
	// Get the state from the ledger

	walletAsBytes, err := stub.GetState(walletKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}
	if walletAsBytes == nil {
		jsonResp := walletKey
		return shim.Error(jsonResp)
	}
	walletModel := DigitalWallet{}
	errWalletUnmarshal := json.Unmarshal(walletAsBytes, &walletModel)
	if errWalletUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}
	PromoKey := args[1]
	RiskRate, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	Interest, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		println("ParseFloat is error" + err.Error())
		return shim.Error("ParseFloat is error" + err.Error())
	}
	Promo := PromotionOrder{
		Address:         PromoKey,
		TransactionName: args[2],
		RiskRate:        int(RiskRate),
		Status:          "Pending",
		Interest:        Interest,
	}
	// validate document existed
	println(">> START validate document existed<<")
	for i := 0; i < len(walletModel.PromotionOrder); i++ {
		println("validate document :" + PromoKey)
		PromoAsBytes, err := stub.GetState(PromoKey)
		if err != nil {
			println("GetState is error" + err.Error())
			return shim.Error("GetState is error" + err.Error())
		}
		if PromoAsBytes != nil {
			println("PromoKey " + PromoKey + " is existed in state Blockchain")
			return shim.Error("PromoKey " + PromoKey + " is existed in state Blockchain")
		}
		println("validate document existed successfully")
	}
	println(">> END validate document existed<<")

	println(">> START parse Token Model to ByteArray <<")
	PromoAsByte, err := json.Marshal(Promo)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")

	//byteArray put tokenAsByte to state blockchain
	println(">> START PromoAsByte PutState to state blockchain <<")
	err = stub.PutState(PromoKey, PromoAsByte)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END PromoAsByte PutState to state blockchain <<")
	walletModel.PromotionOrder = append(walletModel.PromotionOrder, Promo.Address)

	println(">> START parse Token Model to ByteArray <<")
	WalletAsByte, err := json.Marshal(walletModel)
	if err != nil {
		println("Marshal parser token as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse token Model to ByteArray <<")

	//byteArray put tokenAsByte to state blockchain
	println(">> START WalletAsByte PutState to state blockchain <<")
	err = stub.PutState(walletKey, WalletAsByte)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END WalletAsByte PutState to state blockchain <<")
	return shim.Success(PromoAsByte)
}

func (t *SmartContract) LenderGetPromotionOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[LenderGetPromotionOrder]"
	print("==========================================" + functionName + "====================================================")
	println("Input: " + args[0])
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")
	var result []PromotionOrder
	if len(args) != 1 { //wallet address
		println("Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	//Get DigitalWallet
	walletAddress := args[0]
	walletProfileAsByte, err := stub.GetState(walletAddress)
	if walletProfileAsByte == nil {
		println("walletAddress " + walletAddress + " is not defined")
		return shim.Error("walletAddress " + walletAddress + " is not defined")
	}
	if err != nil {
		println("getWalletAddress is error" + err.Error())
		return shim.Error("getWalletAddress is error" + err.Error())
	}

	var walletModel DigitalWallet
	errUnmarshalWalletProfile := json.Unmarshal(walletProfileAsByte, &walletModel)
	if errUnmarshalWalletProfile != nil {
		//error unmarshaling
		println("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
		return shim.Error("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
	}
	if len(walletModel.PromotionOrder) != 0 {
		for _, value := range walletModel.PromotionOrder {
			//Get asset token
			PromoOrderAsByte, err := stub.GetState(value)
			if PromoOrderAsByte == nil {
				println("PromoOrderAsByte " + value + " is not defined")
				return shim.Error("PromoOrderAsByte " + value + " is not defined")
			}
			if err != nil {
				println("PromoOrderAsByte is error" + err.Error())
				return shim.Error("PromoOrderAsByte is error" + err.Error())
			}
			var PromoOrder PromotionOrder
			UnmarshalPromoOrder := json.Unmarshal(PromoOrderAsByte, &PromoOrder)
			if UnmarshalPromoOrder != nil {
				println(" Error " + functionName + " unmarshaling PromoOrder : " + UnmarshalPromoOrder.Error())
				return shim.Error(" Error " + functionName + "  unmarshaling PromoOrder : " + UnmarshalPromoOrder.Error())
			}
			result = append(result, PromoOrder)
		}
	}
	//Parse result
	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		println("Marshal parser result as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser result as Model to ByteArray is error" + err.Error())
	}

	println("query is Successfully ")
	println(functionName + " successfully")
	println("===================================================================================" + functionName + "=============================================================================")
	return shim.Success(resultAsBytes)
}
						
func (t *SmartContract) LenderBuyToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	println("=======================" + "Borrow" + "=======================")

	println("Parse args as string to array")
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")
	println("Parse args as string to array successfully")

	
	var err error
	if len(args) != 3 {
		// 0 Address lender(buyer)	
		// 1 loanDoc				ได้จากmarket
		// 2 token amount			กรอกเอง

 		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}


	walletLenderKey := args[0]
	// Get the state from the ledger

	walletLenderAsBytes, err := stub.GetState(walletLenderKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + walletLenderKey + "\"}"
		return shim.Error(jsonResp)
	}
	if walletLenderAsBytes == nil {
		jsonResp := walletLenderKey
		return shim.Error(jsonResp)
	}
	walletLenderModel := DigitalWallet{}
	errWalletLenderUnmarshal := json.Unmarshal(walletLenderAsBytes, &walletLenderModel)
	if errWalletLenderUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + walletLenderKey + "\"}"
		return shim.Error(jsonResp)
	}

	loanModel := LoanDocument{}
	loanAddress := args[1]

	loanAsBytes, err := stub.GetState(loanAddress)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + loanAddress + "\"}"
		return shim.Error(jsonResp)
	}
	if loanAsBytes == nil {
		jsonResp := loanAddress
		return shim.Error(jsonResp)
	}
	errLoanDocUnmarshal := json.Unmarshal(loanAsBytes, &loanModel)
	if errLoanDocUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for  loan"+"\"}"
		return shim.Error(jsonResp)
	}

	walletBorrowerModel := DigitalWallet{}
	assetBorrowerModel := Asset{}


	tokenAmount, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	var sellingToken []Token
	var tokenModel Token
	var remainDebt float64 = 0

	if (int(tokenAmount) <= len(loanModel.Token)) {
		println("check Token")
		for i := 0; i < int(tokenAmount); i++ {
			TokenAsByte, err := stub.GetState(loanModel.Token[i])
			if err != nil {
				jsonResp := "{\"Error\":\"Failed to get state for token"  + "\"}"
				return shim.Error(jsonResp)
			}
			if TokenAsByte == nil {
				jsonResp := loanModel.Token[i]
				return shim.Error(jsonResp)
			}
			errTokenUnmarshal := json.Unmarshal(TokenAsByte, &tokenModel)
			if errTokenUnmarshal != nil {
				jsonResp := "{\"Error\":\"Failed to get unmarshall for token" + "\"}"
				return shim.Error(jsonResp)
			}
			if(tokenModel.IsSell == true){
				sellingToken = append(sellingToken,tokenModel)
			}
		}

		if (float64(tokenAmount) * sellingToken[0].Rate) <= walletLenderModel.Balance {
			for i:= 0;i< int(tokenAmount); i++ {
				var walletBorrowerKey = ""
				if( sellingToken[i].LenderAddress != ""){
					walletBorrowerKey =  sellingToken[i].LenderAddress
					remainDebt += float64(sellingToken[i].Rate)

				}else {
					walletBorrowerKey =  loanModel.Address					
				}
				walletBorrowerAsBytes, err := stub.GetState(walletBorrowerKey)
					if err != nil {
						jsonResp := "{\"Error\":\"Failed to get state for " + walletBorrowerKey + "\"}"
						return shim.Error(jsonResp)
					}
					if walletBorrowerAsBytes == nil {
						jsonResp := walletBorrowerKey
						return shim.Error(jsonResp)
					}
					errLenderUnmarshal := json.Unmarshal(walletBorrowerAsBytes, &walletBorrowerModel)
					if errLenderUnmarshal != nil {
						jsonResp := "{\"Error\":\"Failed to get unmarshall for " + walletBorrowerKey + "\"}"
						return shim.Error(jsonResp)
					}

					assetBorrowerAddress := sellingToken[i].AssetId
					// Get the state from the ledger
					assetBorrowerAsBytes, err := stub.GetState(assetBorrowerAddress)
					if err != nil {
						jsonResp := "{\"Error\":\"Failed to get state for " + assetBorrowerAddress + "\"}"
						return shim.Error(jsonResp)
					}
					if assetBorrowerAsBytes == nil {
						jsonResp := assetBorrowerAddress
						return shim.Error(jsonResp)
					}
					errAssetLenderUnmarshal := json.Unmarshal(assetBorrowerAsBytes, &assetBorrowerModel)
					if errAssetLenderUnmarshal != nil {
						jsonResp := "{\"Error\":\"Failed to get unmarshall for " + assetBorrowerAddress + "\"}"
						return shim.Error(jsonResp)
					}
					walletLenderModel.Balance -= sellingToken[i].Rate
					walletBorrowerModel.Balance += sellingToken[i].Rate
					assetBorrowerModel.TokenBalance -= 1
					loanModel.RemainDebt += sellingToken[i].Rate
					sellingToken[i].LenderAddress = walletBorrowerModel.Address
					sellingToken[i].IsSell = false
					
					//กันซ้ำ
					// walletLenderModel.LoanDocument = append(walletLenderModel.LoanDocument,loanModel.LoanId)					
					for _, v := range walletLenderModel.LoanDocument {
						if v != loanModel.LoanId {
							walletLenderModel.LoanDocument = append(walletLenderModel.LoanDocument, v)
						}
					}

					//push token
					println(">> START parse Token Model to ByteArray <<")
						tokenAsByte, err := json.Marshal(sellingToken[i])
						if err != nil {
							println("Marshal parser token as Model to ByteArray is error" + err.Error())
							return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
						}
						println(">> END parse token Model to ByteArray <<")

						//byteArray put tokenAsByte to state blockchain
						println(">> START TokenAsByte PutState to state blockchain <<")
						err = stub.PutState(sellingToken[i].TokenId, tokenAsByte)
						if err != nil {
							println("PutState is error" + err.Error())
							return shim.Error("PutState is error" + err.Error())
						}
						println(">> END TokenAsByte PutState to state blockchain <<")

			
				}//loop i 

				loanModel.MinDebtPerMonth = (remainDebt * (float64(100)+loanModel.Interest)) / float64(assetBorrowerModel.LoanDuration)
				//push asset
				println(">> START parse assetBorrowerModel Model to ByteArray <<")
				assetBorrowerAsByte, err := json.Marshal(assetBorrowerModel)
				if err != nil {
					println("Marshal parser assetBorrowerModel as Model to ByteArray is error" + err.Error())
					return shim.Error("Marshal assetBorrowerModel as Model to ByteArray is error" + err.Error())
				}
				println(">> END parse assetBorrowerModel Model to ByteArray <<")

				//byteArray put assetBorrowerAsByte to state blockchain
				println(">> START assetBorrowerAsByte PutState to state blockchain <<")
				err = stub.PutState(assetBorrowerModel.AssetId, assetBorrowerAsByte)
				if err != nil {
					println("PutState is error" + err.Error())
					return shim.Error("PutState is error" + err.Error())
				}
				println(">> END assetBorrowerAsByte PutState to state blockchain <<")

				//push walletlender
				println(">> START parse walletLenderModel Model to ByteArray <<")
				walletLenderAsByte, err := json.Marshal(walletLenderModel)
				if err != nil {
					println("Marshal parser walletLenderModel as Model to ByteArray is error" + err.Error())
					return shim.Error("Marshal walletLenderModel as Model to ByteArray is error" + err.Error())
				}
				println(">> END parse walletLenderModel Model to ByteArray <<")

				//byteArray put walletLenderAsByte to state blockchain
				println(">> START walletLenderAsByte PutState to state blockchain <<")
				err = stub.PutState(walletLenderModel.Address, walletLenderAsByte)
				if err != nil {
					println("PutState is error" + err.Error())
					return shim.Error("PutState is error" + err.Error())
				}
				println(">> END walletLenderAsByte PutState to state blockchain <<")
				
				
				//push walletborrow
				println(">> START parse walletBorrowerModel Model to ByteArray <<")
				walletBorrowerAsByte, err := json.Marshal(walletBorrowerModel)
				if err != nil {
					println("Marshal parser walletBorrowerModel as Model to ByteArray is error" + err.Error())
					return shim.Error("Marshal walletBorrowerModel as Model to ByteArray is error" + err.Error())
				}
				println(">> END parse walletBorrowerModel Model to ByteArray <<")

				//byteArray put walletBorrowerAsByte to state blockchain
				println(">> START walletBorrowerAsByte PutState to state blockchain <<")
				err = stub.PutState(walletBorrowerModel.Address, walletBorrowerAsByte)
				if err != nil {
					println("PutState is error" + err.Error())
					return shim.Error("PutState is error" + err.Error())
				}
				println(">> END walletBorrowerAsByte PutState to state blockchain <<")
				
				//push loan
				println(">> START parse loanModel Model to ByteArray <<")
				loanAsByte, err := json.Marshal(loanModel)
				if err != nil {
					println("Marshal parser loanModel as Model to ByteArray is error" + err.Error())
					return shim.Error("Marshal loanModel as Model to ByteArray is error" + err.Error())
				}
				println(">> END parse loanModel Model to ByteArray <<")

				//byteArray put loanAsByte to state blockchain
				println(">> START loanAsByte PutState to state blockchain <<")
				err = stub.PutState(loanModel.LoanId, loanAsByte)
				if err != nil {
					println("PutState is error" + err.Error())
					return shim.Error("PutState is error" + err.Error())
				}
				println(">> END loanAsByte PutState to state blockchain <<")
			}
			
		}
		println(">> START parse loanModel2 Model to ByteArray <<")
				loanAsByte, err := json.Marshal(loanModel)
				if err != nil {
					println("Marshal parser loanModel as Model to ByteArray is error" + err.Error())
					return shim.Error("Marshal loanModel as Model to ByteArray is error" + err.Error())
				}
				println(">> END parse loanModel2 Model to ByteArray <<")
		return shim.Success(loanAsByte)
	}


	
	//===================== transaction ====================
	// transactionModel := WalletTransaction {
	// 	Address         	:	
	// 	IssueDate       	:
	// 	TransactionName 	:
	// 	TransactionType 	:
	// 	Total           	:
	// 	SnapshotBalance		:
	// }
	// walletLenderModel.Transaction = append(walletModel.Transaction,)
	




func (t *SmartContract) GetMarketplace(stub shim.ChaincodeStubInterface) pb.Response {
	functionName := "[GetMarketplace]"
	var err error
	marketKey := marketKey
	var result MarketAssetList

	//Get Market
	documentMarket, err := stub.GetState(marketKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + marketKey + "\"}"
		return shim.Error(jsonResp)
	}
	if documentMarket == nil {
		jsonResp := marketKey
		return shim.Error(jsonResp)
	}
	var market Market
	UnmarlshalMarket := json.Unmarshal(documentMarket, &market)
	if UnmarlshalMarket != nil {
		println(" Error " + functionName + " unmarshaling MarketModel : " + UnmarlshalMarket.Error())
		return shim.Error(" Error " + functionName + "  unmarshaling MarketModel : " + UnmarlshalMarket.Error())
	}

	//Get loan in market
	for _, value := range market.LoanId {
		var assetList AssetList

		//Get loan
		loanAsByte, err := stub.GetState(value)
		if loanAsByte == nil {
			println("loanAsByte " + value + " is not defined")
			return shim.Error("loanAsByte " + value + " is not defined")
		}
		if err != nil {
			println("getLoanAddress is error" + err.Error())
			return shim.Error("getLoanAddress is error" + err.Error())
		}

		errUnmarshalLoanAsByte := json.Unmarshal(loanAsByte, &assetList.LoanInfo)
		if errUnmarshalLoanAsByte != nil {
			//error unmarshaling
			println("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
			return shim.Error("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
		}

		//Get asset
		AssetAddress := assetList.LoanInfo.AssetId
		AssetAsByte, err := stub.GetState(AssetAddress)
		if AssetAsByte == nil {
			println("AssetAsByte " + AssetAddress + " is not defined")
			return shim.Error("AssetAsByte " + AssetAddress + " is not defined")
		}
		if err != nil {
			println("getAssetAddress is error" + err.Error())
			return shim.Error("getAssetAddress is error" + err.Error())
		}

		errUnmarshalAssetAsByte := json.Unmarshal(AssetAsByte, &assetList.AssetInfo)
		if errUnmarshalAssetAsByte != nil {
			//error unmarshaling
			println("assetModel Error unmarshaling walletAddress:" + errUnmarshalAssetAsByte.Error())
			return shim.Error("assetModel Error unmarshaling walletAddress:" + errUnmarshalAssetAsByte.Error())
		}

		if len(assetList.LoanInfo.Token) != 0 {
			for _, value := range assetList.LoanInfo.Token {
				//Get asset token
				assetTokenAsByte, err := stub.GetState(value)
				if assetTokenAsByte == nil {
					println("assetTokenAsByte " + value + " is not defined")
					return shim.Error("assetTokenAsByte " + value + " is not defined")
				}
				if err != nil {
					println("assetTokenAsByte is error" + err.Error())
					return shim.Error("assetTokenAsByte is error" + err.Error())
				}

				var assetToken Token
				errUnmarshalLoanAsByte := json.Unmarshal(assetTokenAsByte, &assetToken)
				if errUnmarshalLoanAsByte != nil {
					//error unmarshaling
					println("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
					return shim.Error("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
				}

				if assetToken.IsSell == true {
					assetList.AssetToken = append(assetList.AssetToken, assetToken)
				}
			}
		}

		result.AssetList = append(result.AssetList, assetList)

	}

	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		println("Marshal parser result as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser result as Model to ByteArray is error" + err.Error())
	}

	fmt.Printf("Query Response:%s\n", resultAsBytes)
	return shim.Success(resultAsBytes)
}

func (t *SmartContract) BorrowerGetOwnerAssetList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[BorrowerGetOwnerAssetList]"
	var result BorrowerGetOwnerAssetList

	print("==========================================" + functionName + "====================================================")
	println("Input: " + args[0])
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	if len(args) != 1 { //wallet address
		println("Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//Get DigitalWallet
	walletAddress := args[0]
	walletProfileAsByte, err := stub.GetState(walletAddress)
	if walletProfileAsByte == nil {
		println("walletAddress " + walletAddress + " is not defined")
		return shim.Error("walletAddress " + walletAddress + " is not defined")
	}
	if err != nil {
		println("getWalletAddress is error" + err.Error())
		return shim.Error("getWalletAddress is error" + err.Error())
	}

	var walletModel DigitalWallet
	errUnmarshalWalletProfile := json.Unmarshal(walletProfileAsByte, &walletModel)
	if errUnmarshalWalletProfile != nil {
		//error unmarshaling
		println("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
		return shim.Error("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
	}

	for _, value := range walletModel.BorrowerAsset {
		var assetList AssetList
		//Get My Asset
		AssetAddress := value
		AssetAsByte, err := stub.GetState(AssetAddress)
		if AssetAsByte == nil {
			println("AssetAsByte " + AssetAddress + " is not defined")
			return shim.Error("AssetAsByte " + AssetAddress + " is not defined")
		}
		if err != nil {
			println("getAssetAddress is error" + err.Error())
			return shim.Error("getAssetAddress is error" + err.Error())
		}

		errUnmarshalAssetAsByte := json.Unmarshal(AssetAsByte, &assetList.AssetInfo)
		if errUnmarshalAssetAsByte != nil {
			//error unmarshaling
			println("assetModel Error unmarshaling walletAddress:" + errUnmarshalAssetAsByte.Error())
			return shim.Error("assetModel Error unmarshaling walletAddress:" + errUnmarshalAssetAsByte.Error())
		}

		if assetList.AssetInfo.LoanId != "" {
			//Get loan
			println(assetList.AssetInfo.LoanId)
			loanAsByte, err := stub.GetState(assetList.AssetInfo.LoanId)
			if loanAsByte == nil {
				println("loanAsByte " + assetList.AssetInfo.LoanId + " is not defined")
				return shim.Error("loanAsByte " + assetList.AssetInfo.LoanId + " is not defined")
			}
			if err != nil {
				println("getLoanAddress is error" + err.Error())
				return shim.Error("getLoanAddress is error" + err.Error())
			}

			errUnmarshalLoanAsByte := json.Unmarshal(loanAsByte, &assetList.LoanInfo)
			if errUnmarshalLoanAsByte != nil {
				//error unmarshaling
				println("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
				return shim.Error("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
			}

			if len(assetList.LoanInfo.Token) != 0 {
				for _, value := range assetList.LoanInfo.Token {
					//Get asset token
					assetTokenAsByte, err := stub.GetState(value)
					if assetTokenAsByte == nil {
						println("assetTokenAsByte " + value + " is not defined")
						return shim.Error("assetTokenAsByte " + value + " is not defined")
					}
					if err != nil {
						println("assetTokenAsByte is error" + err.Error())
						return shim.Error("assetTokenAsByte is error" + err.Error())
					}

					var assetToken Token
					errUnmarshalLoanAsByte := json.Unmarshal(assetTokenAsByte, &assetToken)
					if errUnmarshalLoanAsByte != nil {
						//error unmarshaling
						println("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
						return shim.Error("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
					}
					assetList.AssetToken = append(assetList.AssetToken, assetToken)
				}
			}
		}

		result.AssetList = append(result.AssetList, assetList)
	}

	//Parse result
	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		println("Marshal parser result as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser result as Model to ByteArray is error" + err.Error())
	}

	println("query is Successfully ")
	println(functionName + " successfully")
	println("===================================================================================" + functionName + "=============================================================================")
	return shim.Success(resultAsBytes)
}

func (t *SmartContract) LenderGetAssetLendingList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[LenderGetAssetLendingList]"
	var result BorrowerGetOwnerAssetList

	print("==========================================" + functionName + "====================================================")
	println("Input: " + args[0])
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	if len(args) != 1 { //wallet address
		println("Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//Get DigitalWallet
	walletAddress := args[0]
	walletProfileAsByte, err := stub.GetState(walletAddress)
	if walletProfileAsByte == nil {
		println("walletAddress " + walletAddress + " is not defined")
		return shim.Error("walletAddress " + walletAddress + " is not defined")
	}
	if err != nil {
		println("getWalletAddress is error" + err.Error())
		return shim.Error("getWalletAddress is error" + err.Error())
	}

	var walletModel DigitalWallet
	errUnmarshalWalletProfile := json.Unmarshal(walletProfileAsByte, &walletModel)
	if errUnmarshalWalletProfile != nil {
		//error unmarshaling
		println("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
		return shim.Error("walletModel Error unmarshaling walletAddress:" + errUnmarshalWalletProfile.Error())
	}

	for _, value := range walletModel.LenderLoan {
		var assetList AssetList

		//Get loan
		loanAsByte, err := stub.GetState(value)
		if loanAsByte == nil {
			println("loanAsByte " + value + " is not defined")
			return shim.Error("loanAsByte " + value + " is not defined")
		}
		if err != nil {
			println("getLoanAddress is error" + err.Error())
			return shim.Error("getLoanAddress is error" + err.Error())
		}

		errUnmarshalLoanAsByte := json.Unmarshal(loanAsByte, &assetList.LoanInfo)
		if errUnmarshalLoanAsByte != nil {
			//error unmarshaling
			println("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
			return shim.Error("loanModel Error unmarshaling LoanDocument:" + errUnmarshalLoanAsByte.Error())
		}

		if len(assetList.LoanInfo.Token) != 0 {
			for _, value := range assetList.LoanInfo.Token {
				//Get asset token
				assetTokenAsByte, err := stub.GetState(value)
				if assetTokenAsByte == nil {
					println("assetTokenAsByte " + value + " is not defined")
					return shim.Error("assetTokenAsByte " + value + " is not defined")
				}
				if err != nil {
					println("assetTokenAsByte is error" + err.Error())
					return shim.Error("assetTokenAsByte is error" + err.Error())
				}

				var assetToken Token
				errUnmarshalLoanAsByte := json.Unmarshal(assetTokenAsByte, &assetToken)
				if errUnmarshalLoanAsByte != nil {
					//error unmarshaling
					println("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
					return shim.Error("assetTokenAsByte Error unmarshaling assetTokenAsByte:" + errUnmarshalLoanAsByte.Error())
				}

				//get my token loan only
				if assetToken.LenderAddress == walletAddress {
					assetList.AssetToken = append(assetList.AssetToken, assetToken)
				}
			}
		}

		//Get My Asset
		AssetAddress := assetList.LoanInfo.AssetId
		println("Get My Asset : " + AssetAddress)
		AssetAsByte, err := stub.GetState(AssetAddress)
		if AssetAsByte == nil {
			println("AssetAsByte " + AssetAddress + " is not defined")
			return shim.Error("AssetAsByte " + AssetAddress + " is not defined")
		}
		if err != nil {
			println("getAssetAddress is error" + err.Error())
			return shim.Error("getAssetAddress is error" + err.Error())
		}

		errUnmarshalAssetAsByte := json.Unmarshal(AssetAsByte, &assetList.AssetInfo)
		if errUnmarshalAssetAsByte != nil {
			//error unmarshaling
			println("assetModel Error unmarshaling assetAddress:" + errUnmarshalAssetAsByte.Error())
			return shim.Error("assetModel Error unmarshaling assetAddress:" + errUnmarshalAssetAsByte.Error())
		}

		result.AssetList = append(result.AssetList, assetList)
	}

	//Parse result
	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		println("Marshal parser result as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser result as Model to ByteArray is error" + err.Error())
	}

	println("query is Successfully ")
	println(functionName + " successfully")
	println("===================================================================================" + functionName + "=============================================================================")
	return shim.Success(resultAsBytes)
}

func (t *SmartContract) Borrow(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	println("=======================" + "Borrow" + "=======================")

	println("Parse args as string to array")
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")
	println("Parse args as string to array successfully")

	var err error
	if len(args) != 4 {
		// 0 Address
		// 1 AssetId
		// 2 loan
		// 3 token amount
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	walletKey := args[0]
	// Get the state from the ledger

	walletAsBytes, err := stub.GetState(walletKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}
	if walletAsBytes == nil {
		jsonResp := walletKey
		return shim.Error(jsonResp)
	}
	walletModel := DigitalWallet{}
	errWalletUnmarshal := json.Unmarshal(walletAsBytes, &walletModel)
	if errWalletUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + walletKey + "\"}"
		return shim.Error(jsonResp)
	}

	assetKey := args[1]
	assetAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + assetKey + "\"}"
		return shim.Error(jsonResp)
	}
	if assetAsBytes == nil {
		jsonResp := assetKey
		return shim.Error(jsonResp)
	}
	assetModel := Asset{}
	errAssetUnmarshal := json.Unmarshal(assetAsBytes, &assetModel)
	if errAssetUnmarshal != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + assetKey + "\"}"
		return shim.Error(jsonResp)
	}

	if assetModel.Status != "approved" {
		jsonResp := "{\"Error\":\"Failed It's Already be collateral for " + assetKey + "\"}"
		return shim.Error(jsonResp)
	}	
	

	loan, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		println("ParseFloat is error" + err.Error())
		return shim.Error("ParseFloat is error" + err.Error())
	}

	tokenAmount, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}

	loanModel := LoanDocument{
		LoanId          : 	"0xS0HuhoMs0fMN8M9xWTtWnsDDzJytVL6B9nHq8OTWmstV1nbZY4sInXXWj39NcR0D",	//ชั่วคราว
		AssetId         :	assetModel.AssetId,
		Address         :	assetModel.Address,
		Loan           	:	loan	,
		RemainDebt      :	0,
		MinDebtPerMonth :	0,
		Interest        :	assetModel.LoanInterest,
		Status          :	"Pending",
		Token           :	nil,
	}
	
	walletModel.LoanDocument = append(walletModel.LoanDocument,loanModel.LoanId)



	

	println("=======================" + "Create token" + "=======================")

	var i int64
	for i = 0; i < tokenAmount; i++ {
		var hashToken = loanModel.LoanId + fmt.Sprintf("%d", i)
		tokenModel := Token{
			TokenId:       hashString(hashToken),
			AssetId:       loanModel.AssetId,
			LenderAddress: "",
			Rate:          loan / float64(tokenAmount),
			IsSell:        true,
			IssueDate:     (time.Now()).Truncate(24 * time.Hour),
		}

		println(">> START parse Token Model to ByteArray <<")
		tokenAsByte, err := json.Marshal(tokenModel)
		if err != nil {
			println("Marshal parser token as Model to ByteArray is error" + err.Error())
			return shim.Error("Marshal token as Model to ByteArray is error" + err.Error())
		}
		println(">> END parse token Model to ByteArray <<")

		//byteArray put tokenAsByte to state blockchain
		println(">> START tokenAsByte PutState to state blockchain <<")
		err = stub.PutState(tokenModel.TokenId, tokenAsByte)
		if err != nil {
			println("PutState is error" + err.Error())
			return shim.Error("PutState is error" + err.Error())
		}
		println(">> END tokenAsByte PutState to state blockchain <<")

		walletModel.BorrowerToken = append(walletModel.BorrowerToken, tokenModel.TokenId)
		loanModel.Token = append(loanModel.Token, tokenModel.TokenId)

	} //loop i

	//Get Market
	documentMarket, err := stub.GetState(marketKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + marketKey + "\"}"
		return shim.Error(jsonResp)
	}
	if documentMarket == nil {
		jsonResp := marketKey
		return shim.Error(jsonResp)
	}
	var market Market
	UnmarlshalMarket := json.Unmarshal(documentMarket, &market)
	if UnmarlshalMarket != nil {
		jsonResp := "{\"Error\":\"Failed to get unmarshall for " + marketKey + "\"}"
		return shim.Error(jsonResp)
	}
	market.LoanId = append(market.LoanId,loanModel.LoanId)


	println(">> START parse market Model to ByteArray <<")
	marketAsBytes,err := json.Marshal(market)
	if err != nil {
		println("Marshal parser market as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser loan as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse market Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START marketAsBytes PutState to state blockchain <<")
	err = stub.PutState(marketKey, marketAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END marketAsBytes PutState to state blockchain <<")



	println(">> START parse loan Model to ByteArray <<")
	loanAsBytes, err := json.Marshal(loanModel)
	if err != nil {
		println("Marshal parser loan as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser loan as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse loan Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START loanAsBytes PutState to state blockchain <<")
	err = stub.PutState(loanModel.LoanId, loanAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END loanAsBytes PutState to state blockchain <<")

	println(">> START parse Wallet Model to ByteArray <<")
	walletAsByte, err := json.Marshal(walletModel)
	if err != nil {
		println("Marshal parser wallet as Model to ByteArray is error" + err.Error())
		return shim.Error("Marshal wallet loan as Model to ByteArray is error" + err.Error())
	}
	println(">> END parse wallet Model to ByteArray <<")

	//byteArray put to state blockchain
	println(">> START walletAsByte PutState to state blockchain <<")
	err = stub.PutState(walletKey, walletAsByte)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}
	println(">> END walletAsByte PutState to state blockchain <<")

	// jsonResp := "Create LoanDocument Success"
	// fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(loanAsBytes)
}

