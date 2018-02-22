package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var data = []struct {
	name string
	cont string
	perm os.FileMode
}{
	{"test1.file", "Hello\nGolang is great", 0666},
	{"test2.file", "Hello\nGolang is great", 0666},
	{"test3.file", "Not matching\nGolang is great\nLast line", 0666},
}

func main() {

	files := []*os.File{}
	for _, fData := range data {
		f, err := os.Create(fData.name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, fData.cont)
		if err != nil {
			panic(err)
		}
		files = append(files, f)
	}

	// Compare by checksum
	checksums := []string{}
	for _, f := range files {
		f.Seek(0, 0) // reset to beginngin of file
		sum, err := getMD5SumString(f)
		if err != nil {
			panic(err)
		}
		checksums = append(checksums, sum)
	}

	fmt.Println("### Comparing by checksum ###")
	compareCheckSum(checksums[0], checksums[1])
	compareCheckSum(checksums[0], checksums[2])

	fmt.Println("### Comparing line by line ###")
	files[0].Seek(0, 0)
	files[2].Seek(0, 0)
	compareFileByLine(files[0], files[2])

	// Cleanup
	for _, val := range data {
		os.Remove(val.name)
	}

}

func getMD5SumString(f *os.File) (string, error) {
	file1Sum := md5.New()
	_, err := io.Copy(file1Sum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", file1Sum.Sum(nil)), nil
}

func compareCheckSum(sum1, sum2 string) {
	match := "match"
	if sum1 != sum2 {
		match = " does not match"
	}
	fmt.Printf("Sum: %s and Sum: %s %s\n", sum1, sum2, match)
}

func compareLines(line1, line2 string) {
	sign := "o"
	if line1 != line2 {
		sign = "x"
	}
	fmt.Printf("%s | %s | %s \n", sign, line1, line2)
}

func compareFileByLine(f1, f2 *os.File) {
	sc1 := bufio.NewScanner(f1)
	sc2 := bufio.NewScanner(f2)
	for {
		sc1Bool := sc1.Scan()
		sc2Bool := sc2.Scan()
		if !sc1Bool && !sc2Bool {
			break
		}
		compareLines(sc1.Text(), sc2.Text())
	}
}
