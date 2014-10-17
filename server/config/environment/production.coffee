"use strict"

# Production specific configuration
# =================================

module.exports =

  # Server IP
  ip: process.env.OPENSHIFT_NODEJS_IP or process.env.IP or `undefined`

  # Server port
  port: process.env.OPENSHIFT_NODEJS_PORT or process.env.PORT or 8080

  # MongoDB connection options
  mongo:
    uri: process.env.MONGOLAB_URI or process.env.MONGOHQ_URL or "mongodb://localhost/feedbag"
    name: "feedbag"
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
    clientID: process.env.GITHUB_CLIENT_ID
    clientSecret: process.env.GITHUB_CLIENT_SECRET
    callbackURL: (process.env.DOMAIN || '') + '/auth/github/callback'
    # clientID: "4bd3b3bac51982f5b341"
    # clientSecret: "5de2fb4e779d3951257ed0faeb19f9e1c901abdf"
    # callbackURL: "http://localhost:5000/auth/github/callback"
