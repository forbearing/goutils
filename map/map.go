package mymap

import (
	"fmt"
	"sort"
	"strings"
)

// String construct map[string]string to string.
func String(mymap map[string]string) string {
	pairs := []string{}

	for k, v := range mymap {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}

// Set map,
// eg: "name=hybfkuf,age=100"
// eg: "name=hybfkuf, age=100"
func Set(mymap map[string]string, value string) error {

	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) == 2 {
			mymap[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		} else {
			mymap[strings.TrimSpace(arr[0])] = ""
		}
	}

	return nil
}
