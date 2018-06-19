FROM python:3.6.5-jessie

RUN mkdir -p /app/templates

COPY requirements.txt /app
COPY apidemo.py /app
COPY templates/index.html /app/templates

WORKDIR /app

RUN pip install -r requirements.txt

EXPOSE 5002/tcp

ENTRYPOINT python apidemo.py
