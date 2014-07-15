/*
 * grunt-concat-css
 * https://github.com/urturn/grunt-concat-css
 *
 * Copyright (c) 2013 Olivier Amblet
 * Licensed under the MIT license.
 */

'use strict';

module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    jshint: {
      all: [
        'Gruntfile.js',
        'tasks/*.js',
        '<%= nodeunit.tests %>'
      ],
      options: {
        jshintrc: '.jshintrc'
      }
    },

    // Before generating any new files, remove any previously-created files.
    clean: {
      tests: ['tmp']
    },

    // Configuration to be run (and then tested).
    concat_css: {
      default_options: {
        options: {
        },
        files: {
          'tmp/default_options.css': ['test/fixtures/test.css', 'test/fixtures/import.css', 'test/fixtures/component_a/css/style.css']
        }
      },
      baseDir: {
        options: {
          baseDir: 'test/fixtures'
        },
        files: {
          'tmp/basedir_options.css': ['test/fixtures/test.css', 'test/fixtures/import.css']
        }
      },
      rebase_urls: {
        options: {
          baseDir: 'test/fixtures',
          assetBaseUrl: 'static/assets/' // trailing / can be omitted
        },
        files: {
          'tmp/rebase_urls.css': ['test/fixtures/test.css', 'test/fixtures/import.css', 'test/fixtures/component_a/css/style.css']
        }
      },
      no_basedir_options: {
        options: {
          assetBaseUrl: '.' // trailing / can be omitted
        },
        files: {
          'tmp/no_basedir_options.css': ['test/fixtures/test.css', 'test/fixtures/import.css', 'test/fixtures/component_a/css/style.css']
        }
      }
    },

    // Unit tests.
    nodeunit: {
      tests: ['test/*_test.js']
    }

  });

  // Actually load this plugin's task(s).
  grunt.loadTasks('tasks');

  // These plugins provide necessary tasks.
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-nodeunit');

  // Whenever the "test" task is run, first clean the "tmp" dir, then run this
  // plugin's task(s), then test the result.
  grunt.registerTask('test', ['clean', 'concat_css', 'nodeunit']);

  // By default, lint and run all tests.
  grunt.registerTask('default', ['jshint', 'test']);

};
