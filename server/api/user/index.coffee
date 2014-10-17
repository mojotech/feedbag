"use strict"

express = require("express")
controller = require("./user.controller")
config = require("../../config/environment")
auth = require("../../auth/auth.service")

router = express.Router()

router.get "/", auth.isAuthenticated(), controller.index
router.delete '/:id', auth.isAuthenticated(), controller.destroy
router.get '/me', auth.isAuthenticated(), controller.me
router.get '/:id', auth.isAuthenticated(), controller.show

module.exports = router
