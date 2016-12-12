package main

import (
  "log"
  "os"
  "os/exec"

  "github.com/fsnotify/fsnotify"
)

func TerraformWatcher() {
  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
  }
  defer watcher.Close()

  done := make(chan bool)

  go func() {
    for {
      select {
      case event := <-watcher.Events:
        if event.Op&fsnotify.Write == fsnotify.Write {
          log.Println("###############################")
          log.Println("modified file:", event.Name)
          cmd := exec.Command("terraform", "validate")
          cmd.Stdout = os.Stdout
          cmd.Stderr = os.Stderr
          log.Println("Validation...")
          if err := cmd.Run(); err != nil {
            log.Printf("Validation failed!\n\n")
          } else {
            log.Printf("No issues found.\n\n")
          }
        }
      case err := <-watcher.Errors:
        log.Println("error:", err)
      }
    }
  }()

  err = watcher.Add(".")
  if err != nil {
    log.Fatal(err)
  }
  <-done
}

func main() {
  TerraformWatcher()
}
