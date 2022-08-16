package flow

import (
  "github.com/konapun/flozart"
)

// This is a flow for hypothetical message board with authenticated users.
// The API flow is as follows:
// 1. Authenticate
//   a. Create account if account doesn't exist
//   b. Log in if account does exist
// 2. Start a thread
// 3. Make a comment on thread
// 4. Edit comment
// 5. Delete comment
// 6. Delete the thread

type Flow struct {

}

func NewFlow() *Flow {
  return &Flow{}
}
