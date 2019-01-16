# go-bufpool

## Description

golang idiom for buffer([]byte/bytes.Buffer) pool.

## Example

```go
package main

import (
    "os"

    bufpool "github.com/shimt/go-bufpool"
)

func main() {
    p := bufpool.NewByteArrayPool(4096, 0)
    b := p.Get()

    b = append(b, []byte("this is example.\n")...)

    if _, e := os.Stdout.Write(b); e != nil {
        panic(e)
    }

    p.Put(b)
}
```

```go
package main

import (
    "os"

    bufpool "github.com/shimt/go-bufpool"
)

func main() {
    p := bufpool.NewBytesBufferPool(4096, 0)
    b := p.Get()

    b.WriteString("this is example.\n")

    if _, e := os.Stdout.Write(b.Bytes()); e != nil {
        panic(e)
    }

    p.Put(b)
}
```