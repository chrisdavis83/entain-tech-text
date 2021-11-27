package db

const (
	racesList = "list"
	findRace  = "find"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				(case when advertised_start_time > DATE('now') then '0' else '1' end) as status
			FROM races
		`,
		findRace: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				(case when advertised_start_time > DATE('now') then '0' else '1' end) as status
			FROM races
		`,
	}
}
