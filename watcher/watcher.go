package watcher

import (
  "log"
  "os"
  "os/exec"

  "github.com/fsnotify/fsnotify"
  "github.com/0xAX/notificator"
)

var notify = notificator.New(notificator.Options{
  // No icon
  DefaultIcon: "",
  AppName: "terraform-watcher",
})

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
            // Empty string in iconPath for no icon
            notify.Push("Validation Failed!", "Modified file: " + event.Name, "", notificator.UR_NORMAL)
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
