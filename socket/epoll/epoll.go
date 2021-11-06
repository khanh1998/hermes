package epoll

import (
	"hermes/socket/utils"
	"log"
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
	s.clanToFds[clan] = fds
	s.fdToConnection[socketFd] = conn
	log.Println("add new conn: ", fd, s.clanToFds[clan])
	return nil
}

func (s *SocketEpoll) RemoveSocket(conn net.Conn, clan int) error {
	// defer conn.Close()
	socketFd := GetFdFromConnection(conn)
	epollFd := s.fd                        // file descriptor of the epoll
	operationCode := syscall.EPOLL_CTL_DEL // operation of remove a fd out of epoll to unwatch them
	fd := socketFd
	log.Println("delete fd: ", fd)
	if err := unix.EpollCtl(epollFd, operationCode, fd, nil); err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.fdToConnection, fd)
	s.clanToFds[clan] = utils.RemoveFromSorted(s.clanToFds[clan], fd)
	log.Println("fd: ", s.clanToFds)
	return nil
}

func (s *SocketEpoll) Wait() ([]net.Conn, error) {
	eventBucket := make([]unix.EpollEvent, 100)
	epollFd := s.fd
	timeout := 100
	n, err := unix.EpollWait(epollFd, eventBucket, timeout)
	if err != nil {
		return nil, err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var readyConns []net.Conn
	for i := 0; i < n; i++ {
		readyFd := eventBucket[i].Fd // file descriptor that ready to read or write
		socketConn := s.fdToConnection[int(readyFd)]
		readyConns = append(readyConns, socketConn)
	}
	return readyConns, nil
}

func (s *SocketEpoll) GetFDByClan(clanId int) []int {
	return s.clanToFds[clanId]
}

func (s *SocketEpoll) GetConnectionByFD(fd int) net.Conn {
	return s.fdToConnection[fd]
}

// kernel maintains a file descriptor for each connection,
// this function get fd number of the input connection.
func GetFdFromConnection(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
