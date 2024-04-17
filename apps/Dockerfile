FROM node:lts-alpine as node
ARG TARGETOS
ARG TARGETARCH

RUN apk add g++ make py3-pip

WORKDIR /frontend

COPY package.json package-lock.json ./
COPY web web

RUN npm install
RUN npm run build

FROM golang:1.21-alpine as golang

WORKDIR /app
COPY . .
COPY --from=node /frontend .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /likeit-service cmd/likeit-service/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=golang /likeit-service .

EXPOSE 8080

CMD ["/likeit-service"]
