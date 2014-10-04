
"use strict";

if ("undefined" == typeof jQuery)
    throw new Error('jQuery is required');

$(document).ready(function() {
    $('#loginModal .login-modal button').click(function(n) {
        n.preventDefault();

        var $n = $('#loginModal #login-holder'),
            $s = parseInt(n.target.value),
            $f = $('#loginForm #form-pin-code'),
            $a = $f.val();

        if ($f.val().length === 5) {
            console.log('Done!');
            return;
        }

        $a = $a + '' + $s;
        $n.append('<kbd>'+$s+'</kbd> ');
        $f.val($a);

        if ($a.length === 5) {
            $('#loginForm').submit();
            return;
        }
    });

    var n = (function(j) {
        var np = undefined, pl = undefined, state = 0, $ = j,
            parse = function(u) {
                var p = document.createElement('a'), h = undefined, o = {};
                p.href = u;

                h = p.hash.substr(3).split('/');
                o.action = h.shift();
                o.option = h.shift();

                if (o.action.length < 1) {
                    o.action = 'home';
                }

                return o;
            };

        return {
            play: function(uuid) {
                console.log('play '+uuid, state);
            },

            previous: function() {
                console.log('previous');
            },

            pause: function() {
                console.log('pause');
            },

            stop: function() {
                console.log('stop');
            },

            next: function() {
                console.log('next');
            },

            queue: function(uuid) {
                console.log('queue '+uuid);
            },

            love: function(uuid) {
                console.log('love '+uuid);
            },

            add: function(uuid) {
                console.log('playlisting '+uuid);
            },

            edit: function(uuid) {
                console.log('edit '+uuid);
            },

            search: function(items) {
                console.log('searching '+items);
            },

            home: function() {
                //
            },

            scan: function(self, url) {
                url = parse(url);
                if (typeof self[url.action] === 'function') {
                    self[url.action](url.option);
                } else {
                    console.log(self, url, 'Invalid ['+url.action+']');
                }
            },

            init: function(self) {
                $(window).bind('popstate', self, function(e) {
                    var self = e.data, url = e.currentTarget.location;
                    self.scan(self, url);
                });
                self.scan(self, window.location.href);
            }
        }
    })(jQuery);

    n.init(n);
});

