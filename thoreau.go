package main

import (
	"fmt"
	"github.com/yimtun/thoreau/auth"
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
	account=auth.GetAccount()
	fmt.Println(account)
	//
	/*
	userAdd("tom", "123")
	userAdd("cat", "456")
	userAdd("yyy", "456")
	userAdd("zzz", "456")
	 */

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
		fmt.Println("user name:：", s.User())
		fmt.Println("user client ip : port", s.RemoteAddr())
		// usrname and usertoken
		cmd := exec.Command("./controller", s.User(),s.User()+"token")
		// cmd := exec.Command("./controller", s.User(),user.Token)
		ptyReq, winCh, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			if err != nil {
                                fmt.Println(err)
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
	err:=ssh.ListenAndServe(":2022", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			//			fmt.Println(ctx.User())
			//	fmt.Println(ctx.User())
			//message code or ldap ...
			fmt.Println("ssh userName:",ctx.User())

			return findUser(ctx.User()) && findPasswd(ctx.User()) == pass

		}),

	)
	if err!=nil{}
	panic(err)

}
