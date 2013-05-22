find-spec-file
==============

A simple Go program to find the filename in the `spec` directory that most closely resembles the given filename.

Building
--------

With Go installed,

```bash
$ go get
$ go build find-spec-file.go
```

Using
-----

With `find-spec-file` somewhere on your path and with the current directory as the root of a typical Rails project,

```bash
$ find-spec-file app/models/user.rb
spec/models/user_spec.rb
```
