package mdb

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Document is an entry in the database
type Document struct {
	Created string `json:"created"`
	Edited  string `json:"edited"`
	Doc     string `json:"doc"`
}

// Database is the database in memory
type Database struct {
	Data map[string]Document `json:"data"`
}

// NewDatabase creates a new empty Database struct
func NewDatabase() *Database {
	return &Database{
		Data: map[string]Document{},
	}
}

// Load loads database file into a Database instance
func Load(db *Database, fn string) error {
	var err error
	var f *os.File
	// if file does not exist, create
	if _, err = os.Stat(fn); err != nil {
		f, err = os.Create(fn)
		if err != nil {
			return err
		}
	} else {
		f, err = os.Open(fn)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	for {
		if err := decoder.Decode(db); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

// Commit writes the Database instance in memory to file
func Commit(db *Database, fn string) error {
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err = encoder.Encode(db); err != nil {
		return err
	}
	return nil
}

// Get fetches a single document from a Database instance
func Get(db *Database, id string) (*Document, error) {
	if val, ok := db.Data[id]; ok {
		return &val, nil
	} else {
		return nil, fmt.Errorf("document id doesn't exist: %s", id)
	}
}

// Put inserts a new Document into a Database
func Put(db *Database, doc string) (string, error) {
	pid, err := uuid()
	if err != nil {
		return "", err
	}
	db.Data[pid] = Document{
		Created: time.Now().UTC().String(),
		Edited:  time.Now().UTC().String(),
		Doc:     doc,
	}
	return pid, nil
}

// Update updates an existing document in a Database
func Update(db *Database, id, doc string) error {
	currData, err := Get(db, id)
	if err != nil {
		return err
	}

	currData.Edited = time.Now().UTC().String()
	currData.Doc = doc

	db.Data[id] = *currData
	return nil
}

// Delete removes a Document from a Database
func Delete(db *Database, id string) error {
	if _, err := Get(db, id); err != nil {
		return err
	}
	delete(db.Data, id)
	return nil
}

// uuid creates a pseudo uuid V4 string
func uuid() (u string, err error) {
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		return
	}

	// 13th character is "4"
	b[6] = (b[6] | 0x40) & 0x4F
	// 17th character is 8, 9, a, or b
	b[8] = (b[8] | 0x80) & 0xBF

	u = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
