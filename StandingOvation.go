package main

import (
        "fmt"
    "bufio"
    "os"
	  "bytes"
    "strings"
    "strconv"
)

    var tests, checkAt int
  //  var input string
	var smax, t bytes.Buffer
  var text = ""
  var testNum = ""

    var friendsInvited = 0
    var standing = 0
    var caseNumber = 0
    var sLevel = 0



func main(){
      getTests()

      if(tests < 1 || tests > 100){
            fmt.Println("Tests are between 1 and 100")
        }
    //createArrays
    inText := make([]string, tests)
    inSMax := make([]int, tests)
    out := make([]int, tests)


    for n := 0; n < tests; n++{
      inText[n] = askInput() //obtains text as Input
      inSMax[n] = getSMax(inText[n])  //gets SMax and determines where other text starts
    }

    for n := 0; n < tests; n++{
      out[n] = peopleStanding(inText[n], inSMax[n]) //obtains value of result
    }

    for n := 0; n < tests; n++{
      fmt.Print("Case #", (n+1) , ": ", out[n], "\n") //obtains value of result
    }

}

func getTests(){
  fmt.Print("Tests: ")
  consoleReader := bufio.NewReader(os.Stdin)
  testNum, _ = consoleReader.ReadString('\n')
  testNum = strings.TrimSuffix(testNum, "\n")


   for i := 0; i < (len((testNum)) - 1) ; i++ {
        if( testNum[i] != ' '){
        runes := []rune(testNum)
        t.WriteString(string(runes[i]))
        }
    }

    tests = getInt(t.String())
}

func askInput()(text string){
      fmt.Print("Input: ")
      consoleReader := bufio.NewReader(os.Stdin)
      text, _ = consoleReader.ReadString('\n')
      text = strings.TrimSuffix(text, "\n")

    return

}

func getSMax(text string)(checkAt int){
   for i := 0; i < (len((text)) - 1) ; i++ {

        if( text[i] != ' '){
          runes := []rune(text)
          smax.WriteString(string(runes[i]))
          checkAt = i+1
        } else{
          checkAt = i+1
          return
        }
    }
    return
  //  fmt.Println(smax.String())
}

func peopleStanding(inputText string, inputCheckAt int)(friendsInvited int){
    friendsInvited = 0
    standing = 0
    sLevel = 0
    runes := []rune(inputText)

    caseNumber++

    for i := inputCheckAt; i < (len((inputText)) - 1); i++ {
        if(i == inputCheckAt || getInt(string(runes[i])) == 0){
        standing += getInt(string(runes[i]))
        sLevel++
        } else if(standing >= sLevel){
        standing += getInt(string(runes[i]))
        sLevel++
        } else{
        friendsInvited += sLevel - standing
        standing += (friendsInvited + getInt(string(runes[i])))
        sLevel++
        }
    }

    return
}

func getInt(x string) (y int){
  y, err := strconv.Atoi(x)
     if err != nil {
             fmt.Println(err)
             return
           }
        return
}
