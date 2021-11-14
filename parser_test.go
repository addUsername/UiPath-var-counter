package main

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestGetDisplayOK(t *testing.T) {

	file, err := ioutil.ReadFile("./test_input/test_displayName_Outter.txt")
	if err != nil {
		t.Fatalf("Eror getting file ")
	}
	asset1 := string(file)

	displayName := "Outter"

	mssg := getDisplayname(asset1, 15)
	if displayName != mssg {
		t.Fatalf(`getDisplayname(xaml,loc) returns %s but it shoud be %s`, mssg, displayName)
	}
}

func TestGetDisplayNameKO(t *testing.T) {

	file, err := ioutil.ReadFile("./test_input/test_displayName_ID.txt")

	if err != nil {
		t.Fatalf("Eror getting file ")
	}
	asset1 := string(file)

	id := "Sequence_1"

	mssg := getDisplayname(asset1, 0)
	if id != mssg {
		t.Fatalf(`getDisplayname(xaml,loc) returns %s but it shoud be %s`, mssg, id)
	}
}

func TestGetMembers(t *testing.T) {

	file, err := ioutil.ReadFile("./test_input/test_getMembers_OK.txt")

	if err != nil {
		t.Fatalf("Eror getting file ")
	}

	asset1 := string(file)

	arg := Member{Name: "in_OrchestratorQueueName", Class: "InArgument(x:String)", Count: 10}

	text := getMembers(asset1)

	var members Xalm
	err = xml.Unmarshal([]byte(text), &members)

	if err != nil {
		t.Fatalf(`Error xml.Unmarshall: \n%s `, err.Error())
	}
	if len(members.Members) != 2 {
		t.Fatal(`It should be members here`)
	}

	if members.Members[0].Name != arg.Name || members.Members[0].Class != arg.Class {
		t.Fatalf(` Name = %s but it shoud be %s AND %s should be %s`, members.Members[0].Name, arg.Name, members.Members[0].Class, arg.Class)
	}
}
