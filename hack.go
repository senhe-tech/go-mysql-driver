package mysql

import "net"

func (mc *mysqlConn) Conn() net.Conn {
	return mc.netConn
}
