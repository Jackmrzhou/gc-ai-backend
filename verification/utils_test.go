package verification

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var testMail = "594780735@qq.com"

func TestGenerateCode(t *testing.T) {
	for i := 0; i < 100; i++{
		if code := generateCode("email"); len(code) != 6{
			t.Errorf("wrong length of code : %s", code)
		}
	}
}

func TestSendCode(t *testing.T) {
	if code,err := SendCode(testMail); err != nil{
		t.Error(err)
	}else{
		c,_ := models.QueryVeriCode(testMail)
		assert.Equal(t, c.Code, code)
		fmt.Println(c)
	}
	models.DeleteVeriCodeDebug(testMail)
}

func TestCheckAndDelCode(t *testing.T) {
	models.CreateVeriCode(testMail, "123456")
	if !CheckAndDelCode(testMail, "123456"){
		t.Error("Verification should passed!")
	}

	models.CreateVeriCodeDebug(&models.VerificationCode{
		Email:testMail,
		Code:"123456",
		Timestamp:time.Now().Add(time.Duration(-11)*time.Minute),
	})
	if CheckAndDelCode(testMail, "123456"){
		t.Error("Verification should failed due to the time limit!")
	}
}

func TestMain(m *testing.M) {
	if err:=conf.LoadConf("../conf/app.conf"); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	models.OpenDB()
	m.Run()
}