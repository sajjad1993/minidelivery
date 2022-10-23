package interactor

import (
	"log"
	"minidelivery/internal/contract"
	"minidelivery/internal/entity"
)

type Interactor struct {
	main   contract.MainControl
	logger *log.Logger
}

func (i *Interactor) Go(dest entity.Location) {
	// get from terminal
	// for example
	_ = i.main.Deliver(dest)
}
func New(main contract.MainControl, logger *log.Logger) *Interactor {
	interactor := Interactor{
		main:   main,
		logger: logger,
	}
	return &interactor
}
