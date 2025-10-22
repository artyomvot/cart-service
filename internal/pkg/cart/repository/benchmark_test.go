package repository

import "testing"

func BenchmarkAddProduct(b *testing.B) {
	repo := NewCartRepository(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.AddProduct(1, int64(i), 1)
	}

}
func BenchmarkDeleteProduct(b *testing.B) {
	repo := NewCartRepository(100)

	for i := 0; i < b.N; i++ {
		repo.AddProduct(1, int64(i), 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.DeleteProduct(1, int64(i))
	}
}
