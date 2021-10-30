package main

func print_name(name string) string {
   return "Hello " + name
}

func main() {
   name := "Haydn"
   test_function := print_name(name)
   println(test_function)
}
