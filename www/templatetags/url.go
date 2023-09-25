package templatetags

import (
	"sync"

	"github.com/flosch/pongo2/v6"
)

var (
	urls = make(map[string]string)
	mu   = sync.RWMutex{}
)

func urlFilter(in, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	return pongo2.AsValue(GetURL(in.String())), nil
}

func AddURL(name, path string) {
	mu.Lock()
	urls[name] = path
	mu.Unlock()
}

func GetURL(name string) string {
	mu.RLock()
	path := urls[name]
	mu.RUnlock()
	return path
}

func init() {
	err := pongo2.RegisterFilter("url", urlFilter)
	if err != nil {
		panic(err)
	}
}
