package utils

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// IdentifyChanges compares two documents and returns a map of the fields that have changed.
func IdentifyChanges(existing, updated interface{}) bson.M {
	// Convert both documents to their reflection values
	existingValue := reflect.ValueOf(existing).Elem()
	updatedValue := reflect.ValueOf(updated).Elem()

	// Prepare a map to store the changed fields
	changes := bson.M{}

	// Iterate over the fields in the updated document
	for i := 0; i < updatedValue.NumField(); i++ {
		// Get the field name and value from both documents
		fieldName := updatedValue.Type().Field(i).Name
		existingField := existingValue.FieldByName(fieldName)
		updatedField := updatedValue.Field(i)

		// Check if the field values are different and the updated field is non-zero
		if !reflect.DeepEqual(existingField.Interface(), updatedField.Interface()) && !isZeroValue(updatedField) {
			// Add the changed field to the changes map
			bsonFieldName := updatedValue.Type().Field(i).Tag.Get("bson")
			if bsonFieldName == "" {
				bsonFieldName = fieldName
			}
			changes[bsonFieldName] = updatedField.Interface()
		}
	}

	return changes
}

// isZeroValue checks if a given value is the zero value for its type
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
