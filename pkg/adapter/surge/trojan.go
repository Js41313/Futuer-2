package surge

import (
	"strconv"
	"strings"

	"github.com/Js41313/Futuer-2/pkg/adapter/proxy"
)

func buildTrojan(data proxy.Proxy, uuid string) string {
	trojan, ok := data.Option.(proxy.Trojan)
	if !ok {
		return ""
	}
	config := []string{
		data.Name + "=trojan",
		data.Server,
		strconv.Itoa(data.Port),
		"password=" + uuid,
		"tfo=true",
		"udp-relay=true",
	}
	if trojan.SecurityConfig.SNI != "" {
		config = append(config, "sni="+trojan.SecurityConfig.SNI)
	}
	if trojan.SecurityConfig.AllowInsecure {
		config = append(config, "skip-cert-verify=true")
	} else {
		config = append(config, "skip-cert-verify=false")
	}
	return strings.Join(config, ",") + "\r\n"
}
