package person

type Person struct {
	Name     string
	Email    string
	Password int  `json:"-"`
	Premium  bool `json:"subscribed?"`
}
