package main

type Mux struct {
	add     chan net.Conn
	remove  chan net.Addr
	sendMsg chan string
}

func (m *Mux) Add(conn net.Conn) {
	m.add <- conn
}

func (m *Mux) Remove(addr net.Addr) {
	m.remove <- add
}
func (m *Mux) SendMsg(msg string) error {
	m.sendMsg <- msg
	return nil
}

func (m *Mux) loop() {
	conns := make(map[net.Addr]net.Conn)
	for {
		select {
		case conn := <-m.add:
			m.conns[conn.RemoteAddr()] = conn
		case addr := <-m.remove:
			delete(m.conns, addr)
		case msg := <-m.sendMsg:
			for _, conn := range m.conns {
				io.WriteString(conn, msg)
			}
		}
	}
}
