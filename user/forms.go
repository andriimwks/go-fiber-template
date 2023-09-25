package user

type SignInForm struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUpForm struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Nickname  string `json:"nickname" form:"nickname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}
