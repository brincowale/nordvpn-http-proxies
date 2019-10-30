package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"nordvpn-proxies/output"
	"nordvpn-proxies/proxies"
	"time"
)

func main() {
	request := gorequest.New().Timeout(30*time.Second).Retry(3, 5*time.Second, http.StatusInternalServerError)
	_, body, _ := request.Get("https://nordvpn.com/wp-admin/admin-ajax.php?action=servers_recommendations&limit=999999").
		EndBytes()
	err := json.Unmarshal(body, &proxies.ProxyList)
	if err != nil {
		fmt.Println(err)
	}
	for _, proxy := range proxies.ProxyList {
		if proxies.IsValidProxy(proxy) {
			output.Write(proxy)
		}
	}
	output.Save()
}


