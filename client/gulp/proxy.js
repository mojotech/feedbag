 /*jshint unused:false */

'use strict';

var gulp = require('gulp');

var paths = gulp.paths;
var ports = gulp.ports;

var http = require('http');

var httpProxy = require('http-proxy');

function proxyServerInit() {
  // Proxy for static asset requests and browser sync socket
  var proxyWeb = new httpProxy.createProxyServer({
    target: {
      host: 'localhost',
      port: ports.static
    }
  });

  // Proxy for server requests and server socket
  var proxyServer = new httpProxy.createProxyServer({
    target: {
      host: 'localhost',
      port: ports.server
    }
  });

  // Create basic http server to use as proxy
  var server = http.createServer(function (req, res) {
    var staticExtensions = /\.(html|css|js|png|jpg|jpeg|gif|ico|xml|rss|txt|eot|svg|ttf|woff|cur|woff2|map)(\?((r|v|rel|rev)=[\-\.\w]*)?)?$/.test(req.url);
    var indexPage = req.url === '/';
    var browserSyncSocket = /^\/browser-sync/.test(req.url);

    if (staticExtensions || indexPage || browserSyncSocket) {
      proxyWeb.web(req, res);
    } else {
      proxyServer.web(req, res);
    }
  });

  // Listen to the `upgrade` event and proxy the socket
  server.on('upgrade', function (req, socket, head) {
    var browserSyncSocket = /^\/browser-sync/.test(req.url);
    if (browserSyncSocket) {
      proxyWeb.ws(req, socket, head);
    } else {
      proxyServer.ws(req, socket, head);
    }
  });

  // Serve it up
  server.listen(ports.proxy);
}

gulp.task('serve', ['browserSync', 'open'], function() {
  proxyServerInit();
});
