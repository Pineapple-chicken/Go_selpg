# Go_selpg
Assignment for service computing : golang for selpg

## Development Environment
* Ubantu:ubuntu-17.04-desktop-amd64
* Golang:go1.9.1.linux-amd64

## Build
```shell
go get github.com/Pineapple-chicken/Go_selpg
```

## Usage
```shell
selpg -s=Number -e=Number [options] [filename]
```
* -s : Start page.
* -e : End page.
* -l : Determine the number of lines per page and default is 72.
* -f : Determine the type and the way to be seprated.
* -d : Determine the destination of output.
* [filename] : Read input from this file.
* If filename is not given, read input from stdin. and Ctrl+D to cut out.

## Example

