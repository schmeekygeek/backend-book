# Basics of System Design

// system design introduction

Now that we've covered most basic backend development concepts, let us try and understand a little about how backend applications are architected. Writing efficient and fast code is important, but it is also necessary to architect software and have a proper high level design that puts everything together such that components act cohesively and are utilised effectively, and efficiently. Let's learn about this process of designing a good backend solution.

## What is System Design?
**System design is the study of the high level design of backend applications.** In other words, it is the process of architecting software and putting components and pieces in the right places to be used in the best way possible.

// TODO: figure add system design diagram

## Monolithic or 3-tier applications
An application that consists of the following three parts: The UI, the backend and the database, is known as a 3-tier application.

3-tier applications are the simplest example of a modern-day, full-stack application. The user pushes a button on a mobile app, it triggers a request to the server, the server reads or writes from the database, depending on the request and returns a response back to the client.
Such a form of application that only has a singular API or data access point and is centralized, is called a *Monolith* or a *Monolithic* application.
The word 'monolith' means any structure carved out of a singular stone. In the context of programming, it is an application that does all the things such as managing users, writing logs, reading and writing to a database, etc.

```{=latex}
\begin{center}
```
![3-tier application](src/book/images/7.1.png){width=60%}
```{=latex}
\end{center}
```

At a glance, this may seem like a backend solution that is good enough and will suffice a lot of business needs, but there are quite a few problems associated with it. Firstly, since these are singular applications, they are not fault-tolerant. Which means that if our application were to go down, every user of the product would have no service. Secondly, it is difficult to scale these kinds of applications on demand. And lastly, they're not very flexible. Let's look at a different architectural approach that solves this problem.

## Microservices
The microservice architecture aims to solve problems with the monolithic architecture by having multiple loosely coupled servers that run independently. Each server or node has its own purpose and multiple nodes communicate with each other via lightweight APIs.

For example, if you have an e-commerce application, you can have multiple services for each task: One for managing users, one for storing application-wide configurations, one for cropping and resizing images, interacting with a database and others. An API gateway sits between these services and will redirect requests to them accordingly.
```{=latex}
\begin{center}
```
![Microservice architecture](src/book/images/7.2.png){width=60%}
```{=latex}
\end{center}
```
One of the biggest advantages of the microservice architecture is flexibility. If you have a change to make, you can simply do it in one of the services, instead of having to change it in the entire application and then stopping and starting everything. It also means that your application will be fault-tolerant, meaning if something goes wrong and one of the servers go down, you will still continue to serve users. Even though some of them will be affected by it, it is still better than not being able to serve users at all.

The microservice architecture, if implemented correctly, is a great way to solve the problems with the monolithic architecture. But implementing the microservices architecture without an appropriate reason will lead to unnecessary complexity and maintenance costs.

We briefly talked about how monolithic applications are difficult to scale in Section 6.3. Let's understand what this means.

## Scaling
In the practical world, scaling is the process of increasing supply to meet growing demands.
In the context of programming and backend development, it is generally used to describe the process of increasing computational power such that demands like, a growing userbase, are met. The ability of a backend application to scale easily on demand is called its scalability. Scalability is also an important aspect to consider when designing backend solutions. Now, there are two types of scaling, **Vertical** and **Horizontal** scaling, with each of them having their own quirks, advantages and disadvantages. Let's go over them.

### Vertical scaling
Vertical scaling is the process of increasing or upgrading a single machine's or hosts' computing capability and strength to meet the increasing demands. If your machine, having 2 cores and a gigabyte of RAM, is struggling to serve, say, 10 users at the same time, you might want to consider adding another gigabyte to the machine to make it more capable than it previously was. Vertical scaling is simple, but it can also get pretty expensive and pricier as you keep climbing up the ladder. There is also a limit as to how much you can scale vertically, due to our current technological limitations. Scaling vertically can be different for different ways of upgrading that may be required. If you are running a database server, you might consider increasing the storage capacity of the disks to which your data is being written. If you want to run a highly responsive server, you should consider increasing its processing power.

***Pros:***

- Great for a single quick upgrade
- Cost effective, for smaller upgrades.

***Cons:***
- There are limitations to how much you can upgrade.
- Can get pricey after a certain threshold.
- May require you to upgrade other parts in order to upgrade the one you want to.

// TODO: figure: sam going to the gym

### Horizontal scaling
Horizontal scaling is the process of increasing the number of host machines or servers instead of increasing a singular machine's computing capacity. A load balancer is another machine used when scaling horizontally. It sits in front of the host machines usually running a software that distributes the workload evenly such that one machine or host is not overwhelmed completely. Horizontal scaling has many advantages since you can have as many computers as you like, doing work. It is also a much viable option in the long run because you can simply add another machine to the group of servers you may be running to ease the workload of the others.

// TODO: figure: many workers
