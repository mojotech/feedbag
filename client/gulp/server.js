'use strict';

var gulp = require('gulp');

var paths = gulp.paths;
var ports = gulp.ports;

var util = require('util');

var browserSync = require('browser-sync');

function browserSyncInit(baseDir, files, browser) {
  browser = browser === undefined ? 'default' : browser;

  var routes = null;
  if(baseDir === paths.src || (util.isArray(baseDir) && baseDir.indexOf(paths.src) !== -1)) {
    routes = {
      '/bower_components': 'bower_components'
    };
  }

  browserSync.instance = browserSync.init(files, {
    startPath: '/',
    server: {
      baseDir: baseDir,
      routes: routes
    },
    port: ports.static,
    open: false,
    browser: browser
  });
}

gulp.task('browserSync', ['watch'], function () {
  browserSyncInit([
    paths.tmp + '/serve',
    paths.src
  ], [
    paths.tmp + '/serve/assets/styles/**/*.css',
    paths.tmp + '/serve/app/**/*.js',
    paths.src + 'src/assets/images/**/*',
    paths.tmp + '/serve/*.html',
    paths.tmp + '/serve/app/**/*.html',
    paths.src + '/app/**/*.html'
  ]);
});

gulp.task('serve:dist', ['build'], function () {
  browserSyncInit(paths.dist);
});
