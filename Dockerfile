FROM golang

WORKDIR /App

COPY go.mod .

RUN go mod download && go mod verify

COPY . .

RUN go get lambo-rizz-bot-go/components

RUN go build -v -o /usr/local/bin/app

CMD [ "app" ]