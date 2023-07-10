package helper

func PanicIfErrors(err error) {
	if err != nil {
		panic(err)
	}
}
