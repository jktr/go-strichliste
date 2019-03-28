package schema

const (
	EndpointUser       = "/user"
	EndpointUserSearch = "/user/search"
)

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsActive    bool      `json:"active"`
	Email       *string   `json:"mailAddress"`
	Balance     int       `json:"balance"`
	TimeCreated Timestamp `json:"created"`
	TimeUpdated Timestamp `json:"updated"`
}

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"mailAddress,omitempty"`
}

type UserUpdateRequest struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"mailAddress,omitempty"`
	SetActive *bool  `json:"active,omitempty"`
}

type SingleUserResponse struct {
	User User `json:"user"`
}

type MultiUserResponse struct {
	Users []User `json:"users"`
}
