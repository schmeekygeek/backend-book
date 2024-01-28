# Introduction to a Database

If you recall, in chapter I, we discussed briefly what a database is. By definition, **a database is an organized collection of structured information, or data, typically stored electronically in a computer system.**[^1] Simply put, it is a software where we store records such as user details.

```{=latex}
\begin{center}
```
![A basic database](src/book/images/6.1.png){width=30%}
```{=latex}
\end{center}
```

You might be wondering why we need such a special type of software to store data, when we could accomplish the same thing with tools like Microsoft Excel or Google Drive. Here are a few reasons why tools like these fail compared to a typical database.

1. Not performant.
2. Not secure.
3. Inability to backup data
4. Limited storage.

Every independant operation performed on a database, or unit of work done, is called **a transaction**. A DBMS (Database Management System) is a software that helps you perform operations such as read and write on a database. A DBMS typically has a set of properties that ensures that any transaction performed will complete with its intended purpose.

## ACID properties
ACID is an acronym that stands for **A**tomicity, **C**onsistency, **I**solation and **D**urability.

[^1]: https://www.oracle.com/database/what-is-database/
