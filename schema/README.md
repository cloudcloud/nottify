#schema
Schema contains versioned copies of all schema modifications within the database. By default, these are
written for PostgreSQL. Each version is named to follow a sequential process, allowing for rolling
upgrades and rebuilds without having conflicts and issues with dependencies within the data.

##naming
Each file contains a versioned copy of any modifications, the version is the first part of the name.
Following on from that, is an arbitrary string that is more as a tl;dr of the content.
```
v[version_number]_[arbitrary_string].sql
```
