package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string
}

func main() {
    // Initialize a new database connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create a new user
    db.Create(&User{Name: "John Doe", Email: "john@example.com"})

    // Read a user
    var user User
    db.First(&user, 1) // Find user with integer primary key
    db.First(&user, "name = ?", "John Doe") // Find user by name

    // Update - update user's email
    db.Model(&user).Update("Email", "john.doe@example.com")

    // Delete - delete user
    db.Delete(&user, 1)
}
