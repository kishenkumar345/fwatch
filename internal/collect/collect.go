package collect

import (
    "fmt"
    "os"
    "log"
    "path/filepath"
)


func CollectFiles() []string {
    // Take command line inputted directory
    watchDirs := os.Args[1:]
    fmt.Println(watchDirs)

    var allFiles []string
    for _, dir := range watchDirs {
        err := filepath.Walk(dir, 
            func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() {
                allFiles = append(allFiles, path)
            }
            return nil
        })

        if err != nil {
            log.Fatal((err))
        }

    }

    return allFiles
}