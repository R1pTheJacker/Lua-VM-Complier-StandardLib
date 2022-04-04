package vm

import . "luago/api"
/* ch09 */
// func getTabUp(i Instruction, vm LuaVM) {
// 	a, _, c := i.ABC()
// 	a += 1

// 	vm.PushGlobalTable()
// 	vm.GetRK(c)
// 	vm.GetTable(-2)
// 	vm.Replace(a)
// 	vm.Pop(1)
// }

func getUpval(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(LuaUpvalueIndex(b), a)
}

func setUpval(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(a, LuaUpvalueIndex(b))
}

func getTabUp(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.GetRK(c)//push key to top
	vm.GetTable(LuaUpvalueIndex(b))//pop key, push value
	vm.Replace(a)//pop value, move to register
}

func setTabUp(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(LuaUpvalueIndex(a))
}