FROM scratch

WORKDIR $GOPATH/src/github.com/jackmrzhou/gc-ai-backend
COPY . $GOPATH/src/github.com/jackmrzhou/gc-ai-backend

EXPOSE 8080
ENTRYPOINT ["./gc_ai"]