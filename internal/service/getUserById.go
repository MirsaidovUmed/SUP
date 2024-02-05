package service

func (s *Service) CheckUserById(userID int) (err error) {
	return s.Repo.CheckUserById(userID)
}
