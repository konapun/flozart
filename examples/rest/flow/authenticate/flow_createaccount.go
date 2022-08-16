package authenticate

import (
  "github.com/konapun/flozart/examples/rest/api"

  "github.com/konapun/flozart"
  "github.com/konapun/flozart/action"
)


var CreateAccount = flozart.RegisterFlow("create-account", createAccount) // or flow.Register(...) with init?

func createAccount(env *flozart.Environment) *flozart.Flow {
  start := &flozart.Flow{
    Action: func(_ ...action.Parameter) error {
      client := api.NewApiClient() // TODO: store in env
      response, err := client.CreateAccount()
    },
  }

  return start
}
