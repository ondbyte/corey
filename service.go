package corey

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddContact(c *Contact) error {
	return s.repo.AddContact(c)
}

func (s *Service) GetContact(id uint) (*Contact, error) {
	return s.repo.GetContact(id)
}

func (s *Service) GetAllContact() ([]*Contact, error) {
	return s.repo.GetAllContacts()
}

func (s *Service) AddTask(c *Task) error {
	return s.repo.AddTask(c)
}

func (s *Service) GetTask(id uint) (*Task, error) {
	return s.repo.GetTask(id)
}

func (s *Service) GetAllTask() ([]*Task, error) {
	return s.repo.GetAllTasks()
}
