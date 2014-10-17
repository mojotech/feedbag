###
Error responses
###
"use strict"

module.exports[500] = internalServerError = (err, req, res, next) ->
  res.status(err.status or 500).json
      status: err.status or 500
      error:
          type: err.message or 'internal_server_error'
          message: err.stack
      request:
          method: req.method
          host: req.headers.host
          path: req._parsedUrl.pathname
          params: req.query


module.exports[401] = unauthorized = (err, req, res, next) ->
  if err.status is 401
    res.status(err.status).json
      status: err.status
      error:
        type: "Authroization"
        message: err.message or 'No Authorized'
      request:
        method: req.method
        host: req.headers.host
        path: req._parsedUrl.pathname
        params: req.query
  else
    next()


module.exports[404] = pageNotFound = (req, res) ->
  viewFilePath = "404"
  statusCode = 404
  result = status: statusCode
  res.status result.status

  res.render viewFilePath, (err) ->
    return res.json(result, result.status)  if err
    res.render viewFilePath
