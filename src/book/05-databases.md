# Introduction to a Database

If you recall, in chapter I, we discussed briefly what a database is. By definition, **a database is an organized collection of structured information, or data, typically stored electronically in a computer system.**[^1] Simply put, it is a place where we store records such as user details.

```{=latex}
\begin{center}
```
![A basic database](src/book/images/5.1.png){width=30%}
```{=latex}
\end{center}
```

## What is a DBMS?
A Database Managment System is a software that allows for interaction with databases.

You might be wondering why we need such a special type of software to store data, when we could accomplish the same thing with tools like Microsoft Excel or Google Drive. Here are a few very basic reasons why tools like these fail compared to a typical database.

1. Not performant.
2. Not secure.
3. Inability to backup data.
4. Limited storage.
5. Not optimal for concurrent operations.

## SQL
SQL stands for Structured Query Language. It is an advanced query language used to query database and perform actions such as read and retrieve data in certain ways, write, back up data, etc. 

## Relational vs Non-Relational Databases

## ACID properties

Every independant operation performed on a database, or unit of work done, is called **a transaction**. A DBMS (Database Management System) is a software that helps you perform operations such as read and write on a database. A DBMS typically has a set of properties that ensures that any transaction performed will complete with its intended purpose.

ACID is an acronym that stands for **A**tomicity, **C**onsistency, **I**solation and **D**urability. These are the characteristics of a transaction performed on a database. Databases whose transactions possesses the above characteristics, are said to be *ACID* compliant.

Let's look at these characteristics one by one.

1. **Atomicity:** Guarantees that transactions are either completed or are discarded. If a transaction or operation begins but fails to conclude successfully, the transaction is rolled back and all changes are discarded.
2. **Consistency:** It refers to the ability of the database to maintain data integrity constraints. So if the constraint of a field is that it should only contain text, writing an integer value to it will fail.
3. **Isolation:** Makes sure that all transactions that run concurrently on a unit of data are isolated from each other. Meaning that if we were to withdraw 100$ and 50$ from an account containing 300\$, they both will run in isolation and the end result would be 150$ instead of 250$ or 200$.
4. **Durability:** It is the ability of the database to persist data safely, even in circumstances such as power outages that are usually unexpected and uncalled for.

Let us now understand what a database schema is.

[^1]: https://www.oracle.com/database/what-is-database/
