func (fd *netFD) Read(p []byte) (n int, err error) {
	// preamble
	for {
		n, err = syscall.Read(fd.sysfd, p) // HL
		if err != nil {
			n = 0
			if err == syscall.EAGAIN {
				if err = fd.pd.WaitRead(); err == nil { // HL
					continue
				}
			}
		}
		err = fd.eofError(n, err)
		break
	}
	// epilog
	return
}
