package xip

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type IpSearch struct {
	prefStart [256]uint32
	prefEnd   [256]uint32
	endArr    []uint32
	addrArr   []string
}

type (
	// IPLocation ip location
	IPLocation struct {
		IP        string `json:"ip"`
		Continent string `json:"continent"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		City      string `json:"city"`
		County    string `json:"county"`
		ISP       string `json:"isp"`
		En        string `json:"en"`
		Zip       string `json:"zip"`
		ShortCode string `json:"shortcode"`
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}
)

var instance *IpSearch
var once sync.Once

func Get(ip string) (ipLocation IPLocation) {
	return GetInstance().Get(ip)
}
func GetInstance() *IpSearch {
	return GetInstanceForPath("")
}
func GetInstanceForPath(path string) *IpSearch {
	once.Do(func() {
		instance = &IpSearch{}
		var err error
		if path == "" {
			path = "./data/ultimate.dat"
		}
		instance, err = LoadDat(path)
		if err != nil {
			log.Fatal("the IP Dat loaded failed!", err)
		}
	})
	return instance
}

func LoadDat(file string) (*IpSearch, error) {
	p := IpSearch{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	for k := 0; k < 256; k++ {
		i := k*8 + 4
		p.prefStart[k] = ReadLittleEndian32(data[i], data[i+1], data[i+2], data[i+3])
		p.prefEnd[k] = ReadLittleEndian32(data[i+4], data[i+5], data[i+6], data[i+7])
	}

	RecordSize := int(ReadLittleEndian32(data[0], data[1], data[2], data[3]))

	p.endArr = make([]uint32, RecordSize)
	p.addrArr = make([]string, RecordSize)
	for i := 0; i < RecordSize; i++ {
		j := 2052 + (i * 8)
		endipnum := ReadLittleEndian32(data[j], data[1+j], data[2+j], data[3+j])
		offset := ReadLittleEndian24(data[4+j], data[5+j], data[6+j])
		length := uint32(data[7+j])
		p.endArr[i] = endipnum
		p.addrArr[i] = string(data[offset:int(offset+length)])
	}
	return &p, err

}

func (p *IpSearch) Get(ip string) (ipLocation IPLocation) {
	ips := strings.Split(ip, ".")
	x, _ := strconv.Atoi(ips[0])
	prefix := uint32(x)
	intIP := ip2Long(ip)

	low := p.prefStart[prefix]
	high := p.prefEnd[prefix]

	var cur uint32
	if low == high {
		cur = low
	} else {
		cur = p.binarySearch(low, high, intIP)
	}
	ils := strings.Split(p.addrArr[cur], "|")
	ipLocation = IPLocation{}
	ipLocation.Continent = ils[0]
	ipLocation.Country = ils[1]
	ipLocation.Province = ils[2]
	ipLocation.City = ils[3]
	ipLocation.County = ils[4]
	ipLocation.IP = ip
	ipLocation.ISP = ils[5]
	ipLocation.Zip = ils[6]
	ipLocation.En = ils[7]
	ipLocation.ShortCode = ils[8]
	ipLocation.Longitude = ils[9]
	ipLocation.Latitude = ils[10]
	return
}

func (p *IpSearch) binarySearch(low uint32, high uint32, k uint32) uint32 {
	var M uint32 = 0
	for low <= high {
		mid := (low + high) / 2
		endipNum := p.endArr[mid]
		if endipNum >= k {
			M = mid
			if mid == 0 {
				break
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return M
}

func ip2Long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func ReadLittleEndian32(a, b, c, d byte) uint32 {
	return (uint32(a) & 0xFF) | ((uint32(b) << 8) & 0xFF00) | ((uint32(c) << 16) & 0xFF0000) | ((uint32(d) << 24) & 0xFF000000)
}

func ReadLittleEndian24(a, b, c byte) uint32 {
	return (uint32(a) & 0xFF) | ((uint32(b) << 8) & 0xFF00) | ((uint32(c) << 16) & 0xFF0000)
}
