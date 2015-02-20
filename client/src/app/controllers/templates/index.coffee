templates = {}

emojify.setConfig(img_dir: "assets/images/emoji")
Handlebars.registerHelper('marked', (text)-> marked(text))
Handlebars.registerHelper('emojify', (text)-> emojify.replace(text))

class Templates
  fetch: (cb) ->
    $.getJSON "templates", (res) ->
      res.forEach (template) ->
        templates[template.id] = template
        templates[template.id].template = Handlebars.compile(template.template)
      cb?()

  get: (id) ->
    return templates unless id
    return templates[id]

module.exports = new Templates
