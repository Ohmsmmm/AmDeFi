var express = require('express')
var bodyParser = require('body-parser')
const cors = require('cors')

const ORG = process.env.ORG
const ROLE = process.env.ROLE || "developer"
module.exports = function()
{
    var app = express()
    app.use(bodyParser.json())
    app.use(bodyParser.urlencoded({ extended: true }))
    app.use(bodyParser.text())
    app.use(cors({ origin: '*' }));

    require('../router/index.js')(app)
    return app
}
