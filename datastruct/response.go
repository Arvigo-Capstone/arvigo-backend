package datastruct

import "time"

type (
	LoginRegisterResponse struct {
		UserID uint64 `json:"user_id"`
		Token  string `json:"token"`
	}

	UserDetailResponse struct {
		UserDetail
		Address                UserAddress                `json:"address"`
		PersonalityPercentages PersonalityPercentagesUser `json:"personality"`
	}

	PersonalityPercentagesUser struct {
		Agreeable     float64 `gorm:"column:agr_result" json:"percentage_of_agreeable"`
		Conscientious float64 `gorm:"column:csn_result" json:"percentage_of_conscientious"`
		Extraversion  float64 `gorm:"column:ext_result" json:"percentage_of_extraversion"`
		Neurotic      float64 `gorm:"column:est_result" json:"percentage_of_neurotic"`
		Openness      float64 `gorm:"column:opn_result" json:"percentage_of_openess"`
	}
	FaceShapeResponse struct {
		ImageUrl string `json:"image_url"`
		Result   string `json:"result"`
	}

	UserDetail struct {
		ID                        uint64     `json:"id"`
		Email                     string     `json:"email"`
		RoleID                    uint64     `json:"role_id"`
		RoleName                  string     `json:"role_name"`
		FullName                  string     `json:"full_name"`
		Gender                    string     `json:"gender"`
		DateOfBirth               *time.Time `json:"date_of_birth"`
		PlaceOfBirth              string     `json:"place_of_birth"`
		IsCompletePersonalityTest bool       `json:"is_complete_personality_test"`
		IsCompleteFaceTest        bool       `json:"is_complete_face_test"`
		IsSubscriptionActive      bool       `json:"is_subscription_active"`
		PersonalityID             uint64     `gorm:"column:personality_id" json:"-"`
		FaceShapeTagID            uint64     `gorm:"column:face_shape_id" json:"-"`
		FaceShape                 *string    `json:"face_shape"`
		PersonalityType           *string    `json:"personality_type"`
		IsVerified                bool       `json:"is_verified"`
		Avatar                    string     `json:"avatar"`
		AddressID                 uint64     `json:"addresses_id"`
		MerchantID                uint64     `json:"merchant_id"`
	}

	UserAddress struct {
		Street      string `json:"street"`
		Province    string `json:"province"`
		City        string `json:"city"`
		District    string `json:"district"`
		SubDistrict string `json:"sub_district"`
		PostalCode  uint64 `json:"postal_code"`
	}

	InitialProductResponse struct {
		InitialProduct
		Images                []string                     `json:"images"`
		Variants              []InitialProductVariant      `json:"variants"`
		ListMarketplace       []ProductMarketplaceWishlist `json:"marketplaces"`
		Tags                  []string                     `json:"tags"`
		RecommendationProduct []RecommendationProductML    `json:"recommendation_product"`
	}

	RecommendationProductML struct {
		ID        uint64 `gorm:"column:id" json:"id"`
		Name      string `gorm:"column:name" json:"name"`
		BrandName string `gorm:"column:brand_name" json:"brand_name"`
		Images    string `gorm:"column:images" json:"image"`
	}

	InitialProduct struct {
		ID                   uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name                 string `gorm:"column:name" json:"name"`
		Description          string `gorm:"column:description" json:"description"`
		LinkExternal         string `gorm:"column:link_external" json:"link_external"`
		CategoryName         string `gorm:"column:category_name" json:"category_name"`
		Status               string `gorm:"column:status" json:"status"`
		IsWishlisted         bool   `gorm:"column:is_wishlisted" json:"is_wishlisted"`
		IsSubscriptionActive bool   `gorm:"column:is_subscription_active" json:"is_subscription_active"`
		RejectedNote         string `gorm:"column:rejected_note" json:"rejected_note"`
		BrandName            string `gorm:"column:brand_name" json:"brand_name"`
		Images               string `gorm:"column:images" json:"-"`
	}

	InitialProductVariant struct {
		Name             string `gorm:"column:name" json:"name"`
		LinkAR           string `gorm:"column:link_ar" json:"link_ar"`
		IsPrimaryVariant bool   `gorm:"column:is_primary_variant" json:"is_primary_variant"`
		ProductID        uint64 `gorm:"column:product_id" json:"-"`
	}

	BrandResponse struct {
		ID         uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name       string `gorm:"column:name" json:"name"`
		Image      string `gorm:"column:image" json:"image"`
		CategoryID uint64 `gorm:"column:category_id" json:"category_id"`
	}

	CategoryResponse struct {
		ID   uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name string `gorm:"column:name" json:"name"`
	}

	ProductRecommendationResponse struct {
		ID          uint64 `gorm:"column:id" json:"id"`
		Name        string `gorm:"column:name" json:"name"`
		Description string `gorm:"column:description" json:"description"`
		Category    string `gorm:"column:category" json:"category"`
		Brand       string `gorm:"column:brand" json:"brand"`
		Tags        string `gorm:"column:tags" json:"tags"`
		Merchants   string `gorm:"column:merchants" json:"merchants"`
		Clicked     uint64 `gorm:"column:clicked" json:"clicked"`
		MerchantIDs string `gorm:"column:merchant_id" json:"-"`
		ProductIDs  string `gorm:"column:linked_product" json:"-"`
	}

	UserWishlistResponse struct {
		Product []ProductWishlist            `json:"products"`
		Store   []ProductMarketplaceWishlist `json:"stores"`
	}

	ProductWishlist struct {
		ID    uint64 `gorm:"column:id" json:"id"`
		Name  string `gorm:"column:name" json:"name"`
		Image string `gorm:"column:images" json:"image"`
		Brand string `gorm:"column:brand" json:"brand"`
	}

	ProductMarketplaceWishlist struct {
		ID              uint64  `gorm:"column:id" json:"id"`
		Name            string  `gorm:"column:name" json:"name"`
		Brand           string  `gorm:"column:brand" json:"brand"`
		Image           string  `gorm:"column:images" json:"image"`
		Price           float64 `gorm:"column:price" json:"price"`
		Merchant        string  `gorm:"column:merchant" json:"merchant"`
		Type            string  `json:"store_type"`
		Address         *string `json:"address"`
		Marketplace     *string `json:"marketplace_name"`
		MarketplaceLink *string `gorm:"column:marketplace_link" json:"marketplace_link"`
		MarketplaceID   uint64  `gorm:"column:marketplace_id" json:"-"`
		AddressID       uint64  `gorm:"column:addresses_id" json:"-"`
	}

	ProductMarketplaceDetail struct {
		ID              uint64                  `gorm:"column:id" json:"id"`
		ProductID       uint64                  `gorm:"column:product_id" json:"-"`
		Name            string                  `gorm:"column:name" json:"name"`
		Description     string                  `gorm:"column:description" json:"description"`
		Brand           string                  `gorm:"column:brand" json:"brand"`
		Image           string                  `gorm:"column:images" json:"-"`
		IsWishlisted    bool                    `gorm:"column:is_wishlisted" json:"is_wishlisted"`
		Images          []string                `json:"images"`
		Price           float64                 `gorm:"column:price" json:"price"`
		Merchant        string                  `gorm:"column:merchant" json:"merchant"`
		MerchantID      uint64                  `gorm:"column:merchant_id" json:"-"`
		Type            string                  `json:"store_type"`
		Address         *string                 `json:"address"`
		Marketplace     *string                 `json:"marketplace_name"`
		MarketplaceLink *string                 `gorm:"column:marketplace_link" json:"marketplace_link"`
		MarketplaceID   uint64                  `gorm:"column:marketplace_id" json:"-"`
		AddressID       uint64                  `gorm:"column:addresses_id" json:"-"`
		Variants        []InitialProductVariant `json:"variants"`
	}

	HomeMerchantResponse struct {
		MerchantID uint64        `gorm:"column:merchant_id" json:"merchant_id"`
		Name       string        `gorm:"column:merchant_name" json:"merchant_name"`
		AddressID  uint64        `gorm:"column:addresses_id" json:"-"`
		Location   string        `json:"location"`
		Product    []interface{} `json:"products"`
	}

	HomeResponse struct {
		Personality    []HomeProduct `json:"personalities"`
		FaceShape      []HomeProduct `json:"face_shapes"`
		Recommendation []HomeProduct `json:"recommendations"`
	}

	HomeProduct struct {
		ID    uint64   `gorm:"column:id" json:"id"`
		Name  string   `gorm:"column:name" json:"name"`
		Brand string   `gorm:"column:brand" json:"brand"`
		Image string   `gorm:"column:images" json:"image"`
		Tags  []string `json:"tags"`
	}

	MerchantHome struct {
		MerchantProduct []MerchantProduct `json:"products"`
		MerchantVisitor MerchantVisitor   `json:"visitors"`
	}

	MerchantHomeDetail struct {
		MerchantProductByID
		Marketplace []MerchantMarketplace `json:"marketplaces"`
	}

	UserSubscriptionResponse struct {
		ID                uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		UserName          string     `gorm:"column:user_name" json:"user_name"`
		UserID            uint64     `gorm:"column:user_id" json:"user_id"`
		Price             uint64     `gorm:"column:price" json:"price"`
		UniqueCode        uint64     `gorm:"column:unique_code" json:"unique_code"`
		SubscriptionStart *time.Time `gorm:"column:subscription_start" json:"subscription_start"`
		SubscriptionEnd   *time.Time `gorm:"column:subscription_end" json:"subscription_end"`
		Status            string     `gorm:"column:status" json:"status"`
		Message           string     `gorm:"column:message" json:"message"`
		Bank              string     `gorm:"column:bank" json:"bank"`
		PaidAt            *time.Time `gorm:"column:paid_at" json:"paid_at"`
		CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
		UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
	}

	MerchantSubscriptionResponse struct {
		ID                uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		UserName          string     `gorm:"column:user_name" json:"user_name"`
		UserID            uint64     `gorm:"column:user_id" json:"user_id"`
		Price             uint64     `gorm:"column:price" json:"price"`
		UniqueCode        uint64     `gorm:"column:unique_code" json:"unique_code"`
		SubscriptionStart *time.Time `gorm:"column:subscription_start" json:"subscription_start"`
		SubscriptionEnd   *time.Time `gorm:"column:subscription_end" json:"subscription_end"`
		Status            string     `gorm:"column:status" json:"status"`
		Message           string     `gorm:"column:message" json:"message"`
		Bank              string     `gorm:"column:bank" json:"bank"`
		Product           string     `gorm:"column:products" json:"products"`
		PaidAt            *time.Time `gorm:"column:paid_at" json:"paid_at"`
		CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
		UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
	}
)
