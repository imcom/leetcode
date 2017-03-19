package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

var hookableSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGTERM,
}

type gracefulServer struct {
	listener *gracefulListener
	wg       sync.WaitGroup
	sigChan  chan os.Signal
	handler  func(net.Conn)
	isChild  bool
}

func newGracefulServer(graceful bool) (*gracefulServer, error) {
	server := &gracefulServer{
		wg:      sync.WaitGroup{},
		sigChan: make(chan os.Signal),
		handler: simpleHandler,
		isChild: graceful,
	}

	if !server.isChild {
		l, err := newGracefulListener("0.0.0.0:8081")
		if err != nil {
			return nil, err
		}
		server.listener = l
		server.listener.srv = server
	} else {
		f := os.NewFile(3, "")
		l, _ := net.FileListener(f)
		server.listener = &gracefulListener{
			Listener: l,
			stopped:  false,
			srv:      server,
		}
	}

	log.Printf("server is listening at 0.0.0.0:8081\n")

	go server.restartHandler()

	if server.isChild {
		parent := syscall.Getppid()
		log.Printf("killing parent process, %d", parent)
		syscall.Kill(parent, syscall.SIGTERM)
	}

	return server, nil
}

func (srv *gracefulServer) fork() (err error) {
	path := os.Args[0]
	args := []string{
		"-graceful"}
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	files := []*os.File{srv.listener.File()}

	cmd.ExtraFiles = files

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Restart: Failed to launch, error: %v", err)
	}

	return
}

func (srv *gracefulServer) restartHandler() {
	signal.Notify(
		srv.sigChan,
		hookableSignals...,
	)

	log.Printf("registered signals handler\n")

	pid := syscall.Getpid()
	for {
		sig := <-srv.sigChan
		switch sig {
		case syscall.SIGHUP:
			log.Printf("gracefully restarting, %d", pid)
			err := srv.fork()
			if err != nil {
				log.Fatal("Fork failed:", err)
			}
		case syscall.SIGTERM:
			log.Printf("shutting down old server, %d", pid)
			err := srv.listener.Close()
			if err != nil {
				log.Fatal("shutdown failed:", err)
			}
		default:
			log.Printf("ignore %v", sig)
		}
	}
}

func (srv *gracefulServer) Serve() error {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			return err
		}

		go srv.handler(conn)
		srv.wg.Add(1)
	}
}

type gracefulConn struct {
	net.Conn
	srv *gracefulServer
}

func (gc gracefulConn) Close() error {
	gc.srv.wg.Done()
	log.Printf("connection closed")
	return gc.Conn.Close()
}

type gracefulListener struct {
	net.Listener
	stopped bool
	srv     *gracefulServer
}

func newGracefulListener(addr string) (*gracefulListener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	gl := &gracefulListener{
		Listener: l,
		stopped:  false,
	}

	return gl, nil
}

func (gl *gracefulListener) Accept() (c net.Conn, err error) {
	c, err = gl.Listener.Accept()
	if err != nil {
		return
	}

	c = gracefulConn{Conn: c, srv: gl.srv}
	// add wait group later
	return
}

func (gl *gracefulListener) Close() error {
	gl.stopped = true

	return gl.Listener.Close()
}

func (gl *gracefulListener) File() *os.File {
	l := gl.Listener.(*net.TCPListener)
	fd, _ := l.File()
	return fd
}

func simpleHandler(conn net.Conn) {
	buf := make([]byte, 4)
	n, err := io.ReadFull(conn, buf)
	if err != nil {
		conn.Close()
	}
	fmt.Printf("%d bytes are read, %s", n, string(buf))
	io.WriteString(conn, "hello world")
	conn.Close()
}

func main() {
	var gracefulChild bool

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")
	flag.Parse()

	server, err := newGracefulServer(gracefulChild)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
	log.Printf("%d server waits for alive connections to close", os.Getpid())

	server.wg.Wait()
	log.Printf("%d server quited with no live connections", os.Getpid())
}
