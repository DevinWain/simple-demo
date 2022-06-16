package test

import (
	"fmt"
	"simple-demo/helper"
	"testing"
)

func TestHelper(t *testing.T) {
	res, _ := helper.AnalyseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IndhaW4iLCJ1c2VySUQiOiI1IiwiZXhwIjoxNjU1MzgyNTY1LCJpYXQiOjE2NTUzNzg5NjUsImlzcyI6IkRvdXlpbiIsInN1YiI6InVzZXJUb2tlbiJ9.s0UZZJ4umHulT9UfWtQkvXq9wWeQz-ZAWxRMx3LN6Aw")
	fmt.Println(res.UserID)
	fmt.Println(res.UserName)
}