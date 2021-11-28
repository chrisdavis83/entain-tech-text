package db

const (
	sportsList = "list"
)

func getSportQueries() map[string]string {
	return map[string]string{
		sportsList: `
			SELECT 
				id, 
				name, 
				sportType, 
				visible, 
				advertised_start_time
			FROM sports
		`,
	}
}
