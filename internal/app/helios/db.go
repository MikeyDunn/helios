package helios

import "encoding/json"

// selectUsers : mock user select
func selectUsers() users {
	mockUsersByte := []byte(`{"users":[{"id":0,"instagram_handle":"dallasdevs"},{"id":1,"instagram_handle":""}]}`)
	mockUsers := users{}

	json.Unmarshal(mockUsersByte, &mockUsers)
	return mockUsers
}
