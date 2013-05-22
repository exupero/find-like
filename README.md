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

Vim
---

When using Vim, calling the following function will open the spec file that corresponds to the current buffer.

```vim
function OpenSpecFile()
  exe "edit ".system("find-spec-file ".expand("%"))
endfunction
```

This is the mapping I use to open the spec file within a split:

```vim
nnoremap <LocalLeader>u :vsp<CR>:call OpenSpecFile()<CR>
```
