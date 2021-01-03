package persistence

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// StoredItem element which is stored in DB
// represents a project
type StoredItem struct {
	ProjectName string `json:"project_name"`
	JiraQueries []string `json:"jira_queries"`
	ImageUrl string `json:"image_url"`
}

// OneFileDB it's just one file which is used for storing information
type OneFileDB struct {
	PathToFile string
}

// Validate tries to open db file
// if can't will return error
// then read file
// if can't will return error
// check length of read data
// if it's 0 will return nil
// then parse data to []StoredItem
// if can't will return error
// otherwise will return nil
func (db *OneFileDB) Validate() error {
	_, err := readFile(db.PathToFile)
	return err
}

// ReadAll returns all projects which are stored in DB
func (db *OneFileDB) ReadAll() ([]StoredItem, error) {
	return readFile(db.PathToFile)
}

func readFile(path string) ([]StoredItem, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("[ERROR] Couldn't close file at path:", path, err.Error())
		}
	}()

	data, res := ioutil.ReadAll(file)

	if res != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var parsed []StoredItem

	if err = json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}