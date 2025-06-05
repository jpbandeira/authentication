# Etapa 1: build da aplicação
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copia os arquivos go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o authentication ./cmd/authentication

# Etapa 2: imagem final
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copia o binário da etapa anterior
COPY --from=builder /app/authentication .

# Porta que o serviço vai escutar
EXPOSE 8083

# Comando para iniciar o serviço
CMD ["./authentication"]