package chain

import (
	"github.com/samber/lo"
)

type SliceChain[T1, T2, T3, T4 any] []T1
type SliceChainComparable[T1 comparable, T2, T3, T4 any] SliceChain[T1, T2, T3, T4]

func NewSliceChain4[T1, T2, T3, T4 any](s []T1) SliceChain[T1, T2, T3, T4] {
	return s
}

func (s SliceChain[T1, T2, T3, T4]) Map(f func(T1, int) T2) SliceChain[T2, T3, T4, any] {
	return NewSliceChain4[T2, T3, T4, any](lo.Map(s, f))
}

/*
func (s SliceChain4[T1, T2, T3, R]) Uniq() SliceChain4[T1, T2, T3, R] { return lo.Uniq(s) }
func (s SliceChain3[T2, T3, R]) Uniq() SliceChain3[T2, T3, R]         { return lo.Uniq(s) }
func (s SliceChain2[T3, R]) Uniq() SliceChain2[T3, R]                 { return lo.Uniq(s) }
func (s SliceChain[R]) Uniq() SliceChain[R]                           { return lo.Uniq(s) }
*/
