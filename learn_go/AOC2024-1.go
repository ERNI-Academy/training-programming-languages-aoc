package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    // Abrir input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("No se pudo abrir input.txt: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var left []int
    var right []int

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }

        // Dividir la línea por espacios (uno o más)
        parts := strings.Fields(line)
        if len(parts) < 2 {
            log.Fatalf("Línea inválida (esperaba dos columnas): %q", line)
        }

        a, err := strconv.Atoi(parts[0])
        if err != nil {
            log.Fatalf("Valor inválido en columna izquierda '%s': %v", parts[0], err)
        }

        b, err := strconv.Atoi(parts[1])
        if err != nil {
            log.Fatalf("Valor inválido en columna derecha '%s': %v", parts[1], err)
        }

        left = append(left, a)
        right = append(right, b)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error leyendo el fichero: %v", err)
    }

    // Ordenar ambas listas
    sort.Ints(left)
    sort.Ints(right)

    // Sumar distancias absolutas
    var total int

    for i := range left {
        diff := int(math.Abs(float64(left[i] - right[i])))
        total += diff
    }

    // Imprimir SOLO el resultado (como pide Advent of Code)
    fmt.Println(total)
}
