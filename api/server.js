
var express = require('./config/express')
const logger = require('./Util/logger.js')
if (process.env.ROLE == "production") {
  var mongoose = require('./config/mongoose')
  var db = mongoose();
}
var app = express()

var service = require('./blockchain/service')
new service().Init()

const port = process.env.PORT || 8000
// const host = "localhost"
app.listen(port,'0.0.0.0', () => {
  logger.info('http://localhost:' + port)
  logger.debug(`[Role] ${process.env.ROLE}`)
  logger.info('Start server at port ' + port)
})