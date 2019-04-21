package schema

const (
	EndpointUser       = "/user"
	EndpointUserSearch = "/user/search"
)

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsActive    bool      `json:"isActive"`
	IsDisabled  bool      `json:"isDisabled"` // TODO
	Email       *string   `json:"email"`
	Balance     int       `json:"balance"`
	TimeCreated Timestamp `json:"created"`
	TimeUpdated Timestamp `json:"updated"`
}

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type UserUpdateRequest struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	SetActive *bool  `json:"active,omitempty"`
}

type SingleUserResponse struct {
	User User `json:"user"`
}

type MultiUserResponse struct {
	Users []User `json:"users"`
}
