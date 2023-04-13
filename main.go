package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
)

type DataCenter struct {
	LocationUid string
	Name        string
}

type PingInfo struct {
	LocationUid string
	Name        string
	SourceIpv4  string
	SourceIpv6  string
	Target      string
	Result      string
}

type NodeInfo struct {
	IPv4 string
	IPv6 string
}

var dataCenters = []DataCenter{
	{"dc116", "Amsterdam (Tier-3)"},
	{"dc124", "Atlanta (Tier-3)"},
	{"dallascolo", "DallasColo (Tier-3)"},
	{"dataline", "DataLine (Tier-3)"},
	{"progress", "Fiord"},
	{"gefaftmcom", "Frankfurt (Tier-3)"},
	{"fnhe4d", "Helsinki (Tier-3)"},
	{"dc123", "Hong Kong (Tier-3)"},
	{"iqdata", "IQ Data (Tier-3)"},
	{"dc125", "Kiev (Tier-3)"},
	{"dc119", "London (Tier-3)"},
	{"uslamcom", "Los Angeles (Tier-3)"},
	{"dc118", "Madrid (Tier-3)"},
	{"dc122", "Palermo (Tier-3)"},
	{"dc120", "Paris (Tier-3)"},
	{"dc121", "Riga (Tier-3)"},
	{"rtcomm", "Rostelecom"},
	{"adman", "Rostelecom (ex Adman)"},
	{"novosibttk", "TTK"},
	{"cato4d", "Toronto (Tier-3)"},
	{"dc117", "Warsaw (Tier-3)"},
}

var (
	pingChan = make(chan PingInfo)
	cookie   = new(string)
)

func main() {
	// Get args
	target := flag.String("target", "", "Target to ping")
	cookie = flag.String("cookie", "", "Cookie. If not set, will get cookie automatically")
	flag.Parse()

	// Check if target is set
	if *target == "" {
		flag.PrintDefaults()
		return
	}

	if *cookie == "" {
		// Get cookie automatically if not set
		*cookie = GetCookie()
	}

	var wg sync.WaitGroup
	// Ping target
	for _, dataCenter := range dataCenters {
		wg.Add(1)
		go func(dataCenter DataCenter) {

			nodeInfo := GetNodeInfo(dataCenter.LocationUid)
			pingInfo := PingInfo{
				LocationUid: dataCenter.LocationUid,
				Name:        dataCenter.Name,
				Target:      *target,
				SourceIpv4:  nodeInfo.IPv4,
				SourceIpv6:  nodeInfo.IPv6,
			}
			if result, err := JustHostPing(pingInfo); err != nil {
				pingChan <- PingInfo{
					LocationUid: pingInfo.LocationUid,
					Name:        pingInfo.Name,
					SourceIpv4:  pingInfo.SourceIpv4,
					SourceIpv6:  pingInfo.SourceIpv6,
					Target:      *target,
					Result:      err.Error(),
				}
				return
			} else {
				pingChan <- PingInfo{
					LocationUid: dataCenter.LocationUid,
					Name:        dataCenter.Name,
					SourceIpv4:  pingInfo.SourceIpv4,
					SourceIpv6:  pingInfo.SourceIpv6,
					Target:      *target,
					Result:      result,
				}
			}
			wg.Done()

		}(dataCenter)
	}

	go func() {
		wg.Wait()
		close(pingChan)
	}()

	for result := range pingChan {
		fmt.Printf("\n\n<======%s======>\nsource(%s / %s) ping target(%s)\n%s\n", result.Name, result.SourceIpv4, result.SourceIpv6, result.Target, result.Result)
	}
}

func JustHostPing(info PingInfo) (ret string, err error) {
	// Ping target

	req, err := http.NewRequest("GET", fmt.Sprintf("https://justhost.asia/en/looking-glass/index?location_uid=%s&target=%s&command=ping", info.LocationUid, info.Target), nil)
	if err != nil {
		return
	}

	// header
	AddEssentialHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ret = gjson.Get(string(body), "content").String()
	// Get Last 2 lines
	lines := strings.Split(ret, "\n")
	if len(lines) > 3 {
		ret = lines[len(lines)-3] + "\n" + lines[len(lines)-2]
	}

	return
}

func GetNodeInfo(locationUID string) NodeInfo {
	// Get node info

	req, err := http.NewRequest("GET", fmt.Sprintf("https://justhost.asia/looking-glass/info?location_uid=%s", locationUID), nil)
	if err != nil {
		return NodeInfo{
			IPv4: "Unknown",
			IPv6: "Unknown",
		}
	}

	// header
	AddEssentialHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return NodeInfo{
			IPv4: "Unknown",
			IPv6: "Unknown",
		}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NodeInfo{
			IPv4: "Unknown",
			IPv6: "Unknown",
		}
	}

	return NodeInfo{
		IPv4: gjson.Get(string(body), "info.ipv4").String(),
		IPv6: gjson.Get(string(body), "info.ipv6").String(),
	}

}

func AddEssentialHeaders(req *http.Request) {
	// Add essential headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", *cookie)
}
