package fizzbuzz

import (
        "testing"
        "os"
)
func TestCheckNumber(t *testing.T) {
        Print()
        var b = make([]byte, 5)
        if _, err := os.Stdout.Read(b); err != nil {
                t.Fatalf("First: %v", err)
        }
        printed1 := string(b[0]) == "1"
        if !printed1 {
                t.Errorf("Printed: %q, want: %v", string(b[0]), "1")
        }
}
