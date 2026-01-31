package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Git status
	fmt.Println("=== Git Status ===")
	cmd := exec.Command("git", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("\n=== Git Add ===")
	// Git add - ask for files
	fmt.Print("Enter files to add (use '.' to add all): ")
	files, _ := reader.ReadString('\n')
	files = strings.TrimSpace(files)

	if files == "" {
		fmt.Println("No files specified. Aborting.")
		return
	}

	cmdAdd := exec.Command("git", "add", files)
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println("Error executing git add:", err)
		return
	}
	fmt.Println("Files added successfully!")

	fmt.Println("\n=== Git Commit ===")
	// Git commit - ask for message
	fmt.Print("Enter commit message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	if message == "" {
		fmt.Println("Empty commit message. Aborting.")
		return
	}

	cmdCommit := exec.Command("git", "commit", "-m", message)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println("Error executing git commit:", err)
		return
	}
	fmt.Println("Commit completed successfully!")

	fmt.Println("\n=== Git Push ===")
	// Git push - default branch main
	fmt.Print("Press 'Enter' for 'main' branch. For another branch, enter the name: ")
	branch, _ := reader.ReadString('\n')
	branch = strings.TrimSpace(branch)

	if branch == "" {
		branch = "main"
	}

	fmt.Printf("Pushing commit to '%s' branch...\n", branch)
	cmdPush := exec.Command("git", "push", "origin", branch)
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println("Error executing git push:", err)
		return
	}
	fmt.Println("Push completed successfully!")
}
