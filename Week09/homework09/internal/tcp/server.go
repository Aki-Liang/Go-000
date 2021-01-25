package tcp

import (
	"bufio"
	"context"
	"fmt"
	"homework09/internal/model"
	"homework09/internal/utils/routine"
	"net"
	"sync"
	"time"

	xerrors "github.com/pkg/errors"
)

type TCPServer struct {
	addr string
	l    net.Listener
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		addr: addr,
	}
}

func (s *TCPServer) Start(ctx context.Context) error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return xerrors.Wrap(err, "tcp: listen failed")
	}
	s.l = listen
	routine.GO(ctx, s.listen)
	return nil
}

func (s *TCPServer) listen(ctx context.Context) {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			fmt.Println("accept error:", err.Error())
		} else {
			fmt.Println("accept success")
			c := &AcceptClient{
				conn: conn,
			}
			c.start(ctx)
		}

		select {
		case <-ctx.Done():
			fmt.Println("tcp listen routine stoped")
			return
		default:
			continue
		}
	}
}

type AcceptClient struct {
	conn    net.Conn
	bStoped bool
	lock    sync.Mutex
	c       chan model.TCPMsg
}

func (c *AcceptClient) start(ctx context.Context) {
	c.c = make(chan model.TCPMsg)
	ctx, cancle := context.WithCancel(ctx)
	go func(ctx context.Context) {
		rd := bufio.NewReader(c.conn)
		for {
			line, _, err := rd.ReadLine()
			if err != nil {
				fmt.Println("read error:", err.Error())
				cancle()
				return
			}
			// fmt.Printf("receive begin")
			// var buf [128]byte
			// n, err := c.conn.Read(buf[:])

			// if err != nil {
			// 	fmt.Printf("read from connect failed, err: %v\n", err)
			// 	break
			// }
			// str := string(buf[:n])
			// fmt.Printf("receive from client, data: %v\n", str)

			fmt.Println("recv client msg:", string(line))

			c.c <- model.TCPMsg{
				Content: fmt.Sprintf("Server msg:%v", time.Now().Format("2006-01-02 15:04:05")),
			}
			select {
			case <-ctx.Done():
				c.stop()
				fmt.Println("accept client read routine stoped")
				return
			default:
				continue
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		wr := bufio.NewWriter(c.conn)
		for {
			select {
			case msg := <-c.c:
				fmt.Println("send server msg:", msg)
				content := append([]byte(msg.Content), '\n')
				_, err := wr.Write(content)
				if err != nil {
					fmt.Println("write string error:", err.Error())
					cancle()
					return
				}
				err = wr.Flush()
				if err != nil {
					fmt.Println("flush conn error:", err.Error())
					cancle()
					return
				}

			case <-ctx.Done():
				c.stop()
				fmt.Println("accept client write routine stoped")
				return
			default:
				continue
			}
		}
	}(ctx)
}

func (c *AcceptClient) stop() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.bStoped {
		c.conn.Close()
		close(c.c)
		c.bStoped = true
	}
}
