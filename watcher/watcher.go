package watcher

import (
  "log"
  "os"
  "os/exec"

  "github.com/fsnotify/fsnotify"
  "github.com/0xAX/notificator"
  "github.com/fatih/color"
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
          log.Println(color.YellowString("Modified file: " + event.Name))
          cmd := exec.Command("terraform", "validate")
          cmd.Stdout = os.Stdout
          cmd.Stderr = os.Stderr
          log.Println(color.YellowString("Validating..."))
          if err := cmd.Run(); err != nil {
            log.Printf(color.RedString("Validation failed! ⚠️\n\n"))
            // Empty string in iconPath for no icon
            notify.Push("Validation Failed!", "Modified file: " + event.Name,
                        "", notificator.UR_NORMAL)
          } else {
            log.Printf(color.GreenString("No issues found 👍\n\n"))
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
