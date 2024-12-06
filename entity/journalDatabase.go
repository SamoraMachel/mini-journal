package entity

type JournalDatabase map[string][]JournalModel

func (j JournalDatabase) CreateJournalRecord(journalKey string) {
	j[journalKey] = make([]JournalModel, 0)
}

func (j JournalDatabase) AddJournalRecord(journalKey string, record JournalModel) error {
	journal, err := j.RetrieveJournalRecord(journalKey)
	if err != nil {
		return ErrJournalNotFound
	}
	newJournal := append(journal, record)
	j[journalKey] = newJournal
	return nil
}

func (j JournalDatabase) RetrieveJournalRecord(journalKey string) ([]JournalModel, error) {
	journal, ok := j[journalKey]
	if !ok {
		return []JournalModel{}, ErrJournalNotFound
	}
	return journal, nil
}

func (j JournalDatabase) DeleteJournalRecord(journalKey string) {
	delete(j, journalKey)
}