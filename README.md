Install
------

```
$ go get github.com/y-matsuwitter/trycatch
```

Example
-----

```
package main

import "github.com/y-matsuwitter/trycatch"

func main() {
	trycatch.TryCatch{}.Try(func() {
		println("do something buggy")
		b := 0
		panic(1 / b)
	}).Catch(func(err error) {
		println("catch error")
		println(err.Error())
	})
}

```
