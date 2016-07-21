package main


import (
    "fmt"
    "github.com/dmulholland/ironclad/irondb"
    "strings"
    "github.com/tonnerre/golang-text"
    "github.com/mitchellh/go-wordwrap"
)


// Print a list of entries in compact format.
func printCompact(list irondb.EntryList, dbsize int) {

    // Bail if we have no entries to display.
    if len(list) == 0 {
        line("-")
        fmt.Println("  No Entries")
        line("-")
        return
    }

    // Header.
    line("-")
    fmt.Println("  ID  |  TITLE")
    line("-")

    // Print the entry listing.
    for _, entry := range list {
        fmt.Printf("%4d  |  %s\n", entry.Id, entry.Title)
    }

    // Footer.
    line("-")
    fmt.Printf("  %d/%d Entries\n", len(list), dbsize)
    line("-")
}


// Print a list of entries in verbose format.
func printVerbose(
    list irondb.EntryList,
    dbsize int,
    key []byte,
    title string,
    cleartext bool) {

    // Bail if we have no entries to display.
    if len(list) == 0 {
        line("-")
        fmt.Println("  No Entries")
        line("-")
        return
    }

    // Header.
    line("-")
    fmt.Println("  " + title)
    line("-")

    // Print the entry listing.
    for _, entry := range list {
        fmt.Printf("  ID:       %d\n", entry.Id)
        fmt.Printf("  Title:    %s\n", entry.Title)

        if entry.Url != "" {
            fmt.Printf("  URL:      %s\n", entry.Url)
        }

        if entry.Username != "" {
            fmt.Printf("  Username: %s\n", entry.Username)
        }

        password, err := entry.GetPassword(key)
        if err != nil {
            exit(err)
        }

        if cleartext {
            fmt.Printf("  Password: %s\n", password)
        } else {
            fmt.Printf("  Password: %s\n", charstr(len([]rune(password)), '*'))
        }

        if entry.Email != "" {
            fmt.Printf("  Email:    %s\n", entry.Email)
        }

        if len(entry.Tags) > 0 {
            fmt.Printf("  Tags:     %s\n", strings.Join(entry.Tags, ", "))
        }

        if entry.Notes != "" {
            iline("~")
            wrapped := wordwrap.WrapString(entry.Notes, 76)
            indented := text.Indent(wrapped, "  ")
            fmt.Println(strings.Trim(indented, "\r\n"))
        }

        line("-")
    }

    // Footer.
    fmt.Printf("  %d/%d Entries\n", len(list), dbsize)
    line("-")
}