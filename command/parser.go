package command

import "github.com/nproc/kb/user"

type Parser func(user.User, string) (interface{}, error)
