package runTests

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type MainResponse struct {
	Output string
}

func TestMainPage(t *testing.T) {
	assert := require.New(t)
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)
	var r MainResponse
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

// func TestLogin(t *testing.T) {
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

// func TestEntry(t *testing.T) {
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
