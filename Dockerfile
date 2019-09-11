FROM golang

ARG app_env

ENV ENV $app_env

COPY . /go/src/deli/exp-service
WORKDIR /go/src/deli/exp-service

RUN make build

#RUN echo $app_env
CMD ./exp-service -E ${ENV}

EXPOSE 8080
