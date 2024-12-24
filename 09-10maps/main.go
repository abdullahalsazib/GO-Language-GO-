package main
import "fmt"

func main() {
    fmt.Println("Maps in go")
    language := make(map[string]string)
    language["JS"] = "javaScript"
    language["JAVA"] = "JAVA"
    language["PY"] = "Python"

    fmt.Println("List of all Language's ", language)
    fmt.Println(language["JS"])

    delete(language, "PY")
    fmt.Println(language)
    
    // loops are interasting in golan

    for key, value := range language{
        fmt.Printf("For kay %v value %v\n", key, value)
    }
}
