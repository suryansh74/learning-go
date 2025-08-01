package learninggo

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestSub(t *testing.T) {
    result := Sub(2, 3)
    expected := -1

    if result != expected {
        t.Errorf("Sub(2, 3) = %d; want %d", result, expected)
    }
}

func TestMul(t *testing.T) {
    result := Mul(2, 3)
    expected := 6

    if result != expected {
        t.Errorf("Mul(2, 3) = %d; want %d", result, expected)
    }
}
func TestDiv(t *testing.T) {
    result := Div(2, 3)
    expected := 0

    if result != expected {
        t.Errorf("Div(2, 3) = %d; want %d", result, expected)
    }
}