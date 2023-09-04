package main

// "fmt"

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rosedblabs/rosedb/v2"
)

func main() {
	opt := rosedb.DefaultOptions
	opt.DirPath = "/tmp/rosedb/"

	key := "null"
	value := "null"

	db, err := rosedb.Open(opt)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	if len(os.Args) >= 3 {
		if os.Args[1] == "Del" || os.Args[1] == "del" || os.Args[1] == "DEL" {
			key = os.Args[2]
			delete(db, key)
		}
		if os.Args[1] == "New" || os.Args[1] == "new" || os.Args[1] == "NEW" {
			key = os.Args[2]
			value = os.Args[3]
			addValue(db, key, value)
		}
	} else if len(os.Args) >= 2 {
		key = os.Args[1]
		read(db, key)
	} else {
		color.Set(color.FgCyan)
		fmt.Printf("To add a new value you need to run the command again, followed by the 'new' keyword and the key that will be used to call the value and the new value\n")
		fmt.Printf("To delete a value you need to run the command again, followed by the 'del' keyword and the key that will be used to call the value\n")
		fmt.Printf("To read a value you need to run the command again, followed by the key that will be used to call the value\n")
		color.Unset()
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
