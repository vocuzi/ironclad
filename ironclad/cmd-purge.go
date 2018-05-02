package main


import "github.com/dmulholland/args"


import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
)


var purgeHelp = fmt.Sprintf(`
Usage: %s purge [FLAGS] [OPTIONS]

  Purge inactive entries from a database.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.

Flags:
  -h, --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


func purgeCallback(parser *args.ArgParser) {
    filename, password, db := loadDB(parser)

    list := db.Inactive()
    if len(list) == 0 {
        exit("no inactive entries to purge")
    }

    printCompact(list, len(list))
    answer := input("  Purge the entries listed above? (y/n): ")
    if strings.ToLower(answer) == "y" {
        db.PurgeInactive()
        line("─")
    } else {
        fmt.Println("  Purge aborted.")
        line("─")
    }

    saveDB(filename, password, db)
}
