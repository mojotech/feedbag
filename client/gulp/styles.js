'use strict';

var gulp = require('gulp');

var paths = gulp.paths;

var $ = require('gulp-load-plugins')();

gulp.task('styles', function () {


  var injectFiles = gulp.src([
    paths.src + '/assets/styles/helpers/**/*.styl', //Ensure that helpers get included first
    paths.src + '/assets/styles/**/*.styl',
    '!' + paths.src + '/assets/styles/main.styl'
  ], { read: false });

  var injectOptions = {
    transform: function(filePath) {
      filePath = filePath.replace(paths.src + '/assets/styles/', '');
      return '@import \'' + filePath + '\';';
    },
    starttag: '// injector',
    endtag: '// endinjector',
    addRootSlash: false
  };

  var indexFilter = $.filter('main.styl');

  return gulp.src([
    paths.src + '/assets/styles/main.styl'
  ])
    .pipe(indexFilter)
    .pipe($.inject(injectFiles, injectOptions))
    .pipe(indexFilter.restore())
    .pipe($.stylus())

  .pipe($.autoprefixer())
    .on('error', function handleError(err) {
      console.error(err.toString());
      this.emit('end');
    })
    .pipe(gulp.dest(paths.tmp + '/serve/assets/styles'));
});
