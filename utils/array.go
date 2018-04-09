package utils

func InStrArray(filed string, arr []string) bool {

	if len(arr) == 0 {
		return false
	}

	for i := 0; i < len(arr); i++   {
		if arr[i] == filed {
			return true
		}
	}
	return false
}
