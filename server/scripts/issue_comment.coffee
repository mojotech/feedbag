###
  Model Name: Issue Comment
  Description:  Creates an event for issue comments
###
'use strict'
_ = require 'lodash'

module.exports = (options, callback) ->

  # Body template to render into the body element in the banner template
  body =
    """
      <div class="test"><%- name %> did something awesome</div>
    """

  # Event description template
  description =
    """
      <a href="<%- profileUrl %>"><%- name %></a> commented on an issue
    """


  # Map events into an array of activity events
  activityEvents = _.map options.events, (event) ->

    if event.type is "IssueCommentEvent" and event.payload.action is "created"

      # Define the variables that will apear in the template
      templateVariables =
        name: options.user.github.raw.name
        profileUrl: options.user.github.raw.html_url

      # Return the activity event hash
      name: "issue_comment"
      template: "banner"
      description: _.template(description, templateVariables)
      icon: "fa-thumbs-up"
      date: event.created_at
      body: _.template(body, templateVariables)


  # Call the callback to indecate the activity script is done processing.
  callback(null, activityEvents)

