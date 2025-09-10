El parser **XJS** no soporta los operadores `===/!==` y transforma los operadores `==/!=` a `===/!===`. Este plugin añade soporte a la desigualdad débil, transformando `~~/!~` a `==/!=`.

Por ejemplo:

```js
a == b // se traduce a `a === b`
a != b // se traduce a `a !== b`
a === b // error! para XJS todas las comparaciones son estrictas y no entiende esto
a !== b // error! para XJS todas las comparaciones son estrictas y no entiende esto
```

Mediante este plugin podemos "enriquecer" el parser y escribir:

```js
a ~~ b // se traduce a `a == b`
a !~ b // se traduce a `a != b`
```