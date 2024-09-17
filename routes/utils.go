package routes

import (
	"errors"
	"reflect"
	"simple-api-app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var requiredFields = []string{"title", "price"}

func CheckFields(data map[string]interface{}) error {
	if err := CheckRequiredFields(data); err != nil {
		return err
	}
	if err := CheckAllowedFields(data); err != nil {
		return err
	}
	return nil
}

func CheckRequiredFields(data map[string]interface{}) error {
	for _, field := range requiredFields {
		value, exists := data[field]
		if !exists || value == nil || value == "" {
			return errors.New("Missing required field: " + field)
		}
	}
	return nil
}

func CheckAllowedFields(data map[string]interface{}) error {
	allowedFields := GetAllowedFields()
	for key := range data {
		if !allowedFields[key] {
			return errors.New("Invalid field: " + key)
		}
	}
	return nil
}

func GetAllowedFields() map[string]bool {
	allowedFields := make(map[string]bool)
	typ := reflect.TypeOf(models.Product{})
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Tag.Get("json") == "_id" {
			continue
		}

		field := typ.Field(i)
		allowedFields[field.Tag.Get("json")] = true
	}
	return allowedFields
}

func MapToStruct(data map[string]interface{}, product *models.Product) error {
	val := reflect.ValueOf(product).Elem()
	typ := val.Type()

	for key, value := range data {
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			if field.Tag.Get("json") == key {
				fieldVal := val.Field(i)

				if fieldVal.CanSet() {
					fieldValue := reflect.ValueOf(value)

					if fieldValue.Type().ConvertibleTo(fieldVal.Type()) {
						fieldVal.Set(fieldValue.Convert(fieldVal.Type()))
					} else {
						return errors.New("Invalid value type for field: " + key)
					}
				}
			}
		}
	}
	return nil
}

func BuildInsertDocument(data map[string]interface{}) bson.M {
	insertDoc := bson.M{}
	for key, value := range data {
		if value != nil && value != "" {
			insertDoc[key] = value
		}
	}
	return insertDoc
}

func AddMetaToProduct(productData map[string]interface{}) {
	if meta, exists := productData["meta"].(map[string]interface{}); exists {
		if createdAt, ok := meta["createdAt"].(string); ok && createdAt != "" {
			meta["updatedAt"] = time.Now().Format(time.RFC3339)
		} else {
			meta["createdAt"] = time.Now().Format(time.RFC3339)
			meta["updatedAt"] = time.Now().Format(time.RFC3339)
		}
	} else {
		productData["meta"] = map[string]interface{}{
			"createdAt": time.Now().Format(time.RFC3339),
			"updatedAt": time.Now().Format(time.RFC3339),
		}
	}
}
