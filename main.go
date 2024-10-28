package main

func main() {
	todos := Todos{}
	todos.add("buy milk")
	todos.add("buy bread")
	todos.add("buy paper")
	todos.add("do homework")
	todos.toggle(0)
	todos.toggle(2)
	todos.print()
}
