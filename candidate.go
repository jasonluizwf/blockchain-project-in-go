package main

type Candidate struct {
	name             string
	politicalPartyId int
}

func (c Candidate) GetName() string {
	return c.name
}

func (c Candidate) SetName(nameCandidate string) {
	c.name = nameCandidate
}

func (c Candidate) GetPoliticalPartyId() int {
	return c.politicalPartyId
}

func (c Candidate) SetPoliticalPartyId(partyId int) {
	c.politicalPartyId = partyId
}
