package main

import (
    "fmt"
    "github.com/casbin/casbin/v2"
)

func main() {
    // Load model and policy from files
    e, err := casbin.NewEnforcer("model.conff", "policy.csv")
    if err != nil {
        panic(err)
    }

    // Check permissions
    perms := []struct {
        sub string
        obj string
        act string
    }{
        {"alice", "data1", "read"},
        {"alice", "data1", "write"},
        {"bob", "data2", "read"},
        {"bob", "data2", "write"},
    }

    for _, perm := range perms {
        hasAccess, err := e.Enforce(perm.sub, perm.obj, perm.act)
        if err != nil {
            fmt.Println("Error enforcing policy:", err)
            continue
        }

        if hasAccess {
            fmt.Printf("Access granted: %s can %s %s\n", perm.sub, perm.act, perm.obj)
        } else {
            fmt.Printf("Access denied: %s cannot %s %s\n", perm.sub, perm.act, perm.obj)
        }
    }
}
