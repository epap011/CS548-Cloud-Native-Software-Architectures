FROM python:3.9-slim

WORKDIR /app

COPY controller.py /app
COPY requirements.txt /app


RUN apt-get update && \ 
    pip install --no-cache-dir -r requirements.txt

EXPOSE 8000

CMD ["python", "controller.py"]
