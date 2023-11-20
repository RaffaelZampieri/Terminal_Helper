package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/rosedblabs/rosedb/v2"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	opt := rosedb.DefaultOptions
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		opt.DirPath = user.HomeDir + "\\rosedb\\"
	}
	if runtime.GOOS == "linux" {
		opt.DirPath = user.HomeDir + "/rosedb/"
	}

	db, err := rosedb.Open(opt)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	for i, arg := range os.Args {
		if i >= 1 {
			if strings.EqualFold(arg, "--del") {
				if i+1 < len(os.Args) {
					delete(db, os.Args[i+1])
				} else {
					color.Set(color.FgRed)
					fmt.Printf("You need to specify a key\n")
					color.Unset()
					return
				}
				break
			} else if strings.EqualFold(arg, "--new") {
				if i+2 < len(os.Args) {
					addValue(db, os.Args[i+1], os.Args[i+2])
				} else {
					color.Set(color.FgRed)
					fmt.Printf("You need to specify a key and a value\n")
					color.Unset()
					return
				}
				break
			} else if strings.EqualFold(arg, "--list") {
				list(db)
				break
			} else if strings.EqualFold(arg, "--help") {
				color.Set(color.FgCyan)
				fmt.Printf("To add a new value you need to run the command again, followed by the '--new' keyword and the key that will be used to call the value and the new value\n")
				fmt.Printf("To delete a value you need to run the command again, followed by the '--del' keyword and the key that will be used to call the value\n")
				fmt.Printf("To read a value you need to run the command again, followed by the key that will be used to call the value\n")
				fmt.Printf("To list all values you need to run the command again, followed by the '--list' keyword\n")
				color.Unset()
				break
			} else {
				read(db, arg)
				break
			}
		}
	}
}

func addValue(db *rosedb.DB, key string, value string) {
	err := db.Put([]byte(key), []byte(value))
	if err != nil {
		panic(err)
	}
	color.Set(color.FgGreen)
	fmt.Println("Saved!")
	color.Unset()
}

func read(db *rosedb.DB, key string) {
	val, err := db.Get([]byte(key))
	if err != nil {
		color.Set(color.FgRed)
		fmt.Printf("Key not found\n")
		color.Unset()
		return
	}
	color.Set(color.FgGreen)
	fmt.Println(string(val))
	color.Unset()
}

func delete(db *rosedb.DB, key string) {
	err := db.Delete([]byte(key))
	if err != nil {
		color.Set(color.FgRed)
		fmt.Printf("Key not found\n")
		color.Unset()
		return
	}
	color.Set(color.FgGreen)
	fmt.Printf("Value deleted\n")
	color.Unset()
}

func list(db *rosedb.DB) {
	i := 1
	db.Ascend(func(k, v []byte) (bool, error) {
		color.Set(color.FgGreen)
		fmt.Printf("%d - Key: %s, Value: %s\n", i, k, v)
		color.Unset()
		i++
		return true, nil
	})
}
