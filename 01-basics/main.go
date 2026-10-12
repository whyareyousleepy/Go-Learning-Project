package main

import "fmt"

type User struct {
    ID   int
    Name string
}

func (u User) Greeting() string {
    return fmt.Sprintf("Hello, %s!", u.Name)
}

func findUser(id int) (User, error) {
    if id != 1 {
        return User{}, fmt.Errorf("user %d not found", id)
    }

    return User{
        ID:   1,
        Name: "Dmitry",
    }, nil
}

func main() {
    user, err := findUser(2)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(user.Greeting())
}