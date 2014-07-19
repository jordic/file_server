module.exports = function(grunt) {
  
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-concat-css');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    
    grunt.initConfig({
    
        pkg: grunt.file.readJSON('package.json'),

        concat: {
            dist: {
                src: ['../js/codemirror.js', '../js/jquery-1.11.0.min.js', 
                    '../js/select2.min.js', '../js/bootstrap.min.js', 
                    '../js/angular.min.js', '../js/angular-sanitize.js',
                    '../js/ui-bootstrap-tpls.min.js', '../js/others.js'
                    
                    ],
                dest: 'libs.js'
            }
        },

        uglify: {
            options: {
                mangle: false
            },
            my_target: {

                files: {
                    '../libs.js': ['libs.js']
                }
            }
        }, 


        concat_css: {
            options: {
          // Task-specific options go here.
            },
            all: {
            src: ["../css/codemirror.css", "../css/monokai.css", "../css/font-awesome.css",
                "../css/bootstrap.css", 
                "../css/select2.css", "../css/ui-select.css"],
            dest: "../styles.css"
            },
        }

        });

  grunt.registerTask('build', 'concat uglify concat_css');

}