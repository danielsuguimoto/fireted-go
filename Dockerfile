FROM golang:1.18

ENV GO111MODULE=off

RUN go get golang.org/x/tools/cmd/present

EXPOSE 3999

WORKDIR /presentation

ENTRYPOINT ["present"]

CMD ["-http=0.0.0.0:3999", "-use_playground"]