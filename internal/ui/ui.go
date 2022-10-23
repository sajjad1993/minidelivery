package ui

import (
	"fmt"
	"minidelivery/internal/entity"
	"minidelivery/internal/interactor"
	"strconv"
)

// get input from somewhere

type Ui struct {
	interactor *interactor.Interactor
}

// listen to terminal to get dest
func (u *Ui) Listen() {

	var xx, yy string
	_, err := fmt.Scanf("%s %s", &xx, &yy)
	x, err := strconv.Atoi(xx)
	y, err := strconv.Atoi(yy)
	if err != nil {
		// ... handle error
		fmt.Printf("\n\n input is wrong \n\n")
		return
	}
	dest := entity.Location{
		X: int8(x),
		Y: int8(y),
	}
	u.interactor.Go(dest)
}
func New(interactor *interactor.Interactor) *Ui {
	return &Ui{interactor: interactor}
}
