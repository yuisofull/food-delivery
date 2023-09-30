package restaurantmodel

type Filter struct {
	OwnerID int `json:"owner_id" form:"owner_id"`
}
