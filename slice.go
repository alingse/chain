package chain

import (
	"github.com/samber/lo"
)

type SliceChain[T1 any] SliceChain4[T1, any, any, any]
type SliceChain2[T1, T2 any] SliceChain4[T1, T2, any, any]
type SliceChain3[T1, T2, T3 any] SliceChain4[T1, T2, T3, any]
type SliceChain4[T1, T2, T3, T4 any] []T1

type MapF[T, R any] func(T, int) R

func (s SliceChain4[T1, T2, T3, R]) Map(f MapF[T1, T2]) SliceChain3[T2, T3, R] {
	return lo.Map(s, f)
}

func (s SliceChain3[T1, T2, R]) Map(f MapF[T1, T2]) SliceChain2[T2, R] {
	return lo.Map(s, f)
}

func (s SliceChain2[T1, R]) Map(f MapF[T1, R]) SliceChain[R] {
	return lo.Map(s, f)
}
