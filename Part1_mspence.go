package main
 
import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
    "math"
)
 
var acceptableVocabulary = []string{"plus", "minus", "times", "divide", "modulo", "power", "(", ")", "{", "}", "[", "]"}
var EOF = -1
var ROMAN_NUMERAL = 1
var PLUS = 2
var MINUS = 3
var TIMES = 4
var DIVIDE = 5
var MODULO = 6
var POWER = 7
var OPEN_PARENTHESIS = 8
var CLOSED_PARENTHESIS = 9
var re = regexp.MustCompile(`^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)


// Helps with checking the input
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}

func convertToToken(str string) int {
    re := regexp.MustCompile(`^(I|V|X|L|C|D|M)+$`)
    if (re.FindStringIndex(str) != nil) {
        return ROMAN_NUMERAL
    } else if (str == "plus") {
        return PLUS
    } else if (str == "minus") {
        return MINUS
    } else if (str == "times") {
        return TIMES
    } else if (str == "divide") {
        return DIVIDE
    } else if (str == "modulo") {
        return MODULO
    } else if (str == "power") {
        return POWER
    } else if (str == "(" || str == "{" || str == "[") {
        return OPEN_PARENTHESIS
    } else if (str == ")" || str == "}" || str == "]") {
        return CLOSED_PARENTHESIS
    } else {
        return EOF
    }
}

var num = map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

var numInv = map[int]string{
    1000: "M",
    900:  "CM",
    500:  "D",
    400:  "CD",
    100:  "C",
    90:   "XC",
    50:   "L",
    40:   "XL",
    10:   "X",
    9:    "IX",
    5:    "V",
    4:    "IV",
    1:    "I",
}

var maxTable = []int{
    1000,
    900,
    500,
    400,
    100,
    90,
    50,
    40,
    10,
    9,
    5,
    4,
    1,
}

type Roman struct{}

func NewRoman() *Roman {
    return &Roman{}
}

func (r *Roman) ToNumber(n string) int {
    out := 0
    ln := len(n)
    for i := 0; i < ln; i++ {
        c := string(n[i])
        vc := num[c]
        if i < ln-1 {
            cnext := string(n[i+1])
            vcnext := num[cnext]
            if vc < vcnext {
                out += vcnext - vc
                i++
            } else {
                out += vc
            }
        } else {
            out += vc
        }
    }
    return out
}

func (r *Roman) ToRoman(n int) string {
    out := ""
    for n > 0 {
        v := highestDecimal(n)
        out += numInv[v]
        n -= v
    }
    return out
}

func highestDecimal(n int) int {
    for _, v := range maxTable {
        if v <= n {
            return v
        }
    }
    return 1
}

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

// 0 <= index <= len(a)
func insert(a []string, index int, value string) []string {
    if len(a) == index { // nil or empty slice or after last element
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...) // index < len(a)
    a[index] = value
    return a
}

// Checks for malformed input
func checkMalformedness(s []string) (bool, string) {
    message := ""
    errorLocation := 0
    par := false
    for i, a := range s[0:] {
        if (!contains(acceptableVocabulary, a) && re.FindStringIndex(a) == nil) {
            errorLocation = i
            for q := range s[0:] {
                if (q > 0 && (s[q-1] != "(" && s[q-1] != "{" && s[q-1] != "[") && (s[q] != ")" && s[q] != "}" && s[q] != "]")) {
                    message += " "
                }
                message += s[q]
            }
            message += "\n"
            if (i > 0) {
                message += " "
            }
            for j := range s[0:] {
                if (j == errorLocation) {
                    if (par) {
                        message += " "
                    }
                    message += "^\n"
                    break
                }
                if (j > 0 && (s[j-1] == "(" || s[j-1] == "{" || s[j-1] == "[")) {
                    par = true
                }
                if (j > 0 && (s[j-1] != "(" && s[j-1] != "{" && s[j-1] != "[") && (s[j] != ")" && s[j] != "}" && s[j] != "]") && (s[j] != "(" && s[j] != "{" && s[j] != "[")) {
                    message += " "
                }
                for p := 0; p < len(s[j]); p += 1 {
                    message += " "
                }
            }
            message += "Quid dicis? You offend Caesar with your sloppy lexical habits!\n"
            return false, message
        }
        i += 1
    }
    return true, message
}

// Substiute input to math
func substituteToMath(str string) string {
    if (re.FindStringIndex(str) != nil) {
        r := NewRoman()
        return strconv.Itoa(r.ToNumber(str))
    } else if (str == "plus") {
        return "+"
    } else if (str == "minus") {
        return "-"
    } else if (str == "times") {
        return "*"
    } else if (str == "divide") {
        return "/"
    } else if (str == "modulo") {
        return "%"
    } else if (str == "power") {
        return "^"
    } else if (str == "{") {
        return "("
    } else if (str == "}") {
        return ")"
    } else if (str == "[") {
        return "("
    } else if (str == "]") {
        return ")"
    } else if (str == "(") {
        return "("
    } else if (str == ")") {
        return ")"
    }  else {
        return "Error!"
    }
}

func setUpInputLexical(str string) []string {
    // Removing two " characters
    str = strings.Replace(str, "\"", "", 2)

    // Add spaces before and after any parenthesis type to make it it's own lexeme
    str = strings.Replace(str, "(", " ( ", -1)
    str = strings.Replace(str, ")", " ) ", -1)
    str = strings.Replace(str, "{", " { ", -1)
    str = strings.Replace(str, "}", " } ", -1)
    str = strings.Replace(str, "[", " [ ", -1)
    str = strings.Replace(str, "]", " ] ", -1)
    
    // Create arguments array which is each lexeme from the input
    temp := strings.Split(str, " ")
    var arguments []string
    index := 0
    for i, a := range temp[0:] {
        if(temp[i] != "") {
            arguments = append(arguments,a)
            index++
        }
    } 
    return arguments
}

func setUpInputSyntactical(s []string) []int {
    // Substitute array
    lexemeArray := s
    lexemeArray = append(lexemeArray, "EOF")

    var tokenArray []int
    for i := range lexemeArray[0:] {
        tokenArray = append(tokenArray, convertToToken(lexemeArray[i]))
    }

    //for i := range tokenArray[0:] {
        //fmt.Printf("%d -> Next token is: %d Next lexeme is %s\n", i, tokenArray[i], lexemeArray[i])
    //}
    //fmt.Printf("\n")
    return tokenArray
}

func syntacticalAnalysis(s []string, raw string) (bool, string) {
    lexemeArray := s
    lexemeArray = append(lexemeArray, "EOF")

    var tokenArray []int
    for i := range lexemeArray[0:] {
        tokenArray = append(tokenArray, convertToToken(lexemeArray[i]))
    }

    var message string
    for q := range tokenArray[0:len(tokenArray)-1] {
        if (q > 0 && (s[q-1] != "(" && s[q-1] != "{" && s[q-1] != "[") && (s[q] != ")" && s[q] != "}" && s[q] != "]")) {
            message += " "
        }
        message += lexemeArray[q]
    }
    message += "\n"
    var spaces string

    if (len(tokenArray) < 2 || len(tokenArray) % 2 == 1) {
        iter := 0
        for iter < len(raw) {
            spaces += " "
            iter++
        }
        return false, raw + "\n" + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
    }

    numParenthesis := 0
    if (tokenArray[0] != 1 && tokenArray[0] != 8) {
        return false, raw + "\n^\nQuid dicis? True Romans would not understand your syntax!\n"
    }
    for i := range tokenArray[0:] {
        for p := 0; p < len(lexemeArray[i]); p += 1 {
            spaces += " "
        }
        if (numParenthesis < 0) {
            spaces += " "
            return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
        }
        if (tokenArray[i] == 1) {
            if (tokenArray[i+1] == 8) {
                if (tokenArray[i+1] != 9) {
                    spaces += " "
                }
                return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
            }
            spaces += " "
        } else if (tokenArray[i] == 2 || tokenArray[i] == 3 || tokenArray[i] == 4 || tokenArray[i] == 5 || tokenArray[i] == 6 || tokenArray[i] == 7) {
            if (tokenArray[i+1] == -1 || tokenArray[i+1] == 2 || tokenArray[i+1] == 3 || tokenArray[i+1] == 4 || tokenArray[i+1] == 5 || tokenArray[i+1] == 6 || tokenArray[i+1] == 7 || tokenArray[i+1] == 9) {
                if (tokenArray[i+1] != 9 && tokenArray[i+1] != -1) {
                    spaces += " "
                }
                return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
            }
            spaces += " "
        } else if (tokenArray[i] == 8) {
            numParenthesis++
            if (tokenArray[i+1] == -1 || tokenArray[i+1] == 2 || tokenArray[i+1] == 3 || tokenArray[i+1] == 4 || tokenArray[i+1] == 5 || tokenArray[i+1] == 6 || tokenArray[i+1] == 7 || tokenArray[i+1] == 8 || tokenArray[i+1] == 9) {
                if (tokenArray[i+1] != 9 && i != 0) {
                    spaces += " "
                }
                return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
            }
            spaces += " "
        } else if (tokenArray[i] == 9) {
            numParenthesis--
            if (tokenArray[i+1] == 1 || tokenArray[i+1] == 8) {
                if (tokenArray[i+1] != 9) {
                    spaces += " "
                }
                return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
            }
            spaces += " "
        } else if (tokenArray[i] == -1) {
            return true, ""
        }
    }
    spaces += " "
    return false, message + spaces + "^\nQuid dicis? True Romans would not understand your syntax!\n"
}

func subCalculate(arguments []string, pureArguments string) (int, int) {
    errorPlacement := 0
    i := 0
    numDivides := 0
    dividesPlaces := 0
    for i := range arguments[0:] {
        if (arguments[i] == "/") {
            numDivides++
        }
        dividesPlaces += len(arguments[i])
    }  
    numMinus := 0
    minusPlaces := 0
    for i := range arguments[0:] {
        if (arguments[i] == "-") {
            numMinus++
        }
        minusPlaces += len(arguments[i])
    }  
    dividePlace := strings.Index(pureArguments, "divide")
    moduloPlace := strings.Index(pureArguments, "modulo")
    minusPlace := strings.Index(pureArguments, "minus")

    for i < len(arguments) {
        if (arguments[i] == "^") {
            if (i < len(arguments)-2 && arguments[i+2] == "^") {
                i += 2
            }
            prev, _ := strconv.ParseFloat(arguments[i-1], 1)
            next, _ := strconv.ParseFloat(arguments[i+1], 1)
            value := math.Pow(prev, next)
            arguments = remove(arguments, i+1)
            arguments = remove(arguments, i)
            arguments = remove(arguments, i-1)
            arguments = insert(arguments, i-1, strconv.Itoa(int(value)))
            i = 0
        }
        i++
    }
    i = 0
    // for i := range arguments[0:] {
    //     fmt.Printf("%s", arguments[i])
    // }  
    // fmt.Printf("\n")

    foundDivides := 0
    for i < len(arguments) {
        errorPlacement += len(arguments[i])
        if (arguments[i] == "*" || arguments[i] == "/" || arguments[i] == "%") {
            prev, _ := strconv.Atoi(arguments[i-1])
            next, _ := strconv.Atoi(arguments[i+1])
            if (arguments[i] == "*") {
                value := prev * next
                arguments = remove(arguments, i+1)
                arguments = remove(arguments, i)
                arguments = remove(arguments, i-1)
                arguments = insert(arguments, i-1, strconv.Itoa(value))
                i = 0
            } else  if (arguments[i] == "/") {
                value := prev / next
                if (value == 0) {
                    return 0, dividePlace
                }
                arguments = remove(arguments, i+1)
                arguments = remove(arguments, i)
                arguments = remove(arguments, i-1)
                arguments = insert(arguments, i-1, strconv.Itoa(value))
                foundDivides++
                i = 0
            } else {
                value := prev % next
                if (value == 0) {
                    return 0, moduloPlace
                }
                arguments = remove(arguments, i+1)
                arguments = remove(arguments, i)
                arguments = remove(arguments, i-1)
                arguments = insert(arguments, i-1, strconv.Itoa(value))
                i = 0
            }
        }
        i++
    }
    i = 0
    errorPlacement = 0
    // for i := range arguments[0:] {
    //     fmt.Printf("%s", arguments[i])
    // }  
    // fmt.Printf("\n")

    foundMinus := 0
    for i < len(arguments) {
        errorPlacement += len(arguments[i])
        if (arguments[i] == "+" || arguments[i] == "-") {
            prev, _ := strconv.Atoi(arguments[i-1])
            next, _ := strconv.Atoi(arguments[i+1])
            if (arguments[i] == "+") {
                value := prev + next
                arguments = remove(arguments, i+1)
                arguments = remove(arguments, i)
                arguments = remove(arguments, i-1)
                arguments = insert(arguments, i-1, strconv.Itoa(value))
                i = 0
            } else {
                value := prev - next
                if (value == 0) {
                    return 0, minusPlace
                } else if (value < 0) {
                    return -1, minusPlace
                }
                arguments = remove(arguments, i+1)
                arguments = remove(arguments, i)
                arguments = remove(arguments, i-1)
                arguments = insert(arguments, i-1, strconv.Itoa(value))
                foundMinus++
                i = 0
            }
        }
        i++
    }
    output, _ := strconv.Atoi(arguments[0])
    return output, -1
}

func calculate(arguments []string, pureArguments string) (string) {
    i := 0
    openIndex := 0
    if (arguments[0] == "+" || arguments[0] == "-" || arguments[0] == "*" || arguments[0] == "/" || arguments[0] == "%" || arguments[0] == "^") {
        return pureArguments + "\n^\nQuid dicis? True Romans would not understand your syntax!\n"
    }
    for i < len(arguments) {
        if (arguments[i] == ")") {
            //fmt.Printf("Found ) at %d\n", i)
            openIndex = i
            prev := arguments[i]
            for prev != "(" {
                openIndex--
                prev = arguments[openIndex]
            }
            //fmt.Printf("Found ( at %d\n", openIndex)
            tempArray := arguments[openIndex+1:i]
            //fmt.Printf("Printing tempArray: \n")
            //for r := range tempArray[0:] {
            //    fmt.Printf("tempArray[r]: %s\n", tempArray[r])
            //}
            parenthesisReturn, errorIndex := subCalculate(tempArray, pureArguments)
                pureArguments += "\n"
                c := 0
                for c < errorIndex {
                    pureArguments += " "
                    c++
                }
                pureArguments += "^\n"
                // var returnString string
                // for r := range pureArguments[0:] {
                //     returnString = returnString + pureArguments[r]
                // }
                // returnString += "\n"
                if (parenthesisReturn == 0) {
                    return pureArguments + "Quid dicis? Arab merchants haven't left for India yet!\n"
                    os.Exit(0)
                } else if (parenthesisReturn < 0) {
                    return pureArguments + "Quid dicis? Caesar demands positive thoughts!\n"
                    os.Exit(0)  
                } else if (parenthesisReturn > 3999) {
                    return "Disregarded. Output is over 3999.\n"
                    os.Exit(0)
                }
            if (errorIndex != -1) {

            }
            //fmt.Printf("parenthesisReturn: %d\n", parenthesisReturn)
            firsthalf := arguments[0:openIndex]
            secondhalf := arguments[i+1:]
            arguments = firsthalf 
            arguments = append(arguments, strconv.Itoa(parenthesisReturn))
            for s := 0; s < len(secondhalf); s++ {
                arguments = append(arguments, secondhalf[s])
            }
            //for u := range arguments[0:] {
            //    fmt.Printf("before insert arguments[%d]: %s\n", u, arguments[u])
            //} 
            //fmt.Printf("\n")
            //insert(arguments, openIndex, strconv.Itoa(parenthesisReturn))
            //for u := range arguments[0:] {
            //    fmt.Printf("arguments[%d]: %s\n", u, arguments[u])
            //} 
            //fmt.Printf("\n")
            if (len(arguments) > 1) {
                //fmt.Printf("len(arguments): %d\n", len(arguments))
                prev = arguments[openIndex]
            }
            i = 0
        }
        i++
    }
    //fmt.Printf("Finished\n\n")
    
    output, errorIndex := subCalculate(arguments, pureArguments)
    pureArguments += "\n"
    c := 0
    for c < errorIndex {
        pureArguments += " "
        c++
    }
    pureArguments += "^\n"
    // var returnString string
    // for r := range pureArguments[0:] {
    //     returnString = returnString + pureArguments[r]
    // }
    // returnString += "\n"
    if (output == 0) {
        return pureArguments + "Quid dicis? Arab merchants haven't left for India yet!\n"
        os.Exit(0)
    } else if (output < 0) {
        return pureArguments + "Quid dicis? Caesar demands positive thoughts!\n"
        os.Exit(0)  
    } else if (output > 3999) {
        return "Disregarded. Output is over 3999.\n"
        os.Exit(0)
    }
    r := NewRoman()
    stringOutput := r.ToRoman(output)
    return stringOutput + "\n"
}

func main() {
    // Check for input
    if (len(os.Args) != 2) {
        fmt.Printf("Incorrect input!\n")
        os.Exit(0)
    }

    arguments := setUpInputLexical(os.Args[1])

    // Checking for malformedness of the inputs
    malformedValue, message := checkMalformedness(arguments)
    if(!malformedValue) {
        fmt.Printf("%s", message)
        os.Exit(0)
    }

    /* Grammar for Programming Project 1:
     * <expr> → <term> {(+ | -) <term>}
     * <term> → <exponent> {(* | / | % ) <exponent>}
     * <exponent> → <factor> {^ <factor>}
     * <factor> → roman_numeral | ( <expr> )
     *
     * Tokens:
     * EOF = -1
     * ROMAN_NUMERAL = 1
     * PLUS = 2
     * MINUS = 3
     * TIMES = 4
     * DIVIDE = 5
     * MODULO = 6
     * POWER = 7
     * OPEN_PARENTHESIS = 8
     * CLOSED_PARENTHESIS = 9
     */

    synAnalysis, synMessage := syntacticalAnalysis(arguments, os.Args[1])
    fmt.Printf("%s", synMessage)
    if(!synAnalysis) {
        os.Exit(0)
    }
    
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }

    fmt.Printf("%s", calculate(arguments, os.Args[1]))
}
