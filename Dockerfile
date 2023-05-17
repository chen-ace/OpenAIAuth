FROM golang

WORKDIR /app

COPY . /app

RUN go build && cp OpenAIAuth /bin/OpenAIAuth

EXPOSE 9090

ENTRYPOINT ["OpenAIAuth"]