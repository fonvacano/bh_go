package main

import "github.com/miekg/dns"

func main() {
	domains := make([]string, 1)
	domains = append(domains, "stacktitan.com", "ya.ru")

	var msg dns.Msg

	for _, d := range domains {
		fqdn := dns.Fqdn(d)
		msg.SetQuestion(fqdn, dns.TypeA)
		dns.Exchange(&msg, "8.8.8.8:53")
	}

}
