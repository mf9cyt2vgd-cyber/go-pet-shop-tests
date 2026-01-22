package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-pet-shop/internal/handlers/product/mocks"
	"go-pet-shop/internal/models"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	mock2 "github.com/stretchr/testify/mock"
)

// Get Product - Ready
func TestGetAllProducts_Success(t *testing.T) {
	// Мокаем storage — он вернёт один продукт.
	mock := mocks.NewProducts(t)
	mock.On("GetAllProducts", mock2.Anything).Return([]models.Product{
		{ID: 1, Name: "Dog Food"},
	}, nil)

	// Создаем HTTP-запрос GET /products
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	// Создаем хендлер с мок-хранилищем
	handler := New(slog.Default(), mock)

	// Вызываем метод GetAllProducts, который является http.HandlerFunc
	handler.GetAllProducts(w, req)

	// Проверяем HTTP-код
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
func TestGetAllProducts_Error(t *testing.T) {
	// Мокаем storage — он будет возвращать ошибку
	//mock := &ProductsMock{
	//	GetAllProductsFunc: func(ctx context.Context) ([]models.Product, error) {
	//		return nil, errors.New("DB error")
	//	},
	//}
	mock := mocks.NewProducts(t)
	mock.On("GetAllProducts", mock2.Anything).Return(nil, errors.New("DB error"))
	// Создаем запрос
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.GetAllProducts(w, req)

	// Ожидаем HTTP 500
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Create Product
// =======================

func TestCreateProduct_Success(t *testing.T) {
	// TODO: Написать Unit-тест для создания продукта (200 OK)
	//mock := &ProductsMock{CreateProductFunc: func(ctx context.Context, product models.Product) (int, error) {
	//	return 1, nil
	//},
	//}
	mock := mocks.NewProducts(t)
	mock.On("CreateProduct", mock2.Anything, mock2.Anything).Return(1, nil)
	body, _ := json.Marshal(map[string]any{
		"name":  "Apple",
		"price": 500,
		"stock": 1,
	})
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)
	handler.CreateProduct(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestCreateProduct_BadRequest(t *testing.T) {
	// TODO: Написать Unit-тест для создания продукта с невалидным JSON (400 Bad Request)
	//mock := &ProductsMock{CreateProductFunc: func(ctx context.Context, product models.Product) (int, error) {
	//	t.Error("Mock should not be called for invalid JSON")
	//	return 0, nil
	//}}
	mock := mocks.NewProducts(t)
	body := []byte("{invalid json}")
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)
	handler.CreateProduct(w, req)
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	mock.AssertNotCalled(t, "CreateProduct", mock2.Anything, mock2.Anything)
}

func TestCreateProduct_Fail(t *testing.T) {
	// TODO: Написать Unit-тест для создания продукта при ошибке сервиса (500 Internal Server Error)
	//mock := &ProductsMock{CreateProductFunc: func(ctx context.Context, product models.Product) (int, error) {
	//	return 0, errors.New("failed to create product")
	//}}
	mock := mocks.NewProducts(t)
	mock.On("CreateProduct", mock2.Anything, mock2.Anything).Return(0, errors.New("failed to create product"))
	body, _ := json.Marshal(map[string]any{
		"name":  "Apple",
		"price": 500,
		"stock": 1,
	})
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)
	handler.CreateProduct(w, req)
	if w.Code != 500 {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Update Product
// =======================

func TestUpdateProduct_Success(t *testing.T) {
	// TODO: Написать Unit-тест для обновления продукта (200 OK)\
	//mock := &ProductsMock{UpdateProductFunc: func(ctx context.Context, product models.Product) error {
	//	return nil
	//}}
	mock := mocks.NewProducts(t)
	mock.On("UpdateProduct", mock2.Anything, mock2.Anything).Return(nil)
	r := chi.NewRouter()
	body, _ := json.Marshal(map[string]any{
		"name":  "Apple",
		"price": 500,
		"stock": 1,
	})
	handler := New(slog.Default(), mock)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.Put("/products/{id}", handler.UpdateProduct)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestUpdateProduct_BadRequest(t *testing.T) {
	// TODO: Написать Unit-тест для обновления продукта с невалидным JSON (400 Bad Request)
	//mock := &ProductsMock{UpdateProductFunc: func(ctx context.Context, product models.Product) error {
	//	return nil
	//}}
	mock := mocks.NewProducts(t)
	r := chi.NewRouter()
	body := []byte("{invalid json}")
	handler := New(slog.Default(), mock)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.Put("/products/{id}", handler.UpdateProduct)
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	mock.AssertNotCalled(t, "UpdateProduct", mock2.Anything, mock2.Anything)
}

func TestUpdateProduct_Fail(t *testing.T) {
	// TODO: Написать Unit-тест для обновления продукта при ошибке сервиса (500 Internal Server Error)
	//mock := &ProductsMock{UpdateProductFunc: func(ctx context.Context, product models.Product) error {
	//	return errors.New("ops")
	//}}
	mock := mocks.NewProducts(t)
	mock.On("UpdateProduct", mock2.Anything, mock2.Anything).Return(errors.New("failed to create product"))
	r := chi.NewRouter()
	body, _ := json.Marshal(map[string]any{
		"name":  "Apple",
		"price": 500,
		"stock": 1,
	})
	handler := New(slog.Default(), mock)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.Put("/products/{id}", handler.UpdateProduct)
	r.ServeHTTP(w, req)
	if w.Code != 500 {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Delete Product
// =======================

func TestDeleteProduct_Success(t *testing.T) {
	// TODO: Написать Unit-тест для удаления продукта (200 OK)
	//mock := &ProductsMock{DeleteProductFunc: func(ctx context.Context, id int) error {
	//	return nil
	//}}
	mock := mocks.NewProducts(t)
	mock.On("DeleteProduct", mock2.Anything, mock2.Anything).Return(nil)
	r := chi.NewRouter()
	handler := New(slog.Default(), mock)
	r.Delete("/products/{id}", handler.DeleteProduct)
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestDeleteProduct_BadRequest(t *testing.T) {
	// TODO: Написать Unit-тест для удаления продукта с пустым id (400 Bad Request)
	//mock := &ProductsMock{DeleteProductFunc: func(ctx context.Context, id int) error {
	//	return nil
	//}}
	mock := mocks.NewProducts(t)
	r := chi.NewRouter()
	handler := New(slog.Default(), mock)
	r.Delete("/products/{id}", handler.DeleteProduct)
	req := httptest.NewRequest(http.MethodDelete, "/products/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	mock.AssertNotCalled(t, "DeleteProduct", mock2.Anything, mock2.Anything)
}

func TestDeleteProduct_Fail(t *testing.T) {
	// TODO: Написать Unit-тест для удаления продукта при ошибке сервиса (500 Internal Server Error)
	//mock := &ProductsMock{DeleteProductFunc: func(ctx context.Context, id int) error {
	//	return errors.New("ops")
	//}}
	mock := mocks.NewProducts(t)
	mock.On("DeleteProduct", mock2.Anything, mock2.Anything).Return(errors.New("failed to create product"))
	r := chi.NewRouter()
	handler := New(slog.Default(), mock)
	r.Delete("/products/{id}", handler.DeleteProduct)
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 500 {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
