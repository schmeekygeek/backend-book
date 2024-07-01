\part{Building WhisperChat}
# Project Introduction and Setup
WhisperChat, the project that we'll build in this book together, is going to be the server for a simple texting platform that will allow users to communicate or chat with someone in their preferred vicinity or range.

We will achieve this by, firstly, taking the location of the client in latitude and longitude and a range (in kms). We will put this client in a waiting room to be matched with other clients shortly. Once a different client connects and provides their details such as location and range, we will put them in the waiting room as well. As soon as we have more than one clients in the waiting room, we will run a matchmaking algorithm that will go over every clients location for each client, calculate the distance between them and check if the distance between them is less than the range that both of them have selected. If it is less, then we match them, or else we go to the next client.

## What you need:
Since we'll be building this project in Go, you need to have it installed. You don't necessarily need to know Go very well. If you have programming experience in other languages, it should be pretty easy to interpret as Go is a really simple and intuitive language.
