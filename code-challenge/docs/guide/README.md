# MELI Technical Challenge

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

## Introducción

Este es un desafío técnico para el proyecto MELI. El objetivo de este desafío es crear una aplicación web que permita al usuario enviar un archivo y recibir una respuesta que lo guarde en una base de datos desde la API de MELI.

La aplicación se creará utilizando el lenguaje de programación Go y utilizará las siguientes tecnologías:

## Primeros pasos

<details>
  <summary><b>&emsp;Revisión e Instalación de Go</b></summary>
  <p>En caso de no tener instalado Go, siga los pasos del link-oficial para instalarlo:  <pre><a href="https://golang.org/doc/install">https://golang.org/doc/install</a></pre></p>
  <p>Si tiene instalado Go, revisar la versión instalada con: <code>go version</code>
  <br>El proyecto esta ejecutándose en la versión <u><b>1.22</b></u> de Go.
  </p>
</details>

<details>
  <summary><b>&emsp;Revisión e Instalación de SQLite</b></summary>
  <p>En caso de no tener instalado SQLite, siga los pasos a continuación para instalarlo:  <a href="https://www.sqlite.org/download.html">https://www.sqlite.org/download.html</a></p>
  <p>Si tiene instalado SQLite, revisar la versión instalada con: <code>sqlite3 --version</code>
  <br>El proyecto esta ejecutándose en la versión <u><b>3.41.3</b></u> de SQLite.
  </p>
</details>

<details>
    <summary><b>&emsp;Revisión e Instalación de Docker</b></summary>
    <p>En caso de no tener instalado Docker, siga los pasos a continuación para instalarlo:  <a href="https://docs.docker.com/engine/install/">https://docs.docker.com/engine/install/</a></p>
    <p>Si tiene instalado Docker, revisar la versión instalada con: <code>docker --version</code>
    <br>El proyecto esta ejecutándose en la versión <u><b>27.3.1</b></u> de Docker.
</details>

## Ejecución y despliegue

Para ejecutar el proyecto, siga los siguientes pasos:

1. Abra una terminal y navega hasta la carpeta del proyecto.

- Se considera que la carpeta `/code-challenge` es la raíz del proyecto.

2. Ejecute el siguiente comando para construir la imagen de Docker con la aplicación:

> [!NOTE]  
> Para ejecutar los comandos del archivo `makefile` puede requerir permisos.

```bash
make start
```

- Al momento de ejecutar el comando, se <u>construirá la imagen de Docker</u> (puede demorar, toma 45s) y se abrirá (en el navegador por defecto) el <u>docsify</u> con estas mismas especificaciones del proyecto en Go.

3. Para correr la aplicación, ejecute el siguiente comando:

```bash
make run
```

> [!NOTE]  
> Puede requerir permisos para ejecutar el comando.

<details>
  <summary>&emsp;<b>Sin permisos o acceso denegado al archivo</b></summary>
  <p>Para correr el proyecto, ejecute el siguiente comando:</p>
  <pre><code>docker run -p 8080:8080 -v $(pwd)/items.db:/app/items.db gin-sqlite-app & open http://localhost:8080/docs/index.html</code></pre>
</details><br>

- Al ejecutar el comando se correrá el contenedor con la imagen de Docker construida previamente. Y se abrirá (en el navegador por defecto) el <u>swagger</u> con los campos para realizar la consulta.

- El swagger se puede acceder en el siguiente enlace, si no se encuentra, abra el navegador en la siguiente dirección:
<pre><a href="http://localhost:8080/docs/index.html">http://localhost:8080/docs/index.html</a></pre>

- Dentro del swagger, acceda al endpoint `/load_file` y realice la petición. Con el Try it out, se puede realizar la consulta.

<details>
  <summary>&emsp;<b>En caso de que lo requiera puede usar - CURL</b></summary>
  <pre><code>curl --location 'localhost:8080/ping' \
--header 'Content-Type: text/plain' \
--data-binary '@'</code></pre>
</details><br>

4. Una vez termine de realizar la petición y se obtenga `Status : 200` se guardará el resultado en la base de datos. Para ver el resultado, ejecute el siguiente comando:

```bash
make get-data-sql
```

5. Para detener el contenedor, ejecute el siguiente comando:

```bash
make stop
```

- También se puede usar los comandos del teclado para detener el contenedor: `Ctrl + C`.

> [!TIP]  
> Se puede abrir el archivo `makefile` para ver los comandos que se utlizaron.
> El PATH de la documentación se puede ver en el archivo `docs/guide/`.

## Paquetes y librerías

<ul>
    <li>GIN-GONIC : <a href="https://pkg.go.dev/github.com/gin-gonic/gin">github.com/gin-gonic/gin v1.10.0</a> : Validado ✅</li>
    <li>SWAGGO : <a href="https://pkg.go.dev/github.com/swaggo/swag/v2">github.com/swaggo/swag v1.16.4</a> : Validado ✅</li>
        <ul>
            <li>SWAGGER</li>
            <li>DOCSIFY</li>
        </ul>
    <li>SQLITE3 : <a href="https://pkg.go.dev/github.com/mattn/go-sqlite3">github.com/mattn/go-sqlite3 v1.14.24</a> : Validado ✅</li>
</ul>
