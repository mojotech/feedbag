'use strict';

var gulp = require('gulp');

var paths = gulp.paths;
var ports = gulp.ports;

var $ = require('gulp-load-plugins')();

gulp.task('watch', ['inject'], function () {
  gulp.watch([
    paths.src + '/*.html',
    paths.src + '/assets/styles/**/*.styl',
    paths.src + '/app/**/*.js',
    paths.src + '/app/**/*.coffee',
    'bower.json',
    '../templates/**/*.tmpl'
  ], ['inject']);
});

gulp.task('open', function() {
  var openOpts = {
    url: 'http://localhost:' + ports.proxy
  };
  gulp.src('./src/index.html')
  .pipe($.open('', openOpts));
});



