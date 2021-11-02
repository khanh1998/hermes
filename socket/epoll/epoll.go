package epoll

import (
	"net"
	"reflect"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

// https://jvns.ca/blog/2017/06/03/async-io-on-linux--select--poll--and-epoll/

type SocketEpoll struct {
	fd             int              // file descriptor of Epoll
	fdToConnection map[int]net.Conn // mapping from socket file descriptor to socket connection
	clanToFds      map[int][]int    // mapping from clan ID to list of file descriptor
	lock           *sync.Mutex
}

func CreateEpoll() (*SocketEpoll, error) {
	epollFD, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &SocketEpoll{
		fd:             epollFD,
		fdToConnection: make(map[int]net.Conn),
		clanToFds:      make(map[int][]int),
		lock:           &sync.Mutex{},
	}, nil
}

func (s *SocketEpoll) AddSocket(conn net.Conn, clan int) error {
	socketFd := GetFdFromConnection(conn)
	epollFd := s.fd                        // file descriptor of the epoll
	operationCode := syscall.EPOLL_CTL_ADD // operation of adding a new fd to epoll to watch
	fd := socketFd                         // the file descriptor to be added
	event := &unix.EpollEvent{
		Events: unix.POLLIN | unix.POLLHUP, // the file descriptor is available to read or write
		Fd:     int32(fd),
	}
	// https://man7.org/linux/man-pages/man2/epoll_ctl.2.html
	// add a new file descriptor to epoll to watch them
	if err := unix.EpollCtl(epollFd, operationCode, fd, event); err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	fds := s.clanToFds[clan]
	fds = append(fds, socketFd)
	s.fdToConnection[socketFd] = conn
	return nil
}

// kernel maintains a file descriptor for each connection,
// this function get fd number of the input connection.
func GetFdFromConnection(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
