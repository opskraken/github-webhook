from flask import Flask, request
import subprocess

app = Flask(__name__)

# Optional: secret to validate GitHub webhook
WEBHOOK_SECRET = "your_secret_token"

@app.route("/deploy", methods=["POST"])
def deploy():
    # You can verify X-Hub-Signature-256 here if needed
    subprocess.Popen(["/app/deploy.sh"])
    return "Deployment started", 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=9000)
