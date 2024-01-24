# Backend Architecture

## Architectural styles

### REST

### GraphQL

### SOAP

### Hateos

Now that we've covered most basic backend development concepts, let us try and understand a little about how backend applications are architectured. Writing efficient and fast code is important, but it is also necessary to have a proper high level design that puts everything together such that components are cohesive and utilised effectively.

## 3-tier or monolithic applications
An application that consists of the following three parts: The UI, the backend and the database, is known as a 3-tier application.

3-tier applications are the simplest example of a modern-day, full-stack application. The user pushes a button on a mobile app, it triggers a request to the server, the server reads or writes from the database, depending on the request and returns a response back to the client. Such a form of application that only has a singular API or data access point, is called a *Monolith* or a *Monolithic* application. The word 'monolith' means any structure carved out of a singular stone. In the context of programming, it is an application that does all the things such as managing users, writing logs, reading and writing to a database, etc.

```{=latex}
\begin{center}
```
![3-tier application](src/book/images/7.1.png){width=60%}
```{=latex}
\end{center}
```

At a glance, this may seem like a backend solution that is good enough and will suffice a lot of business needs, but there are quite a few problems associated with it. Firstly, since these are singular applications, they are not fault-tolerant. Which means that if the server were to go down, every user of the product would have no service. Secondly, it is difficult to scale these kinds of applications on demand. You can scale vertically, but that can only help you for so long. And lastly, they're not very flexible. Let's look at a different architectural approach to solves this problem.

## Microservices
The microservice architecture aims to solve problems with the monolithic architecture by having multiple loosely coupled servers that run independently. Each server or node has its own purpose and communicate with each other via lightweight APIs.

For example, if you have an e-commerce application, you can have multiple services for each task: One for managing users, one for storing application-wide configurations, one for cropping and resizing images, interacting with a database and others. An API gateway sits between these services and will redirect requests to them accordingly.

```{=latex}
\begin{center}
```
![Microservice architecture](src/book/images/7.2.png){width=60%}
```{=latex}
\end{center}
```

One of the biggest advantages of the microservice architecture is flexibility. If you have a change to make, you can simply do it in one of the services, instead of having to change it in the entire application and then stopping and restarting everything. It also means that your application will be fault-tolerant, meaning if something goes wrong and one of the servers go down, you will still continue to serve users. Even though some of them will be affected by it, it is still better than not being able to serve users at all.

The microservice architecture, if implemented correctly, is a great way to solve the problems with the monolithic architecture. But doing so incorrectly will lead to unnecessary complexity and maintenance costs.
