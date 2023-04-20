
# GCP Pub/Sub with Redis

Este proyecto implementa una aplicación en Go que recibe mensajes desde Google Cloud Pub/Sub y los almacena en Redis. Esta aplicación puede ser utilizada como base para construir sistemas de procesamiento de mensajes de alta escalabilidad y disponibilidad.


## Configuración
Antes de ejecutar la aplicación, es necesario configurar los parámetros de conexión a Google Cloud Pub/Sub y Redis. Estos parámetros se pueden configurar a través de variables de entorno, archivos de configuración o argumentos de línea de comandos.

Los siguientes parámetros deben ser configurados:

GOOGLE_CLOUD_PROJECT: El ID del proyecto de Google Cloud donde se encuentra el servicio Pub/Sub.
GOOGLE_APPLICATION_CREDENTIALS: La ruta al archivo JSON que contiene las credenciales de servicio de Google Cloud.

Además, se pueden configurar otros parámetros opcionales, como el nombre del topic de Pub/Sub y el prefijo de las claves en Redis.


## Contribuciones
Si deseas contribuir a este proyecto, ¡eres bienvenido! Para comenzar, puedes hacer un fork del repositorio y enviar pull requests con tus cambios.

