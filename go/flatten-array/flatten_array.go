package flatten

func Flatten(nested interface{}) []interface{} {
	flattenedArray := []interface{}{}

	for _, child := range nested.([]interface{}) /* Not Safe Type Assertion */ {
		switch childType := child.(type) {
		case nil:
			continue
		case []interface{}:
			flattenedArray = append(flattenedArray, Flatten(childType)...) // Recursion
		default:
			flattenedArray = append(flattenedArray, child)
		}
	}
	return flattenedArray
}
