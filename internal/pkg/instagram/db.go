package instagram

import (
	"encoding/json"
	"fmt"
)

// inserData : mock database insert
func insertData(data instagram) {
	dataJSON, _ := json.Marshal(data)

	fmt.Println(string(dataJSON))
}
