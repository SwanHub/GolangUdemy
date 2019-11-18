package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// defer is important for error handling. There are cleanup functions which you may want to run even after the error is hit.

func main() {
	// basicError() // basic error
	// usingScan()            // getting info from the terminal
	// createFileOnComputer() // create file on computer, handle potential errors
	// openFileFromComputer()   // open the file created by "createFileOnComputer"
	// printingOptions()        // different methods of printing errors
	// writingErrorLogsToFile() // writing err logs to a specific log file
	// usingFatal()             // using the fatal function, defer functions are ignored (everything after this point is not run)
	// usingPanicLog() // using log.Panicln (Println + panic())
	// usingPanic()    // using straight up panic (not ideal, though)
	// usingRecover()   // using recover... not sure yet how this works
	// errorsWithInfo()  // adding text to an error clause, saving that text in reusable variable
	customErrStruct() //
}

func basicError() {
	fmt.Println("------using basic errors------")
	i, err := fmt.Println("no error here")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func usingScan() {
	var a, b string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}

	fmt.Println("Age: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		panic(err)
	}
}

func createFileOnComputer() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("New FILE!")

	io.Copy(f, r)
}

func openFileFromComputer() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // "New FILE!"
}

func printingOptions() {
	fmt.Println("------alternate methods of printing errors------")
	_, err := os.Open("non-existent-file.txt")
	if err != nil {
		// fmt.Println("error happened: ", err)
		log.Println("error happened: ", err) // same as fmt. but includes a timestamp
	}
}

func writingErrorLogsToFile() {
	fmt.Println("------setting function output------")
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("oops: ", err)
	}
	defer f.Close()
	log.SetOutput(f) // takes in a writer, will write the log information to that file.

	f2, err := os.Open("non-existent.txt")
	if err != nil {
		log.Println("yikes, error: ", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

func usingFatal() {
	fmt.Println("------using fatal------")
	defer foo()
	_, err := os.Open("non-existent.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is run, deferred functions DO NOT run.")
}

// equivalent to calling Println(), then panic()
func usingPanicLog() {
	fmt.Println("------using panic------")
	_, err := os.Open("non-existent.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func usingPanic() {
	fmt.Println("------using straight panic------")
	_, err := os.Open("non-existent.txt")
	if err != nil {
		panic(err)
	}
}

func usingRecover() {
	fmt.Println("------using recover------")
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g.", i)
	g(i + 1)
}

var errBadMath error = errors.New("bad math: can't take square root of negative number")

func errorsWithInfo() {
	fmt.Printf("Error type: %T\n", errBadMath) // *errors.errorString
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err) // shuts the program down
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errBadMath
		return 0, fmt.Errorf("bad math: cannot take square root of negative number (%v)", f)
	}
	return 42, nil
}

func customErrStruct() {
	fmt.Println("------creating a complex error with struct------")
	_, err := sqrt2(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		bme := fmt.Errorf("bad math: cannot take square root of negative number (%v)", f)
		return 0, badMathError{"58.43.56", "43.43.43", bme}
	}
	return 42, nil
}

type badMathError struct {
	lat  string
	long string
	err  error
}

func (b badMathError) Error() string {
	return fmt.Sprintf("a bad math error occurred: %v %v %v", b.lat, b.long, b.err)
}
