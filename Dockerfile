FROM python:3.6.5-jessie

RUN mkdir /app

COPY requirements.txt /app
COPY apidemo.py /app

WORKDIR /app

RUN pip install -r requirements.txt

EXPOSE 5002/tcp

ENTRYPOINT python apidemo.py
