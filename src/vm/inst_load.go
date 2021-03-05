package vm

import . "api"

// R(A), R(A+l), ... 1 R(A+B) := nil
// luac -1 - : local a,b,c,d,e
func loadNil(i Instruction, vm LuaVM)  {
	a, b, _ := i.ABC()
	a += 1

	vm.PushNil()
	for i := a; i <= a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

// R(A) := (bool)B; if (C) pc++
// luac -1 - : local a,b,c,d,e; c = true
func loadBool(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

func loadK(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	
	vm.GetConst(bx)
	vm.Replace(a)
}

func loadKx(i Instruction, vm LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()

	vm.GetConst(ax)
	vm.Replace(a)
}
