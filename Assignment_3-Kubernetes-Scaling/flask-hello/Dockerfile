FROM python:3.10.3-slim

RUN pip install Flask==3.0.3

COPY . /app
WORKDIR /app

ENV MESSAGE="Dystopia is the only utopia."

CMD python hello.py