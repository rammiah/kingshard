package server

import (
	"errors"
	"net"
	"strings"
	"sync/atomic"

	"github.com/flike/kingshard/mysql"
)

type BoolIndex struct {
	index int32
}

func (b *BoolIndex) Set(index bool) {
	if index {
		atomic.StoreInt32(&b.index, 1)
	} else {
		atomic.StoreInt32(&b.index, 0)
	}
}

func (b *BoolIndex) Get() (int32, int32, bool) {
	index := atomic.LoadInt32(&b.index)
	if index == 1 {
		return 1, 0, true
	} else {
		return 0, 1, false
	}
}

type IPInfo struct {
	info    string
	isIPNet bool
	ip      net.IP
	ipNet   net.IPNet
}

func ParseIPInfo(v string) (IPInfo, error) {
	if ip, ipNet, err := net.ParseCIDR(v); err == nil {
		return IPInfo{
			info:    v,
			isIPNet: true,
			ip:      ip,
			ipNet:   *ipNet,
		}, nil
	}

	if ip := net.ParseIP(v); ip != nil {
		return IPInfo{
			info:    v,
			isIPNet: false,
			ip:      ip,
		}, nil
	}

	return IPInfo{}, errors.New("invalid ip address")
}

func (t *IPInfo) Info() string {
	return t.info
}

func (t *IPInfo) Match(ip net.IP) bool {
	if t.isIPNet {
		return t.ipNet.Contains(ip)
	} else {
		return t.ip.Equal(ip)
	}
}

func GetSqlCommand(token string) string {
	tokenId, ok := mysql.PARSE_TOKEN_MAP[strings.ToLower(token)]
	if !ok {
		return ""
	}
	switch tokenId {
	case mysql.TK_ID_SELECT:
		return mysql.COM_SELECT_STR
	case mysql.TK_ID_DELETE:
		return mysql.COM_DELETE_STR
	case mysql.TK_ID_INSERT:
		return mysql.COM_INSERT_STR
	case mysql.TK_ID_REPLACE:
		return mysql.COM_REPLACE_STR
	case mysql.TK_ID_UPDATE:
		return mysql.COM_UPDATE_STR
	case mysql.TK_ID_SET:
		return mysql.COM_SET_STR
	case mysql.TK_ID_SHOW:
		return mysql.COM_SHOW_STR
	case mysql.TK_ID_TRUNCATE:
		return mysql.COM_TRUNCATE_STR
	default:
		return ""
	}
}
