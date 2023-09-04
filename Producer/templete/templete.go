package templete

import (
	"time"
)

type Product struct {
	ID                   int       `bson:"_id" json:"product_id"`
	UserID               int       `bson:"user_id" json:"user_id"`
	ProductName          string    `bson:"product_name" json:"product_name"`
	ProductDescription   string    `bson:"product_description" json:"product_description"`
	ProductImages        []string  `bson:"product_images" json:"product_images"`
	ProductPrice         float64   `bson:"product_price" json:"product_price"`
	CompressedImagesPath []string  `bson:"compressed_product_images" json:"compressed_product_images"`
	CreatedAt            time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time `bson:"updated_at" json:"updated_at"`
}

type User struct {
	ID        int       `bson:"_id,omitempty" json:"user_id"`
	Name      string    `bson:"name" json:"name"`
	Mobile    string    `bson:"mobile" json:"mobile"`
	Latitude  float64   `bson:"latitude" json:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
