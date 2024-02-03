\part{Fundamentals}
# Introduction to the Backend
Consider this analogy: You go to a restaurant, the waiter greets you, then sits you at a table. He then brings you a menu for you to decide what you want to eat. You spend some time dilly-dallying, overthinking before finally deciding to go with something basic as noodles. The waiter gives you a look of disappointment, writes down your order, then disappears.
```{=latex}
\begin{center}
```
![Sam visits a restaurant](src/book/images/2.1.png "2.1"){width=40%}
```{=latex}
\end{center}
```
He returns after some time to give you the ever-saddening news that the restaurant is out of noodles and asks if you would like some rice instead. You agree, and he serves you rice a few minutes later.

*What just happened?* To put it bluntly, the customer is the user. The waiter is the website that the user is going to interact with, and the kitchen is the server that the waiter will make requests to. When the server receives the request of noodles, it checks the refrigerator to see if it has the stock for said noodles or not. Here the refrigerator is the database, where all the ingredients of the kitchen are stored. This is obviously an oversimplification, but you get the idea.

Backend development is the server-side development of a web application. It mostly involves writing code that is used by other applications and services; Mainly business logic.
```{=latex}
\begin{center}
```
![Business Logic](src/book/images/2.2.png "2.2"){width=60%}
```{=latex}
\end{center}
```

## What is an API?
The term API stands for Application Programming Interface. It is essentially what we've been calling a backend application. You can write an API using any popular backend framework or programming language meant to run on the server. We will learn how to build one in the 3rd part of this book using Go.

## Why do we need an API?
When learning about backend development, it is important to understand why we need to have a backend in our application in the first place in order to truly understand its essence and the problems it solves. You might be thinking why the client-side application can't communicate with the database directly. Well, it can, but it **shouldn't**.

Communicating with the database directly is simply not optimal, as we need to have a central point of data-access that would take a request, know if the request is coming from a person that has the authority to make that request, read some data, run some code, perform some theatrics on that data, read some more data, write some logs to disk, perform other actions to then finally wrap up everything neatly in a response-body to be sent back to the client. 

In the next chapter we will understand how this communication between two machines over the internet takes place.
