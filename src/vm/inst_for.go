package vm

import . "api"

// R(A)-=R(A+2); pc+=sBx
func forPrep(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	// R(A) -= R(A+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(LUA_OPSUB)
	vm.Replace(a)
	// pc+=sBx
	vm.AddPC(sBx)
}

// R(A)+=R(A+2);
// ifR(A) <?= R(A+1)then {
//	   pc+=sBx; R(A+3)=R(A)
// }
// luac -1 -1 -
// for i=l,2,100 do f() end
func forLoop(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	// R(A)+=R(A+2);
	vm.PushValue(a + 2)
	vm.PushValue(a)
	vm.Arith(LUA_OPADD)
	vm.Replace(a)

	isPositiveStep := vm.ToNumber(a+2) >= 0
	if isPositiveStep && vm.Compare(a, a+1, LUA_OPLE) ||
		!isPositiveStep && vm.Compare(a+1, a, LUA_OPLE) {

		// pc+=sBx; R(A+3)=R(A)
		vm.AddPC(sBx)
		vm.Copy(a, a+3)
	}
}
