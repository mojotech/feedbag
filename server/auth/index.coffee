"use strict"
express = require("express")
passport = require("passport")
config = require("../config/environment")
User = require("../api/user/user.model")

# Passport Configuration
require('./github/passport').setup User, config

router = express.Router()

router.use "/github", require("./github")

module.exports = router
