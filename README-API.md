<a name="top"></a>
# Loyalty Service v0.1.0

Microservicio de Loyalty

- [Loyalty](#loyalty)
	- [Consultar Profile](#consultar-profile)
	- [Crear Profile](#crear-profile)
	
- [RabbitMQ_GET](#rabbitmq_get)
	- [Actualizar Profile](#actualizar-profile)
	
- [RabbitMQ_POST](#rabbitmq_post)
	- [Notificacion de LevelUp](#notificacion-de-levelup)
	


# <a name='loyalty'></a> Loyalty

## <a name='consultar-profile'></a> Consultar Profile
[Back to top](#top)

<p>Consulta un profile existente</p>

	GET /v1/loyalty/userProfile





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  Authorization |  | <p>{bearer-token}</p>|


### Success Response

Respuesta

```
HTTP/1.1 200 OK
 {
 "userID": "{userID}",
 "userLevel" : "{userLevel}" ,
 "experience" :    "{experiencia actual}",
 "currentDiscount" : "{descuento actual}"
	}
```


### Error Response

400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```
## <a name='crear-profile'></a> Crear Profile
[Back to top](#top)

<p>Crea y asocia un profile a un nuevo usuario.</p>

	POST /v1/loyalty/userProfile





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  Body |  | <p>{ &quot;userID&quot;: &quot;{userID}&quot; }</p>|


### Success Response

Body

```
HTTP/1.1 200 Ok
```


### Error Response

400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```
# <a name='rabbitmq_get'></a> RabbitMQ_GET

## <a name='actualizar-profile'></a> Actualizar Profile
[Back to top](#top)

<p>Loyalty escucha un mensaje de Cart para actualizar el profile de un user luego del checkout de un carrito.</p>

	DIRECT /loyalty/loyalty



### Examples

Mensaje

```
{
  "cartId"  : "{cartId}",
  "userId"  : "{userId}",
  "articles": [
  "id"	  : "{articleId}",
  "quantity": "{value}"
      }, ...
 ]"
}
```




# <a name='rabbitmq_post'></a> RabbitMQ_POST

## <a name='notificacion-de-levelup'></a> Notificacion de LevelUp
[Back to top](#top)

<p>Loyalty envia un mensaje de catalog para notificarle que un usuario subio de nivel</p>

	DIRECT /loyalty/loyalty_notification



### Examples

Mensaje

```
{
  	"currentLevel"   :  "{nivel actual del usuario}",
      "currentDiscount":  "{descuento actual del usuario}",
}
```




