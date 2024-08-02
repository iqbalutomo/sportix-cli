package constants

type PaymentStatus string

const (
	Owner string = "owner"
	User  string = "user"

	Completed PaymentStatus = "completed"
	Pending   PaymentStatus = "pending"
)

var (
	RoleOptions  = []string{"owner", "user"}
	YesNoOptions = []string{"yes", "no"}
)
