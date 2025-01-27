package net

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type (
	Args struct {
		X int
		Y int
	}

	Result struct {
		Quotient  int
		Remainder int
	}
)

type MathService struct {
}

func (s *MathService) Add(args *Args, result *int) error {
	*result = args.X + args.Y
	return nil
}

func (s *MathService) Divide(args *Args, result *Result) error {
	if args.Y == 0 {
		return errors.New("division by zero")
	}

	result.Quotient = args.X / args.Y
	result.Remainder = args.X % args.Y
	return nil
}

type MathService2 struct {
}

func (s *MathService2) Rand(args *Args, result *int) error {
	*result = rand.IntN(args.Y) + args.X
	return nil
}

func (s *MathService2) NoReply(args *Args, _ *struct{}) error {
	fmt.Println("no reply")
	fmt.Println(args.X, args.Y)
	return nil
}

type MathService3 struct {
}

func (s *MathService3) Rand(args *Args, result *int) error {
	*result = rand.IntN(args.Y) + args.X
	return nil
}
