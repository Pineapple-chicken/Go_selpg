package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type selpg_args struct {
	start_page      int
	end_page        int
	in_filename     string
	dest            string
	page_len        int
	page_type       bool
}

var progname string

func main() {
	progname = os.Args[0]
	var args selpg_args
	Init(&args)
	process_args(&args)
	process_input(&args)
}



func Init(args *selpg_args) {
	flag.Usage = usage
	flag.IntVar(&args.start_page, "s", -1, "start_page")
	flag.IntVar(&args.end_page, "e", -1, "end_page")
	flag.StringVar(&args.dest, "d", "", "Destination.")
	flag.IntVar(&args.page_len, "l", 72, "page_length")
	flag.BoolVar(&args.page_type, "f", false, "page_type")
	flag.Parse()
}

func process_args(args *selpg_args) {
	if args.start_page == -1 || args.end_page == -1 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		if args.end_page != -1 {
			fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n", progname)
		} else {
			fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", progname)
		}
		flag.Usage()
		os.Exit(1)
	}
	if args.start_page < 0 || args.end_page < 0  || args.start_page > args.end_page {
		fmt.Fprintln(os.Stderr, "Invalid arguments")
		flag.Usage()
		os.Exit(2)
	}
}

func process_input(args *selpg_args) {
	var cmd *exec.Cmd
	var cmdIn io.WriteCloser
	var cmdOut io.ReadCloser
	if args.dest != "" {
		cmd = exec.Command("bash", "-c", args.dest)
		cmdIn, _ = cmd.StdinPipe()
		cmdOut, _ = cmd.StdoutPipe()
		cmd.Start()
	}
	if flag.NArg() > 0 {
		args.in_filename = flag.Arg(0)
		fout, perror := os.Open(args.in_filename)
		if perror != nil {
			fmt.Println(perror)
			os.Exit(3)
		}
		fin := bufio.NewReader(fout)
		count := 0
		for {
			line, _, perror := fin.ReadLine()
			if perror != io.EOF && perror != nil {
				fmt.Println(perror)
				os.Exit(4)
			}
			if perror == io.EOF {
				break
			}
			if count/args.page_len >= args.start_page && count/args.page_len <= args.end_page {
				if args.dest == "" {
					fmt.Println(string(line))
				} else {
					fmt.Fprintln(cmdIn, string(line))
				}
			}
			count++
		}
		if args.dest != "" {
			cmdIn.Close()
			cmdBytes, _ := ioutil.ReadAll(cmdOut)
			cmd.Wait()
			fmt.Print(string(cmdBytes))
		}
	} else {
		s := bufio.NewScanner(os.Stdin)
		count := 0
		st := ""
		for s.Scan() {
			line := s.Text()
			line += "\n"
			if count/args.page_len >= args.start_page && count/args.page_len <= args.end_page {
				st += line
			}
			count++
		}
		if args.dest == "" {
			fmt.Print(st)
		} else {
			fmt.Fprint(cmdIn, st)
			cmdIn.Close()
			cmdBytes, _ := ioutil.ReadAll(cmdOut)
			cmd.Wait()
			fmt.Print(string(cmdBytes))
		}
	}
}

func usage() {
	fmt.Println("\nUsage of selpg.\n")
	fmt.Println("\tselpg -s=Number -e=Number [options] [filename]")
	fmt.Println("\t-l:Determine the number of lines per page and default is 72.")
	fmt.Println("\t-f:Determine the type and the way to be seprated.")
	fmt.Println("\t-d:Determine the destination of output.")
	fmt.Println("\t[filename]: Read input from this file.")
	fmt.Println("\tIf filename is not given, read input from stdin. and Ctrl+D to cut out.\n")
}
