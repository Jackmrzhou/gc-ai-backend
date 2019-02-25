FROM gcc:4.9

ENV GOPATH /goworkspace

WORKDIR $GOPATH/bin
COPY ./external-judger $GOPATH/bin
RUN make

FROM scratch

WORKDIR $GOPATH/src/github.com/jackmrzhou/gc-ai-backend
COPY . $GOPATH/src/github.com/jackmrzhou/gc-ai-backend

EXPOSE 8080
ENTRYPOINT ["./gc_ai"]