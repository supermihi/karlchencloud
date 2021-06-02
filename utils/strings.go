package utils

func GetStringWithDefault(value *string, def func() string) string {
	if value != nil && *value != "" {
		return *value
	}
	return def()
}
