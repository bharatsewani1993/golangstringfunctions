package main

import(
  "fmt"
  "strings"
)

func main(){
  var s1 string = "My name is Bharat and I am a programmer."
  var s2 string = "Entrepreneurship teaches you about being humble."

   //compare two string.
    fmt.Println(strings.Compare(s1,s2)) //s2 > s1
    fmt.Println(strings.Compare(s2,s1)) //s1 < s2
    fmt.Println(strings.Compare(" "," ")) // single space == Single space

   //contains a string case sensitive.
    fmt.Println(strings.Contains(s1,"BHARAT")) //do not contain.
    fmt.Println(strings.Contains(s1,"Bharat")) //do contain.

  //contains any of the character case sensitive.
    fmt.Println(strings.ContainsAny(s1,"f z")) //do not contain.
    fmt.Println(strings.ContainsAny(s1,"i a")) //do contain.

  //count the occurrence of string case sensitive.
    fmt.Println(strings.Count(s1,"i"))
    fmt.Println(strings.Count(s1,"m"))

  //string to array
    arr := strings.Fields(s1)
    for _, str := range arr {
      fmt.Println(str)
    }

  //array to string
    fmt.Println(strings.Join(arr," "))

  //return the index of substring
    fmt.Println(strings.Index(s1,"I"))
    fmt.Println(strings.Index(s1,"i"))

  //repeat string
    fmt.Println(strings.Repeat(s1,5))

  //convert string to uppercase
    fmt.Println(strings.ToTitle(s1))

   //convert string to lowercase
    fmt.Println(strings.ToLower(s1))

}
