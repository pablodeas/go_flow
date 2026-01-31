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
	// Git add - perguntar arquivos
	fmt.Print("Digite os arquivos para adicionar (use '.' para adicionar todos): ")
	files, _ := reader.ReadString('\n')
	files = strings.TrimSpace(files)

	if files == "" {
		fmt.Println("Nenhum arquivo especificado. Abortando.")
		return
	}

	cmdAdd := exec.Command("git", "add", files)
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println("Erro ao executar git add:", err)
		return
	}
	fmt.Println("Arquivos adicionados com sucesso!")

	fmt.Println("\n=== Git Commit ===")
	// Git commit - perguntar mensagem
	fmt.Print("Digite a mensagem do commit: ")
	mensagem, _ := reader.ReadString('\n')
	mensagem = strings.TrimSpace(mensagem)

	if mensagem == "" {
		fmt.Println("Mensagem de commit vazia. Abortando.")
		return
	}

	cmdCommit := exec.Command("git", "commit", "-m", mensagem)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println("Erro ao executar git commit:", err)
		return
	}
	fmt.Println("Commit realizado com sucesso!")

	fmt.Println("\n=== Git Push ===")
	// Git push - branch padrão main
	fmt.Print("Para branch 'main' pressione 'Enter'. Para outra branch, digite o nome: ")
	branch, _ := reader.ReadString('\n')
	branch = strings.TrimSpace(branch)

	if branch == "" {
		branch = "main"
	}

	fmt.Printf("Enviando commit para a branch '%s'...\n", branch)
	cmdPush := exec.Command("git", "push", "origin", branch)
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println("Erro ao executar git push:", err)
		return
	}
	fmt.Println("Push realizado com sucesso!")
}
