package action

type ActionParam struct {
  // TODO
}

type Action interface {
  Execute(params ...ActionParam) (any, error)
}
