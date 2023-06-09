package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterProductRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	v1Group.GET("/product-recommendation", getRecommendationProduct, middleware.ApiKeyMiddleware)

	v1Group.GET("/merchants/product", getDashboardMerchant, middleware.AuthMiddleware)
	productGroup := v1Group.Group("/products", middleware.AuthMiddleware)
	productGroup.DELETE("/:id", delProductByID)
	initialProductGroup := productGroup.Group("/initials")
	initialProductGroup.GET("/:id", getInitalProductByID)
	initialProductGroup.GET("/marketplace/:id", getMarketplaceProductByID)

	initialProductGroup.POST("", createInitialProductHandler)
	initialProductGroup.PUT("/:id", updateInitialProductHandler)
	initialProductGroup.GET("/category/:id", getInitalProductByCategoryID)

	merchantProductGroup := productGroup.Group("/merchants")
	merchantProductGroup.POST("", createMerchantProductHandler)
	merchantProductGroup.PUT("", updateMerchantProduct)
	merchantProductGroup.PUT("/verify", verifyMerchantProduct)
}

func createInitialProductHandler(c echo.Context) error {
	var data datastruct.CreateInitialProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Images = images
	statusCode, err := repository.CreateInitialProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed create product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product created", nil, statusCode)
}

func updateInitialProductHandler(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "Invalid product ID", nil, http.StatusBadRequest)
	}
	var data datastruct.CreateInitialProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Images = images
	statusCode, err := repository.UpdateInitialProduct(data, pID)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product updated", nil, statusCode)
}

func createMerchantProductHandler(c echo.Context) error {
	var data datastruct.CreateMerchantProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Images = images

	statusCode, err := repository.CreateMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed create product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product created", nil, statusCode)
}

func getRecommendationProduct(c echo.Context) error {
	data, statusCode, err := repository.GetProductRecommendationMachineLearning()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getDashboardMerchant(c echo.Context) error {
	data, statusCode, err := repository.GetMerchantDashboard()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getInitalProductByCategoryID(c echo.Context) error {
	categoryID := utils.StrToUint64(c.Param("id"), 0)
	if categoryID == 0 {
		return utils.ResponseJSON(c, "Invalid category ID", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetInitialProductByCategoryID(categoryID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), []uint64{}, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getInitalProductByID(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "Invalid product ID", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetInitialProductByID(pID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getMarketplaceProductByID(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "Invalid product ID", nil, http.StatusBadRequest)
	}
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	data, statusCode, err := repository.GetMarketplaceProductByID(pID, userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func verifyMerchantProduct(c echo.Context) error {
	var data datastruct.VerifyProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.VerifyMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product updated", nil, statusCode)
}

func updateMerchantProduct(c echo.Context) error {
	var data datastruct.UpdateProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.UpdateMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product updated", nil, statusCode)
}

func delProductByID(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "Invalid product ID", nil, http.StatusBadRequest)
	}
	statusCode, err := repository.DeleteProduct(pID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", nil, statusCode)
}
