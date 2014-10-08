
"use strict";

if ("undefined" == typeof jQuery)
    throw new Error('jQuery is required');

Handlebars.registerHelper("debug", function(optionalValue) {
    console.log("Current Context");
    console.log("====================");
    console.log(this);
});

$(document).ready(function() {
    var n = (function(j) {
        var nowplaying = undefined, playlist = undefined, state = 0, $ = j, player = undefined,
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
            }, load = function(template, url, container) {
                var htmlContent, dataContent, path = '/public/handlebars/' + template + '.hbs';

                $.ajax(path,{async:false}).done(function(r) { htmlContent = r; });
                $.ajax(url,{async:false}).done(function(r) { dataContent = r; });

                var template = Handlebars.compile(htmlContent);
                $(container).html(template(dataContent));
            }, togglePlay = function(which) {
                var b = $('#pause-button :first-child');
                if (which === "play") {
                    b.removeClass('glyphicon-play').addClass('glyphicon-pause');
                } else {
                    b.removeClass('glyphicon-pause').addClass('glyphicon-play');
                }
            };

        return {
            play: function(uuid) {
                if (typeof player != "object") {
                    player = new Audio();
                } else {
                    this.stop();
                    togglePlay('pause');
                }

                player.setAttribute("src", "/api/song/" + uuid);
                player.load();
                player.play();

                nowplaying = uuid;
                togglePlay('play');
            },

            previous: function() {
                console.log('previous');
            },

            pause: function() {
                if (typeof player == "object") {
                    if (player.paused) {
                        player.play();
                        togglePlay('play');
                    } else {
                        player.pause();
                        togglePlay('pause');
                    }
                }
            },

            stop: function() {
                if (typeof player == "object") {
                    player.pause();
                    player.currentTime = 0;
                }
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
                var url = '/api/songs', template = 'basic-song-list';

                load(template, url, $('#main-body'));
            },

            scan: function(self, url) {
                url = parse(url);
                if (typeof self[url.action] === 'function') {
                    self[url.action](url.option);
                } else {
                    console.log(self, url, 'Invalid ['+url.action+']');
                }
            },

            unload: function(self, e) {
                console.log(self, e);

                return false;
            },

            init: function(self) {
                $(window).bind('popstate', self, function(e) {
                    var self = e.data, url = e.currentTarget.location;
                    self.scan(self, url);
                });
                self.scan(self, window.location.href);

                $(window).unload(self, function(e) {
                    var self = e.data, url = e.currentTarget.location;
                    self.unload(self, e);
                });
            }
        }
    })(jQuery);

    n.init(n);
    window.onunload = function() { return 'Unsaved modifications will be lost.'; };
    window.onbeforeunload = window.onunload;
});

