FROM python:3.9-slim

WORKDIR /app

COPY controller.py /app
COPY requirements.txt /app

RUN apt-get update && \ 
    apt-get install -y curl && \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    pip install --no-cache-dir -r requirements.txt && \
    rm kubectl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

CMD ["python", "controller.py"]
