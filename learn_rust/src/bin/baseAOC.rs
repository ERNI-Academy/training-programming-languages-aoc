use std::fs;
use std::io;

fn main() -> io::Result<()> {
    // 1. Leer todo el fichero como String
    let content = fs::read_to_string("input.txt")?;

    // 2. Convertir en líneas
    let lines: Vec<&str> = content.lines().collect();

    // 3. Procesar líneas y obtener un número
    let result = process_lines(&lines);

    // 4. Mostrar resultado
    println!("Resultado del procesamiento: {}", result);

    Ok(())
}

/// Procesa las líneas y devuelve un resultado numérico.
/// Aquí es donde luego metes la lógica que quieras (TBD).
fn process_lines(lines: &[&str]) -> i64 {
    let mut result: i64 = 0;

    for line in lines {
        let line = line.trim();
        if line.is_empty() {
            continue;
        }

        // TODO: lógica concreta
        // Ejemplo de plantilla:
        // if let Ok(value) = line.parse::<i64>() {
        //     result += value;
        // }

    }

    result
}