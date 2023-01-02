package runTests

import (
	"bytes"
	"encoding/json"
	"example.com/blogArch/gateway/models"
	"example.com/blogArch/gateway/responses"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"testing"
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

func testRegister(assert *require.Assertions) {
	// Register Success
	entry := models.LoginModel{
		Username: "user1",
		Password: "password",
	}
	data, _ := json.Marshal(entry)
	log.Println(string(data))
	res, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	log.Printf("Status: %s\n", r.Status)
	assert.Equal("", r.Status, "they should be equal")
	// TODO .Equal user1 ^^

	// Register Fail
	res, err = http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	log.Printf("Status: %s\n", r.Status)
	assert.Equal("", r.Status, "they should be equal")
}

func testFailEntry(assert *require.Assertions) {
	// Post Entry Auth Error
	entry := models.EntryModel{
		Entry: "This is a positive Entry!",
	}
	data, _ := json.Marshal(entry)
	res, err := http.Post("http://localhost:8080/admin/entry?token=UnauthorizedToken", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	log.Printf("FailedEntry Header %d\n", res.StatusCode)
	err = decoder.Decode(&r)
	assert.Nil(err)
	// assert equality
	assert.Equal(400, res.StatusCode, "they should be equal")
}

func testLogin(assert *require.Assertions) string {
	// Login Fail

	// Login Success
	return ""
}

func testEntry(token string, assert *require.Assertions) {
	// Post Entry Postive

	// Post Entry Negative

	// Post Entry Positive
}

func testProfile(token string, assert *require.Assertions) {
	// Query Profile
}

// func TestLogin(t *testing.T) {
//   assert := require.New(t)
// 	entry := models.LoginModel {
// 		Username: "user1",
// 		Password: "password",
// 	}
// 	data, _ := json.Marshal(entry)
// 	log.Println(string(data))
// 	res, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	var r responses.StatusResponse
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	log.Printf("Status: %s\n", r.Status)
//   assert.Equal("Logged in", r.Status, "they should be equal")

// 	entry.Username = "Not Registered User"
// 	data, _ = json.Marshal(entry)
// 	res, err = http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	decoder = json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	log.Printf("Status: %s\n", r.Status)
// 	assert.Equal("Username or password does not match", r.Status, "they should be equal")

// 	entry.Username = "user1"
// 	entry.Password = "IncorrectPassword"
// 	data, _ = json.Marshal(entry)
// 	res, err = http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	decoder = json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	log.Printf("Status: %s\n", r.Status)
// 	assert.Equal("Username or password does not match", r.Status, "they should be equal")
// }

// func TestEntry(t *testing.T) {
//   assert := require.New(t)
// 	entry := models.EntryModel {
// 		Entry: "This is a positive Entry!",
// 	}
// 	data, _ := json.Marshal(entry)
// 	res, err := http.Post("http://localhost:8080/admin/entry", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	var r responses.StatusResponse
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
//   // assert equality
//   assert.Equal("Inserted entry!", r.Status, "they should be equal")

// 	entry.Entry = "This is a negative entry. You suck!"
// 	data, _ = json.Marshal(entry)
// 	res, err = http.Post("http://localhost:8080/admin/entry", "application/json", bytes.NewBuffer(data))
// 	assert.Nil(err)
// 	decoder = json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	assert.Equal("Entry not inserted. Please refrain from toxic comments.", r.Status,"they should be equal")
// }

// func TestProfile(t *testing.T) {
//   assert := require.New(t)
// 	req, _ := http.NewRequest("GET", "http://localhost:8080/admin/profile", nil)
// 	res, err := http.DefaultClient.Do(req)
// 	assert.Nil(err)
// 	var r responses.ProfileResponse
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&r)
// 	assert.Nil(err)
// 	for _, entry := range r.Entries {
// 		log.Printf("%s", entry)
// 		assert.Equal("This is a positive Entry!", entry, "they should be equal")
// 	}
// }

func TestAll(t *testing.T) {
	assert := require.New(t)
	testRegister(assert)
	testFailEntry(assert)
	// may not be a string
	token := testLogin(assert)
	testEntry(token, assert)
	testProfile(token, assert)
}
