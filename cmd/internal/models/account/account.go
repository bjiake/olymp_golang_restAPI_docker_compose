package account

type Account struct {
	ID        int64  `json:"id" validate:"required" bson:"_id"`
	FirstName string `json:"firstName" validate:"required" bson:"firstName"`
	LastName  string `json:"lastName" validate:"required" bson:"lastName"`
	Email     string `json:"email" validate:"required" bson:"email"`
	Password  string `json:"password" validate:"required" bson:"password"`
}
type AccountInfo struct {
	ID        int64  `json:"id" validate:"required" bson:"_id" `
	FirstName string `json:"firstName" validate:"required" bson:"firstName"`
	LastName  string `json:"lastName" validate:"required" bson:"lastName"`
	Email     string `json:"email" validate:"required" bson:"email"`
}

type AccountLogin struct {
	Email    string `json:"email" validate:"required" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}

type AccountRegistration struct {
	FirstName string `json:"firstName" validate:"required" bson:"firstName"`
	LastName  string `json:"lastName" validate:"required" bson:"lastName"`
	Email     string `json:"email" validate:"required" bson:"email"`
	Password  string `json:"password" validate:"required" bson:"password"`
}
