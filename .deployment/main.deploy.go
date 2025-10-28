//================================================================================================
// Deploy Script - Go HTTP Trigger
//================================================================================================

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

//------------------------------------------------------------------------------------------------
// HandleDeployment
//------------------------------------------------------------------------------------------------

func HandleDeployment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Deployment...")

	repoDir := "/complyage/complyage.com"
	deployDir := filepath.Join(repoDir, ".deployment")
	deployScript := filepath.Join(deployDir, "deploy.sh")

	configSrc := "/complyage/.config/config.json"
	envComplyageSrc := "/complyage/.config/.env"
	configDst := filepath.Join(repoDir, "config.json")
	envComplyageDst := filepath.Join(repoDir, ".env")

	//--------------------------------------------------------------------------------------------
	// Clean + Clone repo fresh
	//--------------------------------------------------------------------------------------------
	fmt.Println("Cleaning and cloning repository...")
	os.RemoveAll(repoDir)
	cmd := exec.Command("git", "clone", "https://github.com/complyage/complyage.com.git", repoDir)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, fmt.Sprintf("git clone failed: %v", err), http.StatusInternalServerError)
		return
	}

	//--------------------------------------------------------------------------------------------
	// Copy config and environment
	//--------------------------------------------------------------------------------------------
	for src, dst := range map[string]string{
		configSrc:       configDst,
		envComplyageSrc: envComplyageDst,
	} {
		fmt.Printf("Copying %s -> %s\n", src, dst)
		if err := CopyFile(src, dst); err != nil {
			http.Error(w, fmt.Sprintf("copy failed %s: %v", src, err), http.StatusInternalServerError)
			return
		}
	}

	//--------------------------------------------------------------------------------------------
	// Make deploy.sh executable
	//--------------------------------------------------------------------------------------------
	fmt.Println("Setting deploy.sh executable...")
	if err := os.Chmod(deployScript, 0755); err != nil {
		http.Error(w, fmt.Sprintf("chmod deploy.sh failed: %v", err), http.StatusInternalServerError)
		return
	}

	//--------------------------------------------------------------------------------------------
	// Execute deploy.sh
	//--------------------------------------------------------------------------------------------
	fmt.Println("Executing deploy.sh...")
	cmd = exec.Command("bash", deployScript)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, fmt.Sprintf("deploy.sh failed: %v", err), http.StatusInternalServerError)
		return
	}

	//--------------------------------------------------------------------------------------------
	// Response
	//--------------------------------------------------------------------------------------------
	fmt.Println("Deployment Completed Successfully.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Deployment Completed Successfully"))
}
