package initHash

import (
	// "fmt"
	"hash/fnv"
	"log"
	"os"
    "time"
)

type File struct {
    hash uint32
    name string
    idx int
}

func Hash(s string) uint32 {
    hash := fnv.New32a()
    hash.Write([]byte(s))
    return hash.Sum32()
}

// TODO: Don't allow hash collision, rehash if collision
func FileHash(allFiles []string) []File {
    var hashedFiles []File

    for i, j := range allFiles {
        read, err := os.ReadFile(j)
        if err != nil {
            log.Fatal(err)
        }

        var file File
        fileHash := Hash(string(read))
        file.idx = i
        file.hash = fileHash
        file.name = j
        hashedFiles = append(hashedFiles, file)
    }

    return hashedFiles
}

func HashCheck(hashedFiles []File) {
    if len(hashedFiles) != 0 {
        log.Println("Checking files...")
        for i := 0; i < len(hashedFiles); i++ {
            read, err := os.ReadFile(hashedFiles[i].name)
            if err != nil {
                log.Fatal(err)
            }
            fileHash := Hash(string(read))
            if fileHash != hashedFiles[i].hash {
                log.Println("File " + hashedFiles[i].name + " has changed!")
            } else {
                log.Println("File " + hashedFiles[i].name + " has not changed")
            }
        }    
    } else {
        log.Println("No files found...")
    }
    time.Sleep(5 * time.Second)
}