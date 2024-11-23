package main

import "testing"

func TestCalculateDiscount(t *testing.T) {
    tests := []struct {
        name           string
        originalPrice  float64
        discountPercent float64
        expected       float64
    }{
        {"Valid discount 10%", 100, 10, 90},
        {"Valid discount 50%", 200, 50, 100},
        {"No discount", 150, 0, 150},
        {"Full discount 100%", 100, 100, 0},
        {"Invalid discount negative", 100, -10, -1},
        {"Invalid discount over 100%", 100, 110, -1},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := CalculateDiscount(tc.originalPrice, tc.discountPercent)
            if result != tc.expected {
                t.Errorf("Expected %f, got %f", tc.expected, result)
            }
        })
    }
}
