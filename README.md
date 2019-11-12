# Loyalty Service en GO

[Microservicios Loyalty](https://github.com/gabriel-giorgi/tp-arquitectura-de-microservicios)

Se encarga de manejar los profiles de los usuarios como asi manejar el descuento que le corresponde dependiendo de su nivel

[Documentación de API](./README-API.md)

La documentación de las api también se pueden consultar desde el home del microservicio
que una vez levantado el servidor se puede navegar en [localhost:4100]http://localhost:4100/v1/loyalty/)

## Requisitos

Go 1.10  [golang.org](https://golang.org/doc/install)

Dep [github.com/golang/dep](https://github.com/golang/dep)


## Instalar Librerías requeridas

```bash
dep ensure
```

Build y ejecución
-Ejecutar main.go

## MongoDB

La base de datos se almacena en MongoDb.

Seguir las guías de instalación de mongo desde el sitio oficial [mongodb.com](https://www.mongodb.com/download-center#community)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## RabbitMQ

Este microservicio notifica los logouts de usuarios con Rabbit.

Seguir los pasos de instalación en la pagina oficial [rabbitmq.com](https://www.rabbitmq.com/)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## Apidoc

Apidoc es una herramienta que genera documentación de apis para proyectos node (ver [Apidoc](http://apidocjs.com/)).

El microservicio muestra la documentación como archivos estáticos si se abre en un browser la raíz del servidor [localhost:3000](http://localhost:3000/)

Ademas se genera la documentación en formato markdown.

Para que funcione correctamente hay que instalarla globalmente con

```bash
npm install apidoc -g
npm install -g apidoc-markdown2
```

La documentación necesita ser generada manualmente ejecutando la siguiente linea en la carpeta loyalty :

```bash
apidoc -o www
apidoc-markdown2 -p www -o README-API.md
```

Esto nos genera una carpeta con la documentación, esta carpeta debe estar presente desde donde se ejecute loyalty, loyalty busca ./www para localizarlo, aunque se puede configurar desde el archivo de properties.

## Archivo config.json

Este archivo permite configurar los parámetros del servidor, ver ejemplos en config-example.json.
El servidor busca el archivo "./config.json". Podemos definir el archivo su ruta completa ejecutando

