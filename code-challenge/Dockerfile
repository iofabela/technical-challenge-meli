# Go imagen base
FROM golang:1.21

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios al contenedor
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el contenido del proyecto al contenedor
COPY . .

# Construye el binario de la aplicación
RUN go build -o /main ./cmd/api

# Exponer el puerto en el que correrá la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/main"]
