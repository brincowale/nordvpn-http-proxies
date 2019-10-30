package proxies

import (
	"nordvpn-proxies/config"
	"strings"
)

var ProxyList []Proxy

type Proxy struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Station   string `json:"station"`
	Hostname  string `json:"hostname"`
	Load      int    `json:"load"`
	Status    string `json:"status"`
	Locations []struct {
		Country   struct {
			Name string `json:"name"`
			Code string `json:"code"`
			City struct {
				Name      string  `json:"name"`
			} `json:"city"`
		} `json:"country"`
	} `json:"locations"`
	Technologies []struct {
		Identifier string `json:"identifier"`
		Pivot      struct {
			Status string `json:"status"`
		} `json:"pivot"`
	} `json:"technologies"`
}

func IsValidProxy(proxy Proxy) bool {
	if !isLoadUnderMax(proxy){
		return false
	}
    if !isCountryValid(proxy) {
    	return false
	}
	if !isProtocolValid(proxy) {
		return false
	}
	return true

}

func isLoadUnderMax(proxy Proxy) bool {
	if proxy.Load > config.Cfg.Load {
		return false
	}
	return true
}

func isCountryValid(proxy Proxy) bool {
	if len(config.Cfg.CountryCode) > 0 {
		for _, country := range config.Cfg.CountryCode {
			if strings.EqualFold(proxy.Locations[0].Country.Code, country) {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

func isProtocolValid(proxy Proxy) bool {
	if config.Cfg.ProxyType != "" {
		for _, protocol := range proxy.Technologies {
			if strings.EqualFold(protocol.Identifier, config.Cfg.ProxyType) && strings.EqualFold(protocol.Pivot.Status, "online") {
				return true
			}
		}
	} else {
		return true
	}
	return false
}