package respone
func ResponseData(name string,value interface{}) map[string]interface{} {
	return map[string]interface{}{
		name : value,
	}
}