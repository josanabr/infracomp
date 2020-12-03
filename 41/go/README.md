# Servidor web en Go y Dockerfile

En este directorio se encuentran los archivos:

* [proyecto.go](proyecto.go), este archivo contiene la implementación de un servidor HTTP en el lenguaje de programación Go.

* [Dockerfile](Dockerfile), este archivo permite la creación de una imagen de Docker que al ejecutarse crea un contenedor que ejecuta un servidor web.

---

## Compilación y ejecución de `proyecto.go`

Para compilar el programa `proyecto.go` se ejecuta el comando:

```
go build proyecto.go
```

Esto genera un archivo llamado `proyecto`. 
Se ejecuta desde la terminal:

```
./proyecto
```

Ahora, usted puede abrir un navegador que vaya al url [http://localhost:8000](http://localhost:8000).

---

## Creación de contenedor `proyecto`

Para crear la imagen que indica la presentación ejecute el siguiente comando:

```
docker build -t proyecto .
```

Para ejecutar un contenedor a partir de la imagen anterior ejecute el siguiente código:

```
docker run proyecto
```




