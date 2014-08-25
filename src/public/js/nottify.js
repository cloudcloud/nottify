
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
});

