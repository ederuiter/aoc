package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Operand1  string
	Operand2  string
	Executed  bool
	Type      string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	registers := map[string]bool{}
	instructions := map[string]*Instruction{}
	parsedInput := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsedInput = true
		} else if !parsedInput {
			parts := strings.SplitN(line, ": ", 2)
			registers[parts[0]] = parts[1] == "1"
		} else {
			parts := strings.SplitN(line, " ", 5)
			instructions[parts[4]] = &Instruction{
				Operation: parts[1],
				Operand1:  parts[0],
				Operand2:  parts[2],
				Executed:  false,
				Type:      "",
			}
		}
	}

	done := false
	for !done {
		done = true
		for output, instruction := range instructions {
			if instruction.Executed {
				continue
			}

			o1, ok1 := registers[instruction.Operand1]
			o2, ok2 := registers[instruction.Operand2]
			if ok1 && ok2 {
				fmt.Printf("%s = %s %s %s\n", output, instruction.Operand1, instruction.Operation, instruction.Operand2)
				switch instruction.Operation {
				case "AND":
					registers[output] = o1 && o2
				case "OR":
					registers[output] = o1 || o2
				case "XOR":
					registers[output] = o1 != o2
				}
				instruction.Executed = true
			} else {
				done = false
			}
		}
	}

	value := 0
	for key, val := range registers {
		if key[0] == 'z' && val {
			pos, _ := strconv.ParseInt(key[1:], 10, 64)
			value += 1 << pos
		}
	}
	fmt.Println(value)

	/*
		  half adder for first bit =>
			- x00 XOR y00 => z00
			  => ok
			- x00 AND y00 => c01
			  => ok

		  full adders for next bits
			- xnn XOR ynn => tnn
			  => all ok
			- tnn XOR cnn => znn
			  => 4 not ok
			  z15 => &{Operation:XOR Operand1:skh Operand2:rkt Executed:true Type:zxx} <= caused by skh being wrong
			  wpd => &{Operation:XOR Operand1:qqw Operand2:gkc Executed:true Type:z11} <= WRONG: wpd,z11
			  mdd => &{Operation:XOR Operand1:wfc Operand2:cmp Executed:true Type:z19} <= WRONG: mdd,z19
			  wts => &{Operation:XOR Operand1:smt Operand2:wpp Executed:true Type:z37} <= WRONG: wts,z37
			- tnn AND cnn => rnn
			  => 1 not ok
			  z11 => &{Operation:AND Operand1:gkc Operand2:qqw Executed:true Type:r11} <= caused by z11 being wrong
			  kjk => &{Operation:AND Operand1:rkt Operand2:skh Executed:true Type:rxx} <= caused by shk being wrong
			- xnn AND ynn => snn
			  => all ok
			- snn OR rnn => c<n+1>
			  => 2 not ok:
			  hvn => &{Operation:OR Operand1:pbb Operand2:mdd Executed:true Type:c00} <= caused by mdd being wrong
			  kbq => &{Operation:OR Operand1:jqf Operand2:kjk Executed:true Type:c00} <= WRONG: jqf, skh <= switched

			  carry out for last full adder goes to Z<num bits+1>

		 => jqf,mdd,skh,wpd,wts,z11,z19,z37
	*/

	for _, instruction := range instructions {
		switch instruction.Operation {
		case "OR":
			instruction.Type = "cxx"
		case "AND":
			if instruction.Operand1 == "x00" || instruction.Operand1 == "y00" {
				instruction.Type = "c01"
			} else if instruction.Operand1[0] == 'x' || instruction.Operand1[0] == 'y' {
				instruction.Type = "s" + string(instruction.Operand1[1:])
			} else {
				instruction.Type = "rxx"
			}
		case "XOR":
			if instruction.Operand1 == "x00" || instruction.Operand1 == "y00" {
				instruction.Type = "z00"
			} else if instruction.Operand1[0] == 'x' || instruction.Operand1[0] == 'y' {
				instruction.Type = "t" + string(instruction.Operand1[1:])
			} else {
				instruction.Type = "zxx"
			}
		}
	}

	for _, instruction := range instructions {
		if instruction.Type[1:] != "xx" {
			continue
		}
		switch instruction.Operation {
		case "OR":
			index := int64(0)
			if instructions[instruction.Operand1].Type[0] == 's' {
				index, _ = strconv.ParseInt(instructions[instruction.Operand1].Operand1[1:], 10, 64)
			} else if instructions[instruction.Operand2].Type[0] == 's' {
				index, _ = strconv.ParseInt(instructions[instruction.Operand2].Operand1[1:], 10, 64)
			} else {
				index = -1
				fmt.Printf("Wrong: %+v\n", instruction)
			}
			instruction.Type = fmt.Sprintf("c%02d", index+1)
		case "AND":
			if instructions[instruction.Operand1].Type[0] == 't' {
				instruction.Type = "r" + instructions[instruction.Operand1].Operand1[1:]
			} else if instructions[instruction.Operand2].Type[0] == 't' {
				instruction.Type = "r" + instructions[instruction.Operand2].Operand1[1:]
			} else {
				fmt.Printf("Wrong: %+v\n", instruction)
			}
		case "XOR":
			if instructions[instruction.Operand1].Type[0] == 't' {
				instruction.Type = "z" + instructions[instruction.Operand1].Operand1[1:]
			} else if instructions[instruction.Operand2].Type[0] == 't' {
				instruction.Type = "z" + instructions[instruction.Operand2].Operand1[1:]
			} else {
				fmt.Printf("Wrong: %+v\n", instruction)
			}
		}
	}

	for i := 0; i <= 45; i++ {
		match := fmt.Sprintf("%02d", i)
		for output, instruction := range instructions {
			if instruction.Type[1:] == match {
				fmt.Printf("%s => %+v\n", output, instruction)
			}
		}
	}
	for output, instruction := range instructions {
		if instruction.Type[1:] == "xx" {
			fmt.Printf("%s => %+v\n", output, instruction)
		}
	}
}
