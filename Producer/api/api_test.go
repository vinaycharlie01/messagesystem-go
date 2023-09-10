package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	m1 "myapp/Producer/mongo1"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	// Connect to MongoDB and initialize collections
	client, _, err := m1.MongoConnect("vinay1", "product")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	gin.SetMode(gin.TestMode)
	samples := []struct {
		inputJSON    string
		statusCode   int
		user_id      int
		product_name string
		errMessage   string
	}{
		{
			inputJSON: `{
				"user_id": 1030,
				"product_name": "Sample Product",
				"product_description": "This is a sample product description.",
				"product_images": [
					"https://images.unsplash.com/5/unsplash-kitsune-4.jpg?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=bc01c83c3da0425e9baa6c7a9204af81",
					"https://media.wired.com/photos/593261cab8eb31692072f129/master/w_2240,c_limit/85120553.jpg"
				],
				"product_price": 29.99
			}`,
			statusCode:   201,
			user_id:      1030,
			product_name: "Sample Product",
			errMessage:   "",
		},
		{
			inputJSON: `{
				"user_id": 1022,
				"product_name": "Sample Product",
				"product_description": "This is a sample product description.",
				"product_images": [
					"https://images.unsplash.com/5/unsplash-kitsune-4.jpg?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=bc01c83c3da0425e9baa6c7a9204af81",
					"https://media.wired.com/photos/593261cab8eb31692072f129/master/w_2240,c_limit/85120553.jpg"
				],
				"product_price": 29.99
			}`,
			statusCode:   500,
			user_id:      0,
			product_name: "",
			errMessage:   "",
		},
		{
			inputJSON: `{
				"user_id":",
				"product_name": "Sample Product",
				"product_description": "This is a sample product description.",
				"product_images": [
					"https://images.unsplash.com/5/unsplash-kitsune-4.jpg?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=bc01c83c3da0425e9baa6c7a9204af81",
					"https://media.wired.com/photos/593261cab8eb31692072f129/master/w_2240,c_limit/85120553.jpg"
				],
				"product_price": 29.99
			}`,
			statusCode:   400,
			user_id:      0,
			product_name: "",
			errMessage:   "invalid character '\\n' in string literal",
		},
	}
	for _, v := range samples {
		r := gin.Default()
		r.POST("/create-product", CreateProduct)
		req, err := http.NewRequest(http.MethodPost, "/create-product", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		fmt.Println("this is the response data: ", responseMap)
		assert.Equal(t, rr.Code, v.statusCode)
		resout, _ := responseMap["product"]
		res, ok := resout.(map[string]interface{})
		if v.statusCode == 201 && ok {
			val, _ := res["user_id"]
			assert.Equal(t, int(val.(float64)), v.user_id)
			assert.Equal(t, res["product_name"], v.product_name)
		}

		if v.statusCode == 400 {
			assert.Equal(t, responseMap["error"], v.errMessage)
		}

	}
}
