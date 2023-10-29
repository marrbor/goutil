package maputil

// HasMapItem returns whether specified key is included specified map or not.
func HasMapItem(params map[string]interface{}, key string) bool {
	return params[key] != nil
}
