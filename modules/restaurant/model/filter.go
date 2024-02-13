package restaurantmodel

type Filter struct {
	FakeOwnerID string `json:"-" form:"owner_id"`
	OwnerID     int    `json:"owner_id,omitempty" form:"-"`
	Status      []int  `json:"-"`
}
