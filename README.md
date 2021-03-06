# Feedbag

Feedbag is a TV first, multiuser Github events dashboard for organizations and teams.

Designed to be flexable, simple, and highly configurable, new widget dashboards can be added by simply adding a template to the `templates/` directory. The config options at the top of the templates define which github event triggers that template to render on the client.

An optional `condition` field can be added to further specify when the template is rendered. ie. `"and .PushEvent (eq .Branch 'master')"` will render the template if there is a force push to "master"

### Version
0.0.1

[![Join the chat at https://gitter.im/mojotech/feedbag](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/mojotech/feedbag?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/mojotech/feedbag?status.svg)](https://godoc.org/github.com/mojotech/feedbag) [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/mojotech/feedbag) [![Build Status](https://travis-ci.org/mojotech/feedbag.svg)](https://travis-ci.org/mojotech/feedbag) [![Coverage](http://gocover.io/_badge/github.com/mojotech/feedbag)](http://gocover.io/github.com/mojotech/feedbag) [![Stories in
Ready](https://badge.waffle.io/mojotech/FeedBag.svg?label=ready&title=Ready)](http://waffle.io/mojotech/FeedBag)

[![Throughput Graph](https://graphs.waffle.io/mojotech/feedbag/throughput.svg)](https://waffle.io/mojotech/feedbag/metrics)

### Tech

Feedbag uses a number of open source projects to work:

* [Go]
* [Gin]
* [Backbone]
* [Gulp]
* [jQuery]

## Developer Guide

### Installation

Optionally, install gin for running the go server

```sh
$ go get github.com/codegangsta/gin
```

```sh
$ git clone [git-repo-url] feedbag
$ cd feedbag
$ go get ./...
```

### Startup

Run the go server

```sh
//Point to the location of the gulp index file
$ export GITHUB_KEY=[your github app key] GITHUB_SECRET=[your github app secret]

$ go build
$ ./feedbag

//Or
$ gin
```

Run the gulp task

```sh
$ cd client
$ npm i
& npm start
```

### Todo's

 - Write tests
 - Add more events and variables
 - Add more template examples
 - Add styleguide and classes
 - Add user control over repo events shown

License
----

MIT


**Free Software, BooYa!**

[Go]:http://golang.org
[Gin]:http://gin-gonic.github.io/gin/
[jQuery]:http://jquery.com
[Backbone]:http://backbonejs.org
[Gulp]:http://gulpjs.com
