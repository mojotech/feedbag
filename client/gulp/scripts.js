'use strict';

var gulp = require('gulp');

var paths = gulp.paths;

var $ = require('gulp-load-plugins')();

var browserify = require('gulp-browserify');
var coffeeify = require('coffeeify')

var browserifyTransform = ['coffeeify'];

gulp.task('scripts', function () {
  return gulp.src(paths.src + '/app/app.coffee', {read: false})
    .pipe($.coffeelint())
    .pipe($.coffeelint.reporter())
    .pipe(browserify({
      transform: browserifyTransform,
      extensions: ['.coffee'],
    }))
    .pipe($.rename('app.js'))
    .pipe(gulp.dest(paths.tmp + '/serve/'))
    .pipe($.size());
});

