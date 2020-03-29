FROM python:2.7
WORKDIR /app
COPY app.py requirements.txt ./
RUN pip install -r requirements.txt
ENTRYPOINT ["python"]
CMD ["app.py"]