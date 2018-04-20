package middleware

import (
	"oversea/db"
	"oversea/utils"
)

type IpBlacklist struct {

}

var blacklistKey = "blacklistKey"

func (this *IpBlacklist) SetBlacklistIP(ip ...string) {
	if len(ip) > 0 {
		for k :=0; k < len(ip); k++ {
			if utils.CheckIpIsValid(ip[k]) {
				db.NewConn().SAdd(blacklistKey, ip[k])
			}
		}
	}
}

func (this *IpBlacklist) IsIPInBacklist(ip string) bool {
	isInBlacklist :=  db.NewConn().SIsMember(blacklistKey, ip)
	return isInBlacklist.Val()
}


func (this *IpBlacklist) DelIPFromBacklist(ip string) bool {
	flag :=  db.NewConn().SRem(blacklistKey, ip)
	return flag.Val() > 0
}