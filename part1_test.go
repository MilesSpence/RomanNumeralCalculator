package main

import "testing"
import "strings"

/***** Testing for lexical errors *****/
func TestMalformednessExample1(t *testing.T) {
	input := "II plus III"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample2(t *testing.T) {
	input := "II plus III times IV minus I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample3(t *testing.T) {
	input := "II times IV minus I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample4(t *testing.T) {
	input := "II times {IV minus I]"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample5(t *testing.T) {
	input := "[V minus {VI minus (III minus {II minus I]}])"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample6(t *testing.T) {
	input := "III plus {IV times II] power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample7(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample8(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExample9(t *testing.T) {
	input := "III plu {IV times II] power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExample10(t *testing.T) {
	input := "I plus III minus VX times VI"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample11(t *testing.T) {
	input := "III plus {IV times II power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample12(t *testing.T) {
	input := "II times (I plus II minus III)"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample13(t *testing.T) {
	input := "II plus III divide IV"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample14(t *testing.T) {
	input := "II plus I times III minus VI"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExample15(t *testing.T) {
	input := "II power III power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleParenthesisError(t *testing.T) {
	input := "III plus {VX times II] power II"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleIIII(t *testing.T) {
	input := "II minus IIII"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleXXXXX(t *testing.T) {
	input := "II minus XXXXX"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleFirst(t *testing.T) {
	input := "VX power I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleLast(t *testing.T) {
	input := "X power VX"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleOutsideVocab(t *testing.T) {
	input := "CD powers XI"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleWithSymobols(t *testing.T) {
	input := "XXV + V"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleMultipleErrors(t *testing.T) {
	input := "XXV + V powers I minus 7"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleParenthesisSpace(t *testing.T) {
	input := "VX minus ( X minus X)"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExampleParenthesisSpaceWorks(t *testing.T) {
	input := "X minus ( X minus X)"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleCCC(t *testing.T) {
	input := "CCCC plus I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleWord(t *testing.T) {
	input := "II plus II minus (XC power X) divide X word I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleNotVocab(t *testing.T) {
	input := "II plus (II pmi II)"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExampleLotsOfBrackets(t *testing.T) {
	input := "II plus {III minus (IV power I] power M)"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExampleSyntactical1(t *testing.T) {
	input := "I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestMalformednessExampleSyntactical2(t *testing.T) {
	input := "I plus X minus"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue != true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleSyntactical3(t *testing.T) {
	input := "IXI plus"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleSyntactical4(t *testing.T) {
	input := "plus IIXVCDM"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleSyntactical5(t *testing.T) {
	input := "I X V C D V IXI plus"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleSyntactical6(t *testing.T) {
	input := "I X V yah C D V"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

// Lexical error case
func TestMalformednessExampleSyntactical7(t *testing.T) {
	input := "I word I plus I"
	arguments := setUpInputLexical(input)
	malformedValue, message := checkMalformedness(arguments)
	if malformedValue == true {
		t.Errorf("\n%s", message)
	}
}

/***** Testing for syntactical errors *****/
func TestSyntacticalExample1(t *testing.T) {
	input := "II plus III"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample2(t *testing.T) {
	input := "II plus III times IV minus I"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample3(t *testing.T) {
	input := "II times IV minus I"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample4(t *testing.T) {
	input := "II times {IV minus I]"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample5(t *testing.T) {
	input := "[V minus {VI minus (III minus {II minus I]}])"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample6(t *testing.T) {
	input := "III plus {IV times II] power II"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample7(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample8(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X power II"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample9(t *testing.T) {
	input := "III plu {IV times II] power II"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample10(t *testing.T) {
	input := "I plus III minus VX times VI"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalExample11(t *testing.T) {
	input := "III plus {IV times II power II"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample12(t *testing.T) {
	input := "II times (I plus II minus III)"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample13(t *testing.T) {
	input := "II plus III divide IV"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample14(t *testing.T) {
	input := "II plus I times III minus VI"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalExample15(t *testing.T) {
	input := "II power III power II"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalParenthesis1(t *testing.T) {
	input := "} minus I plus ("
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalParenthesis2(t *testing.T) {
	input := "{ minus I plus }"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalParenthesis3(t *testing.T) {
	input := "X minus I plus }"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalParenthesis4(t *testing.T) {
	input := "X minus { I plus "
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalEven1(t *testing.T) {
	input := "X minus"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalParenthesis5(t *testing.T) {
	input := "X minus X }"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalEven2(t *testing.T) {
	input := "X minus X minus X minus"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

func TestSyntacticalX(t *testing.T) {
	input := "X"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue != true {
		t.Errorf("\n%s", message)
	}
}

// Syntactical error case
func TestSyntacticalPower(t *testing.T) {
	input := "power"
	arguments := setUpInputLexical(input)
	syntacticalValue, message := syntacticalAnalysis(arguments, input)
	if syntacticalValue == true {
		t.Errorf("\n%s", message)
	}
}

/***** Testing for correct output *****/
func TestOutputExample1(t *testing.T) {
	input := "II plus III"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "V") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample2(t *testing.T) {
	input := "II plus III times IV minus I"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "XIII") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample3(t *testing.T) {
	input := "II times IV minus I"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "VII") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample4(t *testing.T) {
	input := "II times {IV minus I]"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "VI") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample5(t *testing.T) {
	input := "[V minus {VI minus (III minus {II minus I]}])"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "I") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample6(t *testing.T) {
	input := "III plus {IV times II] power II"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "LXVII") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample7(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "X") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputExample8(t *testing.T) {
	input := "{MCMXCVIII divide III divide VI minus XI) divide X power II"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "I") {
		t.Errorf("\n%s", output)
	}
}

// Runtime Error 0
func TestOutputExample12(t *testing.T) {
	input := "II times (I plus II minus III)"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if output != "II times (I plus II minus III)\n                    ^\nQuid dicis? Arab merchants haven't left for India yet!" {
		t.Errorf("\nExpected:\nII times (I plus II minus III)\n                    ^\nQuid dicis? Arab merchants haven't left for India yet!\n\nReceived:\n%s", output)
	}
}

// Runtime Error 0
func TestOutputExample13(t *testing.T) {
	input := "II plus III divide IV"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if output != "II plus III divide IV\n            ^\nQuid dicis? Arab merchants haven't left for India yet!" {
		t.Errorf("\nExpected:\nII plus III divide IV\n            ^\nQuid dicis? Arab merchants haven't left for India yet!\n\nReceived:\n%s", output)
	}
}

// Runtime Error -
func TestOutputExample14(t *testing.T) {
	input := "II plus I times III minus VI"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if output != "II plus I times III minus VI\n                    ^\nQuid dicis? Caesar demands positive thoughts!" {
		t.Errorf("\nExpected:\nII plus I times III minus VI\n                    ^\nQuid dicis? Caesar demands positive thoughts!\n\nReceived:\n%s", output)
	}
}

func TestOutputExample15(t *testing.T) {
	input := "II power III power II"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "DXII") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputMorePower(t *testing.T) {
	input := "II power II power III"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "CCLVI") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputEvenMorePower(t *testing.T) {
	input := "II power I power II power III power I"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "II") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputLastPower(t *testing.T) {
	input := "II power II power III power I"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "CCLVI") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputModulo1(t *testing.T) {
	input := "V modulo II"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "I") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputModulo2(t *testing.T) {
	input := "VIII modulo III"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "II") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputModulo25(t *testing.T) {
	input := "I plus (X minus II) modulo III"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "III") {
		t.Errorf("\n%s", output)
	}
}

func TestOutputModulo3(t *testing.T) {
	input := "III plus IV modulo III times II plus (III power II]"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "XIV") {
		t.Errorf("\n%s", output)
	}
}

// Runtime Error 0
func TestOutputModulo4(t *testing.T) {
	input := "IV modulo II"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if output != "IV modulo II\n   ^\nQuid dicis? Arab merchants haven't left for India yet!" {
		t.Errorf("\nExpected:\nIV modulo II\n   ^\nQuid dicis? Arab merchants haven't left for India yet!\n\nReceived:\n%s", output)
	}
}

func TestOutputModulo5(t *testing.T) {
	input := "(XXXVII modulo XX) plus (III times II)"
	arguments := setUpInputLexical(input)
    for i := range arguments[0:] {
        arguments[i] = substituteToMath(arguments[i])
    }
    output := strings.TrimSpace(calculate(arguments, input))
	if (output != "XXIII") {
		t.Errorf("\n%s", output)
	}
}
