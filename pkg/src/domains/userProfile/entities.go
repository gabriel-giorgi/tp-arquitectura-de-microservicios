package userProfile

type Profile struct {
	UserID               string  `json:"userID"`
	UserLevel            int     `json:"userLevel"`
	Experience int     `json:"experience"`
	CurrentDiscount      float64 `json:"currentDiscount"`
}

type LevelConfig struct {
	experienceForLevelUp int
	discount             float64
}

type User struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Login       string   `json:"login"`
	Permissions []string `json:"permissions"`
}

type CartResponse struct {
	Message struct {
		CartID   string `json:"cartId"`
		UserID   string `json:"userId"`
		Articles []struct {
			ID       string `json:"id"`
			Quantity int    `json:"quantity"`
		} `json:"articles"`
	} `json:"message"`
}

type CatalogResponse struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Enabled     bool    `json:"enabled"`
}
type NotificationLevelUp struct {
	CurrentLevel int `json:"currentLevel"`
	CurrentDiscount float64 `json:"currentDiscount"`
}
