/*
 angular-tablesort v1.0.4
 (c) 2013 Mattias Holmlund, http://mattiash.github.io/angular-tablesort
 License: MIT
*/

var tableSortModule = angular.module( 'tableSort', [] );

tableSortModule.directive('tsWrapper', function( $log, $parse ) {
    'use strict';
    return {
        scope: true,
        controller: function($scope) {
            $scope.sortExpression = [];
            $scope.headings = [];

            var parse_sortexpr = function( expr ) {
                return [$parse( expr ), null, false];
            };

            this.setSortField = function( sortexpr, element ) {
                var i;
                var expr = parse_sortexpr( sortexpr );
                if( $scope.sortExpression.length === 1
                    && $scope.sortExpression[0][0] === expr[0] ) {
                    if( $scope.sortExpression[0][2] ) {
                        element.removeClass( "tablesort-desc" );
                        element.addClass( "tablesort-asc" );
                        $scope.sortExpression[0][2] = false;
                    }
                    else {
                        element.removeClass( "tablesort-asc" );
                        element.addClass( "tablesort-desc" );
                        $scope.sortExpression[0][2] = true;
                    }
                }
                else {
                    for( i=0; i<$scope.headings.length; i=i+1 ) {
                        $scope.headings[i]
                            .removeClass( "tablesort-desc" )
                            .removeClass( "tablesort-asc" );
                    }
                    element.addClass( "tablesort-asc" );
                    $scope.sortExpression = [expr];
                }
            };

            this.addSortField = function( sortexpr, element ) {
                var i;
                var toggle_order = false;
                var expr = parse_sortexpr( sortexpr );
                for( i=0; i<$scope.sortExpression.length; i=i+1 ) {
                    if( $scope.sortExpression[i][0] === expr[0] ) {
                        if( $scope.sortExpression[i][2] ) {
                            element.removeClass( "tablesort-desc" );
                            element.addClass( "tablesort-asc" );
                            $scope.sortExpression[i][2] = false;
                        }
                        else {
                            element.removeClass( "tablesort-asc" );
                            element.addClass( "tablesort-desc" );
                            $scope.sortExpression[i][2] = true;
                        }
                        toggle_order = true;
                    }
                }
                if( !toggle_order ) {
                    element.addClass( "tablesort-asc" );
                    $scope.sortExpression.push( expr );
                }
            };

            this.registerHeading = function( headingelement ) {
                $scope.headings.push( headingelement );
            };

            $scope.sortFun = function( a, b ) {
                var i, aval, bval, descending, filterFun;
                for( i=0; i<$scope.sortExpression.length; i=i+1 ){
                    aval = $scope.sortExpression[i][0](a);
                    bval = $scope.sortExpression[i][0](b);
                    filterFun = b[$scope.sortExpression[i][1]];
                    if( filterFun ) {
                        aval = filterFun( aval );
                        bval = filterFun( bval );
                    }
                    if( aval === undefined ) {
                        aval = "";
                    }
                    if( bval === undefined ) {
                       bval = "";
                    }
                    descending = $scope.sortExpression[i][2];
                    if( aval > bval ) {
                        return descending ? -1 : 1;
                    }
                    else if( aval < bval ) {
                        return descending ? 1 : -1;
                    }
                }
                return 0;
            };
        }
    };
} );

tableSortModule.directive('tsCriteria', function() {
    return {
        require: "^tsWrapper",
        link: function(scope, element, attrs, tsWrapperCtrl) {
            var clickingCallback = function(event) {
                scope.$apply( function() {
                    if( event.shiftKey ) {
                        tsWrapperCtrl.addSortField(attrs.tsCriteria, element);
                    }
                    else {
                        tsWrapperCtrl.setSortField(attrs.tsCriteria, element);
                    }
                } );
            };
            element.bind('click', clickingCallback);
            element.addClass('tablesort-sortable');
            if( "tsDefault" in attrs && attrs.tsDefault !== "0" ) {
                tsWrapperCtrl.addSortField( attrs.tsCriteria, element );
                if( attrs.tsDefault == "descending" ) {
                    tsWrapperCtrl.addSortField( attrs.tsCriteria, element );
                }
            }
            tsWrapperCtrl.registerHeading( element );
        }
    };
});

tableSortModule.directive("tsRepeat", function($compile) {
    return {
        terminal: true,
        require: "^tsWrapper",
        priority: 1000000,
        link: function(scope, element) {
            var clone = element.clone();
            var tdcount = element[0].childElementCount;
            var repeatExpr = clone.attr("ng-repeat");
            repeatExpr = repeatExpr.replace(/^\s*([\s\S]+?)\s+in\s+([\s\S]+?)(\s+track\s+by\s+[\s\S]+?)?\s*$/,
                "$1 in $2 | tablesortOrderBy:sortFun$3");

            element.html("<td colspan='"+tdcount+"'></td>");
            element[0].className += " showIfLast";
            clone.removeAttr("ts-repeat");

            clone.attr("ng-repeat", repeatExpr);
            var clonedElement = $compile(clone)(scope);
            element.after(clonedElement);
        }
    };
} );

tableSortModule.filter( 'tablesortOrderBy', function(){
    return function(array, sortfun ) {
        if(!array) return;
        var arrayCopy = [];
        for ( var i = 0; i < array.length; i++) { arrayCopy.push(array[i]); }
        return arrayCopy.sort( sortfun );
    };
} );

tableSortModule.filter( 'parseInt', function(){
    return function(input) {
        return parseInt( input );
    };
} );

tableSortModule.filter( 'parseFloat', function(){
    return function(input) {
        return parseFloat( input );
    };
} );


'use strict';

/**
 * Binds a CodeMirror widget to a <textarea> element.
 */
angular.module('ui.codemirror', [])
  .constant('uiCodemirrorConfig', {})
  .directive('uiCodemirror', ['uiCodemirrorConfig', function (uiCodemirrorConfig) {

    return {
      restrict: 'EA',
      require: '?ngModel',
      priority: 1,
      compile: function compile() {

        // Require CodeMirror
        if (angular.isUndefined(window.CodeMirror)) {
          throw new Error('ui-codemirror need CodeMirror to work... (o rly?)');
        }

        return function postLink(scope, iElement, iAttrs, ngModel) {


          var options, opts, codeMirror, value;

          value = iElement.text();

          if (iElement[0].tagName === 'TEXTAREA') {
            // Might bug but still ...
            codeMirror = window.CodeMirror.fromTextArea(iElement[0], {
              value: value
            });
          } else {
            iElement.html('');
            codeMirror = new window.CodeMirror(function(cm_el) {
              iElement.append(cm_el);
            }, {
              value: value
            });
          }

          options = uiCodemirrorConfig.codemirror || {};
          opts = angular.extend({}, options, scope.$eval(iAttrs.uiCodemirror), scope.$eval(iAttrs.uiCodemirrorOpts));

          function updateOptions(newValues) {
            for (var key in newValues) {
              if (newValues.hasOwnProperty(key)) {
                codeMirror.setOption(key, newValues[key]);
              }
            }
          }

          updateOptions(opts);

          if (iAttrs.uiCodemirror) {
            scope.$watch(iAttrs.uiCodemirror, updateOptions, true);
          }


          if (ngModel) {
            // CodeMirror expects a string, so make sure it gets one.
            // This does not change the model.
            ngModel.$formatters.push(function (value) {
              if (angular.isUndefined(value) || value === null) {
                return '';
              } else if (angular.isObject(value) || angular.isArray(value)) {
                throw new Error('ui-codemirror cannot use an object or an array as a model');
              }
              return value;
            });


            // Override the ngModelController $render method, which is what gets called when the model is updated.
            // This takes care of the synchronizing the codeMirror element with the underlying model, in the case that it is changed by something else.
            ngModel.$render = function () {
              //Code mirror expects a string so make sure it gets one
              //Although the formatter have already done this, it can be possible that another formatter returns undefined (for example the required directive)
              var safeViewValue = ngModel.$viewValue || '';
              codeMirror.setValue(safeViewValue);
            };


            // Keep the ngModel in sync with changes from CodeMirror
            codeMirror.on('change', function (instance) {
              var newValue = instance.getValue();
              if (newValue !== ngModel.$viewValue) {
                // Changes to the model from a callback need to be wrapped in $apply or angular will not notice them
                scope.$apply(function() {
                  ngModel.$setViewValue(newValue);
                });
              }
            });
          }


          // Watch ui-refresh and refresh the directive
          if (iAttrs.uiRefresh) {
            scope.$watch(iAttrs.uiRefresh, function (newVal, oldVal) {
              // Skip the initial watch firing
              if (newVal !== oldVal) {
                codeMirror.refresh();
              }
            });
          }


          // Allow access to the CodeMirror instance through a broadcasted event
          // eg: $broadcast('CodeMirror', function(cm){...});
          scope.$on('CodeMirror', function(event, callback) {
            if (angular.isFunction(callback)) {
              callback(codeMirror);
            } else {
              throw new Error('the CodeMirror event requires a callback function');
            }
          });


          // onLoad callback
          if (angular.isFunction(opts.onLoad)) {
            opts.onLoad(codeMirror);
          }

        };
      }
    };
  }]);


/*!
 * ui-select
 * http://github.com/angular-ui/ui-select
 * Version: 0.3.1 - 2014-07-12T16:26:10.166Z
 * License: MIT
 */
!function(){"use strict";void 0===angular.element.prototype.querySelectorAll&&(angular.element.prototype.querySelectorAll=function(e){return angular.element(this[0].querySelectorAll(e))}),angular.module("ui.select",[]).constant("uiSelectConfig",{theme:"bootstrap",placeholder:"",refreshDelay:1e3}).service("uiSelectMinErr",function(){var e=angular.$$minErr("ui.select");return function(){var t=e.apply(this,arguments),c=t.message.replace(new RegExp("\nhttp://errors.angularjs.org/.*"),"");return new Error(c)}}).service("RepeatParser",["uiSelectMinErr",function(e){var t=this;t.parse=function(t){if(!t)throw e("repeat","Expected 'repeat' expression.");var c=t.match(/^\s*([\s\S]+?)\s+in\s+([\s\S]+?)(?:\s+track\s+by\s+([\s\S]+?))?\s*$/);if(!c)throw e("iexp","Expected expression in form of '_item_ in _collection_[ track by _id_]' but got '{0}'.",t);var s=c[1],l=c[2],i=c[3];if(c=s.match(/^(?:([\$\w]+)|\(([\$\w]+)\s*,\s*([\$\w]+)\))$/),!c)throw e("iidexp","'_item_' in '_item_ in _collection_' should be an identifier or '(_key_, _value_)' expression, but got '{0}'.",s);return{lhs:s,rhs:l,trackByExp:i}},t.getGroupNgRepeatExpression=function(){return"($group, $items) in $select.groups"},t.getNgRepeatExpression=function(e,t,c,s){var l=e+" in "+(s?"$items":t);return c&&(l+=" track by "+c),l}}]).controller("uiSelectCtrl",["$scope","$element","$timeout","RepeatParser","uiSelectMinErr",function(e,t,c,s,l){function i(){o.resetSearchInput&&(o.search=a,o.selected&&o.items.length&&(o.activeIndex=o.items.indexOf(o.selected)))}function n(e){var t=!0;switch(e){case h.Down:o.activeIndex<o.items.length-1&&o.activeIndex++;break;case h.Up:o.activeIndex>0&&o.activeIndex--;break;case h.Tab:case h.Enter:o.select(o.items[o.activeIndex]);break;case h.Escape:o.close();break;default:t=!1}return t}function r(){var e=t.querySelectorAll(".ui-select-choices-content"),c=e.querySelectorAll(".ui-select-choices-row");if(c.length<1)throw l("choices","Expected multiple .ui-select-choices-row but got '{0}'.",c.length);var s=c[o.activeIndex],i=s.offsetTop+s.clientHeight-e[0].scrollTop,n=e[0].offsetHeight;i>n?e[0].scrollTop+=i-n:i<s.clientHeight&&(o.isGrouped&&0===o.activeIndex?e[0].scrollTop=0:e[0].scrollTop-=s.clientHeight-i)}var o=this,a="";o.placeholder=void 0,o.search=a,o.activeIndex=0,o.items=[],o.selected=void 0,o.open=!1,o.focus=!1,o.focusser=void 0,o.disabled=void 0,o.resetSearchInput=void 0,o.refreshDelay=void 0;var u=t.querySelectorAll("input.ui-select-search");if(1!==u.length)throw l("searchInput","Expected 1 input.ui-select-search but got '{0}'.",u.length);o.activate=function(e){o.disabled||(i(),o.open=!0,c(function(){o.search=e||o.search,u[0].focus()}))},o.parseRepeatAttr=function(t,c){function i(t){o.groups={},angular.forEach(t,function(t){var s=e.$eval(c),l=angular.isFunction(s)?s(t):t[s];o.groups[l]?o.groups[l].push(t):o.groups[l]=[t]}),o.items=[],angular.forEach(Object.keys(o.groups).sort(),function(e){o.items=o.items.concat(o.groups[e])})}function n(e){o.items=e}var r=s.parse(t),a=c?i:n;o.isGrouped=!!c,o.itemProperty=r.lhs,e.$watchCollection(r.rhs,function(e){if(void 0===e||null===e)o.items=[];else{if(!angular.isArray(e))throw l("items","Expected an array but got '{0}'.",e);a(e)}})};var d;o.refresh=function(t){void 0!==t&&(d&&c.cancel(d),d=c(function(){e.$eval(t)},o.refreshDelay))},o.setActiveItem=function(e){o.activeIndex=o.items.indexOf(e)},o.isActive=function(e){return o.items.indexOf(e[o.itemProperty])===o.activeIndex},o.select=function(e){o.selected=e,o.close()},o.close=function(){o.open&&(i(),o.open=!1,o.focusser[0].focus())};var h={Enter:13,Tab:9,Up:38,Down:40,Escape:27};u.on("keydown",function(t){if(o.items&&o.items.length>=0){var c=t.which;switch(e.$apply(function(){var e=n(c);e&&(t.preventDefault(),t.stopPropagation())}),c){case h.Down:case h.Up:r()}}}),e.$on("$destroy",function(){u.off("keydown")})}]).directive("uiSelect",["$document","uiSelectConfig","uiSelectMinErr","$compile",function(e,t,c,s){return{restrict:"EA",templateUrl:function(e,c){var s=c.theme||t.theme;return s+"/select.tpl.html"},replace:!0,transclude:!0,require:["uiSelect","ngModel"],scope:!0,controller:"uiSelectCtrl",controllerAs:"$select",link:function(t,l,i,n,r){function o(e){var c=!1;c=window.jQuery?window.jQuery.contains(l[0],e.target):l[0].contains(e.target),c||(a.close(),t.$digest())}var a=n[0],u=n[1],d=angular.element("<input ng-disabled='$select.disabled' class='ui-select-focusser ui-select-offscreen' type='text' aria-haspopup='true' role='button' />");s(d)(t),a.focusser=d,l.append(d),d.bind("focus",function(){t.$evalAsync(function(){a.focus=!0})}),d.bind("blur",function(){t.$evalAsync(function(){a.focus=!1})}),d.bind("keydown",function(e){e.which===h.TAB||h.isControl(e)||h.isFunctionKey(e)||e.which===h.ESC||((e.which==h.DOWN||e.which==h.UP||e.which==h.ENTER||e.which==h.SPACE)&&(e.preventDefault(),e.stopPropagation(),a.activate()),t.$digest())}),d.bind("keyup input",function(e){e.which===h.TAB||h.isControl(e)||h.isFunctionKey(e)||e.which===h.ESC||e.which==h.ENTER||(a.activate(d.val()),d.val(""),t.$digest())});var h={TAB:9,ENTER:13,ESC:27,SPACE:32,LEFT:37,UP:38,RIGHT:39,DOWN:40,SHIFT:16,CTRL:17,ALT:18,PAGE_UP:33,PAGE_DOWN:34,HOME:36,END:35,BACKSPACE:8,DELETE:46,isArrow:function(e){switch(e=e.which?e.which:e){case h.LEFT:case h.RIGHT:case h.UP:case h.DOWN:return!0}return!1},isControl:function(e){var t=e.which;switch(t){case h.SHIFT:case h.CTRL:case h.ALT:return!0}return e.metaKey?!0:!1},isFunctionKey:function(e){return e=e.which?e.which:e,e>=112&&123>=e}};i.$observe("disabled",function(){a.disabled=void 0!==i.disabled?i.disabled:!1}),i.$observe("resetSearchInput",function(){var e=t.$eval(i.resetSearchInput);a.resetSearchInput=void 0!==e?e:!0}),t.$watch("$select.selected",function(e){u.$viewValue!==e&&u.$setViewValue(e)}),u.$render=function(){a.selected=u.$viewValue},e.on("click",o),t.$on("$destroy",function(){e.off("click",o)}),r(t,function(e){var t=angular.element("<div>").append(e),s=t.querySelectorAll(".ui-select-match");if(1!==s.length)throw c("transcluded","Expected 1 .ui-select-match but got '{0}'.",s.length);l.querySelectorAll(".ui-select-match").replaceWith(s);var i=t.querySelectorAll(".ui-select-choices");if(1!==i.length)throw c("transcluded","Expected 1 .ui-select-choices but got '{0}'.",i.length);l.querySelectorAll(".ui-select-choices").replaceWith(i)})}}}]).directive("uiSelectChoices",["uiSelectConfig","RepeatParser","uiSelectMinErr","$compile",function(e,t,c,s){return{restrict:"EA",require:"^uiSelect",replace:!0,transclude:!0,templateUrl:function(t){var c=t.parent().attr("theme")||e.theme;return c+"/choices.tpl.html"},compile:function(l,i){var n=t.parse(i.repeat),r=i.groupBy;return function(l,i,o,a,u){if(r){var d=i.querySelectorAll(".ui-select-choices-group");if(1!==d.length)throw c("rows","Expected 1 .ui-select-choices-group but got '{0}'.",d.length);d.attr("ng-repeat",t.getGroupNgRepeatExpression())}var h=i.querySelectorAll(".ui-select-choices-row");if(1!==h.length)throw c("rows","Expected 1 .ui-select-choices-row but got '{0}'.",h.length);h.attr("ng-repeat",t.getNgRepeatExpression(n.lhs,"$select.items",n.trackByExp,r)).attr("ng-mouseenter","$select.setActiveItem("+n.lhs+")").attr("ng-click","$select.select("+n.lhs+")"),u(function(e){var t=i.querySelectorAll(".ui-select-choices-row-inner");if(1!==t.length)throw c("rows","Expected 1 .ui-select-choices-row-inner but got '{0}'.",t.length);t.append(e),s(i)(l)}),a.parseRepeatAttr(o.repeat,r),l.$watch("$select.search",function(){a.activeIndex=0,a.refresh(o.refresh)}),o.$observe("refreshDelay",function(){var t=l.$eval(o.refreshDelay);a.refreshDelay=void 0!==t?t:e.refreshDelay})}}}}]).directive("uiSelectMatch",["uiSelectConfig",function(e){return{restrict:"EA",require:"^uiSelect",replace:!0,transclude:!0,templateUrl:function(t){var c=t.parent().attr("theme")||e.theme;return c+"/match.tpl.html"},link:function(t,c,s,l){s.$observe("placeholder",function(t){l.placeholder=void 0!==t?t:e.placeholder})}}}]).filter("highlight",function(){function e(e){return e.replace(/([.?*+^$[\]\\(){}|-])/g,"\\$1")}return function(t,c){return c&&t?t.replace(new RegExp(e(c),"gi"),'<span class="ui-select-highlight">$&</span>'):t}})}(),angular.module("ui.select").run(["$templateCache",function(e){e.put("bootstrap/choices.tpl.html",'<ul class="ui-select-choices ui-select-choices-content dropdown-menu" role="menu" aria-labelledby="dLabel" ng-show="$select.items.length > 0"><li class="ui-select-choices-group"><div class="divider" ng-show="$index > 0"></div><div ng-show="$select.isGrouped" class="ui-select-choices-group-label dropdown-header">{{$group}}</div><div class="ui-select-choices-row" ng-class="{active: $select.isActive(this)}"><a href="javascript:void(0)" class="ui-select-choices-row-inner"></a></div></li></ul>'),e.put("bootstrap/match.tpl.html",'<button type="button" class="btn btn-sm form-control ui-select-match" tabindex="-1" ng-hide="$select.open" ng-disabled="$select.disabled" ng-class="{\'btn-sm-focus\':$select.focus}" ;="" ng-click="$select.activate()"><span ng-hide="$select.selected !== undefined" class="text-muted">{{$select.placeholder}}</span> <span ng-show="$select.selected !== undefined" ng-transclude=""></span> <span class="pt glyphicon glyphicon-search"></span></button>'),e.put("bootstrap/select.tpl.html",'<div class="ui-select-bootstrap dropdown" ng-class="{open: $select.open}"><div class="ui-select-match"></div><input type="text" autocomplete="off" tabindex="-1" class="form-control ui-select-search" placeholder="{{$select.placeholder}}" ng-model="$select.search" ng-show="$select.open"><div class="ui-select-choices"></div></div>'),e.put("select2/choices.tpl.html",'<ul class="ui-select-choices ui-select-choices-content select2-results"><li class="ui-select-choices-group" ng-class="{\'select2-result-with-children\': $select.isGrouped}"><div ng-show="$select.isGrouped" class="ui-select-choices-group-label select2-result-label">{{$group}}</div><ul class="select2-result-sub"><li class="ui-select-choices-row" ng-class="{\'select2-highlighted\': $select.isActive(this)}"><div class="select2-result-label ui-select-choices-row-inner"></div></li></ul></li></ul>'),e.put("select2/match.tpl.html",'<a class="select2-choice ui-select-match" ng-class="{\'select2-default\': $select.selected === undefined}" ng-click="$select.activate()"><span ng-hide="$select.selected !== undefined" class="select2-chosen">{{$select.placeholder}}</span> <span ng-show="$select.selected !== undefined" class="select2-chosen" ng-transclude=""></span> <span class="select2-arrow"><b></b></span></a>'),e.put("select2/select.tpl.html",'<div class="select2 select2-container" ng-class="{\'select2-container-active select2-dropdown-open\': $select.open,\n                \'select2-container-disabled\': $select.disabled,\n                \'select2-container-active\': $select.focus }"><div class="ui-select-match"></div><div class="select2-drop select2-with-searchbox select2-drop-active" ng-class="{\'select2-display-none\': !$select.open}"><div class="select2-search"><input type="text" autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" class="ui-select-search select2-input" ng-model="$select.search"></div><div class="ui-select-choices"></div></div></div>'),e.put("selectize/choices.tpl.html",'<div ng-show="$select.open" class="ui-select-choices selectize-dropdown single"><div class="ui-select-choices-content selectize-dropdown-content"><div class="ui-select-choices-group optgroup"><div ng-show="$select.isGrouped" class="ui-select-choices-group-label optgroup-header">{{$group}}</div><div class="ui-select-choices-row" ng-class="{active: $select.isActive(this)}"><div class="option ui-select-choices-row-inner" data-selectable=""></div></div></div></div></div>'),e.put("selectize/match.tpl.html",'<div ng-hide="$select.open || $select.selected === undefined" class="ui-select-match" ng-transclude=""></div>'),e.put("selectize/select.tpl.html",'<div class="selectize-control single"><div class="selectize-input" ng-class="{\'focus\': $select.open, \'disabled\': $select.disabled, \'selectize-focus\' : $select.focus}" ng-click="$select.activate()"><div class="ui-select-match"></div><input type="text" autocomplete="off" tabindex="-1" class="ui-select-search" placeholder="{{$select.placeholder}}" ng-model="$select.search" ng-hide="$select.selected && !$select.open" ng-disabled="$select.disabled"></div><div class="ui-select-choices"></div></div>')}]);


