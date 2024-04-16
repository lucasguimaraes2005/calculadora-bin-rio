package main

import (
	"fmt"
	"os"
	"strconv"
)

func decToBin(dec int) string {
	if dec == 0 {
		return "0"
	}
	bin := ""
	for dec > 0 {
		rem := dec % 2
		bin = fmt.Sprintf("%d%s", rem, bin)
		dec /= 2
	}
	return bin
}
func binToDec(bin string) int {
	dec := 0
	power := len(bin) - 1
	for _, digit := range bin {
		if digit != '0' && digit != '1' {
			fmt.Println("Número binário inválido")
			os.Exit(1)
		}
		dec += int(digit-'0') * (1 << uint(power))
		power--
	}
	return dec
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso correto: go run main.go [decimal|binario] [número]")
		os.Exit(1)
	}
	input := os.Args[1]
	number := os.Args[2]
	switch input {
	case "decimal":
		bin := decToBin(numberInt(number))
		fmt.Printf("Binario: %s\n", bin)
	case "binario":
		dec := binToDec(number)
		fmt.Printf("Decimal: %d\n", dec)
	default:
		fmt.Println("Formato inválido parceiro")
		os.Exit(1)
	}
}
func numberInt(number string) int {
	dec, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Número inválido")
		os.Exit(1)
	}
	return dec
}
