# STRUCTURE_RUST.md

## Estructura oficial de un proyecto Rust (Cargo)

Rust usa **Cargo** como sistema de construcción y gestor de dependencias. Cargo define una estructura **fija** y **muy consistente** para que todos los proyectos Rust sean predecibles, fáciles de mantener y compatibles con todas las herramientas del ecosistema.

Este documento explica **cómo debe organizarse un proyecto Rust**, qué significa cada carpeta y cómo ejecutar múltiples programas dentro del mismo proyecto.

---

# 1. Estructura mínima de un proyecto

Al crear un proyecto con:

```bash
cargo new mi_proyecto
```

Se genera:

```
mi_proyecto/
 ├── Cargo.toml
 └── src/
     └── main.rs
```

### `Cargo.toml`

Archivo de configuración del proyecto.

### `src/main.rs`

Punto de entrada del **ejecutable principal**. Si haces:

```bash
cargo run
```

se ejecuta este archivo.

---

# 2. Múltiples ejecutables

Cargo permite tener más de un programa dentro del mismo proyecto.

## ✔ Para crear ejecutables adicionales:

Crea archivos dentro de:

```
src/bin/
```

Ejemplo:

```
src/bin/ejercicio1.rs
src/bin/ejercicio2.rs
```

Cada archivo es un **binario independiente** con su propia función `main()`.

## Ejecutar un binario específico:

```bash
cargo run --bin ejercicio1
```

Otro ejemplo:

```bash
cargo run --bin ejercicio2
```

Cargo detecta automáticamente todos los archivos dentro de `src/bin/`.

---

# 3. Librerías en Rust

Si quieres crear una **librería** reutilizable dentro del proyecto:

```
src/lib.rs
```

Esto te permite definir funciones o módulos comunes a varios binarios.

Puedes importar la librería en tus ejecutables así:

```rust
use mi_proyecto::*;
```

---

# 4. Carpetas especiales

Rust reconoce varias carpetas con significado especial:

```
tests/      → tests de integración
examples/   → ejemplos ejecutables
benches/    → benchmarks
```

Ejemplo en `examples/`:

```
examples/demo.rs
```

Ejecutar:

```bash
cargo run --example demo
```

---

# 5. Resumen visual

```
mi_proyecto/
 ├── Cargo.toml         ← configuración
 └── src/
      ├── main.rs       ← ejecutable principal
      ├── lib.rs        ← librería (opcional)
      └── bin/          ← ejecutables adicionales
           ├── ej1.rs
           ├── ej2.rs
           └── aoc1.rs
 ├── tests/             ← tests de integración
 ├── examples/          ← ejecutables de ejemplo
 └── benches/           ← benchmarks
```

---

# 6. Idea clave

Rust es **muy estricto** con la estructura, pero eso evita:

* conflictos entre múltiples `main()`
* estructuras caóticas de carpetas
* comportamientos inconsistentes

A cambio te da:

* proyectos predecibles
* herramientas que siempre funcionan
* soporte excelente en IDEs (VS Code + rust-analyzer)
* compilaciones limpias y reproducibles



