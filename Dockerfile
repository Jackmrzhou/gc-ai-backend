FROM gcc:4.9 as judgerBuilder

ENV GOPATH /goworkspace

WORKDIR $GOPATH/build
COPY ./external-judger $GOPATH/build
RUN make

#FROM ubuntu:16.04
#
#RUN apt-get update
#RUN apt-get install g++ -y

FROM ubuntu_cpp

ENV GOPATH /goworkspace

RUN mkdir -p $GOPATH/bin
COPY --from=judgerBuilder $GOPATH/build/judger $GOPATH/bin

WORKDIR $GOPATH/src/github.com/jackmrzhou/gc-ai-backend
COPY . $GOPATH/src/github.com/jackmrzhou/gc-ai-backend

EXPOSE 8080
ENTRYPOINT ["./gc_ai"]