# INSTALLING_GO.md

## Instalación completa de Go + VS Code + Hello World

Este documento contiene **todos los pasos necesarios**, desde cero, para instalar Go, configurar Visual Studio Code, crear un proyecto y ejecutar un programa *Hello World*.

---

# 1. Instalar Go

Go se instala mediante su distribución oficial.

## 1.1 Descargar Go

1. Ir a: [https://go.dev/dl/](https://go.dev/dl/)
2. Descargar la versión para tu sistema operativo (Linux, macOS o Windows).
3. Ejecutar el instalador.
4. Aceptar opciones por defecto.

## 1.2 Verificar instalación

Cerrar y abrir la terminal, luego ejecutar:

```bash
go version
```

Debería mostrar algo como:

```
go version go1.22.0 linux/amd64
```

---

# 2. Instalar Visual Studio Code

## 2.1 Descargar VS Code

1. Ir a: [https://code.visualstudio.com/](https://code.visualstudio.com/)
2. Descargar la versión adecuada.
3. Instalar.

## 2.2 Instalar extensión oficial de Go

1. Abrir VS Code.
2. Ir a extensiones.
3. Buscar: **Go (Go Team at Google)**.
4. Instalar.

Esta extensión ofrece autocompletado, debug, analizador estático y gestión de módulos.

---

# 3. Crear el proyecto Hello World

Go utiliza una estructura simple basada en módulos.

## 3.1 Crear un nuevo directorio para el proyecto

```bash
mkdir hello_go
cd hello_go
```

## 3.2 Inicializar el módulo Go

```bash
go mod init hello_go
```

Esto genera un archivo:

```
go.mod
```

## 3.3 Crear el archivo principal

Crear `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

## 3.4 Ejecutar el programa

```bash
go run .
```

Salida esperada:

```
Hello, world!
```

---

# 4. Estructura del proyecto Go

Un proyecto mínimo luce así:

```
hello_go/
 ├── go.mod
 └── main.go
```

Go no requiere carpetas rígidas como Rust; en su lugar utiliza módulos.

---

# 5. Ejecutar un fichero que no sea el principal

Go permite ejecutar directamente un fichero específico:

```bash
go run ejercicio1.go
```

O múltiples ficheros a la vez:

```bash
go run ejercicio1.go utilidades.go
```

Go **no restringe ubicaciones**: puedes organizar tus ejercicios así:

```
.
├── go.mod
├── main.go
├── ejercicio1.go
├── ejercicio2.go
└── ejercicios/
    ├── e1.go
    ├── e2.go
    └── e3.go
```

Y ejecutarlos así:

```bash
go run ejercicios/e1.go
```

---

# 6. Recomendación para cursos y ejercicios

Go permite múltiples ficheros ejecutables sin estructura rígida.

Organización recomendada:

```
hello_go/
 ├── go.mod
 ├── main.go               ← hello world
 ├── lectura.go            ← ejercicio leer fichero
 ├── aoc2024_day1.go       ← advent of code
 └── ejercicios/
      ├── ej1.go
      ├── ej2.go
      └── ej3.go
```

Ejecutar cualquier ejercicio:

```bash
go run lectura.go
```

O:

```bash
go run ejercicios/ej2.go
```

---

# 7. Proyecto Go listo

Has instalado Go, configurado VS Code y creado un proyecto completamente funcional con soporte para múltiples ejercicios ejecutables.
