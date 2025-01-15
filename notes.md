## Go Basics
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

## Package vs Module
- A **package** is a collection of Go source files in the same directory that are compiled together
- Packages are used to organize code and provide modularity and reusability within a codebase
- Every Go file begins with a package declaration
- Example packages:
    ```
    package logcolor

    import "fmt"

    func PrintGreen(message string) {
        fmt.Println("\033[32m" + message + "\033[0m") // Print message in green
    }
    ```
    ```
    package main

    import "path/to/logcolor"

    func main() {
        logcolor.PrintGreen("Hello, Go!")
    }
    ```
- A **module** is defined by a go.mod file in its root directory
- The go.mod file specifies the module path, required dependencies and their versions
- A module can contain multiple packages
- Example of go.mod:
    ```
    module gitlab.pipelines.git

    go 1.20

    require (
        github.com/fatih/color v1.13.0
    )
    ```
- In this case, the module is identified as **gitlab.pipelines.git**

## Variables 
- Golang variables are mutable (can change) after they are declared
- Use `var` to declare them
    `var x int = 10`
- Can also use `:=` to declare them (cannot use this for constants or if you want to explicitly define a type)
    `x := 10`
- Use `=` for assigning values to variables
- Use `==` for comparing of 2 values

## Constants 
- Golang constants cannot be changed once they are declared
- Use `const` to declare them
    ```
    const pi = 3.14
    pi = 3.14159 // Not allowed
    ```

## Data Types
- In Golang, all values have data types
- Each language has different data types that it supports
- When declaring a variable, you need to tell the Go Compiler the data type
- ex. `var userName string` indicates a string type, just like `var userName = "Tom"` indicates a string type
- ex. `var userTickets int` indicates an integer type, just like `var userTickets = 34` indicated an integer type
- Optionally, you may explicitly define the data types in the variable
    - Ex. `var conferenceName string = "Go Conference"` or `const conferenceTickets int = 50`
- Numbers can be whole numbers, decimals, or negative numbers, so you may want to explicitly define the data type when using these
- Floating point types can contain a decimal component
- Int types cannot contain a decimal point, but they can be whole numbers or negative numbers


## User Input
- Can use fmt.Scan() to collect user input data and read formatted input
- Must create a pointer to save the variable in
- When defining variables, the variable value gets saved in memory on your computer
- A **pointer** (or special variable) is a variable that points to the memory address of another variable
- Use `&` before a variable to print the memory location of that variable
    - Ex. `fmt.Println(&remainingTickets)` would print the memory location (ex. 0xc00014098)

## Number Compilations
- When doing compilations with numbers, they must be of the same type
- Ex. Cannot do a - b if a is type `int` and b is type `uint`

## Arrays
- When tracking multiple users' data, you want to save that in a list
- Arrays in go have a fixed size
- Use square [] brackets to define an array
- The size of the array is defined in the square brackets
- Ex. `var bookings = [50]string{}`
- Define the data type that the array will contain
- Each array can contain only one data type

## Slices
- More commonly used than arrays
- What if you don't know the size of the array when you create it?
    - Create a list/array where you don't need to specify the size at the beginning
    - The array should automatically expand when new elements are added
- A **slice** is an abstraction of an array and is a more dynamic array
- Slices can have variable length or get a sub-array of its own
- Slices are index-based and have a size, but is resized when needed
- Ex. `var bookings []string`
- Ex. `var bookings = []string{}`
- Ex. `bookings := []string{}`

## For Loops
- Go only has the `for` loop
- Declare variables inside a function but outside the for loop that you don't want to be reset by the loop
- Iterate through a list
    - Ex. `for index, booking := range bookings
    - Range interates over elements for different data structures
- strings.Fields() splits a string with white space as a separator
    - Ex. Splits string "Nchole Smith" to slice ["Nicole", "Smith"]
- Underscores `_` are blank identifiers that are used when you want to ignore a variable you don't want to use
- Use a **break** statement to terminate the for loop. Once the for loop is broken, the application continues with the code right after the break (if there is any more code)
- Use **continue** statement to skip the remainder of the body of a for loop and immediately retest the for loop's condition
- Can use if-else condition statements within the for loop
    ```
    if condition {
        // Code to be executed if condition is true
    } else if condition {
        // Code to be executed if condition is true
    } else {
        // Code to be executed if both conditions are false
    }
    ```
- Can only have one "if" statement and one "else" statement, but can have as many "if else" statements in between those two as you want
- You can use true/false conditionals in for loops
- **Infinite Loop** - a loop that repeast endlessly as the condition is always **true**

## If-else for User Input Validation
- `len()` returns the length of a variable, according to its type
    - For arrays and slices, the size of the list or number of elements is returned
    - For strings, the number of characters is returned

## Validation
- Use validation variables to validate that user input is acceptable
- Can chain validations together with `&&`
- Can use a logical `OR` operator if any of the conditions (not all) must be true for the whole expression to be true
- `||` means `or`
- Ex. `var isValidCity = city == "Singapore" || city == "London"`
- Ex. `var isInvalidCity = city != "Singpaore" && city != "London"`
- Can use a `NOT` operator to negate/reverse a boolean value
    - Ex. `var isValidCity = city == "Singapore" || city == "London"`
    - Ex. !isValidCity

## Switch Statements
- Switch statement allows a variable to be tested for equality against a list of values
- Default case handles the case if no match is found
    ```
    switch city {
        case "New York":
            // Execute code for booking New York conference tickets
        case "Singapore":
            // Execute code for booking Singapore conference tickets
        case "London", "Berlin":
            // same logic for booking these two cities
        case "Mexico City", "Hong Kong":
            // same logic for booking these two cities
        default:
            fmt.Println("No valid city selected.")
    }
    ```

## Multiple Functions
- The first function is the main function, and there can be only one main function
- The main function is the entrypoint for the application
- All subsequent functions are only executed when called
- You can call a function as many times as you want, thus functions can be used to reduce code duplication
- Must call the function in main for it to run via `<function_name()>`
- Functions can be reused in your application by calling the function name
- When defining functions, you must include the parameter name and type
- When calling functions, you only include the parameters (not types)

## Parameters
- Information can be passed into functions as parameters
- Parameters are defined in the function signature (ex. `func greetUsers(name string)`)
- Arguments are the actual values passed when calling the function (ex. in `greetUsers("John")`, "John" is the argument)
- When creating a function that uses a previously defined variable, pass that variable into the function as a parameter and in the main function where the new function is called
    - Call the new function in the main function via `<function_name()>`
    - Ex. `greetUsers()`
    - The greetUsers function uses some variables defined in the main function, so you must pass those variables as parameters both in the greetUsers() call and the greetUsers function
    - Ex. `greetUsers(conferenceName, conferenceTickets, remainingTickets)`
    - Ex. `func greetUsers(conferenceName, conferenceTickets, remainingTickets)`

## Return
- Use **return** to return an output within a function that can be used in the main function or somewhere else
- A function can take an input and return an output
- Ex. `return firstNames` 
- The output of this will get printed in a prior function
- In Go you have to define the input and output parameters including its type explicitly
- You need to tell the function what **data type** the return tool is returning, so add `[]string` as a parameter of the `printFirstNames` function, so it knows it is returning a list of strings
- Ex. `func printFirstNames(bookings2 []string) []string {`
- Notice the input parameters are within () and the output parameter is outside of () if returning a single value
- When a function returns something to you, you can save that return as a variable
    - Ex. `firstNames := getFirstNames(bookings)`
    - Ex. `var firstNames = getFirstNames(bookings)`
- Unlike most programming languages, Go allows you to return multiple values (not just one)
- When returning multiple values, they need to be inside ()
    ```
    func printFirstnames(bookings2 []string) (bool, bool, string) {
        if len(bookings2) == 0 {
            return false, false, ""
        }
    ```

## Package Level Variables
- You can define variables that are shared among multiple functions
- Package level variables are defined at the top outside all functions
- Package level variables cannot be created using `:=` syntax
- Create package level variables with `var` syntax
- If you make some of your variables package level variables, you do not need to pass them as parameters to your functions
- Best practice is to define a variable as locally as possible, so if it is only being used within a single function, define it within that function, not at the package level

## Multiple Packages
- If your application is large, you may want to organize it with multiple packages
- For example, if a booking application is available for each city, maybe each city's code is its own package (main.go, london.go, singapore.go, etc) and you put shared logic in a common package
- Variables and functions defined outside any function can be accessed in all other files **within the same package**
- When running `go run main.go` and using helper or shared functions, you also have to include where those shared functions live
    - Ex. `go run main.go` will give you an error "undefined: <shared_function_name>
    - The correct command would be `go run main.go <shared_function_file>`
    - Ex. `go run main.go shared.go`
- If you don't want to specify the shared file (you may have many), you can instead specify a folder that contains those files
    - Ex. `go run .` will run all the files in the current folder
- When you have multiple packages, you should have a folder for each package, and all the files pertaining to that package reside in the package's folder
- If referencing a package within another package, you must import it into the package that wants to use it
- The `go.mod` file defines the module (Ex. `booking-app`), which is the import path for all the packages in your application
    - Ex. `import "booking-app/shared"`, NOT `import "shared"`
- After importing the other package, you call it in the main package via `<package_name>.<function_name>`
    - Ex. `shared.validateUserInput`
- Also have to explicitly export the function that is in the other package
- Exporting a function makes it available for use in other pacakges (It's not enough just to define the function in its own pacakge. You must also export it.)
- Export a function by CAPITALIZING the first letter in the function
    - Ex. `func ValidateUserInput` instead of `func validateUserInput`
- You can export variables, constants, functions, and types all by simply capitalizing the first letter

## Variable Scope
- Package, global, and local
- Local 
    - Declaration within a function can be used only within that function
    - Declaration within a block (Ex. for, if-else) can only be used within that block
- Package
    - Declaration outside all functions can be used everywhere in the same package
- Global
    - Declaration with capitalized variable name can be used everywhere in the module

## Maps
- Use a map data type to store key: value pairs
- Ex. You want to only print the firstName and lastName user input, but you want to save the firstName, lastName, email, and userTickets input somewhere. Use a map to do this
- Key: value pair would look like:
    ```
    firstName: "John"
    lastName: "Doe"
    email: "john@email.com"
    userTickets: 5
    ```
- When creating maps, all keys have the same data type and all values have the same data type
- Make an empty map with `var userData = make(map[string]string)`
- In the above example, `[string]` is the data type of the key and `string` is the data type of the value
- If you have a 2nd data type, like int or unit that you want to include in the map, you can convert it to a string using `strconv.FormatInt` or `strconv.FormatUint`
    - Ex. `strconv.FormatUint(uint64(userTickets), 10)`
    - This takes your uint value and formats it to a string with a decimal number
    - The 10 is for base 10, which represents decimal numbers
    - To save that value in the map, you would use `userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)`

## Collect Different Data Types of Data
- Example of different types of data:
    birthdate - Date - 01.02.1990
    newsletter - bool - true
    attendanceNames - make([]map[string]string, 0) - slice of maps where each map holds the name of the attendant and email, like [{name:"Mike", email "mike@gmail.com"}, {name: "Sara", email: "sara@yahoo.com"}]
- Remember maps only allow a single data type
- How can you save mixed data types?

Structs
- Allows you to save different data types as key: value pairs
- For each key in the struct, define the data type of the value
    ```
    type UserData struct {
        firstName string
        lastName string
        email string
        numberOfTickets uint
    }
    ```

- **Type** keyword creates a new type, with the name you specify (UserData)
- You can then use the struct
    - Ex. `var bookings = make([]UserData, 0)`
    - Ex. 
    ```
    var userData = UserData {
        firstName: firstName,
        lastName: lasName,
        email: email,
        numberOfTickets: userTickets,
    }
    ```

## Concurrency
- fmt.Sprintf allows you to print a value and save it into a string variable
- Concurrency allows you to make a program more efficient, by allowing multiple processes to run at the same time
- Single thread execution - subsequent code is blocked until previous code completes
- You can break out code from the main thread and execute it in a separate thread to make your code more efficient
- Ex. If 20 users book the ticket at the same time, 20 new separate threads will break off to complete those bookings; once the break off thread is completed, the thread will be deleted
- To make a function run with concurrent threads, simply put the `go` keyword in front of the function call; this starts a new goroutine
- Ex. `go sendTicket(userTickets, firstName, lastName, email)`
- You can take concurrency out of for lop by removing the loop and break, and then telling your function to wait until the go routine is completed
- Use WaitGroup to wait for the launched goroutine to finish
- WaitGroup is included in the sync package, which provides basic synchronization functionality
- Ex. `var wg = sync.WaitGroup{}`
- WaitGroup `Add` sets the number of goroutines to wait for (increases the counter by the provided number)
- Ex. `wg.Add(1)` if you have 1 go routine defined
- WaitGroup `Wait` needs to be executed at the end of the main function thread, and it waits for all the threads to be done before it exists the application
- Ex. `wg.Wait()` placed at end of function with the go routine in it
- WaitGroup `Done` decrements the WaitGroup counter by 1, and is called by the goroutine to indicate it is finished
- Ex. If the go routine is for the sendTicket function (i.e. `go sendTicket(userTickets, firstName, lastName, email)`) then within the `func sendTicket`, add `wg.Done()` at the end of the function
    ```
    func sendTicket(userTickets uint, firstName string, lastName string, email string) {
        fmt.Println("Processing...")
        time.Sleep(10 * time.Second) // Sleep for 10 seconds
        var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
        fmt.Println("#######") // Adds a visual divider
        fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
        fmt.Println("#######")
        wg.Done()
    }
    ```
## Error Printing
- `t.Errorf` method of t will print out a message and fail the test, but continue subsequent tests
- `f` stands for format, allowing you to build a string with values inserted into the placeholder values `%q`
- `t.Fatalf` method of t will print a message and fail the test, but does not continue subsequent tests
- `%v` = value in a default format; when printing structs, `%tv` adds field names
- `%t` = true or false
- `%q` = a double-quoted string safely escaped with Go syntax
- `%s` = uninterpreted bytes of a string or slice
- More at pkg.go.dev/fmt#hdr-printing










