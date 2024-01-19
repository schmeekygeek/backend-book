# Communication Protocols
Consider this scenario: You see an old woman at the supermarket that is having trouble carrying her basket. So you offer to help, but to your surprise, she doesn't speak the same language as you! How would you ask her if she needs any help?

```{=latex}
\begin{center}
```
![Sam has trouble talking to an old woman](src/book/images/3.1.png "3.1"){width=40%}
```{=latex}
\end{center}
```

In order to verbally communicate with a person effectively, it is necessary that they speak the same language as you. Pretty much the same goes for computers trying to communicate with each other over the internet. If one speaks a language the other doesn't, they won't be able to exchange information at all.

Basically, these 'languages' are called communication protocols. Communication protocols determine who speaks, what is being said, what can be said, how long it should take, etc. Computers over the internet communicate over a particular protocol. They are a pre-defined set of rules that needs to be followed in order to exchange information successfully.

Let's look at one of the most fundamental protocols of the internet, or, the **Internet Protocol** suite, a set of communication protocols, the transmission control protocol.

## Transmission Control Protocol (TCP)
Transmission Control Protocol, or simply, TCP, is a highly-reliable communication protocol which is connection oriented. It needs a stateful handshake to occur between two machines before packets of data can be exchanged.

### The 3-Way Handshake
Two hosts over a network may begin communicating over TCP after they've made a successful handshake. This handshake occurs in 3 steps. Which is why it is called the 3-way handshake.

```{=latex}
\begin{center}
```
![The 3-Way Handshake](src/book/images/3.2.png "3.2"){width=80%}
```{=latex}
\end{center}
```

**Step 1:** Host #1 sends a message with the SYN (Synchronise Sequence Number) flag set.

**Step 2:** Host #2 responds with its own message with the SYN flag set, along with an ACK (acknowledgement) that the SYN of the client was received successfully.

**Step 3:** Host #1 receives the SYN+ACK, then responds back with another ACK.

### What makes TCP so reliable?
Every single time a data packet is sent over TCP, an acknowledgement message is sent by the receiver to the sender indicating that the message sent was received successfully. If the sender doesn't get the acknowledgement message after a set period of time, it simply sends the data packet again. This makes TCP highly reliable in environments where data is critical and cannot be lost. There may be instances of TCP not being reliable, but these are exceptionally rare.

Let's look at our next protocol based on TCP itself, HTTP.

## HTTP
HTTP stands for Hyper Text Transfer Protocol. It is the most used protocol over the internet.
