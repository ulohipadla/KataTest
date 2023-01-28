package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func add(x int, y int) string {
	return strconv.Itoa((x + y))
}

func mult(x, y int) string {
	return strconv.Itoa((x * y))
}

func div(x, y int) string {
	return strconv.Itoa((x / y))
}

func der(x, y int) string {
	return strconv.Itoa((x - y))
}

func checkSystems(x, y string) int {
	romean := "IVX"
	if strings.ContainsAny(x, romean) && strings.ContainsAny(y, romean) {
		return 1
	} else if strings.ContainsAny(x, romean) == false && strings.ContainsAny(y, romean) == false {
		return 2
	} else {
		fmt.Println("Error. Different types of numbers.")
		return 3
	}
}

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func operationBlock(x string) string {
	var operationalMassive = strings.Split(strings.ToUpper(x), " ")
	if len(operationalMassive) < 3 {
		return "Error. Not a full statement"
	}
	if len(operationalMassive) > 3 {
		return "Error. To much statements"
	}
	if operationalMassive[1] == "+" {
		switch checkSystems(operationalMassive[0], operationalMassive[2]) {
		case 1:
			tempString1 := romanToInt(operationalMassive[0])
			tempString2 := romanToInt(operationalMassive[2])
			result, err := strconv.Atoi(add(tempString1, tempString2))
			if err != nil {
				log.Fatal(err)
			}
			return intToRoman(result)
		case 2:
			tempString1 := operationalMassive[0]
			tempString2 := operationalMassive[2]
			tempInt, errAtoi1 := strconv.Atoi(tempString1)
			tempInt2, errAtoi2 := strconv.Atoi(tempString2)
			if (errAtoi1 != nil) && (errAtoi2 != nil) {
				log.Fatal(errAtoi1)
				log.Fatal(errAtoi2)
				return ""
			}
			if (tempInt < 1 || tempInt > 10) || (tempInt2 < 1 || tempInt2 > 10) {
				return "Error. Arguments are not in range [1;10]"
			}
			return add(tempInt, tempInt2)
		case 3:
			return ""
		}

	}

	if operationalMassive[1] == "-" {
		switch checkSystems(operationalMassive[0], operationalMassive[2]) {
		case 1:
			tempString1 := romanToInt(operationalMassive[0])
			tempString2 := romanToInt(operationalMassive[2])
			result, err := strconv.Atoi(der(tempString1, tempString2))
			if result < 1 {
				fmt.Println("Error. Romans can't be negative or equal to zero")
				return ""
			}
			if err != nil {
				log.Fatal(err)
			}
			return intToRoman(result)
		case 2:
			tempString1 := operationalMassive[0]
			tempString2 := operationalMassive[2]
			tempInt, errAtoi1 := strconv.Atoi(tempString1)
			tempInt2, errAtoi2 := strconv.Atoi(tempString2)
			if (errAtoi1 != nil) && (errAtoi2 != nil) {
				log.Fatal(errAtoi1)
				log.Fatal(errAtoi2)
				return ""
			}
			if (tempInt < 1 || tempInt > 10) || (tempInt2 < 1 || tempInt2 > 10) {
				return "Error. Arguments are not in range [1;10]"
			}
			return der(tempInt, tempInt2)
		case 3:
			return ""
		}
	}

	if operationalMassive[1] == "*" {
		switch checkSystems(operationalMassive[0], operationalMassive[2]) {
		case 1:
			tempString1 := romanToInt(operationalMassive[0])
			tempString2 := romanToInt(operationalMassive[2])
			result, err := strconv.Atoi(mult(tempString1, tempString2))
			if err != nil {
				log.Fatal(err)
			}
			return intToRoman(result)
		case 2:
			tempString1 := operationalMassive[0]
			tempString2 := operationalMassive[2]
			tempInt, errAtoi1 := strconv.Atoi(tempString1)
			tempInt2, errAtoi2 := strconv.Atoi(tempString2)
			if (errAtoi1 != nil) && (errAtoi2 != nil) {
				log.Fatal(errAtoi1)
				log.Fatal(errAtoi2)
				return ""
			}
			if (tempInt < 1 || tempInt > 10) || (tempInt2 < 1 || tempInt2 > 10) {
				return "Error. Arguments are not in range [1;10]"
			}
			return mult(tempInt, tempInt2)
		case 3:
			return ""
		}
	}

	if operationalMassive[1] == "/" {
		switch checkSystems(operationalMassive[0], operationalMassive[2]) {
		case 1:
			tempString1 := romanToInt(operationalMassive[0])
			tempString2 := romanToInt(operationalMassive[2])
			result, errAtoi := strconv.Atoi(div(tempString1, tempString2))
			if errAtoi != nil {
				log.Fatal(errAtoi)
			}
			if result < 1 {
				fmt.Println("Error. Romans can't be negative or equal to zero")
				return ""
			}
			return intToRoman(result)
		case 2:
			tempString1 := operationalMassive[0]
			tempString2 := operationalMassive[2]
			tempInt, errAtoi1 := strconv.Atoi(tempString1)
			tempInt2, errAtoi2 := strconv.Atoi(tempString2)
			if (errAtoi1 != nil) && (errAtoi2 != nil) {
				log.Fatal(errAtoi1)
				log.Fatal(errAtoi2)
				return ""
			}
			if (tempInt < 1 || tempInt > 10) || (tempInt2 < 1 || tempInt2 > 10) {
				return "Error. Arguments are not in range [1;10]"
			}
			return div(tempInt, tempInt2)
		case 3:
			return ""
		}
	}

	return "Error. No operator"
}

func main() {
	for true {
		fmt.Println("Type in an expression:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		result := operationBlock(scanner.Text())
		if result != "" && strings.Contains(result, "Error.") == false {
			fmt.Println("Answer:", result)
		} else {
			fmt.Println(result)
			fmt.Println("Executing programm.")
			break
		}
	}
}
