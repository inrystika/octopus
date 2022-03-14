FROM python:3.7
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN pip install --upgrade pip
RUN apt update
RUN mkdir -p /app/models
VOLUME  /app/models
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
EXPOSE 5000

# Define environment variable
ENV MODEL_NAME PytorchServer
ENV SERVICE_TYPE MODEL
ENV PERSISTENCE 0


CMD exec seldon-core-microservice $MODEL_NAME --service-type $SERVICE_TYPE --persistence $PERSISTENCE
