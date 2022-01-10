# RomanNumeralCalculator

### The following are the instructions for the project:

### Overview

You are a Royal Information Officer in the Court of a Roman Emperor. To facilitate commerce, the Emperor ordered your team to produce a calculator. Your team is responsible for implementing the calculator’s backend, which would take as input a string containing an arithmetic expression and then would either report an error or output the result. Given the realities of your time and geographic location, the calculator will have to abide by the following conditions:
* All numbers will be Roman numerals (e.g., I, II, XXX, etc.). You should follow the rules specified here: Roman Numerals
* All operators will be Latin words rather than their corresponding modern math symbols. Specifically, your calculator is to support: plus, minus, times, divide, modulo, power, and parenthesis (see item D below). 
  * The operators are to follow their modern precedence and associativity, such as in JavaScript expressions (i.e., multiplication, division, and modulo have a higher precedence than addition and subtraction, with parenthesis used to override the built-in precedence). 
  * All operators are left-associative, with the exception of exponentiation.
* The input expressions are case sensitive: all numerals are upper case, while all operators are lower case.
* Rather than only parentheses (i.e., ‘(‘ and ‘)’), the input expressions can include brackets (i.e., ‘[‘ and ‘]’)  and curly braces (i.e., ‘{‘ and ‘}’). Parthehesis, brackets, and curly braces have the traditional semantic of parentheses, but all of them can be mixed arbitrarily as long as they are matched. (The thing is that the Roman keyboards are not that precise when it comes to entering these operators, so the calculator must be able to accept all their possible mixes.)
* The mathematical operations have modern semantics, with the exception that there are no zero and no negative numbers. The concepts of zero and negative numbers were not known during these times. Any operations involving zero or negative numbers are undefined and should raise an error. 
  * The calculator must detect errors and report them via the caret operator (‘^’) to point to the location of the reported error in the expression.
    * When encountering a lexical error, the calculator is to terminate with the error message: Quid dicis? You offend Caesar with your sloppy lexical habits!
    * When encountering a syntactic error, the calculator is to terminate with the error message: Quid dicis? True Romans would not understand your syntax!
  * According to historical reconstructions, the concept of zero was brought to Europe from India by Arab merchants long after the collapse of the Roman Empire. So if any computation produces zero (even as an intermediate result), the calculator is to terminate with the error message: Quid dicis? Arab merchants haven't left for India yet!
  * During the Roman times, nobody had yet thought of negative numbers as a concept. So if any computation produces a negative number (even as an intermediate result), the calculator is to terminate with the error message: Quid dicis? Caesar demands positive thoughts!
### Implementation Requirements
* A Royal Decree commands that you implement your solution in the Go language.
* You are to implement the calculator as a command line program that takes an expression surrounded by double quotes as an argument to the main function.

When grading your assignment, we will test your solution against the following examples and also several other examples, which will not be revealed. Assume that your Go executable is named “part1”. So your solution would be expected to work as follows:

      part1 "II plus III"
      V

      part1 "II plus III times IV minus I"
      XIII

      part1 "II times IV minus I"
      VII

      part1 "II times {IV minus I]"
      VI
      part1 "[V minus {VI minus (III minus {II minus I]}])"
      I

      part1 "III plus {IV times II] power II"
      LXVII

      part1 "{MCMXCVIII divide III divide VI minus XI) divide X"
      X

      part1 "{MCMXCVIII divide III divide VI minus XI) divide X power II"
      I
      part1 "III plu {IV times II] power II"
      III plu {IV times II] power II
            ^
      Quid dicis? You offend Caesar with your sloppy lexical habits!

      part1 "I plus III minus VX times VI"
      I plus III minus VX times VI
                                ^
      Quid dicis? You offend Caesar with your sloppy lexical habits!

      part1 "III plus {IV times II power II"
      III plus {IV times II power II
                                                      ^
      Quid dicis? True Romans would not understand your syntax!

      part1 "II times (I plus II minus III)"
      II times (I plus II minus III)
                                    ^
      Quid dicis? Arab merchants haven't left for India yet!

      part1 "II plus III divide IV"
      II plus III divide IV
                         ^
      Quid dicis? Arab merchants haven't left for India yet!

      part1 "II plus I times III minus VI"
      II plus I times III minus VI
                                     ^
      Quid dicis? Caesar demands positive thoughts!
      part1 "II power III power II"
      DXII

Notes: 

There are many possible ways to implement this calculator. You can choose a top-down or a bottom-up parser, implemented in a variety of strategies, for your calculator. You may be tempted to quickly put your solution together, but to succeed, you’d have to apply some of the theoretical concepts we have studied in class.

Extra Test Cases:

These are the additional test cases I created:

Tests additional incorrect roman numerals: 

    II minus IIII 

             ^ 

    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests additional incorrect roman numerals:

    II minus XXXXX

             ^

    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests having a lexical error inside parentheses:

    III plus {VX times II] power II

             ^

    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests having a lexical error first in the input:

    VX power I

    ^

    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests having a lexical error last in the input:

    X power VX

            ^
               
    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests if a symbol would work in the input:

    XXV + V

        ^

    Quid dicis? You offend Caesar with your sloppy lexical habits!

Tests for incorrect facing parentheses:

    } minus I plus (

    ^

    Quid dicis? True Romans would not understand your syntax!

Tests for a lack of roman numerals to calculate:

    { minus I plus }

    ^

    Quid dicis? True Romans would not understand your syntax!

Tests for only one closed parenthesis:

    X minus I plus }

                     ^

    Quid dicis? True Romans would not understand your syntax!

Tests for only one open parenthesis:

    X minus { I plus

              ^
              
    Quid dicis? True Romans would not understand your syntax!

Tests for missing roman numeral:

    X minus

            ^

    Quid dicis? True Romans would not understand your syntax!

Tests what occurs for a single term:

    X

    X

An additional test to check if power performs correctly:

    II power I power II power III power I

    II

Test if modulo works properly:

    V modulo II

    I

Tests for the case where zero is output, but from modulo, not divide:

    IV modulo II

    ^
    Quid dicis? Arab merchants haven't left for India yet!
