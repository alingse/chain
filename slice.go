package chain

import (
	"github.com/samber/lo"
)

type SliceChain[R any] []R
type SliceChain2[T3, R any] []T3
type SliceChain3[T2, T3, R any] []T2
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

type PredicateF[V any] func(V, int) bool

func (s SliceChain4[T1, T2, T3, R]) Filter(predicate PredicateF[T1]) SliceChain4[T1, T2, T3, R] {
	return lo.Filter(s, predicate)
}

func (s SliceChain3[T2, T3, R]) Filter(predicate PredicateF[T2]) SliceChain3[T2, T3, R] {
	return lo.Filter(s, predicate)
}

func (s SliceChain2[T3, R]) Filter(predicate PredicateF[T3]) SliceChain2[T3, R] {
	return lo.Filter(s, predicate)
}
func (s SliceChain[R]) Filter(predicate PredicateF[R]) SliceChain[R] { return lo.Filter(s, predicate) }

/*
func (s SliceChain4[T1, T2, T3, R]) Uniq() SliceChain4[T1, T2, T3, R] { return lo.Uniq(s) }
func (s SliceChain3[T2, T3, R]) Uniq() SliceChain3[T2, T3, R]         { return lo.Uniq(s) }
func (s SliceChain2[T3, R]) Uniq() SliceChain2[T3, R]                 { return lo.Uniq(s) }
func (s SliceChain[R]) Uniq() SliceChain[R]                           { return lo.Uniq(s) }
*/
