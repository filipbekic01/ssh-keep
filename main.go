package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/manifoldco/promptui"
)

func main() {
	var args []string
	var argsLen int
	var cmd string

	args = os.Args[1:]
	argsLen = len(args)

	if argsLen == 0 {
		list()
		return
	}

	cmd = args[0]

	if cmd == "help" {
		fmt.Println("Edit ~/.ssh_keep.conf file.")
		return
	}
}

func list() {

	var items []string

	// Get user
	osUser, err := user.Current()
	check(err)

	// Open file
	file, err := os.Open(osUser.HomeDir + "/.ssh-keep.conf")
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	items = append(items, "exit")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Template
	templates := &promptui.SelectTemplates{
		Help:     "🏰 ssh-keep",
		Label:    "Open new SSH tunnel to:",
		Inactive: "  {{ . }}",
		Active:   "➤ {{ . | green }}",
	}

	// Select
	prompt := promptui.Select{
		Items:     items,
		Templates: templates,
	}

	_, result, err := prompt.Run()
	check(err)

	if result != "exit" {
		ssh(result)
	}
}

func ssh(cmdTest string) {
	usr, err := user.Current()
	check(err)

	uid, err := strconv.Atoi(usr.Uid)
	check(err)

	gid, err := strconv.Atoi(usr.Gid)
	check(err)

	// The Credential fields are used to set UID, GID and attitional GIDS of the process
	// You need to run the program as  root to do this
	var cred = &syscall.Credential{
		Uid:         uint32(uid),
		Gid:         uint32(gid),
		Groups:      []uint32{},
		NoSetGroups: true,
	}

	// the Noctty flag is used to detach the process from parent tty
	var sysproc = &syscall.SysProcAttr{Credential: cred, Noctty: true}
	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},

		Sys: sysproc,
	}

	// Start process
	process, err := os.StartProcess("/bin/ssh", []string{"/bin/ssh", "projectnelth.com"}, &attr)
	if err != nil {
		fmt.Println(err.Error())
	}

	process.Wait()
}
