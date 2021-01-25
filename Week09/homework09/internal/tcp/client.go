package tcp

import (
	"bufio"
	"context"
	"fmt"
	"homework09/internal/model"
	"net"
	"sync"

	xerrors "github.com/pkg/errors"
)

type TCPClient struct {
	addr     string
	conn     net.Conn
	bStoped  bool
	lock     sync.Mutex
	sendChan chan model.TCPMsg
	recvChan chan model.TCPMsg
}

func NewTCPClient(addr string) *TCPClient {
	return &TCPClient{
		addr: addr,
	}
}

func (c *TCPClient) Start(ctx context.Context) (func(), chan model.TCPMsg, chan model.TCPMsg, error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, nil, nil, xerrors.Wrap(err, "tcp: connect failed")
	}
	c.conn = conn
	ctx, cancle := context.WithCancel(ctx)
	c.recvChan = make(chan model.TCPMsg)
	c.sendChan = make(chan model.TCPMsg)
	go func(ctx context.Context) {
		rd := bufio.NewReader(c.conn)
		for {
			line, _, err := rd.ReadLine()
			if err != nil {
				fmt.Println("read error:", err.Error())
				cancle()
				return
			}

			c.recvChan <- model.TCPMsg{
				Content: string(line),
			}
			select {
			case <-ctx.Done():
				c.stop()
				fmt.Println("client read routine stoped")
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
			case msg := <-c.sendChan:
				fmt.Println("send msg:", msg)
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
				fmt.Println("send ok")
			case <-ctx.Done():
				c.stop()
				fmt.Println("client write routine stoped")
				return
			default:
				continue
			}
		}
	}(ctx)

	return func() {
		cancle()
	}, c.recvChan, c.sendChan, nil
}

func (c *TCPClient) stop() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.bStoped {
		c.conn.Close()
		close(c.sendChan)
		close(c.recvChan)
		c.bStoped = true
	}
}
