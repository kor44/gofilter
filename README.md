# Gofilter

Gofilter is a [Go](http://golang.org/) implementation of the
[Wireshark](https://wireshark.org/) display filter for type Message (map[string]interface{}).

Yacc and Ragel are used for filter parser.

### Installation

    go get github.com/kor44/gofilter
	
## Filter syntax

### Comparison operators
```
eq, ==    Equal
ne, !=    Not Equal
gt, >     Greater Than
lt, <     Less Than
ge, >=    Greater than or Equal to
le, <=    Less than or Equal to
```
	
### Search and match operators
```
contains  Does the protocol, field or slice contain a value
matches   Does the protocol or text string match the given Perl
          regular expression
```			  
### Logical expressions
```
and, &&   Logical AND
or,  ||   Logical OR
not, !    Logical NOT
```

### Unsupported features

Functions to convert strings:
```
upper(string-field) - converts a string field to uppercase
lower(string-field) - converts a string field to lowercase
```
The slice operator:
```
eth.src[0:3] == 00:00:83
```
Bit field operations:
```
bitwise_and, & - Bitwise AND
```	
	
### Fields type
```go
FT_BOOL   ftenum = iota // bool
FT_STRING ftenum = iota // string
FT_INT    ftenum = iota // int
FT_UINT   ftenum = iota // uint

FT_UINT8  ftenum = iota // uint8
FT_UINT16 ftenum = iota // uint16
FT_UINT24 ftenum = iota // uint24
FT_UINT32 ftenum = iota // uint32
FT_UINT64 ftenum = iota // uint64
FT_INT8   ftenum = iota // int8
FT_INT16  ftenum = iota // int16
FT_INT24  ftenum = iota // int24
FT_INT32  ftenum = iota // int32
FT_INT64  ftenum = iota // int64

FT_FLOAT32 ftenum = iota // float32
FT_FLOAT64 ftenum = iota // float64

FT_BYTES ftenum = iota // []byte

FT_IP  ftenum = iota // net.IP
FT_MAC ftenum = iota // net.HardwareAddr
```	

## Usage example
```go
package main

import (
	"fmt"
	"github.com/le0developer/gofilter"
	"net"
)

func main() {
	ctx := gofilter.CreateContext()
	ctx.RegisterField("ip.src", gofilter.FT_IP)
	ctx.RegisterField("ip.dst", gofilter.FT_IP)

	f, err := ctx.NewFilter("ip.src == 192.168.0.0/24 and ip.dst == 192.168.0.1")
	if err != nil {
		fmt.Printf("Filter parse error: %s", err)
	}

	msg := gofilter.Message{
		"ip.src": net.ParseIP("192.168.0.100"),
		"ip.dst": net.ParseIP("192.168.0.1"),
	}

	if f.Apply(msg) {
		fmt.Println("Message pass")
	} else {
		fmt.Println("Message not pass")
	}
}
```

## Building

1. Install ragel somehow.
2. Install goyacc with `go install golang.org/x/tools/cmd/goyacc`
3. `go generate`
