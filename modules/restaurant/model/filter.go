package restaurantmodel

type Filter struct {
	OwnerID string `json:"owner_id,omitempty" form:"owner_id"`
	Status  []int  `json:"-"`
}
