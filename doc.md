# Event-loop, concurrencia y paralelismo en JS

## Concurrencia

Es cuando 2 o mas procesos pueden empezar, correr, y completarse en periodos compartidos de tiempo. No necesariamente significa que van a correr al mismo tiempo, puede correr uno primero, luego el otro hasta que terminan.

## Paralelismo

Es cuando 2 o mas procesos se ejecutan a la misma vez, en paralelo, uno no bloquea al otro.

## JS

JavaScript es un lenguaje single threaded, asincrono y concurrente. V8 que es un runtime de JavaScript, tiene un call stack y un heap.

Heap es la parte de la memoria donde existe tu app.

JavaScript es single threaded, tiene 1 solo call stack y solo puede hacer solo 1 cosa a la vez.

### Call Stack

El stack en JavaScript es un FIFO (first in first out), se inserta una funcion desde la parte alta del stack, y se saca cuando su execution termina. El stacktrace te da el current state del stack, por ejemplo cuando un error pasa, para que sea mas facil de debuggear.

Mientras el call stack este procesando una funcion, nada puede pasar hasta que esta funcion retorne. Si una funcion es blocking, it stops the world hasta que termine.

En el caso de los browsers, en el thread donde se corre JavaScript se corre tambien el render del UI, lo que hace que tu browser se congele* (make sure this is true)

Cual es la solucion sa esto?, usar APIs asincronas. Usando APIs asincronas, nuestro llamado a functiones asincronas se ejecuta inmediatamente y se colocan en el task queue, luego que el callback esta listo para ejecutarse, se mueve al task queue.

```js
(function () {
   console.log('Start');
   
	setTimeout(function () {
	  console.log('Not yet')
	}, 1000);
	
   console.log('End');
}())
```

### Event-Loop

El event loop es la manera de JavaScript de ser concurrente. El event-loop tiene 1 solo trabajo, mover cosas del **callback queue** al stack. Para que esto sea posible, el stack tiene que estar vacio, asi que cuando haces una operacion asincrona, se te garantiza que primero tienes que vaciar el stack antes de que una funcion asincrona se llame

```js
(function () {   
	setTimeout(function () {
	  console.log('Not yet')
	}, 0);
	
   console.log('Start');
	
   console.log('Running');
   console.log('Running');
   console.log('Running');
   console.log('Running');
   console.log('End');
}())
```

### Callback Queue

El callback queue, es el lugar donde los callbacks de las WebAPIs viven luego que estan listos para ejectutarse. Vale resaltar que es solo WebAPIs pueden poner cosas en este queue.

```js
(function () {
  [1,2,3].forEach(function log(number) {
    console.log(number);
  });
  
  console.log('End');
}())
```

### Concurrencia

JavaScript usa el event loop, el callback stack y las APIs asincronas para manejar concurrencia

### Paralelismo

JavaScript per se, no puede ser paralelo, ya que tiene un solo thread. Asi que una cosa se procesa a la vez.

Para tener paralelimo en JavaScript, debemos ejecutar multiples procesos.

Web Workers es otra Web API que nos da concurrencia, pero no paralelismo, podemos procesar cosas en paralelo, pero a la hora de que llegan a JavaScript, encolan en el callback stack sus respectivos callbacks

Node.js tiene una libreria para hacer eso, [Cluster](https://nodejs.org/api/cluster.html). Pero vamos a usar una mas sencilla llamada [Throng](https://github.com/hunterloftis/throng)

### Demo

1. Crear un grid en el browser de 8 x 8
2. Tenemos un array de longitud 8x8 (64) que contiene mensajes en JSON
3. Conectamos 1 socket por proceso desde el browser al backend, el backend on connection empieza a despachar
4. En el caso de node, vamos a tener que dividir este array entre el numero de procesos, no podemos compartir el socket connection asi que vamos a tener que abrir mas de un connection

## Golang

El modelo de concurrencia de Go esta basado en los Go routines, los go routines. Los go routines son funciones que se ejectutan concurrentemente manejadas por el Go routine, comparten el mismo espacio de memoria, asi que una go routine puede acceder a datos de su padre (closures).

Las goroutines son distribuidas en varios threads del OS, asi que si una se bloquea, las otras pueden seguir corriendo.

Tener este paralelismo y acceso a memoria compartida tambien tiene sus cosas divertidas, como data races.

```go
ejemplo de data race usando channels
```

Go tambien tiene formas de compartir data entre go routines, usando Channels.

```go 
var codes := []string{
  "#color1",
  "#color2",
  "#color3",
  "#color4",  
}

for _, value := range codes {
  log.Println(value)
}
```

```go 
var codes := []string{
  "#color1",
  "#color2",
  "#color3",
  "#color4",  
}

for _, value := range codes {
  go log.Println(value)
}
```

```go
ejemplo con channels
```

### Demo
1. Crear un grid en el browser de 8 x 8
2. Tenemos un array de longitud 8x8 (64) que contiene mensajes en JSON
3. Conectamos 1 socket por proceso del browser al backend, en el caso de go podemos compartir el socket entre varios go routines
4. Mandamos en cada ticket un mensaje por el socket, el browser renderea lo que dice el socket en el grid
5. cada step del array va a tener un delay de x segundos, para fakear que el proceso toma tiempo


## Resources

* https://vimeo.com/49718712 (Rob Pike - ‘Concurrency Is Not Parallelism’)
* Concurrency / Paralellism (http://stackoverflow.com/a/1050257/536984)
* Philip Roberts: What the heck is the event loop anyway? (https://www.youtube.com/watch?v=8aGhZQkoFbQ)
* JavaScript's Call Stack (https://egghead.io/lessons/javascript-call-stack)
