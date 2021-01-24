FROM golang:1.15-alpine as build 

ENV COARCH=amd64 \
GOOS=linux \
CGO_ENABLED=0

WORKDIR /src
COPY . .

RUN go build \
  -ldflags "-s -w -extldflags '-static'" \
  -o /bin/action .

FROM python:3.8-slim

ARG secret
COPY --from=build /bin/action /app
COPY *.ipynb .
COPY --from=build /src/requirements.txt .
ENV secret=$secret
RUN pip install -r requirements.txt 
ENTRYPOINT [ "/app" ] 