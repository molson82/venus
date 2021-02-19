package data

import "github.com/molson82/venus/models"

// UserMockData : returns mock user data
func UserMockData() []models.User {
	users := []models.User{
		{
			ID:        1,
			Firstname: "John",
			LastName:  "Smith",
			Address: models.Address{
				Street1: "123 Maple Street",
				Street2: "",
				City:    "Chicago",
				State:   "IL",
				ZipCode: 34988,
			},
			Phone: "111-867-3099",
		},
		{
			ID:        2,
			Firstname: "Mike",
			LastName:  "Johnson",
			Address: models.Address{
				Street1: "8 Cactus Lane",
				Street2: "Apt 44",
				City:    "Austin",
				State:   "TX",
				ZipCode: 44444,
			},
			Phone: "987-354-6687",
		},
		{
			ID:        3,
			Firstname: "Sarah",
			LastName:  "Claire",
			Address: models.Address{
				Street1: "98 Nova Circle",
				Street2: "Apt 35",
				City:    "New York",
				State:   "NY",
				ZipCode: 23877,
			},
			Phone: "713-999-4453",
		},
		{
			ID:        4,
			Firstname: "Brady",
			LastName:  "Bunch",
			Address: models.Address{
				Street1: "8615 Broadway",
				Street2: "",
				City:    "Hollywood",
				State:   "CA",
				ZipCode: 90210,
			},
			Phone: "444-908-1234",
		},
	}

	return users
}
