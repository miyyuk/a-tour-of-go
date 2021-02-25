// https://go-tour-jp.appspot.com/methods/18

// 参考
// https://qiita.com/tenntenn/items/453a09c4c6d7f580d0ab

package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	for i, v := range ip {
		s[i] = fmt.Sprint(int(v))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// >result
// loopback: 127.0.0.1
// googleDNS: 8.8.8.8
