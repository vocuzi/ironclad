package main


import "github.com/dmulholland/args"


import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
)


var restoreHelp = fmt.Sprintf(`
Usage: %s restore [FLAGS] [OPTIONS] ARGUMENTS

  Restore one or more deleted (i.e. inactive) entries. Entries to restore
  should be specified by ID.

Arguments:
  <entries>                 List of entry IDs.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.

Flags:
  -h, --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


func restoreCallback(parser *args.ArgParser) {

    // Check that at least one entry argument has been supplied.
    if !parser.HasArgs() {
        exit("you must specify at least one entry to restore")
    }

    // Load the database.
    filename, password, db := loadDB(parser)

    // Grab the entries to restore.
    list := db.Inactive().FilterByIDString(parser.GetArgs()...)
    if len(list) == 0 {
        exit("no matching entries")
    }

    // Print a listing and request confirmation.
    printCompact(list, len(db.Inactive()))
    answer := input("  Restore the entries listed above? (y/n): ")
    if strings.ToLower(answer) == "y" {
        for _, entry := range list {
            db.SetActive(entry.Id)
        }
        fmt.Println("  Entries restored.")
        line("─")
    } else {
        fmt.Println("  Restore aborted.")
        line("─")
    }

    // Save the updated database to disk.
    saveDB(filename, password, db)
}