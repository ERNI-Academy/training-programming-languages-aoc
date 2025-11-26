var path = "Input/Aoc2024Day01.txt";

string[] lines = await File.ReadAllLinesAsync(path);
int[] leftLines = new int[lines.Length];
int[] rightLines = new int[lines.Length];
for (int i = 0; i < lines.Length; i++)
{
    var parts = lines[i].Split("  ");
    if (parts.Length >= 2)
    {
        leftLines[i] = int.Parse(parts[0]);
        rightLines[i] = int.Parse(parts[1]);
    }
}

SolvePart1(leftLines, rightLines);

SolvePart2(leftLines, rightLines);

static void SolvePart1(int[] leftLines, int[] rightLines)
{
    var totalDistance = 0;
    Array.Sort(leftLines);
    Array.Sort(rightLines);

    for (int i = 0; i < leftLines.Length; i++)
    {
        totalDistance += Math.Abs(leftLines[i] - rightLines[i]);
    }

    Console.WriteLine($"Day 1 part 1: {totalDistance}");
}

static void SolvePart2(int[] leftLines, int[] rightLines)
{
    var totalSimilarityScore = 0;
    Dictionary<int, int> rightLineCounts = [];

    for (int i = 0; i < rightLines.Length; i++)
    {
        rightLineCounts[rightLines[i]] = rightLineCounts.ContainsKey(rightLines[i]) ? rightLineCounts[rightLines[i]] + 1 : 1;
    }

    for (int i = 0; i < leftLines.Length; i++)
    {
        var similarityScore = rightLineCounts.ContainsKey(leftLines[i]) ? rightLineCounts[leftLines[i]] * leftLines[i] : 0;
        totalSimilarityScore += similarityScore;
    }

    Console.WriteLine($"Day 1 part 2: {totalSimilarityScore}");
}