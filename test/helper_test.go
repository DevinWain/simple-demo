package test

import (
	"fmt"
	"simple-demo/helper"
	"testing"
)

func TestHelper(t *testing.T) {
	res, _ := helper.AnalyseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImxpdXppa2FuZyIsInVzZXJJRCI6IjAiLCJleHAiOjE2NTUzNjY5MDgsImlhdCI6MTY1NTM2MzMwOCwiaXNzIjoiRG91eWluIiwic3ViIjoidXNlclRva2VuIn0.pSajjDPVhxHdFQBy9c_LXgUb1szS5aG-37Sfm1Lh9GU")
	fmt.Println(res.UserID)
	fmt.Println(res.UserName)
}