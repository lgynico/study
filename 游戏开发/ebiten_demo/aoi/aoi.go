package main

type AOI interface {
	OnEnter(*Entity)
	OnExit(*Entity)
	OnMove(*Entity, *Point)
}
