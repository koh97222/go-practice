package main

// Rectangle 四角形の構造体
type Rectangle struct {
	width  float64
	height float64
}

// Perimeter Perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Area 四角形の面積
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
