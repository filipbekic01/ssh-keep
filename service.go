package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/manifoldco/promptui"
)

func list() {
	var items []string

	// Open file
	file, err := os.Open(USER.HomeDir + "/" + CONF_FILE)
	if err != nil {
		os.Create(USER.HomeDir + "/" + CONF_FILE)
		fmt.Println(INFO)
		return
	}
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	items = append(items, EXIT)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Template
	templates := &promptui.SelectTemplates{
		Help:     HELP,
		Label:    LABEL,
		Inactive: INACTIVE,
		Active:   ACTIVE,
	}

	// Select
	prompt := promptui.Select{
		Items:     items,
		Templates: templates,
	}

	_, command, err := prompt.Run()
	check(err)

	if command != EXIT {
		run(command)
	}
}

func run(command string) {
	uid, err := strconv.Atoi(USER.Uid)
	check(err)

	gid, err := strconv.Atoi(USER.Gid)
	check(err)

	cred := &syscall.Credential{
		Uid:         uint32(uid),
		Gid:         uint32(gid),
		Groups:      []uint32{},
		NoSetGroups: true,
	}

	sysproc := &syscall.SysProcAttr{Credential: cred, Noctty: true}

	attr := os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},

		Sys: sysproc,
	}

	params := strings.Split(command, " ")
	params = append([]string{SSH_PATH}, params...)

	process, err := os.StartProcess(SSH_PATH, params, &attr)
	if err != nil {
		fmt.Println(err.Error())
	}

	process.Wait()
}

func input(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	return text
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
