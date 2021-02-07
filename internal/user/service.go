package user

type Service struct {
	users Users
}

func NewService() Service {
	return Service{
		users: make(Users, 0),
	}
}

func (s *Service) AddUser(user User) {
	s.users = append(s.users, user)
}

func (s Service) GetUsers() Users {
	return s.users
}

func (s Service) GetUsersByID(id int) User {
	for _, user := range s.users {
		if user.ID == id {
			return user
		}
	}

	return User{}
}

func (s *Service) UpdateUser(id int, newUser User) {
	for i := 0; i < len(s.users); i++ {
		if s.users[i].ID == id {
			s.users[i].Name = newUser.Name
			s.users[i].Age = newUser.Age
			break
		}
	}
}

func (s *Service) DeleteUser(id int) {
	for i := 0; i < len(s.users); i++ {
		if s.users[i].ID == id {
			s.users[i] = s.users[len(s.users)-1]
			s.users[len(s.users)-1] = User{}
			s.users = s.users[:len(s.users)-1]
			break
		}
	}
}
