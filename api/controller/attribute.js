const logger = require('../Util/logger.js');

    function ParseLenderBuyTokenAttr(unparsedAttrs) {
        let functionName = '[ParseLenderBuyTokenAttr]';
        return new Promise(async function (resolve, reject) {
            try {
                var ParsedAttrs = {
                    wallet_address : unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`),
                    loan_id : unparsedAttrs.loan_id ||  reject(`${functionName}.loan_id is required`),
                    token_amount : unparsedAttrs.token_amount ||  reject(`${functionName}.token_amount is required`),
                }

                resolve([
                    ParsedAttrs.wallet_address.toString().trim(),
                    ParsedAttrs.loan_id.toString().trim(),
                    ParsedAttrs.token_amount.toString().trim(),
                ])
            }
            catch (error) {
                console.log(`${functionName} Parsing treatment failed ${error}`)
                reject(`${functionName} Parsing treatment failed ${error}`)
            }
        })
    }


    function ParseLenderSellTokenAttr(unparsedAttrs) {
        let functionName = '[ParseLenderSellTokenAttr]';
        return new Promise(async function (resolve, reject) {
            try {
                var ParsedAttrs = {
                    wallet_address : unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`),
                    loan_id : unparsedAttrs.loan_id || reject(`${functionName}.loan_document is required`),
                    token_amount : unparsedAttrs.token_amount || reject(`${functionName}.token_amount is required`),
                }

                resolve([
                    ParsedAttrs.wallet_address.toString().trim(),
                    ParsedAttrs.loan_id.toString().trim(),
                    ParsedAttrs.token_amount.toString().trim(),
                ])
            }
            catch (error) {
                console.log(`${functionName} Parsing treatment failed ${error}`)
                reject(`${functionName} Parsing treatment failed ${error}`)
            }
        })
    }

            
function ParseIssueBorrowAttr(unparsedAttrs) {
    let functionName = '[ParseIssueBorrowAttr]';
    return new Promise(async function (resolve, reject) {
        try {
            var ParsedAttrs = {
                
                wallet_address : unparsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`),
                asset_id : unparsedAttrs.asset_id ||  reject(`${functionName}.asset_id is required`),
                loan : unparsedAttrs.loan ||  reject(`${functionName}.loan is required`),
                token_amount : unparsedAttrs.token_amount ||  reject(`${functionName}.token_amount is required`),
            }

            resolve([
                ParsedAttrs.wallet_address.toString().trim(),
                ParsedAttrs.asset_id.toString().trim(),
                ParsedAttrs.loan.toString().trim(),
                ParsedAttrs.token_amount.toString().trim(),
            ])
        }
        catch (error) {
            logger.error(`${functionName} Parsing treatment failed ${error}`)
            reject(`${functionName} Parsing treatment failed ${error}`)
        }
    })
}


// -----------------------------ParseIssuePromotionOrderAttr---------------------------------------------------------- 
function ParseIssuePromotionOrderAttr(unParsedAttrs, hash_promotion) {
    let functionName = '[ParseIssuePromotionAttr(unParsedAttrs, hash_promotion)]';
    return new Promise(async function (resolve, reject) {
        try {
            var ParsedAttrs = {
                wallet_address: unParsedAttrs.wallet_address || reject(`${functionName}.wallet_address is required`),
                value: parseFloat(unParsedAttrs.value) || reject(`${functionName}.value is required`),
                risk_rate: parseInt(unParsedAttrs.risk_rate) || reject(`Issue${functionName}.risk_rate is required`),
                interest: parseFloat(unParsedAttrs.interest) || reject(`Issue${functionName}.interest is required`),
            }

            resolve([
                ParsedAttrs.wallet_address.toString().trim(),
                hash_promotion.toString().trim(),
                ParsedAttrs.value.toString().trim(),
                ParsedAttrs.risk_rate.toString().trim(),
                ParsedAttrs.interest.toString().trim()
            ])
        }
        catch (error) {
            logger.error(`${functionName} Parsing treatment failed ${error}`)
            reject(`${functionName} Parsing treatment failed ${error}`)
        }
    })
}

module.exports = {
    ParseIssuePromotionOrderAttr: ParseIssuePromotionOrderAttr,
    ParseIssueBorrowAttr: ParseIssueBorrowAttr,
    ParseLenderBuyTokenAttr:ParseLenderBuyTokenAttr,
    ParseLenderSellTokenAttr: ParseLenderSellTokenAttr,
}
