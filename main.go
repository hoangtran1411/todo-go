package main

func main() {
	todos := Todos{}
	storage := NewStogare[Todos]("todos.json")
	storage.Load(&todos)

	todos.Add("Buy milk")
	todos.Add("Buy eggs")
	todos.Add("Buy bread")

	todos.Toggle(0)

	todos.Print()
	storage.Save(todos)
	//
}
