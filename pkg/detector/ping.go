package detector

import (
	"errors"
	"fmt"
	"github.com/sptuan/stargazer/internal/entity"
	"github.com/sptuan/stargazer/internal/model"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"strings"
	"time"
)

// PingDetector ping target with ICMP ping
//
// Note: ICMP ping requires root privilege!
type PingDetector struct {
	BaseDetector
}

func NewPingDetector(task *entity.Task) (*PingDetector, error) {
	tp := model.ProbeType(task.Type)
	if tp != model.PING {
		return nil, errors.New(fmt.Sprintf("invalid probe type: %v", tp))
	}
	return &PingDetector{
		BaseDetector: BaseDetector{
			Name:       task.Name,
			Type:       model.ProbeType(task.Type),
			Target:     task.Target,
			Timeout:    time.Duration(task.Timeout) * time.Millisecond,
			Interval:   time.Duration(task.Interval) * time.Millisecond,
			UpdateTime: time.Now(),
			CreateTime: time.Now(),
		},
	}, nil
}

// Detect
//
// WARNING: Ping ICMP requires root privilege!
func (d *PingDetector) Detect() (entity.DataLog, error) {
	d.UpdateTime = time.Now()
	msg := newMessage()
	pinger := fastping.NewPinger()
	//_, err := pinger.Network("udp")
	//if err != nil {
	//	msg.Level = string(model.ERROR)
	//	msg.Message = fmt.Sprintf("make request error: %v", err)
	//	return msg, err
	//}

	netProto := "ip4:icmp"
	if strings.Index(d.Target, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, d.Target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type response struct {
		addr *net.IPAddr
		rtt  time.Duration
	}
	results := make(map[string]*response)
	results[ra.String()] = nil
	pinger.AddIPAddr(ra)

	onRecv, onIdle := make(chan *response), make(chan bool)
	pinger.OnRecv = func(addr *net.IPAddr, t time.Duration) {
		onRecv <- &response{addr: addr, rtt: t}
	}
	pinger.OnIdle = func() {
		onIdle <- true
	}
	pinger.MaxRTT = time.Second
	pinger.RunLoop()
	defer func() {
		pinger.Stop()
	}()
loop:
	for {
		select {
		case res := <-onRecv:
			if _, ok := results[res.addr.String()]; ok {
				results[res.addr.String()] = res
			}
			break loop
		case <-onIdle:
			for host, r := range results {
				if r == nil {
					msg.Level = string(model.ERROR)
					msg.Message = fmt.Sprintf("ping: %s : unreachable after %v", host, pinger.MaxRTT)
					return msg, errors.New(msg.Message)
				}
			}
			break loop
		case <-time.After(d.Timeout):
			msg.Level = string(model.ERROR)
			msg.Message = fmt.Sprintf("ping: %s : timeout, reach after timeout: %s. Target not reachable. Or are you root user? (ICMP on linux accquires sudo!)", d.Target, d.Timeout)
			return msg, errors.New(msg.Message)
		}
	}

	for host, r := range results {
		if r == nil {
			msg.Level = string(model.ERROR)
			msg.Message = fmt.Sprintf("ping: %s : unreachable after %v", d.Target, pinger.MaxRTT)
		} else if r.rtt > d.Timeout {
			msg.Level = string(model.ERROR)
			msg.Message = fmt.Sprintf("ping: %s : %v, threshold: %v ", d.Target, r.rtt, d.Timeout)
		} else {
			msg.Level = string(model.INFO)
			msg.Message = fmt.Sprintf("ping: %s : %v", host, r.rtt)
		}
	}

	//msg.Message = fmt.Sprintf("ping success, target: %v, rtt: %v", d.Target, end.Sub(start))
	return msg, nil
}

// NOTE: github.com/go-ping/ping has issue on my mac, so I can't test it
//func (d *PingDetector) Detect() (entity.DataLog, error) {
//	// ping target
//	msg := newMessage()
//	pinger, err := ping.NewPinger(d.Target)
//	if err != nil {
//		msg.Level = string(model.ERROR)
//		msg.Message = fmt.Sprintf("make pinger error: %v", err)
//		return msg, err
//	}
//	pinger.Count = 1
//	err = pinger.Run()
//	if err != nil {
//		msg.Level = string(model.ERROR)
//		msg.Message = fmt.Sprintf("ping request error: %v", err)
//		return msg, err
//	}
//
//	s := pinger.Statistics()
//	if s.PacketLoss != 0 {
//		msg.Level = string(model.ERROR)
//		msg.Message = fmt.Sprintf("ping packet loss: %v", s.PacketLoss)
//		return msg, err
//	}
//	if s.MaxRtt > d.Timeout {
//		msg.Level = string(model.ERROR)
//		msg.Message = fmt.Sprintf("ping timeout: %v", err)
//		return msg, err
//	}
//
//	msg.Message = fmt.Sprintf("ping success, target: %v, rtt: %v", d.Target, s.MaxRtt)
//	return msg, nil
//}
