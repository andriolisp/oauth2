package redis

import "net"

func (c *baseClient) Pool() pool {
	return c.connPool
}

func (cn *conn) SetNetConn(netcn net.Conn) {
	cn.netcn = netcn
}

func HashSlot(key string) int {
	return hashSlot(key)
}
