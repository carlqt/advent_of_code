package main

import "log"

type Instruction struct {
	Opcode         int
	Value1         int
	Value2         int
	PositionResult int
	inputs         []int
}

func (i *Instruction) Halt() bool {
	if i.Opcode == 99 {
		return true
	}

	return false
}

func NewInstruction(inputs []int, memory []int) Instruction {
	var instruction Instruction
	index1 := inputs[1]
	index2 := inputs[2]

	instruction.inputs = inputs
	instruction.Opcode = inputs[0]
	instruction.PositionResult = inputs[3]
	instruction.Value1 = memory[index1]
	instruction.Value2 = memory[index2]

	return instruction
}

func (i *Instruction) Execute(memory []int) {
	if !i.hasValidOpcode() {
		log.Fatalln("Unknown Opcode")
	}

	if i.Opcode == 99 {
		return
	}

	log.Printf("Executing %v - %d\n", i.inputs, i.result())

	memory[i.PositionResult] = i.result()
}

func (i *Instruction) result() int {
	if i.Opcode == 1 {
		return i.Value1 + i.Value2
	}

	return i.Value1 * i.Value2
}

func (i *Instruction) hasValidOpcode() bool {
	validOpcodes := [3]int{1, 2, 99}

	for _, opcode := range validOpcodes {
		if i.Opcode == opcode {
			return true
		}
	}

	return false
}
