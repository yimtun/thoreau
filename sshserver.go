package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/sirupsen/logrus"

	"github.com/gliderlabs/ssh"
	"github.com/yimtun/kr-pty"
)

var (
	account = make(map[string]string)
)

func userAdd(user, passwd string) {
	account[user] = passwd

}

func init() {
	userAdd("user1", "123.123")
	userAdd("xxx", "123")
	userAdd("yyy", "456")
	userAdd("zzz", "456")
	//	fmt.Println(account["user1"])
	///	fmt.Println(findUser("user1"))
	//	fmt.Println(FindPasswd("user1"))
}

func findUser(user string) bool {

	if _, ok := account[user]; ok {
		return true
	} else {
		return false
	}
}

func findPasswd(user string) string {
	return account[user]

}

func main() {
	SshserverStart()
}

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func SshserverStart() {
	ssh.Handle(func(s ssh.Session) {
		fmt.Println("获取到的用户：", s.User())
		fmt.Println("获取的客户端地址:", s.RemoteAddr())

		cmd := exec.Command("./controller", s.User())
		ptyReq, winCh, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			if err != nil {
                                fmt.Println("报错")
				fmt.Println(err)
				//panic(err)
			}
			go func() {
				for win := range winCh {
					setWinsize(f, win.Width, win.Height)
				}
			}()
			go func() {
				io.Copy(f, s) // stdin
			}()
			io.Copy(s, f) // stdout

			err = cmd.Wait()
			if err != nil {
				logrus.Warn(err)
			}
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}
	})

	log.Println("starting ssh server on port 2022...")
	ssh.ListenAndServe(":2022", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			//			fmt.Println(ctx.User())
			//	fmt.Println(ctx.User())
			return findUser(ctx.User()) && findPasswd(ctx.User()) == pass
		}),
	)
}
