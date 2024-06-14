package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	deletelog "github.com/hellden/delete-docker-log/internal/delete-log"
	log "github.com/hellden/delete-docker-log/internal/logging"
)

/* ------------------------------ | INPUT USER | ------------------------------ */
func InputUser() string {
	if len(os.Args) < 2 {
		fmt.Println("Specify container name")
		os.Exit(1)
	}
	return os.Args[1]
}

/* ------------------------------ | GET CONTAINER ID Optinnal| ------------------------------ */
func GetContainerId(containerName string) (string, error) {
	// Create buffer
	var out bytes.Buffer
	var stderr bytes.Buffer

	// Execute commande Pwsh -> docker ps -q --filter name="container name"
	cmd := exec.Command("docker", "ps", "-q", "--filter", fmt.Sprintf("name=%s", containerName))
	fmt.Printf("%v", &out)

	// Param out
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("erreur : %v: %s", err, stderr.String())
	}

	containerId := strings.TrimSpace(out.String())
	if containerId == "" {
		return "", fmt.Errorf("no container name %s", containerName)
	}

	return containerId, nil
}

func GetPathContainer(containerName string) (string, error) {
	// Success
	var out bytes.Buffer

	// Error
	var stderr bytes.Buffer
	cmd := exec.Command("docker", "inspect", "--format='{{.LogPath}}'", InputUser())
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("erreur : %v: %s", err, stderr.String())
	}

	pathLogContainer := strings.TrimSpace(out.String())

	return pathLogContainer, nil
}

func main() {
	log.Init()

	pathLogContainre, err := GetPathContainer(InputUser())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deletelog.DeleteLog(pathLogContainre)
	// fmt.Println("ID container :", containerID)
	fmt.Printf("Log path : %s", pathLogContainre)
}
