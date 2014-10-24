###
Socket.io configuration
###
"use strict"

config = require("./environment")

# When the user disconnects.. perform this
onDisconnect = (socket) ->

# When the user connects.. perform this
onConnect = (socket) ->

  # When the client emits 'info', this listens and executes
  socket.on "info", (data) ->
    console.info "#{socket.address} #{JSON.stringify(data, null, 2)}"


  # Insert sockets below
  require("../api/activity/activity.socket").register socket

module.exports = (socketio) ->

  # socket.io (v1.x.x) is powered by debug.
  # In order to see all the debug output, set DEBUG (in server/config/local.env.js) to including the desired scope.
  #
  # ex: DEBUG: "http*,socket.io:socket"

  # We can authenticate socket.io users and access their token through socket.handshake.decoded_token
  #
  # 1. You will need to send the token in `client/components/socket/socket.service.js`
  #
  # 2. Require authentication here:
  # socketio.use(require('socketio-jwt').authorize({
  #   secret: config.secrets.session,
  #   handshake: true
  # }));
  socketio.on "connection", (socket) ->
    socket.address = (if socket.handshake.address isnt null then socket.handshake.address.address + ":" + socket.handshake.address.port else process.env.DOMAIN)
    socket.connectedAt = new Date()

    # Call onDisconnect.
    socket.on "disconnect", ->
      onDisconnect socket
      console.info "[#{socket.address}] DISCONNECTED"

    # Call onConnect.
    onConnect socket
    console.info "[#{socket.address}] CONNECTED"
