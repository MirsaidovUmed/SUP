package service

func (s *Service) RoleExists(roleName string) (err error) {
	err = s.Repo.RoleExists(roleName)
	if err != nil {
		return err
	}
	return
}
