package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	domains := make([]string, 0)
	domains = append(domains, "ya.ru", "stacktitan.com")

	var msg dns.Msg
	mp := make(map[string][]dns.RR)

	for _, d := range domains {
		mp[d] = make([]dns.RR, 0)
		fqdn := dns.Fqdn(d)
		msg.SetQuestion(fqdn, dns.TypeA)
		exchange, err := dns.Exchange(&msg, "8.8.8.8:53")
		if err != nil {
			panic(err)
		}
		if len(exchange.Answer) < 1 {
			fmt.Println("No Records")
			continue
		}

		for _, answer := range exchange.Answer {
			if a, ok := answer.(*dns.A); ok {
				fmt.Println(a.A)
			}
			mp[d] = append(mp[d], answer)
		}
	}
	fmt.Println(mp)

}
