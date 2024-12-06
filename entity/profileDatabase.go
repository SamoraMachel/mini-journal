package entity


type ProfileDatabase map[string]Profile

func (p ProfileDatabase) AddProfile(hashKey string, profile Profile) {
	p[hashKey] = profile
}

func (p ProfileDatabase) DeleteProfile(hashKey string) {
	delete(p, hashKey)
}

func (p ProfileDatabase) SearchProfile(hashKey string) (Profile, error) {
	profile, ok := p[hashKey]
	if !ok {
		return Profile{}, ErrProfileNotFound 
	}
	return profile, nil
}