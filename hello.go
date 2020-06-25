package hellogomod

import "fmt"

func Hello(who, t string) string {
	switch t {
    case "en":
        return fmt.Sprintf("Hello, %s!", who), nil
    cse "cn":
        return fmt.Sprintf("你好, %s!", who), nil
    default:
        return "", errors.New("unknown language")
    }
}
