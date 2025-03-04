# Usa uma imagem oficial do Go como base
FROM golang:1.18-alpine

# Instala dependências do sistema (GCC, build-base para compilar SQLite)
RUN apk add --no-cache gcc musl-dev

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para o container
COPY . .

# Instala as dependências do Go
RUN go mod tidy

# Expõe a porta usada pelo app
EXPOSE 9000

# Comando para rodar a aplicação
CMD ["go", "run", "."]
