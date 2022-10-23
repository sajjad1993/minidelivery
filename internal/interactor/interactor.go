package interactor

import (
	"context"
	"fmt"
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
	ctx := context.Background()
	err := i.main.Deliver(dest, ctx)
	if err != nil {
		fmt.Printf("some error occurd %s \n ", err)
	}
}
func New(main contract.MainControl, logger *log.Logger) *Interactor {
	interactor := Interactor{
		main:   main,
		logger: logger,
	}
	return &interactor
}
