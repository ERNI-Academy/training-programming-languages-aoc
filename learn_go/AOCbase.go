package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// Este programa lee un fichero input.txt, procesa sus líneas
// y devuelve un resultado numérico por pantalla.
func main() {
    // 1. Abrir el fichero
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("No se pudo abrir input.txt: %v", err)
    }
    defer file.Close()

    // 2. Leer líneas una a una
    scanner := bufio.NewScanner(file)

    var numbers []int

    for scanner.Scan() {
        line := scanner.Text()
        // Ignorar líneas vacías
        if line == "" {
            continue
        }

        // Ejemplo de parseo numérico (puedes cambiar esta lógica luego)
        n, err := strconv.Atoi(line)
        if err != nil {
            log.Fatalf("No se pudo convertir '%s' a número: %v", line, err)
        }
        numbers = append(numbers, n)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error leyendo el fichero: %v", err)
    }

    // 3. Procesar las líneas (aquí va la lógica del ejercicio)
    result := processNumbers(numbers)

    // 4. Mostrar el resultado
    fmt.Println("Resultado del procesamiento:", result)
}

// processNumbers es el lugar donde irá la lógica del ejercicio.
// De momento es solo un esqueleto.
func processNumbers(nums []int) int {
    result := 0

    for _, n := range nums {
        // TODO: lógica real
        // Ejemplo placeholder: sumar todos
        // result += n

        _ = n // evitar warning mientras está TODO
    }

    return result
}
