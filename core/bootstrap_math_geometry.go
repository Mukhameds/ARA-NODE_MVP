// core/bootstrap_math_geometry.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathGeometry загружает геометрические понятия и сигналы формы, размера и пространства
func BootstrapMathGeometry(mem *MemoryEngine) {
	geometry := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Базовые элементы
		{"geo_point", "a point has no size and defines a position", []string{"math", "geometry", "point", "core", "bootstrap"}},
		{"geo_line", "a line is a straight path with infinite length", []string{"math", "geometry", "line", "core", "bootstrap"}},
		{"geo_segment", "a line segment has two endpoints", []string{"math", "geometry", "segment", "core", "bootstrap"}},
		{"geo_ray", "a ray starts at one point and extends infinitely in one direction", []string{"math", "geometry", "ray", "core", "bootstrap"}},

		// Углы и формы
		{"geo_angle", "an angle is formed by two rays with a common endpoint", []string{"math", "geometry", "angle", "core", "bootstrap"}},
		{"geo_triangle", "a triangle has three sides and three angles", []string{"math", "geometry", "triangle", "core", "bootstrap"}},
		{"geo_circle", "a circle is a set of points equidistant from a center", []string{"math", "geometry", "circle", "core", "bootstrap"}},
		{"geo_polygon", "a polygon is a closed figure with straight sides", []string{"math", "geometry", "polygon", "core", "bootstrap"}},

		// Размеры и расстояния
		{"geo_length", "length measures distance between two points", []string{"math", "geometry", "length", "core", "bootstrap"}},
		{"geo_area", "area is the size of a surface", []string{"math", "geometry", "area", "core", "bootstrap"}},
		{"geo_perimeter", "perimeter is the distance around a figure", []string{"math", "geometry", "perimeter", "core", "bootstrap"}},
		{"geo_volume", "volume is the space an object occupies", []string{"math", "geometry", "volume", "core", "bootstrap"}},

		// Координаты
		{"geo_coord_plane", "a coordinate plane defines position using (x, y)", []string{"math", "geometry", "coordinate", "core", "bootstrap"}},
		{"geo_origin", "the origin is the point (0,0) in the coordinate plane", []string{"math", "geometry", "origin", "core", "bootstrap"}},
		{"geo_quadrants", "the plane is divided into four quadrants", []string{"math", "geometry", "quadrant", "core", "bootstrap"}},

		// Свойства и теоремы
		{"geo_right_angle", "a right angle measures 90 degrees", []string{"math", "geometry", "angle", "core", "bootstrap"}},
		{"geo_pythagorean", "a² + b² = c² in a right triangle", []string{"math", "geometry", "theorem", "pythagorean", "core", "bootstrap"}},
	}

	for _, g := range geometry {
		q := QBit{
			ID:        g.ID,
			Content:   g.Content,
			Tags:      g.Tags,
			Phase:     0.88,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📐 [Bootstrap] Geometry concepts and spatial logic loaded.")
}
