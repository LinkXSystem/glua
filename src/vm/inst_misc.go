package vm

import . "api"

// R(A) := R(B)
// luac -1 - : local a,b,c,d,e; d = b
func move(i Instruction, vm LuaVM)  {
	a, b, _ := i.ABC()
	a += 1;
	b += 1;
	vm.Copy(b, a)
}

// R(A) := R(B)
// luac -1 - : ::LOOP:: goto LOOP 
func jmp(i Instruction, vm LuaVM)  {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}