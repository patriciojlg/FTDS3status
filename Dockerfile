# Etapa 1: Build del binario
FROM golang:1.22-alpine AS builder

# Instala herramientas necesarias
RUN apk add --no-cache git

# Define directorio de trabajo
WORKDIR /app

# Copia go.mod y go.sum, descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código completo
COPY . .

# Compila el binario para Lambda (nombre obligatorio: bootstrap)
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/lambda

# Etapa 2: imagen final para AWS Lambda runtime
FROM public.ecr.aws/lambda/provided:al2

# Copia el binario a la ruta esperada por Lambda
COPY --from=builder /app/bootstrap /var/runtime/bootstrap

# (Opcional) Define el comando de inicio del runtime Lambda
# Aunque AWS Lambda lo usa automáticamente, esto puede ayudar si ejecutas localmente o depuras
CMD ["/var/runtime/bootstrap"]