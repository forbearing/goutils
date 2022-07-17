package mymap

import "testing"

func TestMymap(t *testing.T) {
	mymap := map[string]string{
		"name": "hybfkuf",
		"age":  "10000",
	}

	t.Log(String(mymap))

	val := "addr=bulabula, hehe=haha"
	t.Log(Set(mymap, val))
	t.Log(String(mymap))

	// Output:

	//=== RUN   TestMymap
	//    map_test.go:11: age=10000,name=hybfkuf
	//    map_test.go:14: <nil>
	//    map_test.go:15: addr=bulabula,age=10000,hehe=haha,name=hybfkuf
	//--- PASS: TestMymap (0.00s)
	//PASS
	//ok      github.com/forbearing/goutils/map       0.131s
}
