passport = require("passport")
GitHubStrategy = require("passport-github").Strategy

exports.setup = (User, config) ->

  passport.use new GitHubStrategy
    clientID: config.github.clientID
    clientSecret: config.github.clientSecret
    callbackURL: config.github.callbackURL

  , (accessToken, refreshToken, profile, done) ->

    User.findOne
      "github.id": profile.id
    , (err, user) ->
      unless user
        user = new User(
          name: profile.displayName
          email: profile.emails[0].value
          role: "user"
          username: profile.username
          provider: "github"
          github:
            id: profile.id
            username: profile._json.login
            raw: profile._json
            access_token: accessToken
            avatar_url: profile._json.avatar_url
        )
        user.save (err) ->
          done err  if err
          done err, user

      else
        user.github.raw = profile._json
        user.github.access_token = accessToken
        user.save (err) ->
          done err if err
          done err, user
