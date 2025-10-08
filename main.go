package main

func main() {
	todos := Todos{}
	todos.Add("Buy milk")
	todos.Add("Buy eggs")
	todos.Add("Buy bread")

	todos.Toggle(0)

	todos.Print()
}
