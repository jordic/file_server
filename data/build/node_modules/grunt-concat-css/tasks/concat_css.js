/*
 * grunt-concat-css
 * https://github.com/urturn/grunt-concat-css
 *
 * Copyright (c) 2013 Olivier Amblet
 * Licensed under the MIT license.
 */

'use strict';

var path = require('path');

module.exports = function(grunt) {

  // Please see the Grunt documentation for more information regarding task
  // creation: http://gruntjs.com/creating-tasks

  grunt.registerMultiTask('concat_css', 'Your task description goes here.', function() {
    // Merge task-specific and/or target-specific options with these defaults.
    var options = this.options({
      assetBaseUrl: false,
      debugMode: false
    });

    // Iterate over all specified file groups.
    this.files.forEach(function(f) {
      var importStatements = [];
      var sources = [];

      var extractImportStatements = function (data){
        var rex = /@import\s.+\;/gim;
        var replaceRex = /[\s]*@import\s.+\;/gim;
        var matches = data.css.match(rex);
        if(matches && matches.length > 0){
          grunt.log.write("Found " + matches.join(", "));
          data.imports = matches;
          data.css = data.css.replace(replaceRex, '');
        }
      };

      var rebaseUrls = function (data) {
        function dirname(csspath){
          var splits = csspath.split('/');
          splits.pop();
          return splits.join('/');
        }

        // Rebase any url('someUrl') variation
        function dataTransformUrlFunc(basedir) {
          var bd = basedir;
          return function(_, b) {
            return "url('"+normalize([bd, b].join('/'))+"')";
          };
        }

        // Rebase @import 'someUrl' exception
        function dataTransformImportAlternateFunc(basedir) {
          var bd = basedir;
          return function(_, b) {
            return "@import url('"+normalize([bd, b].join('/'))+"')";
          };
        }

        /**
         * remove upFolder(..) part of an URL
         */
        function normalize(url) {
          var computedParts = [];
          var parts = url.split('/');
          for (var i in parts){
            if (parts[i] === '..') {
              computedParts.pop();
            } else {
              computedParts.push(parts[i]);
            }
          }
          return computedParts.join('/');
        }

        function computeBaseUrl() {
          var re = new RegExp('(?:^'+ options.baseDir +'\/?)?(.*)\/[^\/]+$');
          var re_no_base_dir = /(.*)\/[^\/]+$/;
          var relativePath = options.baseDir ? data.path.replace(re, '$1') : data.path.replace(re_no_base_dir, '$1');
          return options.assetBaseUrl.replace(/\/$/, '') + (relativePath ? '/' + relativePath : '');
        }

        var baseUrl = computeBaseUrl();

        data.css = data.css.replace(/url\(['\"]?([^'\"\:]+)['\"]?\)/gm, dataTransformUrlFunc(baseUrl));
        data.css = data.css.replace(/@import\s+['\"]([^'\"\:]+)['\"]/gm, dataTransformImportAlternateFunc(baseUrl));
      };

      var imports = "";
      var cssFragments = [];

      options.debugMode && console.log(f.src, f.dest);
      // Concat specified files.
      var results = f.src.map(function(filepath) {
        // Read file source.
        var data = {
          path: filepath,
          css: grunt.file.read(filepath)
        };
        options.debugMode && console.log(data);
        if ([false, undefined].indexOf(options.assetBaseUrl) === -1) {
          rebaseUrls(data);
        }
        extractImportStatements(data);
        if (data.imports) {
          imports += data.imports.join("\n") + "\n";
        }
        cssFragments.push(data.css.replace(/(^\s+|\s+$)/g,''));
        return data;
      });

      // Write the destination file.
      grunt.file.write(f.dest, imports + cssFragments.join('\n') + '\n');

      // Print a success message.
      grunt.log.writeln('File "' + f.dest + '" created.');
    });
  });

};
