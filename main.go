package main

import (
	"fmt"
	"os"
	"os/user"
)

const CONF_FILE string = ".ssh-keep.conf"
const HELP string = "🏰 ssh-keep"
const LABEL string = "Open new SSH tunnel to:"
const INACTIVE string = "  {{ . }}"
const ACTIVE string = "➤ {{ . | green }}"
const EXE string = "bash"
const EXIT string = "exit"
const INFO string = `You're missing the configuration file.
1) Open ~/.ssh-keep.conf file")
2) Add SSH connection lines, for example:
  user@host
  user2@host2
  user3@host3
3) Happy tunneling!
  Check https://github.com/filipbekic01/ssh-keep for more configuration information.`

var USER *user.User
var ERR error

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen == 0 {
		list()
		return
	}

	if argsLen > 1 {
		fmt.Print(INFO)
	}

	cmd := args[0]

	if cmd == "help" {
		fmt.Println(INFO)
		return
	}
}

func loadUser() {
	USER, ERR = user.Current()
	check(ERR)
}
