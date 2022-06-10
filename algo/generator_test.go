package algo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	
)

const userID = "e9wd9w0q-w9s9dsacx-n324n325a"

func TestGenerator(t *testing.T) {
	initialURL_1 := "https://github.com/solarcreature"
	shortURL_1 := GenerateURL(initialURL_1, userID)
	
	initialURL_2 := "https://www.linkedin.com/in/sanjeev-kotha/"
	shortURL_2 := GenerateURL(initialURL_2, userID)
	
	initialURL_3 := "https://everythingcomputerscience.com/"
	shortURL_3 := GenerateURL(initialURL_3, userID)
	
	assert.Equal(t, shortURL_1, "cJBqP5jx")
	assert.Equal(t, shortURL_2, "VACXg1aB")
	assert.Equal(t, shortURL_3, "YiqMMZ1c")
}
