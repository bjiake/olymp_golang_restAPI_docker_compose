var account = db.getCollection("account")
var account1 = {
    "_id": 1,
    firstName: "John",
    lastName: "Doe",
    email: "john.doe@example.com",
    password: "secret123"
  }
var account2 = {
    "_id": 2,
    firstName: "Jane",
    lastName: "Doe",
    email: "jane.doe@example.com",
    password: "password456"
}
var account3 ={
    "_id": 3,
    firstName: "Bob",
    lastName: "Smith",
    email: "bob.smith@example.com",
    password: "qwerty789"
  }
var account4 = {
    "_id": 4,
    firstName: "Alice",
    lastName: "Johnson",
    email: "alice.johnson@example.com",
    password: "letmein123"
  }
var account5 = {
    "_id": 5,
    firstName: "Charlie",
    lastName: "Brown",
    email: "charlie.brown@example.com",
    password: "12345678"
  }
  
account.insertMany([account1,account2,account3,account4,account5])