package arithmetic

import "strconv"

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func FormatResult(res float64) string {
	if res == float64(int(res)) {
		return strconv.Itoa(int(res))
	}
	return strconv.FormatFloat(res, 'f', 2, 64)
}
