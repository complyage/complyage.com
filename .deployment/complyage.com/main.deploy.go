package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleDeployment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Deployment...")

	//----------------------------------------------------------------------------------------------
	// Repo Directories
	//----------------------------------------------------------------------------------------------

	repoDir := "/complyage/complyage.com"
	configSrc := "/complyage/.config/config.json"
	envDeploySrc := "/complyage/.config/.env.deployment"
	envComplyageSrc := "/complyage/.config/.env.complyage"
	envDeployDst := filepath.Join(repoDir, "deployment/.env")
	envComplyageDst := filepath.Join(repoDir, ".env")

	//----------------------------------------------------------------------------------------------
	// Check if repo folder exists; if not, clone it
	//----------------------------------------------------------------------------------------------

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Repo not found, cloning...")
		cmd := exec.Command("git", "clone", "https://github.com/complyage/complyage.com.git", repoDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			http.Error(w, fmt.Sprintf("git clone failed: %v", err), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("Repo already exists.")
	}

	//----------------------------------------------------------------------------------------------
	// Copy config.json
	//----------------------------------------------------------------------------------------------

	if err := CopyFile(configSrc, filepath.Join(repoDir, "config.json")); err != nil {
		http.Error(w, fmt.Sprintf("copy config.json failed: %v", err), http.StatusInternalServerError)
		return
	}

	//----------------------------------------------------------------------------------------------
	// Copy environment files
	//----------------------------------------------------------------------------------------------

	if err := CopyFile(envDeploySrc, envDeployDst); err != nil {
		http.Error(w, fmt.Sprintf("copy .env.deployment failed: %v", err), http.StatusInternalServerError)
		return
	}

	if err := CopyFile(envComplyageSrc, envComplyageDst); err != nil {
		http.Error(w, fmt.Sprintf("copy .env.complyage failed: %v", err), http.StatusInternalServerError)
		return
	}

	//----------------------------------------------------------------------------------------------
	// Build and restart Docker Compose
	//----------------------------------------------------------------------------------------------

	fmt.Println("Rebuilding and restarting Docker stack...")
	cmd := exec.Command("docker", "compose", "-f", filepath.Join(repoDir, "docker-compose.yml"), "up", "--build", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, fmt.Sprintf("docker build failed: %v", err), http.StatusInternalServerError)
		return
	}

	//----------------------------------------------------------------------------------------------
	// Response
	//----------------------------------------------------------------------------------------------

	fmt.Println("Deployment Completed Successfully.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Deployment Completed Successfully"))
}
