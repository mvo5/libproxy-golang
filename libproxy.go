package libproxy

// #include<proxy.h>
// #include<stdlib.h>
// #cgo LDFLAGS: -lproxy
import "C"

import (
	"runtime"
	"unsafe"
)

type ProxyFactory struct {
	fac *C.pxProxyFactory
}

func NewProxyFactory() *ProxyFactory {
	pf := &ProxyFactory{fac: C.px_proxy_factory_new()}
	runtime.SetFinalizer(pf, func(pf *ProxyFactory) {
		C.px_proxy_factory_free(pf.fac)
	})

	return pf
}

func (pf *ProxyFactory) Get(url string) (proxies []string) {
	p := C.px_proxy_factory_get_proxies(pf.fac, C.CString(url))
	// go provides no pointer arithmetic, so make a new huge
	// go slice of char* and iterate over that
	tmpSlice := (*[10 * 1024]*C.char)(unsafe.Pointer(p))

	for i := 0; tmpSlice[i] != nil; i++ {
		proxies = append(proxies, C.GoString(tmpSlice[i]))
		C.free(unsafe.Pointer(tmpSlice[i]))
	}
	C.free(unsafe.Pointer(p))

	return proxies
}
