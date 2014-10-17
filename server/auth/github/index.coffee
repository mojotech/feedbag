"use strict"

express = require("express")
passport = require("passport")
auth = require("../auth.service")
router = express.Router()

router.get("/", passport.authenticate "github",
    scope: 'repo'
    failureRedirect: "/login"
    session: false

).get "/callback", passport.authenticate("github",
    scope: 'repo'
    failureRedirect: "/login"
    session: false
  ), auth.setTokenCookie


module.exports = router
