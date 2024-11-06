- Create main.go and place code in it
- Must initialize the module with `go mod init <name_of_directory_with_main.go>`
    - `go mod init booking-app`
- All code must belong to a **package**
    - At top of main.go, define keyword "package" and the name of the package that your code will be a part of 
    - Standard name of the main package is "main"
- Must declare an entrypoint for your application
- A function can only have 1 "main" function, because you can only have 1 entrypoint
- Must **import** any functions/libraries you are going to use
<img src="booking-app/images/package.png" width="200" height="100">
- Run `go run main.go` to execute the file

**Variables** 
- Golang variables are mutable (can change) after they are declared
- Use `var` to declare them
    `var x int = 10`
- Can also use `:=` to declare them (cannot use this for constants or if you want to explicitly define a type)
    `x := 10`
- Use `=` for assigning values to variables
- Use `==` for comparing of 2 values

**Constants** 
- Golang constants cannot be changed once they are declared
- Use `const` to declare them
    ```
    const pi = 3.14
    pi = 3.14159 // Not allowed
    ```
**Date Types**
- In Golang, all values have data types
- Each language has different data types that it supports
- When declaring a variable, you need to tell the Go Compiler the data type
- ex. `var userName string` indicates a string type, just like `var userName = "Tom"` indicates a string type
- ex. `var userTickets int` indicates an integer type, just like `var userTickets = 34` indicated an integer type
- Optionally, you may explicitly define the data types in the variable
    - Ex. `var conferenceName string = "Go Conference"` or `const conferenceTickets int = 50`
- Integers can be whole numbers, decimals, or negative numbers, so you may want to explicitly define the data type when using these
- Floating point types can contain a decimal component

**User Input**
- Can use fmt.Scan() to collect user input data and read formatted input
- Must create a pointer to save the variable in
- When defining variables, the variable value gets saved in memory on your computer
- A **pointer** (or special variable) is a variable that points to the memory address of another variable
- Use `&` before a variable to print the memory location of that variable
    - Ex. `fmt.Println(&remainingTickets)` would print the memory location (ex. 0xc00014098)

**Number Compilations**
- When doing compilations with numbers, they must be of the same type
- Ex. Cannot do a - b if a is type `int` and b is type `uint`

**Arrays**
- When tracking multiple users' data, you want to save that in a list
- Arrays in go have a fixed size
- Use square [] brackets to define an array
- The size of the array is defined in the square brackets
- Ex. `var bookings = [50]string{}`
- Define the data type that the array will contain
- Each array can contain only one data type

**Slices**
- What if you don't know the size of the array when you create it?
    - Create a list/array where you don't need to specify the size at the beginning
    - The array should automatically expand when new elements are added
- A **slice** is an abstraction of an array and is a more dynamic array
- Slices can have variable length or get a sub-array of its own
- Slices are index-based and have a size, but is resized when needed
- Ex. `var bookings []string`
- Ex. `var bookings = []string{}`
- Ex. `bookings := []string{}`

**For Loops**
- Go only has the `for` loop
- Declare variables inside a function but outside the for loop that you don't want to be reset by the loop
- Iterate through a list
    - Ex. `for index, booking := range bookings
    - Range interates over elements for different data structures
- strings.Fields() splits a string with white space as a separator
    - Ex. Splits string "Nchole Smith" to slice ["Nicole", "Smith"]
- Underscores `_` are blank identifiers that are used when you want to ignore a variable you don't want to use
- Use a **break** statement to terminate the for loop. Once the for loop is broken, the application continues with the code right after the break (if there is any more code)




