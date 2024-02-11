package restaurantmodel

import (
	"errors"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr"  gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	UserID          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;foreignKey:Id"`
	LikedCount      int                `json:"liked_count" gorm:"-"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(false)
	}
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr"  gorm:"column:addr"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type UpdateRestaurant struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr"  gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (UpdateRestaurant) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
