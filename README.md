#nottify
Personal "cloud"-based music streaming service. Taking the ease of HTML5, with the ease of a PHP web-server, provide a simple mechanism to stream music as if it were from a local player. Assuming a local collection of music files, __nottify__ can index and provide over streaming the content to any browser or device with HTML5 capabilities.

##configuration
* __deploy.location__: Location at which deployment will take place. The apache document root would then be __deploy.location__/web
* __db.path__: A sub-path for the sqlite database to be stored within the __deploy.location__.
* __debug__: Enabling extra debug output.
* __cookie.secret__: Key for encryption of all cookie content.

##requirements
```
apache-2.2
get-id3
mustache
php-5.4
php_json
slim
sqlite
```
Composer will provide the acquisition of all PHP-based requirements. The assumption of PHP and Apache2 being readily available is the basis for all builds and deployments. There shouldn't be issues with other setups, though this isn't readily available.

##build
Ant is used for the build system. There are general targets to complete test suite usage, along with file placement and composer usage. Specific file directories are assumed, but easily modified by editing the file directly. Common targets are listed below for ease of reference.

```
ant clean
    Remove the results directory from tests.

ant simple
    Copy all appropriately used files into the deployment location. This will not interact with composer at all, and is simply a way to provide file-only incremental updates.

ant update
    Copy latest file updates, and then use composer to upgrade all existing libraries.

ant deploy
    Remove the current deployment folder, and complete a full build.

ant apache2
    Move the server-based files into general known locations and restart apache2.
    This assumes a basic general setup, as what is provided with Ubuntu. This can be configured easily for different systems.

ant ingest
    Complete the database ingest from the local file system of music files. This is a simple way to complete the update without requiring full command memorising.
```

##license
Copyright 2014 Allan Shone and other contributors.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in the
    documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

