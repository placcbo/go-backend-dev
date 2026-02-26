package main

import (
    "errors"
    "fmt"
)

// ── Sentinel errors (named error values) ─────────────────────────
var (
    ErrDivisionByZero  = errors.New("division by zero")
    ErrUnknownOperator = errors.New("unknown operator")
)

// ── Calculation holds one completed calculation ───────────────────
type Calculation struct {
    A, B       float64
    Op         string
    Result     float64
    Expression string
}

// ── calculate performs a single operation ─────────────────────────
func calculate(a float64, op string, b float64) (float64, error) {
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, fmt.Errorf("%.0f / 0: %w", a, ErrDivisionByZero)
        }
        return a / b, nil
    default:
        return 0, fmt.Errorf("operator %q: %w", op, ErrUnknownOperator)
    }
}

// ── printHistory prints all past calculations ─────────────────────
func printHistory(history []Calculation) {
    fmt.Println("\n── Calculation History ───────────────────────")
    if len(history) == 0 {
        fmt.Println("  No calculations yet.")
        return
    }
    for i, c := range history {
        fmt.Printf("  [%02d] %s\n", i+1, c.Expression)
    }
}

// ── stats computes and prints aggregate stats ─────────────────────
func stats(history []Calculation) {
    fmt.Println("\n── Stats ─────────────────────────────────────")
    if len(history) == 0 {
        fmt.Println("  No data to compute.")
        return
    }

    // seed min/max with the first result
    min := history[0].Result
    max := history[0].Result
    sum := 0.0

    for _, c := range history {
        sum += c.Result
        if c.Result < min { min = c.Result }
        if c.Result > max { max = c.Result }
    }

    fmt.Printf("  Calculations : %d\n", len(history))
    fmt.Printf("  Sum of results: %.2f\n", sum)
    fmt.Printf("  Average result: %.2f\n", sum/float64(len(history)))
    fmt.Printf("  Highest result: %.2f\n", max)
    fmt.Printf("  Lowest result : %.2f\n", min)
}

func main() {
    // ── Define all the operations to run ──────────────────────────
    type op struct { a float64; operator string; b float64 }

    operations := []op{
        {10, "+", 5},
        {20, "-", 8},
        {6,  "*", 7},
        {100, "/", 4},
        {9,  "/", 0},   // division by zero — should error
        {15, "^", 2},   // unknown operator — should error
        {50, "+", 50},
        {1000, "-", 999},
    }

    var history []Calculation    // nil slice — safe to append to

    fmt.Println("── Running Calculations ──────────────────────")

    for _, o := range operations {
        result, err := calculate(o.a, o.operator, o.b)
        if err != nil {
            fmt.Printf("  ❌ ERROR: %.0f %s %.0f → %v\n", o.a, o.operator, o.b, err)
            continue  // skip failed calculations — don't add to history
        }
        expr := fmt.Sprintf("%.2f %s %.2f = %.2f", o.a, o.operator, o.b, result)
        fmt.Printf("  ✅ %s\n", expr)
        history = append(history, Calculation{
            A: o.a, B: o.b, Op: o.operator,
            Result: result, Expression: expr,
        })
    }

    printHistory(history)
    stats(history)
}