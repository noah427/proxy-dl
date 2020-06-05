package checkp

import (
	// "../log"

	"fmt"

	"github.com/imroc/req"
)

func CheckProxy(proxy string) bool {
	req.SetProxyUrl(fmt.Sprintf("http://%s", proxy))

	resp, err := req.Get("https://google.com")

	if err != nil {
		return false
	}

	if resp.Response().StatusCode != 200 {
		return false
	}

	return true
}
