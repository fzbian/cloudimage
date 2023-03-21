# CloudImage

Una API que funciona bajo un servidor el cual proporcionara la infraestructura para permitir alojar imagenes y requerirlas por medio de el link.

## Requisitos

Antes de usar la API, asegúrate de tener lo siguiente:

 - **Go (golang) 1.16 o superior:** La API está escrita en Golang, por lo que necesitarás tener Python 3.8 o superior instalado en tu máquina.
 - **Paqueteria actualizada:** La API usa algunos paquetes de Golang. Para instalarlos, ejecuta el siguiente comando en tu terminal:
    ```bash
    go get .
    ```
 - **Archivo de configuración:** La API utiliza un archivo de configuración para almacenar credenciales y configuraciones específicas. Debes crear un archivo `.env` en la raíz del proyecto y configurar las siguientes variables de entorno:
    ```bash
   DOMAIN=http://domain.com/
   PORT=8080
    ```

Aunque si vas a **compilar el binario** y usarlo directamente en tu servidor recuerda compilarlo adecuadamente para el sistema operativo adecuado, el sistema operativo del servidor.

## Documentación de la API

La API proporciona acceso a los recursos del sistema a través de una interfaz HTTP. Todas las solicitudes requieren autenticación y deben usar HTTPS.

### Autenticación

La API no usa ningun medio de autenticación

### Recursos disponibles
La API proporciona los siguientes recursos:

#### Health

`GET /api/health`

Devuelve un código de estado HTTP 200 si la API está en línea y es accesible.

**Respuesta exitosa**

Cuando se realiza una solicitud exitosa al recurso health, el servidor responderá con un código de estado HTTP 200.

```http request
HTTP/1.1 200 OK
```

#### POST Image

`POST /api/upload`

Crea una nueva imagen en el servidor. La imagen debe ser proporcionada en el cuerpo de la solicitud como un archivo de tipo imagen.

**Parámetros**

| Parametro | Descripcion                   |
|-----------|-------------------------------|
| `file`    | La imagen que se va a cargar  |

**Respuesta exitosa**

Cuando se carga correctamente una imagen en el servidor, el servidor responderá con un código de estado HTTP 201 y un cuerpo de respuesta con detalles sobre la imagen cargada.

```http request
HTTP/1.1 201 Created
Content-Type: application/json

{
  "link": https://example.com/image.png,
}

```

**Respuesta de error**

Cuando se produce un error en la solicitud, el servidor responderá con un código de estado HTTP de error y un cuerpo de respuesta con detalles sobre el error.

```http request
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": "Se requiere una imagen para cargar."
}

```


## Contribucion

Me encantaría que nos ayudaras a mejorar este proyecto. Para contribuir, sigue estos pasos:

1. Haz un fork de este repositorio y clónalo en tu máquina local.
2. Crea una nueva rama desde la rama principal para tu contribución:
   ```bash
   git checkout -b nombre-de-tu-rama
   ```
3. Haz tus cambios y comprueba que todo funcione correctamente.
4. Haz commit de tus cambios con un mensaje claro y descriptivo:
   ```bash
   git commit -m "Agregar nueva funcionalidad para X"
   ```
5. Sube tu rama a tu fork en GitHub:
   ```bash
   git push origin nombre-de-tu-rama
   ```
6. Crea un pull request a la rama principal de este repositorio.
7. Espera comentarios y feedback.
8. Si se solicita algún cambio, hazlos y vuelve al paso 4.
9. Si tu pull request es aceptado, tu contribución será mergeada a la rama principal del proyecto.

Gracias por contribuir a este proyecto. ¡Espero ver tus cambios pronto!