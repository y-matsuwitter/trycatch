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

type MyError struct {
	error
}

func main() {
	trycatch.TryCatch{}.Try(func() {
		println("do something buggy")
		panic(MyError{})
	}).Catch(MyError{}, func(err error) {
		println("catch MyError")
	}).CatchAll(func(err error) {
		println("catch error")
	}).Finally(func() {
		println("finally do something")
	})
	println("done")
}

```
