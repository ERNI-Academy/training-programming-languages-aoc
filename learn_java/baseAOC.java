import java.nio.file.*;
import java.io.IOException;
import java.util.List;

public class FileProcessor {

    public static void main(String[] args) {
        try {
            // 1. Cargar todas las líneas del fichero
            List<String> lines = Files.readAllLines(Path.of("input.txt"));

            // 2. Procesar las líneas para obtener un resultado numérico
            int result = processLines(lines);

            // 3. Mostrar el resultado numérico por pantalla
            System.out.println("Resultado del procesamiento: " + result);

        } catch (IOException e) {
            System.out.println("Error leyendo el fichero: " + e.getMessage());
        }
    }

    /**
     * Procesa las líneas del fichero y devuelve un resultado numérico.
     */
    private static int processLines(List<String> lines) {
         List<Integer> left = new ArrayList<>();
        List<Integer> right = new ArrayList<>();

        // Parsear cada línea en dos columnas
        for (String line : lines) {
            if (line.isBlank()) {
                continue; // por si hay líneas vacías
            }

            // Dividir por espacios (uno o más)
            String[] parts = line.trim().split("\\s+");
            int leftValue = Integer.parseInt(parts[0]);
            int rightValue = Integer.parseInt(parts[1]);

            left.add(leftValue);
            right.add(rightValue);
        }

        // Ordenar ambas listas
        Collections.sort(left);
        Collections.sort(right);

        // Sumar las distancias absolutas
        long result = 0L;
        for (int i = 0; i < left.size(); i++) {
            int a = left.get(i);
            int b = right.get(i);
            result += Math.abs(a - b);
        }

        return result;
    }
}
