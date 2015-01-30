'use strict';

var gulp = require('gulp');

var paths = gulp.paths;

gulp.task('watch', ['inject'], function () {
  gulp.watch([
    paths.src + '/*.html',
    paths.src + '/assets/styles/**/*.styl',
    paths.src + '/app/**/*.js',
    paths.src + '/app/**/*.coffee',
    'bower.json'
  ], ['inject']);
});
