1.Login: http://localhost:8080/api/v1/users/login
{
"username": "john_doe",
"password": "password123"
}
If not logged in, an "Unauthorized" notification will appear. If the username is not found or the password is incorrect, the output will be:
"error": "User not found"
"error": "Invalid password"


2.Payment: http://localhost:8080/api/v1/payment/
{
"account_number": "1234567890",
"amount": "5000"
}
If the customer is not found or the "account_number" does not match the data, the error will be:
"error": "Customer not found"


3.Transfer: http://localhost:8080/api/v1/transfer/transfers
{
"sender_account": {"account_number":"1234567890"},
"receiver_account": {"account_number":"0987654321"},
"amount": 500000
}
If either the sender or receiver account is not registered, the error will be:
"error": "Sender account not registered"
"error": "Receiver account not registered"


4. History: http://localhost:8080/api/v1/activities/all
This will print all the history:
[
{
"id": 1,
"description": "Payment from 1234567890 , Amount: 5000"
},
{
"id": 2,
"description": "TF from 1234567890 to 0987654321, Amount: 500000"
},
{
"id": 3,
"description": "TF from 1234567890 to 0987654321, Amount: 500000"
}
]




