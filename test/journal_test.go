package test

import (
	"testing"
	"journal/journal"
	"journal/test"
)

const (
	testJournalName = "Test Journal"
	testUserHash = "testUserHash"
)


func TestCreateJournal(t *testing.T) {	
	err := journal.CreateJournal(testUserHash, testJournalName)
	test.AssertNoError(t, err)
}

func TestRetrieveJournal(t *testing.T) {

}

func TestDeleteJournal(t *testing.T) {

}

func TestUpdateJournal(t *testing.T) {

}
