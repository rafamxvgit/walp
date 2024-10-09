package utils

/*
dados um conjunto (data, err), a função expect consome o erro e retorna
somente o dado
*/
func Expect[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
