# INSTALLING_RUST.md

## Instalación completa de Rust + VS Code + Hello World

Este documento contiene **todos los pasos necesarios**, desde cero, para instalar Rust, configurar Visual Studio Code, crear un proyecto y ejecutar un programa *Hello World*.

---

# 1. Instalar Rust

La forma oficial y recomendada de instalar Rust es mediante **rustup**.

## 1.1 Instalar en Linux o macOS

Ejecutar en la terminal:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## 1.2 Instalar en Windows

1. Ir a: [https://rustup.rs](https://rustup.rs)
2. Descargar el ejecutable.
3. Ejecutarlo y seguir las instrucciones.

## 1.3 Verificar instalación

Cerrar y abrir la terminal, luego ejecutar:

```bash
rustc --version
cargo --version
```

Deberías ver versiones recientes de Rust y Cargo.

---

# 2. Instalar Visual Studio Code

## 2.1 Descargar VS Code

1. Ir a: [https://code.visualstudio.com/](https://code.visualstudio.com/)
2. Descargar la versión adecuada.
3. Instalar con valores por defecto.

## 2.2 Instalar soporte para Rust en VS Code

1. Abrir VS Code.
2. Ir al menú de Extensiones.
3. Buscar: **rust-analyzer**.
4. Instalar.

Esta extensión proporciona autocompletado, análisis, errores en tiempo real y ejecución integrada.

---

# 3. Crear el proyecto Hello World

Rust usa `cargo`, su sistema de build y gestión de dependencias.

## 3.1 Crear un nuevo proyecto

```bash
cargo new hello_rust
cd hello_rust
```

Esto crea:

* `Cargo.toml`
* `src/main.rs`

## 3.2 Código de Hello World (ya creado por defecto)

`src/main.rs`:

```rust
fn main() {
    println!("Hello, world!");
}
```

## 3.3 Ejecutar el proyecto

```bash
cargo run
```

Salida esperada:

```
Hello, world!
```

---

# 4. Estructura del proyecto

* `Cargo.toml` → archivo de configuración del proyecto.
* `src/main.rs` → punto de entrada.

Cargo se encarga de:

* Compilar
* Resolver dependencias
* Ejecutar
* Formatear código (`cargo fmt`)
* Analizar código (`cargo clippy`)

---

# 5. Proyecto listo

El entorno Rust está instalado, VS Code tiene soporte completo, y tu primer programa *Hello World* ha sido ejecutado correctamente.
