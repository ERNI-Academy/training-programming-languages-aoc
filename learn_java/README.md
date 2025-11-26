# INSTALLING.md



## Instalación completa de Java + Visual Studio Code + Hello World



Este documento contiene **todos los pasos necesarios**, desde cero, para instalar Java, configurar Visual Studio Code, crear un proyecto y ejecutar un programa *Hello World*.



---



# 1. Instalar Java (JDK)



## 1.1 Descargar e instalar Temurin (Eclipse Adoptium)



1. Abrir un navegador e ir a: \[https://adoptium.net/](https://adoptium.net/)

2. Seleccionar **Temurin 21 (LTS)**.

3. Descargar la versión correspondiente a tu sistema operativo.

4. Ejecutar el instalador.

5. Aceptar todas las opciones por defecto hasta finalizar.



## 1.2 Verificar instalación



Abrir una terminal y ejecutar:



```bash

java -version

```



Debe aparecer algo similar a:



```

openjdk version "21" ...

```



Esto confirma que Java está instalado correctamente.



---



# 2. Instalar Visual Studio Code



## 2.1 Descargar VS Code



1. Ir a: [https://code.visualstudio.com/](https://code.visualstudio.com/)

2. Descargar la versión para tu sistema operativo.

3. Ejecutar el instalador.

4. Aceptar las opciones por defecto.



---



# 3. Instalar soporte completo para Java en VS Code



## 3.1 Abrir Visual Studio Code



Simplemente ejecuta la aplicación instalada.



## 3.2 Instalar el **Extension Pack for Java**



1. En la barra lateral izquierda, hacer clic en el icono de **Extensiones** (cuadrado de cuatro puntos).

2. En la barra de búsqueda, escribir:

&nbsp;  **Extension Pack for Java**

3. Hacer clic en **Install**.

4. Esperar que finalice la instalación.



El paquete instala:



* Herramientas de lenguaje Java

* Depuración

* Soporte para proyectos

* Funciones automáticas de compilación y ejecución



---



# 4. Crear el proyecto Hello World



## 4.1 Crear la carpeta del proyecto



1. Crear manualmente una carpeta llamada **hello** en cualquier ubicación.



## 4.2 Abrir la carpeta en VS Code



1. En VS Code, ir a: **File → Open Folder…**

2. Seleccionar la carpeta `hello`.

3. Confirmar la apertura.



## 4.3 Crear el archivo de código



1. En la barra lateral de archivos, clic derecho → **New File**

2. Nombre: **Hello.java**

3. Pegar este contenido:



```java

public class Hello {

&nbsp;   public static void main(String[] args) {

&nbsp;       System.out.println("Hello, world!");

&nbsp;   }

}

```



Guardar el archivo.



---



# 5. Ejecutar el Hello World dentro de VS Code



## 5.1 Ejecutar el programa



1. Abrir el archivo `Hello.java`.

2. En la parte superior del editor, aparecerá un botón **Run** (▶️).

3. Hacer clic en **Run**.



## 5.2 Ver salida



VS Code abrirá una terminal integrada y mostrará:



```

Hello, world!

```



---



# 6. Ejecutar el programa desde la terminal (alternativa)



Con la carpeta `hello` abierta en una terminal:



## 6.1 Compilar



```bash

javac Hello.java

```



Esto generará un archivo `Hello.class`.



## 6.2 Ejecutar



```bash

java Hello

```



Debería mostrar:



```

Hello, world!

```



---



# 7. Instalación y ejecución completadas



Tu entorno Java está ya configurado, Visual Studio Code está preparado

con soporte completo para Java, y tu primer programa Hello World está compilado y funcionando.



Fin del proceso.



