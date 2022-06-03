package utils

func ConvertMapInterfaceToString(input map[string]interface{}) map[string]string {
	if input == nil {
		return nil
	}
	res := make(map[string]string)
	for k, v := range input {
		res[k] = ToString(v)
	}
	return res
}
