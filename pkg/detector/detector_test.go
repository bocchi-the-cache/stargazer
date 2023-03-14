package detector

import (
	"encoding/json"
	"fmt"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestHttpDetector_Detect(t *testing.T) {
	d, err := NewHttpDetector(
		&entity.Task{
			Name:        "test",
			Description: "test",
			Type:        string(model.HTTPS),
			Target:      "www.baidu.com",
			Timeout:     10000,
			Interval:    1000,
			SslExpire:   true,
			SslVerify:   true,
		})
	if err != nil {
		t.Error(err)
	}
	msg, err := d.Detect()
	if err != nil {
		t.Error(err)
	}
	t.Log(msg)

}

func TestHttpDetector_DetectTimeout(t *testing.T) {
	d, err := NewHttpDetector(
		&entity.Task{
			Name:        "test",
			Description: "test",
			Type:        string(model.HTTPS),
			Target:      "www.baidu.com",
			Timeout:     1,
			Interval:    1000,
			SslExpire:   true,
			SslVerify:   true,
		})
	if err != nil {
		t.Error(err)
	}
	msg, err := d.Detect()
	if err != nil {
		t.Error(err)
	}
	t.Log(msg)

}

func TestPing(t *testing.T) {
	type response struct {
		addr *net.IPAddr
		rtt  time.Duration
	}
	hostname := "www.baidu.com"
	p := fastping.NewPinger()
	//p.Network("icmp")

	netProto := "ip4:icmp"
	if strings.Index(hostname, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, hostname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results := make(map[string]*response)
	results[ra.String()] = nil
	p.AddIPAddr(ra)

	onRecv, onIdle := make(chan *response), make(chan bool)
	p.OnRecv = func(addr *net.IPAddr, t time.Duration) {
		onRecv <- &response{addr: addr, rtt: t}
	}
	p.OnIdle = func() {
		onIdle <- true
	}

	p.MaxRTT = time.Second
	p.RunLoop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

loop:
	for {
		select {
		case <-c:
			fmt.Println("get interrupted")
			break loop
		case res := <-onRecv:
			if _, ok := results[res.addr.String()]; ok {
				results[res.addr.String()] = res
			}
		case <-onIdle:
			for host, r := range results {
				if r == nil {
					fmt.Printf("%s : unreachable %v\n", host, time.Now())
				} else {
					fmt.Printf("%s : %v %v\n", host, r.rtt, time.Now())
				}
				results[host] = nil
			}
		case <-p.Done():
			if err = p.Err(); err != nil {
				fmt.Println("Ping failed:", err)
			}
			break loop
		}
	}
	signal.Stop(c)
	p.Stop()
}

func TestPingDetector_Detect(t *testing.T) {
	d, err := NewPingDetector(
		&entity.Task{
			Name:        "test",
			Description: "test",
			Type:        string(model.PING),
			Target:      "www.baidu.com",
			Timeout:     1000,
			Interval:    1000,
		})
	if err != nil {
		t.Error(err)
	}
	msg, err := d.Detect()
	if err != nil {
		t.Error(err)
	}
	msgJson, _ := json.MarshalIndent(msg, "", "  ")

	t.Log(string(msgJson))
}
