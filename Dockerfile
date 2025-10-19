FROM python:3.12-slim

# Install git, curl, and docker-compose
RUN apt-get update && \
    apt-get install -y git curl docker-compose && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy webhook listener
COPY webhook_listener.py /app/webhook_listener.py

# Copy deploy script and make it executable
COPY deploy_vps_sakamoto.sh /app/deploy.sh
RUN chmod +x /app/deploy.sh

# Install Flask
RUN pip install --no-cache-dir flask

EXPOSE 9000

CMD ["python3", "webhook_listener.py"]
