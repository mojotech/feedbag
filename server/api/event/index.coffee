"use strict"

express = require("express")
controller = require("./event.controller")
auth = require("../../auth/auth.service")

router = express.Router()

router.get "/", controller.index
router.get "/populate", controller.test_populate_events
router.get "/clear", controller.clear

module.exports = router
