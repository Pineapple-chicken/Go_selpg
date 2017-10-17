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
* ```./selpg -s 0 -e 0 -l 10 text.txt```
```shell
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
This is the line 6
This is the line 7
This is the line 8
This is the line 9
This is the line 10
```

* ```./selpg -s 1 -e 1 -l 4 < text.txt```
```shell
This is the line 5
This is the line 6
This is the line 7
This is the line 8
```

* ```./selpg -s 1 -e 1 -l 4 text.txt >output.txt```
```shell
output.txt:
This is the line 5
This is the line 6
This is the line 7
This is the line 8
```

* ```./selpg -s 0 -e 0 -f text.txt```
```shell
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
This is the line 6
This is the line 7
This is the line 8
This is the line 9
This is the line 10
This is the line 11
This is the line 12
This is the line 13
This is the line 14
This is the line 15
This is the line 16
This is the line 17
This is the line 18
This is the line 19
This is the line 20
This is the line 21
This is the line 22
This is the line 23
This is the line 24
This is the line 25
This is the line 26
This is the line 27
This is the line 28
This is the line 29
This is the line 30
```

* ```./selpg -s 1 -e 2 -l 2 -d printer test.txt```
（这里将-d变为cat -n进行测试）
```shell
This is the line 3
This is the line 4
This is the line 5
This is the line 6
```

* ```./selpg -s 1 -l 3 test.txt>output 2 > error```
（显示错误信息）
```shell
not enough arguments
2nd arg should be -eend_page
```

* ```./selpg -s 4 -e 1 -l 10 test.txt>output 2 > error```
（显示错误信息）
```shell
The range of the page is invalid
```
