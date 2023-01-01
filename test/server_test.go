package runTests

import (
	"example.com/blogArch/gateway/models"
	"example.com/blogArch/gateway/responses"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"bytes"
	"log"
)


func TestMainPage(t *testing.T) {
	assert := require.New(t)
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	assert.Equal(r.Status, "In Main Page!", "they should be equal")
}

// func TestRegister(t *testing.T) {
//   assert := require.New(t)
// 	entry := models.LoginModel {
// 		Username: "user1",
// 		Password: "password",
// 	}
// 	data, _ := json.Marshal(entry)
// 	log.Println(string(data))
// 	res, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	var r responses.StatusResponse
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	log.Printf("Status: %s\n", r.Status)
//   assert.Equal("user1", r.Status, "they should be equal")

// 	res, err = http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	decoder = json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	log.Printf("Status: %s\n", r.Status)
// 	assert.Equal("", r.Status, "they should be equal")
// }

func TestLogin(t *testing.T) {
  assert := require.New(t)
	entry := models.LoginModel {
		Username: "user1",
		Password: "password",
	}
	data, _ := json.Marshal(entry)
	log.Println(string(data))
	res, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	log.Printf("Status: %s\n", r.Status)
  assert.Equal("Logged in", r.Status, "they should be equal")

	entry.Username = "Not Registered User"
	data, _ = json.Marshal(entry)
	res, err = http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	log.Printf("Status: %s\n", r.Status)
	assert.Equal("Username or password does not match", r.Status, "they should be equal")

	entry.Username = "user1"
	entry.Password = "IncorrectPassword"
	data, _ = json.Marshal(entry)
	res, err = http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	log.Printf("Status: %s\n", r.Status)
	assert.Equal("Username or password does not match", r.Status, "they should be equal")
}

func TestEntry(t *testing.T) {
  assert := require.New(t)
	entry := models.EntryModel {
		Entry: "This is a positive Entry!",
	}
	data, _ := json.Marshal(entry)
	res, err := http.Post("http://localhost:8080/entry", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
  // assert equality
  assert.Equal("Inserted entry!", r.Status, "they should be equal")

	entry.Entry = "This is a negative entry. You suck!"
	data, _ = json.Marshal(entry)
	res, err = http.Post("http://localhost:8080/entry", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	assert.Equal("Entry not inserted. Please refrain from toxic comments.", r.Status,"they should be equal")
}

func TestProfile(t *testing.T) {
  assert := require.New(t)
	req, _ := http.NewRequest("GET", "http://localhost:8080/profile", nil)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)
	var r responses.ProfileResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	for _, entry := range r.Entries {
		log.Printf("%s", entry)
		assert.Equal("This is a positive Entry!", entry, "they should be equal")
	}
}


// func TestAll(t *testing.T) {
//   assert := require.New(t)

//   // assert equality
//   assert.Equal(123, 123, "they should be equal")

//   // assert inequality
//   assert.NotEqual(123, 456, "they should not be equal")
//   // assert for nil (good for errors)

//   // assert for not nil (good when you expect something)
//   // if assert.NotNil(nil) {
// 		// log.Println("Hi")
//     // now we know that object isn't nil, we are safe to make
//     // further assertions without causing any errors
//     assert.Equal("Something", "Something")
//   // }
// }