package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 12: Leonardo's Monorail ---

You finally reach the top floor of this building: a garden with a slanted glass ceiling. Looks like there are no more stars to be had.

While sitting on a nearby bench amidst some tiger lilies, you manage to decrypt some of the files you extracted from the servers downstairs.

According to these documents, Easter Bunny HQ isn't just this building - it's a collection of buildings in the nearby area. They're all connected by a local monorail, and there's another building not far from here! Unfortunately, being night, the monorail is currently not operating.

You remotely connect to the monorail control systems and discover that the boot sequence expects a password. The password-checking logic (your puzzle input) is easy to extract, but the code it uses is strange: it's assembunny code designed for the new computer you just assembled. You'll have to execute the code and get the password.

The assembunny code you've extracted operates on four registers (a, b, c, and d) that start at 0 and can hold any integer. However, it seems to make use of only a few instructions:

cpy x y copies x (either an integer or the value of a register) into register y.
inc x increases the value of register x by one.
dec x decreases the value of register x by one.
jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
The jnz instruction moves relative to itself: an offset of -1 would continue at the previous instruction, while an offset of 2 would skip over the next instruction.

For example:

cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
The above code would set register a to 41, increase its value by 2, decrease its value by 1, and then skip the last dec a (because a is not zero, so the jnz a 2 skips it), leaving register a at 42. When you move past the last instruction, the program halts.

After executing the assembunny code in your puzzle input, what value is left in register a?
*/

const INPUT = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 17 c
cpy 18 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

func main() {
	fmt.Println("Day 12 Part 1")
	instructions := strings.Split(INPUT, "\n")
	ExecuteInstructions(instructions)
}

type Registers map[byte]int

func CreateRegisters() Registers {
	return make(Registers)
}

func (r Registers) Execute(instruction string) int {
	fields := strings.Fields(instruction)
	index := 1
	switch fields[0] {
	case "cpy":
		r.Cpy(fields[1], fields[2])
	case "inc":
		r.Inc(fields[1])
	case "dec":
		r.Dec(fields[1])
	case "jnz":
		index = r.Jnz(fields[1], fields[2])
	}
	return index
}

func (r Registers) Cpy(arg1, arg2 string) {
	r[arg2[0]] = r.extractValue(arg1)
}

func (r Registers) Inc(arg string) {
	r[arg[0]]++
}

func (r Registers) Dec(arg string) {
	r[arg[0]]--
}

func (r Registers) Jnz(arg1, arg2 string) int {
	if r.extractValue(arg1) != 0 {
		jump, _ := strconv.Atoi(arg2)
		return jump
	}

	return 1
}

func ExecuteInstructions(instructions []string) {
	var i int
	registers := make(Registers)
	for i < len(instructions) {
		i += registers.Execute(instructions[i])
	}

	fmt.Println(registers)
}

func isRegister(s string) bool {
	return len(s) == 1 && !(s[0] >= '0' && s[0] <= '9')
}

func (r Registers) extractValue(arg string) int {
	var value int
	if isRegister(arg) {
		value = r[arg[0]]
	} else {
		value, _ = strconv.Atoi(arg)
	}

	return value
}
