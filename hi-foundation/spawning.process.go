package main

import "fmt"
import "io/ioutil"
import "os/exec"
import "syscall"
import "os"

func main() {
	// Sometimes our Go programs need to spawn other, non-Go processes. For example, the syntax highlighting on this site is implemented by spawning a pygmentize process from a Go program.
	// We do this when we need an external process accessible to a running Go process.
	dateCmd := exec.Command("date")  // We’ll start with a simple command that takes no arguments or input and just prints something to stdout. The exec.Command helper creates an object to represent this external process.
	dateOut, err := dateCmd.Output() // .Output is another helper that handles the common case of running a command, waiting for it to finish, and collecting its output. If there were no errors, dateOut will hold bytes with the date info.
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello") // Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.
	grepIn, _ := grepCmd.StdinPipe()         // Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello") // We ommited error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h") // Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	// Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic exec function.
	// Note that Go does not offer a classic Unix fork function. Usually this isn’t an issue though, since starting goroutines, spawning processes, and exec’ing processes covers most use cases for fork.
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}   // Exec requires arguments in slice form (as apposed to one big string). We’ll give ls a few common arguments. Note that the first argument should be the program name.
	env := os.Environ()                        // Exec also needs a set of environment variables to use. Here we just provide our current environment.
	execErr := syscall.Exec(binary, args, env) // Here’s the actual syscall.Exec call. If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process. If there is an error we’ll get a return value.
	if execErr != nil {
		panic(execErr)
	}
}

// Command
// $ go run spawning-processes.go
// > date
// Wed Oct 10 09:53:11 PDT 2012
// > grep hello
// hello grep
// > ls -a -l -h
