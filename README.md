# desc

tree command implementation in go.

# files

~~~
$ ./tree .
.
|-- README.md
|-- cmd
|   `-- tree
|       `-- main.go
|-- go.mod
|-- makefile
|-- test
|   |-- dir1
|   |   |-- dir11
|   |   |   `-- file3
|   |   `-- dir12
|   |       |-- file4
|   |       `-- file5
|   |-- dir2
|   |   |-- file21
|   |   `-- file22
|   `-- dir3
|       |-- dir31
|       |   `-- dir32
|       |       `-- dir33
|       |-- file31
|       `-- file32
|-- tree
|-- tree.gif
|-- tree.go
`-- tree_test.go
~~~

# usage

## synopsis

~~~
$ tree [-h] [-c] [-d depth] [directory ...]
~~~

## options

### `-h`

show help message

### `-c`

enable colorized output

### `-d`

set max depth
(-1 for unlimit)

# build options

## lint

run `go vet` and `golint`

## build

create executable binary

## test

test tree_test.go

## coverage/cprofile/mprofile

output coverage and profile 

