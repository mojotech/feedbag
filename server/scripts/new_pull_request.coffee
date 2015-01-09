###
  Model Name: New Pull Request
  Description:  Provides an example of how activity scripts are built and functions
                as the default activity processor for new pull requests
###
'use strict'
_ = require 'lodash'


###

  Define script as an AMD modele that is passed `user`, `events`, and `callback`

  @params
    options [Object] =
      user [Object] - The document for the user who's events were just processed.
      events [Array] - An array of the newest events added for this user.
      Event [Model] - The mongoose model used to querying all events in the db.
    callback [Function] - Must be called to indecate that the script has finished processing.


    Callback:
      Callback requires an optional error message and an optional array param containing the activity item to
      save to the db.

      activity [Object] =
        name [String] - The name of the activity module creating the event.
        template [String] - Template to display for this activity. ["blank", "banner", ... more to come?]
        description [String] - 20-60 character short event description
        icon [String] [optional] - any fontawesome icon to be used with the template
        date [Date] [optional] - date to display for activity (does not affect the order that the activity items are rendered)
        body [String] - Text or HTML to render within the {{body}} tag of the template
        css [Object] - CSS to be applied to the body and description attributes within the template

    Querying:
      ie. options.Event.find user: options.user._id, (err, events) ->
        //events is an array of all users events in the db

###
module.exports = (options, callback) ->

  # Body template to render into the body element in the banner template
  body =
    """
      <div>Created Pull-Request <a href="<%- prUrl %>"><%- pr %></a></div>
    """

  # Event description template
  description =
    """
      <a href="<%- profileUrl %>"><%- name %></a> created pull-request <a href="<%- prUrl %>"><%- prNumber %></a>
    """

  # Map events into an array of activity events
  activityEvents = _.map options.events, (event) ->

    if event.type is "PullRequestEvent" and event.payload.action is "opened"

      # Define the variables that will apear in the template
      templateVariables =
        name: options.user.github.raw.name
        pr: event.payload.pull_request?.title ? ""
        prNumber: event.payload.pull_request?.number ? ""
        prUrl: event.payload.pull_request?.html_url ? ""
        profileUrl: options.user.github.raw.html_url

      # Return the activity event hash
      name: "new_pull_request"
      template: "banner"
      description: _.template(description, templateVariables)
      icon: "fa-thumbs-up"
      date: event.created_at
      body: _.template(body, templateVariables)

  # Call the callback to indecate the activity script is done processing.
  callback(null, activityEvents)

