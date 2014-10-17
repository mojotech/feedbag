"use strict"

# Development specific configuration
# ==================================
module.exports =

  # MongoDB connection options
  mongo:
    uri: "mongodb://localhost/feedbag-dev"
    name: "feedbag-dev"
    host: "localhost"
    port: null
    username: null
    password: null

  # RabbitMQ connection options
  rabbit:
    uri: "amqp://localhost"
    workerQueues: ["jobs.github", "activity.create", "events.create"]

  # Github api details
  github:
    clientID: "882c03c2bb5a47c6831a"
    clientSecret: "1a13fbe5c2a11a2067312be4b43d783e8fedf70b"
    callbackURL: "http://localhost:9000/auth/github/callback"
