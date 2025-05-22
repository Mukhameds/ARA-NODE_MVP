// core/bootstrap_math_linear_algebra.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathLinearAlgebra загружает ключевые понятия линейной алгебры
func BootstrapMathLinearAlgebra(mem *MemoryEngine) {
	linear := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Векторы
		{"la_vector", "a vector is a quantity with both magnitude and direction", []string{"math", "linear-algebra", "vector", "core", "bootstrap"}},
		{"la_vector_2d", "a 2D vector is represented as (x, y)", []string{"math", "linear-algebra", "vector", "core", "bootstrap"}},
		{"la_vector_add", "vector addition combines two vectors component-wise", []string{"math", "linear-algebra", "vector", "operation", "core", "bootstrap"}},
		{"la_vector_scalar_mul", "scalar multiplication stretches or shrinks a vector", []string{"math", "linear-algebra", "vector", "operation", "core", "bootstrap"}},

		// Матрицы
		{"la_matrix", "a matrix is a rectangular array of numbers", []string{"math", "linear-algebra", "matrix", "core", "bootstrap"}},
		{"la_matrix_size", "a matrix has size rows × columns", []string{"math", "linear-algebra", "matrix", "size", "core", "bootstrap"}},
		{"la_matrix_mult", "matrix multiplication combines rows and columns", []string{"math", "linear-algebra", "matrix", "operation", "core", "bootstrap"}},
		{"la_matrix_identity", "identity matrix leaves vectors unchanged", []string{"math", "linear-algebra", "matrix", "identity", "core", "bootstrap"}},

		// Системы уравнений
		{"la_linear_system", "a linear system is a set of linear equations", []string{"math", "linear-algebra", "system", "core", "bootstrap"}},
		{"la_solution_system", "solving a linear system finds variable values that satisfy all equations", []string{"math", "linear-algebra", "solution", "core", "bootstrap"}},

		// Определитель и инверсия
		{"la_determinant", "the determinant is a scalar value describing a matrix's transformation", []string{"math", "linear-algebra", "determinant", "core", "bootstrap"}},
		{"la_matrix_inverse", "an inverse matrix undoes the transformation of the original", []string{"math", "linear-algebra", "inverse", "core", "bootstrap"}},

		// Базис и пространство
		{"la_basis", "a basis is a minimal set of vectors that span a space", []string{"math", "linear-algebra", "basis", "core", "bootstrap"}},
		{"la_dimension", "dimension is the number of basis vectors in a space", []string{"math", "linear-algebra", "dimension", "core", "bootstrap"}},
	}

	for _, l := range linear {
		q := QBit{
			ID:        l.ID,
			Content:   l.Content,
			Tags:      l.Tags,
			Phase:     0.88,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📊 [Bootstrap] Linear algebra concepts loaded.")
}
