# documento_programa_crud
API de documentos de soportes para inscripción de programas académicos

Integración con
 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `documento_programa_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación:
Para usar el API, usar el comando:

 - `go get github.com/planesticud/documento_programa_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `API_DOCUMENTO_PROGRAMA_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `DOCUMENTO_PROGRAMA_CRUD__PGUSER`: Usuario de la base de datos
 - `DOCUMENTO_PROGRAMA_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `DOCUMENTO_PROGRAMA_CRUD__PGURLS`: Host de conexión
 - `DOCUMENTO_PROGRAMA_CRUD__PGDB`: Nombre de la base de datos
 - `DOCUMENTO_PROGRAMA_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo:
API_DOCUMENTO_PROGRAMA_HTTP_PORT=9014 DOCUMENTO_PROGRAMA_CRUD__PGUSER=user DOCUMENTO_PROGRAMA_CRUD__PGPASS=password DOCUMENTO_PROGRAMA_CRUD__PGURLS=localhost DOCUMENTO_PROGRAMA_CRUD__PGDB=bd DOCUMENTO_PROGRAMA_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/documento_programa_crud/blob/develop/modelo_documento_programa_crud.png).
