package main

type Voter struct {
	voterId int
}

func (v Voter) GetVoterId() int {
	return v.voterId
}

func (v Voter) SetVoterId(voterId int) {
	v.voterId = voterId
}
