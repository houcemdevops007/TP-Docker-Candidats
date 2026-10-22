FROM python:3.14-slim
WORKDIR /srv
COPY app.py .
USER nobody
EXPOSE 8000
CMD ["python", "server.py"]
