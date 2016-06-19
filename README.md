#nottify
Audio streaming and cataloguing application built in Go.

[![GoDoc](https://godoc.org/github.com/cloudcloud/nottify?status.svg)](https://godoc.org/github.com/cloudcloud/nottify)
[![Circle CI](https://circleci.com/gh/cloudcloud/nottify.svg?style=svg)](https://circleci.com/gh/cloudcloud/nottify)
[![codecov](https://codecov.io/gh/cloudcloud/nottify/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudcloud/nottify)

Nottify is a self-managed music streaming service. For those with many media files, and multiple devices, each without
mass storage capabilities, this provides a way to stream on demand. A command line interface to easily configure and
script the installation opens the way to use the clean and simple web interface.

##requirements
Using ``go get`` all requirements are automatically loaded. Vendoring is used to keep the explicit copy of a
particular dependency available, and managed using the [govendor](https://github.com/kardianos/govendor) tool.

##usage
Initial usage requires processing of base content, after setting up configuration requirements. This
initial process will also run-through base configuration set, and instantiation of the database. Of
course, all configuration can be modified at a later date.

###configuration
Configuration is stored within a **yaml** file. This can be edited by hand, or the CLI can be used
instead. Any changes made whilst Nottify is running will require the process to be restarted.

###tests
Tests are provided for individual files and methods, along with generic usability and stability tests.
These can be used within a CI pipeline for end-to-end testing requirements.

##commands
A series of commands is implemented in the CLI to provide easy usage and modification for Nottify.

###init
**init** will run through a series of steps to help configure an initial installation. This includes
the setup of any database, location of audio files, and various web or Nottify settings.

###config
**config** provides methods to read and write to the configuration file.

###ingest
Nottify relies heavily on working with the filesystem to find and read audio files. To help with
scripting and debugging, this command is provided to asyncronously walk through the provided directories
and find, process, and store information about files.

###search
Useful for scripting, search within the meta database within just the CLI environment. Whilst search
is also provided through the web, the CLI also provides the capability.

###clear
When the file system changes heavily, it may be useful to clear existing data and allow for ``ingest``
to perform fresh analysis. This is also useful for emptying cached content, when manual changes are
being made or configurations updated.

###start
Open the http process, and begin serving http requests.
