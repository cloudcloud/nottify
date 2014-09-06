#nottify
Nottify is a self-managed replacement for Spotify. For those with many media files, and multiple devices, each without mass storage capabilities, this provides a way to stream on demand.

##requirements
```
revel/cmd
revel/revel
go-sql-driver/mysql
```

##usage
Initial usage requires processing of base content, after setting up configuration requirements.
A local mysql database is required, with the DSN being set in the local configuration. This configuration is used all through the application.

To complete the ingestion of files, ready to be used for the web interface, the ``ingest`` command should be run.
On the initial run, this command will also set up the necessary tables as is required.

###configuration
* ``base_path`` The initial path for media files.
* ``pin_code`` The login PIN for actual usage.

