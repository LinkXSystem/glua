package vm

import . "api"

// R(A) := UpValue[B]
// local u,v,w
// function f() local a,b,c,d,e = 1,2,u,v,w end
func getUpval(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(LuaUpvalueIndex(b), a)
}

// UpValue[B] := R(A)
// local u,v,w
// function f() local a,b,c; u,v,w = a,b,c end
func setUpval(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(a, LuaUpvalueIndex(b))
}

// R(A) := UpValue[B][RK(C)]
// local u
// function f() local a,b,k,d,e; d = u(k]; d = u[lOO] end
func getTabUp(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.GetRK(c)
	vm.GetTable(LuaUpvalueIndex(b))
	vm.Replace(a)
}

// UpValue[A][RK(B)] := RK(C)
// local u
// function f() local a,b,k,v,e; u[k] = v; u[lOO] =”fo。” end
func setTabUp(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(LuaUpvalueIndex(a))
}
