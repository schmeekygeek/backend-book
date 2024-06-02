# Communication Protocols
Consider this scenario: You see an old woman at the supermarket that is having trouble carrying her basket. So you offer to help, but to your surprise, she doesn't speak the same language as you! How would you ask her if she needs any help?

```{=latex}
\begin{center}
```
![Sam has trouble talking to an old woman](src/book/images/3.1.png){width=30%}
```{=latex}
\end{center}
```

In order to verbally communicate with a person effectively, it is necessary that they speak the same language as you. Pretty much the same goes for computers trying to communicate with each other over the internet. If one speaks a language the other doesn't, they won't be able to exchange information at all.

Basically, these 'languages' are called communication protocols. Communication protocols determine who speaks, what is being said, what can be said, how long it should take, etc. Computers over the internet communicate over a particular protocol. They are a pre-defined set of rules that need to be followed in order for information to be exchanged successfully.

Let's look at one of the most fundamental protocols of the internet, or, the **Internet Protocol** suite, a set of communication protocols, the transmission control protocol.

## Transmission Control Protocol (TCP)
Transmission Control Protocol, or simply TCP, is a highly-reliable communication protocol which is connection oriented.  Many protocols are based on it, such as HTTP, FTP, Websockets, etc. It needs a stateful handshake to occur between two machines before packets of data can be exchanged. Let's learn how a TCP connection is made between two hosts.

### The 3-Way Handshake
Two hosts over a network may begin communicating over TCP after they've made a successful handshake. This handshake occurs in 3 steps. Which is why it is called the 3-way handshake.

```{=latex}
\begin{center}
```
![The 3-Way Handshake](src/book/images/3.2.png){width=80%}
```{=latex}
\end{center}
```

**Step 1:** Host #1 sends a message with the SYN (Synchronize Sequence Number) flag set.

**Step 2:** Host #2 responds with its own message with the SYN flag set, along with an ACK (acknowledgement) that the SYN of the client was received successfully.

**Step 3:** Host #1 receives the SYN+ACK, then responds back with another ACK.

### What makes TCP so reliable?
Every single time a data packet is sent over TCP, an acknowledgement message is sent by the receiver to the sender indicating that the message sent was received successfully. If the sender doesn't get the acknowledgement message after a set period of time, it simply sends the data packet again. This makes TCP highly reliable in environments where data is critical and cannot be lost. There may be instances of TCP not being reliable, but they are exceptionally rare.

Now, for our next protocol based on TCP itself, HTTP.

## Hyper Text Transfer Protocol (HTTP)
HTTP stands for Hyper Text Transfer Protocol.
It is the foundational protocol for all types of data exchange on the web. Since HTTP connections are based on the request-response model of the internet, they are unidirectional. Which is why an HTTP connection, or an HTTP request, is always made by the client that is requesting a resource from the server. The server will then respond back with the requested resource, depending on whether the client is supposed to see it or not, in which case it will reject the request for said resource.

```{=latex}
\begin{center}
```
![Request-Response model](src/book/images/3.3.png){width=60%}
```{=latex}
\end{center}
```

There exists many other protocols such as FTP and MTP, but they are beyond the scope of what's required in this book.

### URL and URI

You have most likely seen this long string of *mumbo jumbo*. Let's break it down to try and understand what these frightening symbols and characters actually mean.

```{=latex}
\begin{center}
```
![A typical URL](src/book/images/3.4.png){width=60%}
```{=latex}
\end{center}
```

1. **Protocol:** The protocol for this URL
2. **Host:** This is the host, or the server that we are making the request to.
3. **Path:** The path, or the location of the resource on the host. Here, `/images` is the path, and `image.jpg` is the resource we're requesting.
4. **Query Parameters:** These are query parameters. Their purpose is to modify variables that the endpoint might expect to fine-tune our request. They begin with a *?* (Question mark). If there are two or more such values, they are separated by an '*&*' (Amper's And).

What we just saw above was an example of a URL. **A URL stands for Uniform Resource Locator**. It *locates* the resource that we're trying to fetch from the host. Another term similar to URL is URI that stands for **Uniform Resource Identifier**. It is an identifier, or a name, for any resource that we're trying to request.

At first glance, it may seem that both of these terms mean the same thing: To identify our resource, which will also be used to locate it. A URI identifies, and a URL locates. But a URL is also a URI, since it is also used to identify. But not every URI is a URL, since we usually have no idea how to locate the identifier.
```{=latex}
\begin{center}
```
![URL-URI Venn diagram](src/book/images/3.5.png){width=20%}
```{=latex}
\end{center}
```
In other words, **A URL is a subset of a URI.**

### Anatomy of an HTTP request
An HTTP request sent by a client contains many parts that define its structure. Let's try and understand them.

Here is a response that is received upon calling the API *https://httpbin.org/get* using `curl`, a command-line tool used to transfer data from or to a server.

```{=latex}
\begin{center}
```
![A basic HTTP request's response](src/book/images/3.6.png){width=75%}
```{=latex}
\end{center}
```

1. **HTTP version:** The protocol version of HTTP being used.
2. **Status Code:** The status code or response code, sent by the server, indicating the status of the request sent.
3. **Headers:** This is the metadata.
4. **Payload:** The data itself. In our case, the format of the data is `json`

The type of request that we sent above was a `GET` request. `GET` is also called an HTTP verb. Let's look at some of these attributes of HTTP requests.

### HTTP Verbs
HTTP verbs are request methods that let us specify certain actions to be performed with our HTTP requests. They essentially are a vague type of the request that we are sending. For example, if we are trying to fetch a resource, we will use the `GET` method in the request, and make our backend accept the same. If we are trying to send some kind of data, we may have to use the `POST` method. There are many HTTP verbs, each having their own significant meaning. However, there are four fundamental HTTP verbs that are most commonly used in backend development. They are:

1. **GET:** Get or fetch a certain resource from an API endpoint.
2. **POST:** Send some data to the server.
3. **PUT:** Update a resource on the server.
4. **DELETE:** Delete a resource on the server.

These HTTP verbs are also called as CRUD operations where CRUD stands for CREATE (POST), READ (GET), UPDATE (PUT) and DELETE. Some other HTTP verbs include, CONNECT, HEAD, OPTIONS, PATCH, etc.

### HTTP Headers
HTTP headers let the client and servers share information additional to the payload when communicating over HTTP. They are also called metadata because they contain the data that is used to describe the payload and/or used to set other options and flags. Headers are set as key-value pairs separated by a ':' (colon).

There are 3 types of headers:

1. **Request headers:** Headers set by the client containing information about the client or more information about the resource being requested.
2. **Response headers:** Headers set by the server that contain details about the server or time, etc.
3. **Payload headers:** Client or server independent headers that provide details about the payload.

Here are some of the most common HTTP headers:

- `user-agent`: Client application or web browser identifier.
- `authorization`: Used to set any credential or bearer token that the server might need whenever accessing a protected resource. Usually a JWT (JSON Web Token).
- `content-type`: Indicates the media type of the resource.
- `content-length`: Length of the content (in bytes).
- `cookie`: Sent by the client to help the server recognize a session or a user.
- `accept`: Sent by the client to indicate media types that are acceptable.

### HTTP Status Codes
When we interact with APIs, it is necessary for us to know what happened with the request that was sent. Has it completed successfully, did something go wrong, if yes, what went wrong? Is the request being processed, am I being redirected, etc., are many such questions that arise when a request is made. To address this problem, every HTTP response has a status code attached to it that determines if the request was a success or not. These status codes are also called response codes. They exist to inform the client about the status of the request. You have most likely heard about or seen the error: '404 - not found'. This is one of the most common HTTP status codes that informs the client that the resource you were looking for couldn't be found.

A status code is a three-digit number. The first digit signals the general status or success/failure of a response. The next two digits are more specific as to what exactly happened. Here is a range for the status codes.

Range          | Message
-------------- | -------
100 - 199      | Informational
200 - 299      | Success or OK
300 - 399      | Redirection
400 - 499      | Client-side error
500 - 599      | Server-side error

Here are some of the most common status codes used today:

Code|Description
--- |----------
200 | `ok`
201 | `created`
400 | `bad request`
401 | `unauthorized`
403 | `forbidden`
404 | `not found`
500 | `internal server error`

### The problem with HTTP
Now that we've discussed all the quirks of HTTP, there is still one small caveat about it. HTTP is based on a request-response model, mostly used in RESTful APIs (more on this later) meaning that you can't get data from a server without requesting for it initially.
Suppose, you are trying to build a video streaming platform where you can have users stream some kind of audio or video and have clients that are feeded that data.
You can't optimally achieve this by using traditional HTTP because you'll have to poll the server and request it to send you the frames of data. Ideally you'd want the server to stream the data to you until the client connection ends.
To give you another example, say you want to build a chat application. You can have a server that you send details such as the content of the message, to whom the message is meant for, etc. The server cannot directly send this message to the client for whom the message is meant to go to. The recipient would need to poll the server at regular intervals and ask if there are any new messages, which is not very ideal.

```{=latex}
\begin{center}
```
![The server is annoyed at Sam's requests](src/book/images/3.7.png){width=75%}
```{=latex}
\end{center}
```

You'd be wasting precious resources of the server and the user, requesting data that simply doesn't exist. Another solution would be to make use of a technique called long-polling, where you send a request but instead of receiving a response immediately, you keep the connection open so that the server can respond whenever it receives a message. But all of these are just work-arounds and not actual solutions. This is where a special protocol, Websocket, comes in.

## Websocket
Websocket is a special type of communication protocol built on top of HTTP that allows for bi-directional communication between the client and the server. A websocket server can have multiple websocket connections open at a time to which it can send and receive data from.
You can think of a websocket connection as a pipeline where the data is not synchronized and any party can send and receive data whenever they want, as long as the connection is open. This protocol becomes an ideal choice for applications where the server also has a lot of data to send like a video streaming platform such as YouTube, or doesn't know when new data will be available for a client. Websocket connections are made by establishing a handshake, similar to TCP, between two hosts. It is established after upgrading an initially made HTTP connection to a websocket connection with the use of HTTP headers.

In the next chapter we shall learn about two basic security mechanisms that APIs implement in order to serve the right client.
