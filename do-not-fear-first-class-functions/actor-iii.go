package main

func (m *Mux) PrivateMsg(addr net.Addr, msg string) error {
	result := make(chan net.Conn, 1)
	m.ops <- func(m map[net.Addr]net.Conn) {
		result <- m[addr]
	}

	conn := <-result
	if conn == nil {
		return errors.Errorf("client %v not registered", addr)
	}
	return io.WriteString(conn, msg)
}
