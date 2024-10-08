# Backend Security and Cryptography
When we write backend applications, it is necessary to have good security around it. Building performant, fault-tolerant applications that are efficient and serve our users is one thing, but it is also necessary to make our application secure so that critical data is not compromised. As the complexity of our application rises, so does the vulnerabilities. Addressing these vulnerabilities and finding solutions to them is vital. Securing a backend application is a formidable task, especially when not equipped with the knowledge of the best practices, severity of attacks, and how someone can use them to gain information that they're not supposed to possess. But thankfully, due to our current understanding of backend application security, there are quite a few things that we can do in order to make sure those pesky attackers stay away. Let's go over them.

## Authentication and Authorisation

When we write APIs, it is necessary to know if the client that we are serving, or going to serve, is someone who has the privilege of seeing or modifying any resource on our server. For this reason, we have security mechanisms, where, if someone makes a request, help us identify who we are communicating with and if we should be doing so.

### Authentication
If you've ever tried to sign into a website, it asks for a few details, one of which is sure to be your password. This is done to make sure that it is indeed *you* that is trying to sign in and not an impersonator. This is pretty much what authentication is. The server is authenticating if it really is you to grant you the access to your account.

Now, a server may have multiple layers and authentication checks before simply handing an account access to you. Let's go over them.

**2FA (2-Factor Authentication)**

2FA is a type of authentication, where, in addition to providing your password, you need another piece of evidence in order to prove ownership of an account. You may set an email or phone number when you create your account, and every time you try to authenticate, the API will try and verify that it is really you using your email or phone number, either by sending you a one-time-password that you would need to enter, or an authentication link that you will need to go to.

```{=latex}
\begin{center}
```
![Sam fails an authentication check](src/book/images/4.2.png){width=30%}
```{=latex}
\end{center}
```

**Multi-factor authentication**

It is similar to 2FA, but you have to provide more than two pieces of evidence, such as answering security questions that you may have set when you created your account. A business may add as many layers as they desire in order to safeguard critical data.

### Authorisation
After successfully authentication we still need to make sure if a user is allowed to perform a certain action or not. By definition, **authorisation is the process of assigning access rights and privileges to someone.** For example, in most countries, its civilians are not allowed to get on a military base without proper authorisation; They simply don't have the privilege to do so. In other words, they are unauthorised to perform the action of entering a military base.

**Role-based authorisation**

Role-based authorisation is the most basic method of authorisation. We assign a role to every user in our system, and depending on those roles, we permit or deny a particular request. For example, in an e-commerce application, we may have three roles: **User**, **Seller** and **Admin**.

A **User** would have the authority to browse and buy products, save them to their wishlists, share products, etc.

A **Seller** would have the authority of doing everything a **User** can do, with the added privileges of adding and removing products, etc.

And finally, the **Admin** could be someone who'd have the authority to do many other things such as remove products that are offensive, ban users, etc.

### Authentication vs. Authorisation

Many a times, the terms authentication and authorisation are confused with each other and used interchangeably, but they're two very different terms. 

**Authentication** is the process of determining if a person is the one he claims to be.

**Authorisation** is the process of determining if a person is allowed to perform a particular action or not.

We can ask ourselves these questions to help understand them better:

**Authentication** | *Q. Are you really X? Prove it.*
------------------ |   --------------------------
**Authorisation**  | *Q. Are you allowed to do X?*


## Hashing
Hashing in cryptography and backend security, is the process of converting data, or a piece of text into an encrypted, fixed-length string of characters by making use of an algorithm, also called as a hash function. A hash of something is guaranteed to not change unless the data itself was modified, making it one of the most reliable methods to send data with no corruption, maintaining integrity. By converting your data into its hash with the help of a hash function, you can safely send it across the internet and have the recipient generate the hash of that data, and compare it with the hash that you generated. If it equals, then the data is unadulterated and safe to use, but if it is not, then it was most likely tampered with. But you might be asking: "What does this have to do with how we store passwords in our databases?"

### Password hashing
When we store login information, most of the time, we will have data such as emails, phone numbers, and others. Out of these credentials, one of it is sure to be the password of a user, atleast when you're building an authentication system for your application from the ground up, and not relying on other tools such as Okta, and Google sign-in. This password is crucial to grant access to a user's account. Such a credential is not stored in plain-text as it cannot be compromised. Instead of storing passwords itself as plain-text, we store its hash.


```{=latex}
\begin{center}
```
![Hashing in action](src/book/images/4.1.png){width=80%}
```{=latex}
\end{center}
```

When we get a password to store, we generate its hash, and then store that hash instead of the password in plain-text. So, whenever a sign-in request arrives at our  backend, we get an email and a password generally. What we do is, we generate the hash of the password that was sent to us in the request, and then compare that hash with the hash that is stored on our side. If it equals, then the password is correct, and we grant the user the access to the account. A hash function will generate the same hash for the same unmodified data or text. One of the things to note about hashing is that you cannot decode data back from its hash, making it very secure for storing passwords and making sure that data is untouched.

You might be wondering how secure is a popular hashing algorithm such as SHA256 (Simple Hashing Algorithm 256) really is. If you use SHA256, you have a chance of 1 in $2^{256}$ possibilities to guess the data correctly. $2^{256}$ is such an unfathombly large number, that it contains 77 digits after its most significant digit!

## Types of attacks
When we design backend applications, it is necessary to protect it from malpractices and certain actions from attackers that may compromise our server's data and our users' confidentiality. Let's go over some of these attacks and ways we can secure our backend application against.

### DoS and Brute-force attacks
*DoS* attacks or *Denial of Service* attacks are attacks that are meant to overwhelm or overpower a remote server so as to render it useless for actual users who rely on it.
These kinds of attacks are usually performed by attackers who may have ill intent towards a particular business or a particular person. Denial of Service is achieved by spamming a service with many meaningless requests and worthless data with the help of automated scripts and bot applications. Another similar attack is a brute-force attack meant to steal the password of another user by simply trying multiple possible combinations. All of these attacks whether brute-forcing or Denial of Service can be prevented by utilizing a security mechanism called rate-limiting.

// TODO: image of brute-forcing

**Rate limiting**

Rate limiting is the process of limiting or throttling the amount of incoming requests so as to prevent the server from being overwhelmed and continue to serve our users. 

> You must've experienced something like this when trying to find the pin or password of a mobile phone by trying multiple combinations. The phone would simply prevent you from entering anything for some time or would even lock you out.

There are multiple ways to limit or throttle incoming requests, or, rate-limiting algorithms that are present in practice today. Let's go over a few of the most used ones:

### Man in the middle attack
### Packet sniffing attack
