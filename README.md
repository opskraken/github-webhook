# 🚀 GitHub Webhook Auto-Deploy Listener

This project is a lightweight **webhook listener** built in Go that automatically:

1. Receives push events from a GitHub repository ✅
2. Clones or pulls the latest code from the repo 🧰
3. Runs `docker compose up -d` to redeploy the application 🐳

Perfect for **simple CI/CD deployments** without heavy tooling.

---

## 📁 How It Works

### 1. Webhook Trigger

* When you push to a GitHub repo,
* GitHub sends a POST request to `http://<your-server>:8080/webhook`.

### 2. Clone or Pull the Repo

* If the repo is already cloned in `/app/hello-world`:

  * It runs `git pull`.
* If not cloned yet:

  * It runs `git clone <repo-url> /app/hello-world`.

### 3. Run `docker compose`

* Once the code is updated,
* It executes:

  ```bash
  docker compose up -d
  ```

  inside the repository folder.

---

## 🧰 Project Structure

```
.
├── main.go                 # Webhook listener source code
├── Dockerfile              # Container build file
└── README.md               # Documentation
```

---

## 🐳 Build & Run the Webhook Listener (Docker)

### 1. Build the Docker image

```bash
docker build -t webhook-listener .
```

### 2. Run the container

```bash
docker run -p 8080:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /home/iqbal/hello-world:/app/hello-world \
  --name webhook-listener \
  webhook-listener
```

* `-p 8080:8080` → Exposes the webhook listener on port 8080
* `-v /var/run/docker.sock:/var/run/docker.sock` → Allows container to control Docker on the host
* `-v /home/iqbal/hello-world:/app/hello-world` → Persists your app code on the host

---

## 🌐 Expose Your Local Server (Optional)

If you want GitHub to reach your local machine:

```bash
ngrok http 8080
```

Copy the HTTPS forwarding URL from ngrok and use it as your webhook URL.

---

## ⚙️ Configure GitHub Webhook

1. Go to your GitHub repository → **Settings** → **Webhooks**.
2. Click **“Add webhook”**.
3. Set:

   * **Payload URL:** `https://<ngrok-url>/webhook`
   * **Content type:** `application/json`
   * Select event: “Just the push event”.
4. Save the webhook.

✅ Now every push triggers the listener → pulls latest code → redeploys with Docker Compose.

---

## 🧪 Test Manually

You can also test the webhook endpoint manually with:

```bash
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{"repository": {"clone_url":"https://github.com/your/repo.git","name":"hello-world"}, "ref":"refs/heads/main"}'
```

---

## 🧼 Logs & Monitoring

View logs from the running container:

```bash
docker logs -f webhook-listener
```

If something goes wrong (e.g., Git clone or Docker compose fails), the error is printed in the logs.

---

## 🧰 Prerequisites

* Git installed inside the container (already included in the Dockerfile)
* Docker & Docker Compose installed on the host
* The container must have access to the host’s Docker socket:

  ```bash
  -v /var/run/docker.sock:/var/run/docker.sock
  ```

---

## 🧭 Environment Overview

| Component        | Purpose                           |
| ---------------- | --------------------------------- |
| Go Web Server    | Receives and processes webhooks   |
| Git              | Clones and pulls repository       |
| Docker Compose   | Runs application after deployment |
| ngrok (optional) | Exposes local server to GitHub    |

---

## 🛑 Stopping the Listener

```bash
docker stop webhook-listener
docker rm webhook-listener
```

---

## ✨ Future Improvements

* Add support for branch-based deployments (e.g., only `main`)
* Add authentication or secret token validation
* Add error reporting or Slack notifications
* Add rollback mechanism if deploy fails

---

✅ **You now have a simple auto-deploy system using Go + Git + Docker Compose.**

---