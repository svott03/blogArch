package runTests

import (
	"example.com/blogArch/gateway/models"
	"example.com/blogArch/gateway/responses"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"bytes"
	"fmt"
)


func TestMainPage(t *testing.T) {
	assert := require.New(t)
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)
	var r responses.MainResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	assert.Equal(r.Output, "In Main Page!", "they should be equal")
}

// func TestRegister(t *testing.T) {
//   assert := require.New(t)

//   // assert equality
//   assert.Equal(123, 123, "they should be equal")

//   // assert inequality
//   assert.NotEqual(123, 456, "they should not be equal")
//   // assert for nil (good for errors)

//   // assert for not nil (good when you expect something)
//   // if assert.NotNil(nil) {
// 		// fmt.Println("Hi")
//     // now we know that object isn't nil, we are safe to make
//     // further assertions without causing any errors
//     assert.Equal("Something", "Something")
//   // }
// }

func TestRegister(t *testing.T) {
  assert := require.New(t)
	entry := models.LoginModel {
		Username: "user1",
		Password: "password",
	}
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
	res, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	var r responses.StatusResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	fmt.Printf("Status: %s\n", r.Status)
  assert.Equal("user1", r.Status, "they should be equal")



	res, err = http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	fmt.Printf("Status: %s\n", r.Status)
	assert.Equal("", r.Status, "they should be equal")
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
  assert.Equal("POSITIVE\n", r.Status, "they should be equal")

	entry.Entry = "This is a negative entry. You suck!"
	data, _ = json.Marshal(entry)
	res, err = http.Post("http://localhost:8080/entry", "application/json", bytes.NewBuffer(data))
	assert.Nil(err)
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	assert.Nil(err)
	assert.Equal("NEGATIVE\n", r.Status,"they should be equal")
}

// func TestProfile(t *testing.T) {
//   assert := require.New(t)

//   // assert equality
//   assert.Equal(123, 123, "they should be equal")

//   // assert inequality
//   assert.NotEqual(123, 456, "they should not be equal")
//   // assert for nil (good for errors)

//   // assert for not nil (good when you expect something)
//   // if assert.NotNil(nil) {
// 		// fmt.Println("Hi")
//     // now we know that object isn't nil, we are safe to make
//     // further assertions without causing any errors
//     assert.Equal("Something", "Something")
//   // }
// }


// func TestAll(t *testing.T) {
//   assert := require.New(t)

//   // assert equality
//   assert.Equal(123, 123, "they should be equal")

//   // assert inequality
//   assert.NotEqual(123, 456, "they should not be equal")
//   // assert for nil (good for errors)

//   // assert for not nil (good when you expect something)
//   // if assert.NotNil(nil) {
// 		// fmt.Println("Hi")
//     // now we know that object isn't nil, we are safe to make
//     // further assertions without causing any errors
//     assert.Equal("Something", "Something")
//   // }
// }