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

Basically, these 'languages' are called communication protocols. Communication protocols determine who speaks, what is being said, what can be said, how long it should take, etc. Computers over the internet communicate over a particular protocol. They are a pre-defined set of rules that needs to be followed in order for information to be exchanged successfully.

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
An HTTP request contains many parts that define

### HTTP Headers

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

Code | Description
---  | ----------
200  | *OK* - The request was successful.
201  | *CREATED* - Created a new resource, such as an account.
400  | *BAD REQUEST* - Invalid or malformed request
401  | *UNAUTHORIZED* - Client is unauthorized, needs authentication.
403  | *FORBIDDEN* - Client is forbidden to access the resource.
404  | *NOT FOUND* - The resource/s requested couldn't be found
500  | *INTERNAL SERVER ERROR* - Something went wrong at the server side

In the next chapter we shall learn about two basic security mechanisms that APIs implement in order to serve the right client.
