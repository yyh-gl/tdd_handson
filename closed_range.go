package main

import (
	"errors"
	"fmt"
)

type ClosedRange struct {
	lowerEndpoint int
	upperEndpoint int
}

func NewClosedRange(lowerEndpoint int, upperEndpoint int) (*ClosedRange, error) {
	if upperEndpoint < lowerEndpoint {
		return nil, errors.New("下端点は上端点以下にしてください。")
	}

	return &ClosedRange {
		lowerEndpoint: lowerEndpoint,
		upperEndpoint: upperEndpoint,
	}, nil
}

func (cr ClosedRange) StringRange() string {
	return fmt.Sprintf("[%d,%d]", cr.lowerEndpoint, cr.upperEndpoint)
}

func (cr ClosedRange) ContainsNumber(num int) bool {
	return cr.lowerEndpoint <= num && num <= cr.upperEndpoint
}

// FIXME: targetっていう名前どうしよう
func (cr ClosedRange) Equals(target ClosedRange) bool {
	return cr == target
}

// FIXME: targetっていう名前どうしよう
func (cr ClosedRange) ContainsRange(target ClosedRange) bool {
	return cr.lowerEndpoint <= target.lowerEndpoint && target.upperEndpoint <= cr.upperEndpoint
}
