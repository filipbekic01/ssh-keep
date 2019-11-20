package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

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
		help()
		return
	}

	if cmd == "add" {
		add()
		return
	}
}

func list() {

	var items []string

	// Get user
	osUser, err := user.Current()
	check(err)

	// Open file
	file, err := os.Open(osUser.HomeDir + "/ssh-keep.conf")
	check(err)
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Template
	templates := &promptui.SelectTemplates{
		Help:     "🏰 ssh-keep",
		Label:    "select ssh connection",
		Inactive: "{{ . }}",
		Active:   "> {{ . | green | bold }}",
	}

	// Select
	prompt := promptui.Select{
		Label:     "select ssh connection",
		Items:     items,
		Templates: templates,
	}

	_, result, err := prompt.Run()
	check(err)

	_ = result
}

func add() {
	// Get user
	osUser, err := user.Current()
	check(err)

	// Get inputs
	publicKeyPath := input("Public key path (" + osUser.HomeDir + "/.ssh/id_rsa.pub): ")
	if len(publicKeyPath) == 0 {
		publicKeyPath = osUser.HomeDir + "/.ssh/id_rsa.pub"
	}

	user := input("User (" + osUser.Username + "): ")
	if len(user) == 0 {
		user = osUser.Username
	}

	host := input("Host: ")
	if len(host) == 0 {
		host = "localhost"
	}

	port := input("Port (22): ")
	if len(port) == 0 {
		port = "22"
	}

	// Format final string
	final := "ssh " + user + "@" + host + " -i " + publicKeyPath + " -p " + port + "\n"

	// Open file
	f, err := os.OpenFile(osUser.HomeDir+"/ssh-keep.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	// Write to file
	f.WriteString(final)
	f.Sync()
}

func remove() {

}

func help() {
	fmt.Println("1) ssh-keep help")
	fmt.Println("2) ssh-keep add")
}
