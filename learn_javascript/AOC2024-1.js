const fs = require("fs");
const path = require("path");

// --- Input file management via command line argument ---

const file = process.argv[2];

if (!file) {
  console.error("You must provide an input file as argument. Example: ");
  console.error("   node day1.js input.txt");
  process.exit(1);
}

const inputPath = path.join(__dirname, file);

if (!fs.existsSync(inputPath)) {
  console.error(`File does not exist: ${inputPath}`);
  process.exit(1);
}

const rawInput = fs.readFileSync(inputPath, "utf8");
const { left, right } = parseInput(rawInput);

const part1 = totalDistance(left, right);
const part2 = similarityScore(left, right);

console.log("Part 1 - Total distance:", part1);
console.log("Part 2 - Similarity score:", part2);


/**
 * Parses the puzzle input.
 * Each line has two numbers separated by one or more spaces.
 */
function parseInput(input) {
  const left = [];
  const right = [];

  const lines = input
    .trim()
    .split("\n")
    .map((l) => l.trim())
    .filter((l) => l.length > 0);

  for (const line of lines) {
    const parts = line.split(/\s+/);
    left.push(Number(parts[0]));
    right.push(Number(parts[1]));
  }

  return { left, right };
}

/**
 * PART 1
 * Calculates the total distance between the two lists:
 * - Sorts each list separately
 * - Sums |left[i] - right[i]| for each i
 */
function totalDistance(left, right) {
  const sortedLeft = [...left].sort((a, b) => a - b);
  const sortedRight = [...right].sort((a, b) => a - b);

  let sum = 0;
  for (let i = 0; i < sortedLeft.length; i++) {
    sum += Math.abs(sortedLeft[i] - sortedRight[i]);
  }

  return sum;
}

/**
 * PART 2
 * Calculates the "similarity score":
 * - Counts how many times each number in the right list appears
 * - For each number n in left, sums n * (times it appears in right)
 */
function similarityScore(left, right) {
  const freq = new Map();

  for (const n of right) {
    freq.set(n, (freq.get(n) || 0) + 1);
  }

  let score = 0;
  for (const n of left) {
    const count = freq.get(n) || 0;
    score += n * count;
  }

  return score;
}

