package main_control

type ErrorNoDelivery struct {
}

func (e ErrorNoDelivery) Error() string {
	return "no delivery is free"
}
