activity = require("../activity/activity.collection")

class Socket
  constructor: -> @initialize()

  initialize: ->
    # Connect to socket server
    @socket = io("http://localhost:3000")

    @socket.on "activity", @processActivity

  processActivity: (activities) -> activity.add(activities)

module.exports = new Socket()
