# q
Simple FIFO queue. With 100% test coverage.


## Usage


```go
package main

import (
    "github.com/libds/q"
    "fmt"
)

func main() {
    qu := q.New()
    qu.Enq("job1")
    qu.Enq("job2")
    fmt.Println(qu.Size()) // Output: 2
    job, err := qu.Deq()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(job) // Output: job1
}
```


## API Docs

https://godoc.org/github.com/libds/q


## License

The MIT License (MIT)

Copyright (c) 2015 LibDS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
