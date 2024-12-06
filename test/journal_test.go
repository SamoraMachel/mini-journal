package test

import (
	"testing"
	"journal/journal"
	"journal/lib"
	"journal/entity"
)

const (
	testUserHash = "testUserHash"
)
var testSampleJournal = entity.JournalModel{Title: "Sample", Date: "Today", Text: "Text"}


func TestCreateJournal(t *testing.T) {	
	
	t.Run("create a journal for the first time", func(t *testing.T) {
		err := journal.CreateJournal(testUserHash)
		lib.AssertNoError(t, err)
	})

	t.Run("create a journal when journal is already available", func (t *testing.T){
		err := journal.CreateJournal(testUserHash)
		lib.AssertError(t, err, entity.ErrJournalAlreadyCreated)
	})
}

func TestAddingToJournal(t *testing.T) {
	t.Run("add to an existing journal", func(t *testing.T) {
		err := journal.AddToJournal(testUserHash, testSampleJournal)
		lib.AssertNoError(t, err)

		journalList, retrievalError := journal.RetrieveJournal(testUserHash)
		lib.AssertNoError(t, retrievalError)

		if len(journalList) < 1 {
			t.Error("journal entry does not reflect in record")
		}
	})

	t.Run("add to an non-existing journal", func(t *testing.T) {
		err := journal.AddToJournal("random_hash", testSampleJournal)
		lib.AssertError(t, err, entity.ErrJournalNotFound)
	})
}

func TestRetrieveJournal(t *testing.T) {
	t.Run("retrieve existing journal", func(t *testing.T) {
		_, err := journal.RetrieveJournal(testUserHash)
		lib.AssertNoError(t, err)

	})

	t.Run("retrieve non-existing journal", func(t *testing.T) {
		_, err := journal.RetrieveJournal("random_hash")
		lib.AssertError(t, err, entity.ErrJournalNotFound)
	})
}

func TestDeleteJournal(t *testing.T) {
	t.Run("delete existing journal", func(t *testing.T) {
		err := journal.DeleteJournal(testUserHash)
		lib.AssertNoError(t, err)
	})

	t.Run("delete a non-existing journal", func(t *testing.T) {
		err := journal.DeleteJournal(testUserHash)
		lib.AssertError(t, err, entity.ErrJournalNotFound)
	})
}
