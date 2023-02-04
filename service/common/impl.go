package common

func (CommonServiceImpl) Home() string {
	return "server is running 🙂"
}

func (CommonServiceImpl) NotFound() string {
	return "route not found"
}
