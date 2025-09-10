# Weak Equality Parser Plugin for XJS

The **XJS** parser does not support the `===` or `!==` operators and automatically transforms `==` and `!=` into `===` and `!==`. This plugin adds support for weak equality by converting `~~` and `!~` into `==` and `!=`, respectively.

**Examples:**

```js
a == b   // Translated to: a === b
a != b   // Translated to: a !== b
a === b  // Error! XJS only supports strict comparisons.
a !== b  // Error! XJS only supports strict comparisons.
```

With this plugin, you can extend the parser to write:

```js
a ~~ b   // Translated to: a == b
a !~ b   // Translated to: a != b
```