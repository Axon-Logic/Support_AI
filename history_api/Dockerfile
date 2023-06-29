FROM python:3.8.7-slim

# set work directory
WORKDIR /history_api

COPY . .

# install dependencies
RUN pip install --upgrade pip

RUN pip install -r requirements.txt

EXPOSE 5679

CMD ["python", "main.py"]
RUN pwd
RUN ls
