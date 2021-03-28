const service = require('../blockchain/service.js')
const attribute = require('./attribute')
const converthash = require('../Util/hash256.js')
const logger = require('../Util/logger.js');
const uuid = require('../Util/uuid.js')
const CC_NAME_ISSUE_GARDEN = "IssueGarden"
const CC_NAME_ISSUE_Borrow = "IssueBorrow"
const CC_NAME_ISSUE_PROMOTION_ORDER = "IssuePromotionOrder"
const CC_NAME_LENDER_SELL_TOKEN = "LenderSellToken"
const CC_NAME_LENDER_BUY_TOKEN = "LenderBuyToken"

class request {
    // -----------------------------LenderSellToken------------------------------------------------------------
    async LenderSellToken(unparsedAttrs){
        let functionName = `[LenderSellToken]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let parsedAttrs = await attribute.ParseLenderSellTokenAttr(unparsedAttrs)
                var result = await new service().invoke(bc_user.toString().trim(), CC_NAME_LENDER_SELL_TOKEN, parsedAttrs)
                let message = {
                    statusCode: 201,
                    message: {
                        sell_list : JSON.parse(result.toString())
                    }
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }
    // -----------------------------LenderGetPromotionOrder------------------------------------------------------------
    async LenderGetPromotionOrder(unparsedAttrs){
        let functionName = `[LenderGetPromotionOrder]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let wallet_address = unparsedAttrs.wallet_address || reject(`${functionName}.address is required`)
                
                var result = await new service().queryWithArg(bc_user.toString().trim(), "LenderGetPromotionOrder", [wallet_address])
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }
    // -----------------------------IssuePromotionOrder------------------------------------------------------------
    async IssuePromotionOrder(unparsedAttrs) {
        var functionName = '[IssuePromotionOrder]'
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                var bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                var id = uuid.uuid()
                var hash_promotion = converthash.hash(`${id} + ${bc_user.toString().trim()}`)
                var address = "0x" + hash_promotion
                var parsedAttrs = await attribute.ParseIssuePromotionOrderAttr(unparsedAttrs, address)
                var result = await new service().invoke(bc_user.toString().trim(), CC_NAME_ISSUE_PROMOTION_ORDER, parsedAttrs)
                let message = {
                    statusCode: 201,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed : [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }
    async IssueBorrow(unparsedAttrs){
        let functionName = `[IssueBorrow]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                
                var parsedAttrs = await attribute.ParseIssueBorrowAttr(unparsedAttrs)
                var result = await new service().invoke(bc_user.toString().trim(), CC_NAME_ISSUE_Borrow, parsedAttrs)
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async GetLoanDoc(unparsedAttrs){
        let functionName = `[GetLoanDoc]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let loan_id = unparsedAttrs.loan_id || reject(`${functionName}.loan_id is required`)
                
                var result = await new service().query(bc_user.trim(), loan_id)
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async GetTokenInfo(unparsedAttrs){
        let functionName = `[GetTokenInfo]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let token_id = unparsedAttrs.loan_id || reject(`${functionName}.token_id is required`)
                
                var result = await new service().query(bc_user.trim(), token_id)
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async LenderBuyToken(unparsedAttrs){
        let functionName = `[LenderBuyToken]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                var parsedAttrs = await attribute.ParseLenderBuyTokenAttr(unparsedAttrs)

                var result = await new service().invoke(bc_user.toString().trim(), CC_NAME_LENDER_BUY_TOKEN, parsedAttrs)                
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed :  [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async BorrowerGetOwnerAssetList(unparsedAttrs){
        let functionName = `[BorrowerGetOwnerAssetList]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let wallet_address = unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`)

                var result = await new service().queryWithArg(bc_user.toString().trim(), "BorrowerGetOwnerAssetList", [wallet_address])
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed : Address=${Address} [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async LenderGetAssetLendingList(unparsedAttrs){
        let functionName = `[LenderGetAssetLendingList]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                let wallet_address = unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`)

                var result = await new service().queryWithArg(bc_user.toString().trim(), "LenderGetAssetLendingList", [wallet_address])
                let message = {
                    statusCode: 200,
                    message: JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed : Address=${Address} [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }

    async GetWallet(unparsedAttrs){
        let functionName = `[GetWallet]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                var Address = unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`)
                var result = await new service().query(bc_user.trim(), Address)
                // result = JSON.parse(result.toString())
                let message = {
                    statusCode: 200,
                    message : JSON.parse(result.toString())
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed : Address=${Address} [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }
    // -----------------------------GetMarketplace----------------------
    async GetMarketplace(unparsedAttrs){
        let functionName = `[GetMarketplace]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                let bc_user = unparsedAttrs.bc_user || reject(`${functionName}.bc_user is required`)
                var result = await new service().queryWithArg(bc_user.toString().trim(), "GetMarketplace", [])
                // result = JSON.parse(result.toString())
                let message = {
                    statusCode: 200,
                    Message: JSON.parse(result.toString())
                }
                // let message
                // result.forEach(element => {
                //     message = {
                //         AssetName :   element.asset_name,
                //         Amount :   element.token_amount,
                //         Balance  :   element.token_balance
                //     }
                //     msg.Message.push(message)
                // });
                
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} failed : Address=${Address} [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }





    //=====================================================================================================

    // -----------------------------registerUser----------------------
    async registerUser(unparsedAttrs) {
        let functionName = `[registerUser]`
        logger.info(functionName)
        return new Promise(async function (resolve, reject) {
            try {
                var bc_user = unparsedAttrs.bc_user || reject(`registerUser.bc_user is required`)
                var OrgDepartment = unparsedAttrs.OrgDepartment || reject(`registerUser.OrgDepartment is required`)
                var result = await new service().registerUser(bc_user.toString().trim(), OrgDepartment.toString().trim())
                let message = {
                    statusCode: 201,
                    message: result
                }
                resolve(message)
            } catch (error) {
                let messageError = {
                    statusCode: error.statusCode || 400,
                    message: error.message || `${functionName} registerUser failed : bc_user=${bc_user} ,OrgDepartment=${OrgDepartment} [Error] ${error}`
                }
                logger.error(messageError.message)
                reject(messageError)
            }
        })
    }
}
module.exports = request
