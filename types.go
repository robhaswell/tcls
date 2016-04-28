package main

type Connection struct {
	Dest System
	Sig  Sig
}

type System struct {
	Name string
}

type Sig struct {
	Sig string
}
