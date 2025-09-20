FROM golang:1.24

#RUN go get github.com/gin-gonic/gin

ENV REPO_URL=github.com/virtouso/go-challenge-timely-tag-system
ENV port=8090
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/github.com/go-challenge-timely-tag-system

ENV WORKPATH=$APP_PATH/src

COPY src $WORKPATH
workdir $WORKPATH

RUN go build -o go-challenge-timely-tag-system

expose 8090

CMD ["./go-challenge-timely-tag-system"]