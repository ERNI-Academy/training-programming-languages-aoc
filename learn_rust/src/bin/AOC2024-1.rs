use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let content = fs::read_to_string("input.txt")?;
    let lines: Vec<&str> = content.lines().collect();

    let result = solve_part1(&lines);

    println!("{}", result);

    Ok(())
}

fn solve_part1(lines: &[&str]) -> i64 {
    let mut left: Vec<i64> = Vec::new();
    let mut right: Vec<i64> = Vec::new();

    for line in lines {
        let line = line.trim();
        if line.is_empty() {
            continue;
        }

        // separar por espacios múltiples
        let parts: Vec<&str> = line.split_whitespace().collect();
        let a: i64 = parts[0].parse().expect("número izquierdo inválido");
        let b: i64 = parts[1].parse().expect("número derecho inválido");

        left.push(a);
        right.push(b);
    }

    // ordenar listas
    left.sort();
    right.sort();

    // sumar distancias absolutas
    left.iter()
        .zip(right.iter())
        .map(|(a, b)| (a - b).abs())
        .sum()
}