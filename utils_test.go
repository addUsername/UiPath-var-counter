package main

import (
	"fmt"
	"testing"
)

func TestXalmToJson(t *testing.T) {

	asset1 := "./test_input/GeneralTestCase.xaml"

	value := `[{"Scopes":[{"Vars":[{"Class":"x:String","Name":"Output","Default":"","Count":0}],"Name":"GeneralTestCase"}],"Filename":"./test_input/GeneralTestCase.xaml","Arguments":{"Arguments":null}},{"Scopes":[{"Vars":[{"Class":"x:String","Name":"Output","Default":"","Count":0}],"Name":"GeneralTestCase"}],"Filename":"./test_input/GeneralTestCase.xaml","Arguments":{"Arguments":null}}]`
	xaml1 := ParseFile(asset1)
	xalms := [2]Xalm{xaml1, xaml1}

	out := xamlToJson(xalms[:])

	fmt.Println(out)

	if value != out {
		t.Fatal("test failed")
	}

}
