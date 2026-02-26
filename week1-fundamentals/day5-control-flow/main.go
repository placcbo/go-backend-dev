package main

import "fmt"

type Student struct {
    Name   string
    Scores []int
}

// ── Average of a score list ───────────────────────────────────────
func average(scores []int) float64 {
    if len(scores) == 0 {
        return 0.0
    }
    sum := 0
    for _, s := range scores {
        sum += s
    }
    return float64(sum) / float64(len(scores))
}

// ── Grade from average (expression switch) ────────────────────────
func grade(avg float64) string {
    switch {
    case avg >= 90:
        return "A"
    case avg >= 80:
        return "B"
    case avg >= 70:
        return "C"
    case avg >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
    students := []Student{
        {"Alice",   []int{92, 88, 95, 91}},
        {"Bob",     []int{74, 68, 72, 80}},
        {"Carol",   []int{55, 60, 58, 52}},
        {"Dave",    []int{83, 79, 86, 90}},
        {"Eve",     []int{100, 97, 98}},
    }

    // ── Print report card ─────────────────────────────────────────
    fmt.Println("╔══════════════════════════════════╗")
    fmt.Println("║         REPORT CARD              ║")
    fmt.Println("╠══════════════════════════════════╣")
    fmt.Printf("║  %-12s %8s %6s   ║\n", "Name", "Average", "Grade")
    fmt.Println("╠══════════════════════════════════╣")

    var totalAvg float64
    topAvg := 0.0
    topName := ""

    for _, s := range students {
        avg := average(s.Scores)
        g := grade(avg)
        fmt.Printf("║  %-12s %7.2f%% %6s   ║\n", s.Name, avg, g)
        totalAvg += avg
        if avg > topAvg {
            topAvg = avg
            topName = s.Name
        }
    }

    classAvg := totalAvg / float64(len(students))
    fmt.Println("╠══════════════════════════════════╣")
    fmt.Printf("║  Class Average:      %7.2f%%       ║\n", classAvg)
    fmt.Printf("║  Top scorer: %-12s (%s)  ║\n", topName, grade(topAvg))
    fmt.Println("╚══════════════════════════════════╝")
}
