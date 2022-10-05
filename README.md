# goodle - A Moodle Web Service Client for Go
*Work in progress*
## Usage
```go
import "github.com/nguyendhst/goodle"

func main() {
    // Create a new Moodle client
    client := goodle.NewClient("http://moodle.example.com/", "token")
    // Get Site information
    info, _ := client.GetSiteInfo()
    // Get Unread conversations
    conv, _ := client.GetUnreadConversationsCount()
}
```