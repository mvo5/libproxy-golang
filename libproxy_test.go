package libproxy_test

import (
	"testing"

	"github.com/mvo5/libproxy-golang"
)

func TestLibproxy(t *testing.T) {
	pf := libproxy.NewProxyFactory()
	proxies := pf.Get("http://example.com")
	if len(proxies) != 1 {
		t.Fatalf("unexpected result %v", proxies)
	}
}
