package publicip

import "net"

var (
	networks []*net.IPNet

	cidrs = [...]string{
		// https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.0.0.0/29",
		"192.0.0.8/32",
		"192.0.0.9/32",
		"192.0.0.10/32",
		"192.0.0.170/32",
		"192.0.0.171/32",
		"192.0.2.0/24",
		"192.31.196.0/24",
		"192.52.193.0/24",
		"192.88.99.0/24",
		"192.168.0.0/16",
		"192.175.48.0/24",
		"198.18.0.0/15",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"240.0.0.0/4",
		"255.255.255.255/32",

		// https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry.xhtml
		"::1/128",
		"::/128",
		"::ffff:0:0/96",
		"64:ff9b::/96",
		"64:ff9b:1::/48",
		"100::/64",
		"2001::/23",
		"2001::/32",
		"2001:1::1/128",
		"2001:1::2/128",
		"2001:2::/48",
		"2001:3::/32",
		"2001:4:112::/48",
		"2001:10::/28",
		"2001:20::/28",
		"2001:db8::/32",
		"2002::/16",
		"2620:4f:8000::/48",
		"fc00::/7",
		"fe80::/10",
	}
)

func Is(ip net.IP) bool {
	if ip == nil {
		return false
	}
	for _, network := range networks {
		if network.Contains(ip) {
			return false
		}
	}
	return true
}

func init() {
	networks = make([]*net.IPNet, len(cidrs))
	for index, cidr := range cidrs {
		_, networks[index], _ = net.ParseCIDR(cidr)
	}
}
