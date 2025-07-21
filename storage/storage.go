package storage

type UserPreference struct {
	UserID    int64
	Latitude  float32
	Longitude float32
	Artist    string
}

var userPrefs = make(map[int64]UserPreference)

func SaveUserLocation(userID int64, lat, lon float32) {
	pref := userPrefs[userID]
	pref.Latitude = lat
	pref.Longitude = lon
	userPrefs[userID] = pref
}

func SaveUserArtist(userID int64, artist string) {
	pref := userPrefs[userID]
	pref.Artist = artist
	userPrefs[userID] = pref
}

func GetAllUsers() []UserPreference {
	users := make([]UserPreference, 0, len(userPrefs))
	for _, u := range userPrefs {
		users = append(users, u)
	}
	return users
}
