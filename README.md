[![Build Status][travis-image]][travis-url]
# libproxy-golang - cgo wrappera around libproxy

A cgo wrapper around libproxy to query for the proxy to use.

Example:
```golang
package main

import (
	"fmt"

	"github.com/mvo5/libproxy-golang"
)

func main() {
	pf := libproxy.NewProxyFactory()
	proxies := pf.Get("http://example.com")
	for _, proxy := range proxies {
		fmt.Println(proxy)
	}
}
```

[travis-image]: https://travis-ci.org/mvo5/libproxy-golang.svg?branch=master
[travis-url]: https://travis-ci.org/mvo5/libproxy-golang.svg?branch=master

