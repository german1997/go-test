
FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o index .
CMD ["/app/index"]


# Comandos

# Para crear la imagen del servicio: docker build -t go-test .
# Para ver las imagenes creadas (verificar que este creada): docker images
# Ejecutar imagen: docker run -it -p 8080:8080 go-test
# Abrir navegador: http://localhost:8080/beers