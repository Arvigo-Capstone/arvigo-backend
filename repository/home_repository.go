package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/utils"
)

func GetHome(userID uint64) (res datastruct.HomeResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var (
		user                  datastruct.UserWithPersonalityTag
		faceShapeProduct      []datastruct.HomeProduct
		personalityProduct    []datastruct.HomeProduct
		recommendationProduct []datastruct.HomeProduct
	)

	if err := db.Table("users").
		Select([]string{
			"users.*",
			"up.tag_ids",
		}).
		Where("users.id = ?", userID).
		Joins("LEFT JOIN user_personalities up on users.personality_id = up.id").
		First(&user).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	if user.TagID != "" {
		tags := strings.Split(user.TagID, ",")
		for _, v := range tags {
			user.TagIDs = append(user.TagIDs, utils.StrToUint64(v, 0))
		}
	}

	// faceshape
	if user.IsCompleteFaceTest {
		if err := db.Table("products p").
			Select([]string{
				"p.id",
				"p.name",
				"p.images",
				"b.name as brand",
			}).
			Joins("LEFT JOIN brands b on b.id = p.brand_id").
			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.GlassesCategoryID, constant.GetFaceShapeTags[user.FaceShapeID]).
			Find(&faceShapeProduct).
			Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for i, v := range faceShapeProduct {
			faceShapeProduct[i].Image = strings.Split(v.Image, ",")[0]
			var tagIDs []uint64
			var tags []string
			if err := db.Table("detail_product_tags").
				Select([]string{
					"tag_id",
				}).
				Where("product_id = ?", v.ID).
				Find(&tagIDs).
				Error; err != nil {
				return res, http.StatusInternalServerError, err
			}

			for _, v := range tagIDs {
				tags = append(tags, constant.GetTagNameByDetailTag[v]...)
			}
			faceShapeProduct[i].Tags = utils.RemoveDuplicates(tags)
		}
	}

	if user.IsCompletePersonalityTest {
		if err := db.Table("products p").
			Select([]string{
				"p.id",
				"p.name",
				"p.images",
				"b.name as brand",
			}).
			Joins("LEFT JOIN brands b on b.id = p.brand_id").
			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.MakeupCategoryID, user.TagIDs).
			Find(&personalityProduct).
			Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for i, v := range personalityProduct {
			personalityProduct[i].Image = strings.Split(v.Image, ",")[0]
			var tagIDs []uint64
			var tags []string
			if err := db.Table("detail_product_tags").
				Select([]string{
					"tag_id",
				}).
				Where("product_id = ?", v.ID).
				Find(&tagIDs).
				Error; err != nil {
				return res, http.StatusInternalServerError, err
			}

			for _, v := range tagIDs {
				tags = append(tags, constant.GetTagNameByDetailTag[v]...)
			}
			personalityProduct[i].Tags = utils.RemoveDuplicates(tags)
		}
	}

	// recommendation TODO: integrate with ML/sort subs
	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.images",
			"b.name as brand",
		}).
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		Where("p.merchant_id = 0").
		Find(&recommendationProduct).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range recommendationProduct {
		recommendationProduct[i].Image = strings.Split(v.Image, ",")[0]
		var tagIDs []uint64
		var tags []string
		if err := db.Table("detail_product_tags").
			Select([]string{
				"tag_id",
			}).
			Where("product_id = ?", v.ID).
			Find(&tagIDs).
			Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for _, v := range tagIDs {
			tags = append(tags, constant.GetTagNameByDetailTag[v]...)
		}
		recommendationProduct[i].Tags = utils.RemoveDuplicates(tags)
	}

	res = datastruct.HomeResponse{
		Personality:    personalityProduct,
		FaceShape:      faceShapeProduct,
		Recommendation: recommendationProduct,
	}
	return
}

func GetHomeSearch(search string) (res []datastruct.HomeProduct, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var pIDs []string
	response, err := utils.FetchMachineLearningAPI("GET", fmt.Sprintf("/product_search?query=%s", search), nil)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	var products []datastruct.ProductFromML
	err = json.Unmarshal(response, &products)
	if err != nil {
		fmt.Printf("Error unmarshaling response body: %v", err)
		return
	}

	if len(products) == 0 {
		statusCode = http.StatusNotFound
		return
	}

	for _, v := range products {
		pIDs = append(pIDs, v.ID)
	}

	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.images",
			"b.name as brand",
		}).
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		Where("p.id IN (?)", pIDs).
		Find(&res).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range res {
		res[i].Image = strings.Split(v.Image, ",")[0]
	}

	return
}

func GetHomeMerchant() (merchants []datastruct.HomeMerchantResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Table("users").
		Select([]string{
			"addresses_id",
			"m.name merchant_name",
			"m.id merchant_id",
		}).
		Joins("join merchants m on users.merchant_id = m.id").
		Where("merchant_id != 0").
		Find(&merchants).
		Error; err != nil {
		return merchants, http.StatusInternalServerError, err
	}

	for i, v := range merchants {
		_, location, _, _ := GetAddressByID(v.AddressID)
		merchants[i].Location = location

		var (
			merchantProduct []datastruct.ProductMarketplaceWishlist
		)
		if err = db.Table("detail_product_marketplaces").
			Select([]string{
				"detail_product_marketplaces.id AS id",
				"products.name",
				"brands.name AS brand",
				"products.images",
				"products.price",
				"merchants.name AS merchant",
				"detail_product_marketplaces.link AS marketplace_link",
				"detail_product_marketplaces.marketplace_id",
				"detail_product_marketplaces.addresses_id",
			}).
			Joins("LEFT JOIN products ON products.id = detail_product_marketplaces.product_id").
			Joins("LEFT JOIN brands ON brands.id = products.brand_id").
			Joins("LEFT JOIN merchants ON products.merchant_id = merchants.id").
			Where("products.merchant_id = ? AND products.status IN (?)", v.MerchantID, []string{constant.StatusApproved, constant.StatusSubscribed}).
			Order("products.is_subscription_active DESC").
			Find(&merchantProduct).Error; err != nil {
			return merchants, http.StatusInternalServerError, err
		}

		for i, v := range merchantProduct {
			merchantProduct[i].Image = strings.Split(v.Image, ",")[0]

			if v.AddressID != 0 {
				merchantProduct[i].Type = "offline"
				addr, _, _, err := GetAddressByID(v.AddressID)
				if err == nil {
					merchantProduct[i].Address = &addr
				}
				continue
			}

			if v.MarketplaceID != 0 {
				merchantProduct[i].Type = "online"
				marketplaceName := constant.Marketplace[v.MarketplaceID]
				merchantProduct[i].Marketplace = &marketplaceName
			}
		}

		var interfaceSlice []interface{} = make([]interface{}, len(merchantProduct))
		for i, v := range merchantProduct {
			interfaceSlice[i] = v
		}
		merchants[i].Product = interfaceSlice
	}

	return
}
