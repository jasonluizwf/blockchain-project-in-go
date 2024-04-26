package main

type PoliticalParty struct {
	name   string
	number int
}

func (p PoliticalParty) GetName() string {
	return p.name
}

func (p PoliticalParty) SetName(name string) {
	p.name = name
}

func (p PoliticalParty) GetNumber() int {
	return p.number
}

func (p PoliticalParty) SetNumber(number int) {
	p.number = number
}
