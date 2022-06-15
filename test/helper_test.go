package test

import (
	"fmt"
	"simple-demo/helper"
	"testing"
)

func TestHelper(t *testing.T) {
	res, _ := helper.AnalyseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6InRlc3QiLCJ1c2VySUQiOiIzIiwiZXhwIjoxNjU1MzE0MjA5LCJpYXQiOjE2NTUzMTA2MDksImlzcyI6IkRvdXlpbiIsInN1YiI6InVzZXJUb2tlbiJ9.WONUq8YCC4KgU9_iQIJqqZIl1Q3lC8kGL4qEMyY2FvU")
	fmt.Println(res.UserID)
	fmt.Println(res.UserName)
}