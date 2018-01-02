package controller

// Controller specifies the behaviour of a system which can be setup initially and is able to loop a specific business logic.
type Controller interface {
	Loop() error
	Setup() error
}
