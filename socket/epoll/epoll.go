package epoll

import (
	"hermes/socket/httpclient"
	"hermes/socket/utils"
	"log"
	"net"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

// https://jvns.ca/blog/2017/06/03/async-io-on-linux--select--poll--and-epoll/

type SocketEpoll struct {
	fd             int              // file descriptor of Epoll
	fdToConnection map[int]net.Conn // mapping from socket file descriptor to socket connection
	fdToUser       map[int]*httpclient.User
	clanToFds      map[int][]int // mapping from clan ID to list of file descriptor
	lock           *sync.Mutex
}

func CreateEpoll() (*SocketEpoll, error) {
	epollFD, err := unix.EpollCreate1(0)
	log.Println("epoll FD: ", epollFD)
	if err != nil {
		return nil, err
	}
	return &SocketEpoll{
		fd:             epollFD,
		fdToConnection: make(map[int]net.Conn),
		fdToUser:       make(map[int]*httpclient.User),
		clanToFds:      make(map[int][]int),
		lock:           &sync.Mutex{},
	}, nil
}

func (s *SocketEpoll) AddSocket(conn net.Conn, user *httpclient.User) error {
	log.Println("add user: ", user)
	s.lock.Lock()
	defer s.lock.Unlock()
	socketFd := utils.GetFdFromConnection(conn)
	epollFd := s.fd                        // file descriptor of the epoll
	operationCode := syscall.EPOLL_CTL_ADD // operation of adding a new fd to epoll to watch
	fd := socketFd                         // the file descriptor to be added
	event := &unix.EpollEvent{
		Events: unix.EPOLLIN | unix.EPOLLHUP, // the file descriptor is available to read or write
		Fd:     int32(fd),
	}
	// https://man7.org/linux/man-pages/man2/epoll_ctl.2.html
	// add a new file descriptor to epoll to watch them
	if err := unix.EpollCtl(epollFd, operationCode, fd, event); err != nil {
		return err
	}
	for _, clan := range user.Clans {
		fds := s.clanToFds[clan.ID]
		fds = append(fds, socketFd)
		s.clanToFds[clan.ID] = fds
		log.Println("add new conn: ", fd, s.clanToFds[clan.ID])
	}
	log.Println("fd: ", s.clanToFds)
	s.fdToConnection[socketFd] = conn
	s.fdToUser[fd] = user
	return nil
}

func (s *SocketEpoll) RemoveSocket(conn net.Conn) error {
	defer conn.Close()
	s.lock.Lock()
	defer s.lock.Unlock()
	socketFd := utils.GetFdFromConnection(conn)
	epollFd := s.fd                        // file descriptor of the epoll
	operationCode := syscall.EPOLL_CTL_DEL // operation of remove a fd out of epoll to unwatch them
	fd := socketFd
	log.Println("delete fd: ", fd)
	if err := unix.EpollCtl(epollFd, operationCode, fd, nil); err != nil {
		return err
	}
	user := s.fdToUser[fd]
	for clan := range user.Clans {
		s.clanToFds[clan] = utils.RemoveFromSorted(s.clanToFds[clan], fd)
	}
	log.Println("fd: ", s.clanToFds)
	delete(s.fdToConnection, fd)
	delete(s.fdToUser, fd)
	return nil
}

func (s *SocketEpoll) Wait() ([]net.Conn, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	eventBucket := make([]unix.EpollEvent, 100)
	epollFd := s.fd
	timeout := 100
	n, err := unix.EpollWait(epollFd, eventBucket, timeout)
	if err != nil {
		return nil, err
	}
	var readyConns []net.Conn
	for i := 0; i < n; i++ {
		readyFd := eventBucket[i].Fd // file descriptor that ready to read or write
		socketConn := s.fdToConnection[int(readyFd)]
		log.Println("0. --------- new event from epoll --------")
		log.Println("1. event code: ", eventBucket[i].Events, "file descriptor id: ", readyFd)
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
