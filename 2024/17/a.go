package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	OP_ADV = 0
	OP_BXL = 1
	OP_BST = 2
	OP_JNZ = 3
	OP_BXC = 4
	OP_OUT = 5
	OP_BDV = 6
	OP_CDV = 7
)

type CPU struct {
	Registers map[byte]int
	Memory    []byte
	PC        int
	Out       []byte
}

func (cpu *CPU) GetOperandValue(opcode byte, operand byte) int {
	if opcode == OP_BXL || opcode == OP_JNZ || operand <= 3 {
		// literal value
		return int(operand)
	} else if operand == 4 {
		return cpu.Registers['A']
	} else if operand == 5 {
		return cpu.Registers['B']
	} else if operand == 6 {
		return cpu.Registers['C']
	} else {
		panic("reserved")
	}
}

func (cpu *CPU) Dump() {
	for pc := 0; pc < len(cpu.Memory); pc += 2 {
		opcode := cpu.Memory[pc]
		operand := cpu.Memory[pc+1]
		val := ""
		if opcode == OP_BXL || opcode == OP_JNZ || operand <= 3 {
			// literal value
			val = fmt.Sprintf("%d", operand)
		} else if operand == 4 {
			val = "A"
		} else if operand == 5 {
			val = "B"
		} else if operand == 6 {
			val = "C"
		} else {
			panic("reserved")
		}
		switch opcode {
		case OP_ADV:
			numerator := "A"
			denominator := "(2^" + val + ")"
			fmt.Printf("A := %s/%s\n", numerator, denominator)
		case OP_BDV:
			numerator := "A"
			denominator := "(2^" + val + ")"
			fmt.Printf("B := %s/%s\n", numerator, denominator)
		case OP_CDV:
			numerator := "A"
			denominator := "(2^" + val + ")"
			fmt.Printf("C := %s/%s\n", numerator, denominator)
		case OP_BXL:
			fmt.Printf("B := B XOR %s\n", val)
		case OP_BST:
			fmt.Printf("B := %s %% 8\n", val)
		case OP_JNZ:
			fmt.Printf("JNZ A => %s\n", val)
		case OP_BXC:
			fmt.Printf("B := B XOR C\n")
		case OP_OUT:
			fmt.Printf("PRINT %s %% 8\n", val)
		}
	}
}

func (cpu *CPU) Loop() bool {
	pc := cpu.PC
	for cpu.Step() && cpu.PC != pc {
	}
	if cpu.PC+1 >= int(len(cpu.Memory)) {
		return false
	}
	return true
}

func (cpu *CPU) Step() bool {
	if cpu.PC+1 >= int(len(cpu.Memory)) {
		return false
	}

	incr := int(2)
	opcode := cpu.Memory[cpu.PC]
	operand := cpu.Memory[cpu.PC+1]

	operandValue := cpu.GetOperandValue(opcode, operand)
	switch opcode {
	case OP_ADV:
		numerator := int(cpu.Registers['A'])
		denominator := 1 << operandValue
		//fmt.Printf("ADV %d: %d/%d\n", operandValue, numerator, denominator)
		cpu.Registers['A'] = int(numerator / denominator)
	case OP_BDV:
		numerator := int(cpu.Registers['A'])
		denominator := 1 << operandValue
		//fmt.Printf("BDV %d: %d/%d\n", operandValue, numerator, denominator)
		cpu.Registers['B'] = int(numerator / denominator)
	case OP_CDV:
		numerator := int(cpu.Registers['A'])
		denominator := 1 << operandValue
		//fmt.Printf("CDV %d: %d/%d\n", operandValue, numerator, denominator)
		cpu.Registers['C'] = int(numerator / denominator)
	case OP_BXL:
		//fmt.Printf("BXL %d: %d ^ %d\n", operandValue, cpu.Registers['B'], operandValue)
		cpu.Registers['B'] = cpu.Registers['B'] ^ operandValue
	case OP_BST:
		//fmt.Printf("BST %d: %d %% 8\n", operandValue, cpu.Registers['B'])
		cpu.Registers['B'] = operandValue % 8
	case OP_JNZ:
		//fmt.Printf("JNZ %d: %d != 0\n", operandValue, cpu.Registers['A'])
		if cpu.Registers['A'] != 0 {
			incr = 0
			cpu.PC = operandValue
		}
	case OP_BXC:
		//fmt.Printf("BXL %d: %d ^ %d\n", operandValue, cpu.Registers['B'], cpu.Registers['C'])
		cpu.Registers['B'] = cpu.Registers['B'] ^ cpu.Registers['C']
	case OP_OUT:
		//fmt.Printf("OUT %d: %d %% 8\n", operandValue, operandValue)
		cpu.Out = append(cpu.Out, byte(operandValue%8))
	}
	//fmt.Println(cpu.Registers)
	cpu.PC += incr
	return true
}

func (cpu *CPU) Reset() {
	for k, _ := range cpu.Registers {
		cpu.Registers[k] = 0
	}
	cpu.PC = 0
}

func (cpu *CPU) ConsumeOut() byte {
	res := cpu.Out[0]
	cpu.Out = cpu.Out[1:]
	return res
}

func generateProgram(a int, out []byte) int {
	var m byte
	m, out = out[len(out)-1], out[:len(out)-1]
	for b := 0; b < 8; b++ {
		newA := (a << 3) + b
		newB := b ^ 6
		c := newA >> newB
		newB = ((newB ^ c) ^ 7) & 7
		if newB == int(m) {
			if len(out) == 0 {
				return newA
			}
			newA = generateProgram(newA, out)
			if newA >= 0 {
				return newA
			}
		}
	}
	return -1
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cpu := CPU{
		Registers: map[byte]int{'A': 0, 'B': 0, 'C': 0},
		Memory:    []byte{},
		PC:        0,
		Out:       []byte{},
	}

	re := regexp.MustCompile(`Register ([A-Z]): (\d+)`)
	parsedRegisters := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsedRegisters = true
		} else if !parsedRegisters {
			match := re.FindStringSubmatch(line)
			reg := match[1][0]
			value, _ := strconv.ParseInt(match[2], 10, 64)
			cpu.Registers[reg] = int(value)
		} else {
			program := strings.Split(line[9:], ",")
			for _, m := range program {
				cpu.Memory = append(cpu.Memory, m[0]-'0')
			}
		}
	}
	fmt.Printf("%+v\n", cpu)
	for cpu.Step() {
	}
	fmt.Printf("Output: ")
	for len(cpu.Out) > 0 {
		fmt.Printf("%d,", cpu.ConsumeOut())
	}
	fmt.Println("")

	cpu.Dump()

	a := generateProgram(0, cpu.Memory)
	fmt.Println(a)

	cpu.Reset()
	cpu.Registers['A'] = a
	for cpu.Loop() {
		fmt.Printf("%+v\n", cpu)
	}
	fmt.Printf("Output: ")
	for len(cpu.Out) > 0 {
		fmt.Printf("%d,", cpu.ConsumeOut())
	}
	fmt.Println("")
}
