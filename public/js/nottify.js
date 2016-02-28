
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
        var playlist = undefined, state = 0, $ = j,
            player = undefined, songlist = [], playposition = 0, repeat = 0, shuffle = 0,
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
            }, nowplaying = function(uuid) {
                // highlight the item in the list
                $('#main-body .song').removeClass('playing');
                $('#song-'+uuid).addClass('playing');

                // update the text box
            };

        return {
            play: function(uuid) {
                // construct the audio element itself
                if (typeof player != "object") {
                    player = new Audio();
                } else {
                    // clearing out an existing song
                    this.stop();
                    togglePlay('pause');
                }

                // set up the player
                player.setAttribute("src", "/api/song/" + uuid);
                player.load();
                player.play();

                // set the currently playing track
                nowplaying(uuid);
                // add the current song to the queue
                songlist.push(uuid);
                // bump the position within the queue
                playposition++;

                // update the play button
                togglePlay('play');

                // first time through, allow
                $('#pause-button').removeClass('disabled');
                $('#stop-button').removeClass('disabled');

                // check for the other buttons
                if (songlist.length > 1) {
                    $('#previous-button').removeClass('disabled');
                }
            },

            previous: function() {
                //
                // take frm the current place in the list
                //  then move backwards if there's another one
                //  or simply begin the song again
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

                    togglePlay('pause');
                }
            },

            next: function() {
                // similar to the previous
                // but with the exception of moving to the next song in the
                //  list
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

            artist: function(artist) {
                var template = 'basic-artist-view', url = '/api/artist/'+artist;
                load(template, url, $('#main-body'));
            },

            shuffle: function() {
                var a = $('#shuffle-button'), b = $('#shuffle-button :first-child');
                if (shuffle === 0) {
                    shuffle = 1;
                    a.attr('title', 'Shuffle!');
                    b.removeClass('glyphicon-arrow-right').addClass('glyphicon-random');
                } else {
                    shuffle = 0;
                    a.attr('title', 'No Shuffle');
                    b.removeClass('glyphicon-random').addClass('glyphicon-arrow-right');
                }
            },

            repeat: function() {
                var a = $('#repeat-button'), b = $('#repeat-button :first-child');
                if (repeat === 0) {
                    // change from through to repeat single
                    repeat = 1;
                    a.attr('title', 'Repeat Song');
                    b.removeClass('glyphicon-play-circle').addClass('glyphicon-repeat');
                } else if (repeat === 1) {
                    // change from repeat single to repeat all
                    repeat = 2;
                    a.attr('title', 'Repeat List');
                    b.removeClass('glyphicon-repeat').addClass('glyphicon-refresh');
                } else {
                    // change back to through
                    repeat = 0;
                    a.attr('title', 'No Repeat');
                    b.removeClass('glyphicon-refresh').addClass('glyphicon-play-circle');
                }
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

                $(document).on({
                    ajaxStart: function() { $('body').addClass('loading'); },
                    ajaxStop: function() { $('body').removeClass('loading'); }
                });
            }
        }
    })(jQuery);

    n.init(n);
    window.onunload = function() { return 'Unsaved modifications will be lost.'; };
    window.onbeforeunload = window.onunload;
});

