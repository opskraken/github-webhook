package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type WebhookPayload struct {
	Repository struct {
		CloneURL string `json:"clone_url"`
		Name     string `json:"name"`
	} `json:"repository"`
	Ref string `json:"ref"`
}

const repoDir = "/app/hello-world" // where we clone/pull

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusBadRequest)
		return
	}

	// Parse payload
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Println("❌ Failed to parse JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("🔔 Push event for repo: %s, ref: %s\n", payload.Repository.Name, payload.Ref)

	// If repoDir exists, do git pull; otherwise, git clone
	cmd := exec.Command("git", "-C", repoDir, "pull")
	if err := cmd.Run(); err != nil {
		fmt.Println("Repo not found locally, cloning...")
		cmd = exec.Command("git", "clone", payload.Repository.CloneURL, repoDir)
		if out, err := cmd.CombinedOutput(); err != nil {
			fmt.Println("❌ Git clone failed:", string(out))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Run docker-compose up
	cmd = exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = repoDir
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("❌ Docker compose failed:", string(out))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("✅ Docker compose output:\n", string(out))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook processed successfully")
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Webhook listener is running 🚀")
	})
	fmt.Println("🚀 Webhook listener running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
