package admin

func (as AdminServiceImpl) DeleteUser(email string) error {
	err := as.UserRepository.Delete(email)
	return err
}