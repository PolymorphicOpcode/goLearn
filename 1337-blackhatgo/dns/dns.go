package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("tryhackme.com")
	msg.SetQuestion(fqdn, dns.TypeMX)
	in, err := dns.Exchange(&msg, "1.1.1.1:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if mx, ok := answer.(*dns.MX); ok {
			fmt.Println(mx.Mx)
		}
	}
}
