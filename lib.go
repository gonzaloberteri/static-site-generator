package main

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
