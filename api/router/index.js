module.exports = function (app) {
    var request = require('../controller/operation')
    var verifyAPIkey = require('../Util/verifyAPIkey')
    const logger = require('../Util/logger.js');

    app.post('/IssuePromotionOrder', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().IssuePromotionOrder(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    app.post('/LenderBuyToken', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().LenderBuyToken(req.body))
            res.status(201)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    app.post('/LenderSellToken', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().LenderSellToken(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    app.post('/BorrowerGetOwnerAssetList', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().BorrowerGetOwnerAssetList(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })

    app.post('/LenderGetAssetLendingList', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().LenderGetAssetLendingList(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    app.post('/LenderGetPromotionOrder', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().LenderGetPromotionOrder(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    app.post('/GetWallet', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().GetWallet(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    // -----------------------------GetMarketplace------------------------
    app.post('/GetMarketplace', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().GetMarketplace(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    // -----------------------------registerUser------------------------
    app.post('/registerUser', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().registerUser(req.body))
            res.status(result.statusCode)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    // -----------------------------IssueBorrow------------------------
    app.post('/IssueBorrow', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().IssueBorrow(req.body))
            res.status(201)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
    // -----------------------------GetLoanDoc------------------------
    app.post('/GetLoanDoc', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().GetLoanDoc(req.body))
            res.status(200)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })

    app.post('/GetTokenInfo', verifyAPIkey, async (req, res) => {
        try {
            var result = (await new request().GetTokenInfo(req.body))
            res.status(200)
            res.json(result)
        } catch (error) {
            let messageError = {
                statusCode: error.statusCode || 400,
                message: error.message || error
            }
            logger.error(messageError.message)

            res.status(messageError.statusCode)
            res.json(messageError)
        }
    })
}

