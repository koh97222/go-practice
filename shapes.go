package main

import "math"

// Rectangle 四角形の構造体
type Rectangle struct {
	width  float64
	height float64
}

// Circle 円の構造体
type Circle struct {
	radius float64
}

// Triangle 三角形の構造体
type Triangle struct {
	Base   float64
	Height float64
}

// Shape 図形のインターフェース
// インターフェースを宣言することにより、
// ヘルパーは具象型から切り離（Decoupling）され、その機能を実行するために必要なメソッドのみを持ちます。
type Shape interface {
	Area() float64
}

// Area 四角形の面積
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Area 円の面積
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area 三角形の面積
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter Perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Area 四角形の面積
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
