
package main

 import (
    "github.com/Com1Software/C1S"
     "bufio"
	"fmt"
	"os"
	"strings"
 )
 func main() {
xfile:="main.go"
	cline := -1
		dline := 0
		linemax := 40
		if _, err := os.Stat(xfile); err == nil {
			lines, err := readLines(xfile)
			if err != nil {
				fmt.Printf("Load Error : ( %s )\n", err)
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("-------------------------------- Edit\n")
			fmt.Printf("-- Loaded File ( %s )\n", xfile)
			fmt.Printf("-- Lines : ( %d )\n", len(lines))
			for {
				cline++
				dline++
				if dline > len(lines) {
					cline = 0
					dline = 1
				}
				fmt.Printf("-- %s", lines[cline])
				for i := len(lines[cline]); i < linemax; i++ {
					fmt.Printf(" ")
				}

				fmt.Printf(" --<Line %d>. ", dline)
				text, _ := reader.ReadString('\n')
				//---------------------------------------- convert CRLF to LF
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				//----------------------------------------------------------

				//   fmt.Printf("text=%s\n",text)
				//   fmt.Println(len(text))

				//-------------------------------------- Edit Help
				if text == "?" {
					fmt.Println("----------------------------- Edit Help")
					fmt.Println("-- ? - Displays available Edit Commands")
					fmt.Println("-- 0 - Enter a line number to Edit a specifc line")
					fmt.Println("-- a - Add a line after the current line")
					fmt.Println("-- e - Edit current line")
					fmt.Println("-- l - Lists loaded command script")
					fmt.Println("-- r - Runs loaded command script")
					fmt.Println("-- q - Quits Edit")
					cline--
					dline--
				}

				//-------------------------------------- List
				if text == "l" {
					fmt.Printf("\n")
					fmt.Printf("-------------------------------- Listing for : %s\n", xfile)
					for i := 0; i < len(lines); i++ {
						fmt.Printf("-- Line %d>   ", i+1)
						fmt.Printf("%s\n", lines[i])
					}
					fmt.Printf("-------------------------------- End of Listing\n\n")
					cline--
					dline--
				}

				//-------------------------------------- Quit
				if text == "q" {
					fmt.Println("Quit X12Tool Edit")
					return
				}
				//-------------------------------------- Test
				if text == "hi" {
					fmt.Println("Hello, Yourself")
					cline--
					dline--
				}
			}
        }
  fmt.Println("Test")
  c1s.Cmd("hey there")
 }

 //------------------------------------------------------------- Read file to lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//------------------------------------------------------------- Write line to files
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

