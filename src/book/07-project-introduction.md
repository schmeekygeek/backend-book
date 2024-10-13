\part{Project BlogSpace}

# Project Introduction and Setup
*BlogSpace* is a simple blogging platform that we will be building together in this book. Building a project after you've learnt some concepts is a great way to cement the knowledge and grok the true meaning and purpose of a concept inside your mind. We are going to build *BlogSpace* step-by-step, covering each detail, feature implementation and code, line-by-line. Most of the endpoints that we'll build, will be simple CRUD operations so that you pick them up easily. Note that we won't be building the frontend of the entire application as that is beyond the scope of this book. However, we will be testing and trying out our application's API endpoints using tools such as *Postman* (An API testing application) and `curl`. Let's go over the features that we will be implementing in *BlogSpace.*

> You may choose to add your own features on the project or skip the ones that you feel are redundant. I will also be leaving a few exercises for you to complete after the project is concluded.

## Features
- **User login and signup** - We'll implement the total flow of creating
  a user account using details such as email, verifying the email of the
  account and adding login. Before the user is logged in and their email is verified, we will reject all requests from them.
- **Creating blog posts** - We will allow users to write as many blogs as they would like.
- **Like blog posts** - Blog posts can be liked by another user other than the author himself.
- **Search for blogs** - We will do keyword matching and search blog posts for the user.
- **Delete blog posts** - Blog posts can be deleted by the author of the post.

## Features summary
- *User account creation*
- *User email verification*
- *User login*
- *User profile edit*
- *Blog creation*
- *Blog editing*
- *Like blogs*
- *Search for blogs*

## Setting the scene
For this project we'll be using Go (or Golang), specifically version 1.21. We won't be using any library or framework in our project except for a tiny one called `sqlx` which is a superset of the standard `database/sql` library in Go. Meaning it has everything that `database/sql` has and a few extra extensions such as prepared statements. That's it! We won't be needing any massive framework or library, or install huge dependencies to get things done, we just need an editor where we can write your code and the command-line where we'll execute the project. The reason for doing this is to help you understand how applications work at a lower level instead of abstracting everything away to a framework. Frameworks are great, but when your purpose is to learn, explore and understand, you should do so with as little help as possible.

### Why Go?
You might be wondering why I chose Go for this project over something like Python, or maybe JavaScript.

The simple answer is, **I just like Go.** 

It is an extremely simple programming language that is very readable even if you're not familiar with it. One of the biggest perks of Go is that it has pretty much everything you'd need to build applications in its standard library.
You mostly don't need to install huge and bulky dependencies to setup a simple project to try some things out. Also, did I mention that it is one of the very few programming languages that are *blazingly fast*?
Writing Go is fun and I would highly recommend anyone passionate about programming to learn it.


