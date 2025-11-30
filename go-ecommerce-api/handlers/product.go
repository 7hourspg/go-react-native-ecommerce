package handlers

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductServices *services.ProductServices
}

func NewProductHandler(s *services.ProductServices) *ProductHandler {
	return &ProductHandler{
		ProductServices: s,
	}
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieve a specific product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.ProductResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	product, err := h.ProductServices.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Retrieve all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ProductsResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.ProductServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetFeaturedProducts godoc
// @Summary      Get featured products
// @Description  Retrieve all featured products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ProductsResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/featured [get]
func (h *ProductHandler) GetFeaturedProducts(c *gin.Context) {
	products, err := h.ProductServices.GetFeatured()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProductsByCategory godoc
// @Summary      Get products by category
// @Description  Retrieve products filtered by category
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        category   path      string  true  "Product Category"
// @Success      200  {object}  models.ProductsResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/category/{category} [get]
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	products, err := h.ProductServices.GetByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// SearchProducts godoc
// @Summary      Search products
// @Description  Search products by name or description
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        query   query     string  true  "Search query"
// @Success      200  {object}  models.ProductsResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/search [get]
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query parameter is required",
		})
		return
	}

	products, err := h.ProductServices.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product (Admin only). Valid categories: Electronics, Accessories, Home, Office
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product data"
// @Success      201  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product data",
			"error":   err.Error(),
		})
		return
	}

	if err := h.ProductServices.Create(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    product,
	})
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update an existing product (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Param        product  body      models.Product  true  "Updated product data"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product data",
		})
		return
	}

	if err := h.ProductServices.Update(uint(id), &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by ID (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	if err := h.ProductServices.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// BulkCreateProducts godoc
// @Summary      Create multiple products
// @Description  Create multiple products at once (Admin only). Valid categories: Electronics, Accessories, Home, Office
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        products  body      models.BulkCreateProductsRequest  true  "Array of products to create"
// @Success      201  {object}  models.ProductsResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/products/bulk [post]
func (h *ProductHandler) BulkCreateProducts(c *gin.Context) {
	var req models.BulkCreateProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	if len(req.Products) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "At least one product is required",
		})
		return
	}

	products, err := h.ProductServices.BulkCreate(req.Products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create products",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Products created successfully",
		"data":    products,
		"count":   len(products),
	})
}

