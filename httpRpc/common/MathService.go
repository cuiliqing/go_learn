package common

import (
	"errors"
)

type Args struct {
	A, B float32
}

type Result struct {
	Value float32
}

type MathService struct {

}

func (s *MathService) Add(args *Args, result *Result) error {
	result.Value = args.A + args.B
	return nil
}

func (s *MathService) Divide(args *Args,result *Result) error {
	if args.B == 0 {
		return errors.New("divider is not zero!")
	}
	result.Value = args.A / args.B
	return nil
}