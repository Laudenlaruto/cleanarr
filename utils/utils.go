package utils

func ConvertBytestoGibiBytes(x int) int {
	y := x / 1024 / 1024 / 1024
	return y
}

func ConvertBytestoGigaBytes(x int) int {
	y := x / 1000 / 1000 / 1000
	return y
}
