requiredProcessEnv = (name) ->
  throw new Error("You must set the " + name + " environment variable")  unless process.env[name]
  process.env[name]

"use strict"

path = require("path")
_ = require("lodash")
glob = require('glob')

# All configurations will extend these options
# ============================================
all =
  env: process.env.NODE_ENV or "development"

  # Root path of server
  root: path.normalize(__dirname + "/../../..")

  # Server port
  port: process.env.PORT or 9000

  # Secret for session, you will want to change this and make it an environment variable
  secrets:
    session: "Feedbag-secret"


  # List of user roles
  userRoles: [
    "guest"
    "user"
    "admin"
  ]

  # MongoDB connection options
  mongo:
    options:
      db:
        safe: true

  getGlobbedFiles: (globPatterns, removeRoot) ->

    # URL paths regex
    urlRegex = new RegExp("^(?:[a-z]+:)?//", "i")

    # The output array
    output = []

    # If glob pattern is array so we use each pattern in a recursive way, otherwise we use glob
    if _.isArray(globPatterns)
      globPatterns.forEach (globPattern) ->
        output = _.union(output, @getGlobbedFiles(globPattern, removeRoot))

    else if _.isString(globPatterns)
      if urlRegex.test(globPatterns)
        output.push globPatterns
      else
        glob globPatterns,
          sync: true
        , (err, files) ->
          if removeRoot
            files = files.map((file) ->
              file.replace removeRoot, ""
            )
          output = _.union(output, files)

    return output


# Export the config object based on the NODE_ENV
# ==============================================
module.exports = _.merge(all, require("./" + all.env) or {})
