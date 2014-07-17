'use strict';

var grunt = require('grunt');

/*
  ======== A Handy Little Nodeunit Reference ========
  https://github.com/caolan/nodeunit

  Test methods:
    test.expect(numAssertions)
    test.done()
  Test assertions:
    test.ok(value, [message])
    test.equal(actual, expected, [message])
    test.notEqual(actual, expected, [message])
    test.deepEqual(actual, expected, [message])
    test.notDeepEqual(actual, expected, [message])
    test.strictEqual(actual, expected, [message])
    test.notStrictEqual(actual, expected, [message])
    test.throws(block, [error], [message])
    test.doesNotThrow(block, [error], [message])
    test.ifError(value)
*/

exports.concat_css = {
  setUp: function(done) {
    // setup here if necessary
    done();
  },

  default_options: function(test) {
    var actual = grunt.file.read('tmp/default_options.css').split('\n');
    var expected = grunt.file.read('test/expected/default_options.css').split('\n');
    test.deepEqual(actual, expected, 'The two files should be concatanated together.');
    test.done();
  },

  rebase_urls: function(test) {
    var actual = grunt.file.read('tmp/rebase_urls.css').split('\n');
    var expected = grunt.file.read('test/expected/rebase_urls.css').split('\n');
    test.deepEqual(actual, expected, 'All URLs should have been rebased');
    test.done();
  },

  basedir: function(test) {
    var actual = grunt.file.read('tmp/basedir_options.css').split('\n');
    var expected = grunt.file.read('test/expected/basedir_options.css').split('\n');
    test.deepEqual(actual, expected, 'All URLs should have been rebased');
    test.done();
  },

  no_basedir_options: function(test) {
    var actual = grunt.file.read('tmp/no_basedir_options.css').split('\n');
    var expected = grunt.file.read('test/expected/no_basedir_options.css').split('\n');
    test.deepEqual(actual, expected, 'All URLs should have been rebased');
    test.done();
  }
};
