FROM golang:1.16

WORKDIR /app
COPY go.mod . 
COPY go.sum . 
RUN go mod download

COPY . . 

# 자동 리로드
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# CMD ["go", "run", "main.go"]

CMD ["air"]