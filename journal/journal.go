package journal

import (
	"journal/entity"
)

var journalDatabase = make(entity.JournalDatabase)

func CreateJournal(profileKey string) error {
	_, err := journalDatabase.RetrieveJournalRecord(profileKey)
	if err == nil {
		return entity.ErrJournalAlreadyCreated
	}
	journalDatabase.CreateJournalRecord(profileKey)
	return nil
}

func AddToJournal(profileKey string, journalModel entity.JournalModel) error {
	return journalDatabase.AddJournalRecord(profileKey, journalModel)
}

func RetrieveJournal(profileKey string) ([]entity.JournalModel, error) {
	return journalDatabase.RetrieveJournalRecord(profileKey)
}

func DeleteJournal(profileKey string) error {
	_, err := journalDatabase.RetrieveJournalRecord(profileKey)
	
	if err != nil {
		return entity.ErrJournalNotFound
	}

	journalDatabase.DeleteJournalRecord(profileKey)
	return nil
}