# Authentication and Authorisation

When we write APIs, it is necessary to know if the client that we are serving, or going to serve, is someone who has the privilege of seeing or modifying any resource on our server. For this reason, we have security mechanisms that when someone requests something, help us identify who we are communicating with and if they have the authority to be making that request.

## Authentication
If you've ever tried to sign into a website, it asks for a few details, one of which is sure to be your password. This is done to make sure that it is indeed *you* that is trying to sign in and not an impersonator. This is pretty much what authentication is. The server is authenticating if it really is you.

Now, a server may have multiple layers and authentication checks before simply handing an account access to you. Let's go over them.

### 2FA (2-Factor Authentication)
2FA is a type of authentication, where, in addition to providing your password, you need another piece of evidence in order to prove ownership of an account. You may set an email or phone number when you create your account, and every time you try to authenticate, the API will try and verify that it is really you using your email or phone number, either by sending you a one-time-password that you would need to enter, or an authentication link that you will need to go to

```{=latex}
\begin{center}
```
![Sam fails an authentication check](src/book/images/4.1.png){width=30%}
```{=latex}
\end{center}
```

### Multi-factor authentication
It is similar to 2FA, but you have to provide more than two pieces of evidence, such as answering security questions that you may have set when you created your account. The goal is to add more layers of security depending on the business and the criticality of data.

## Authorisation

## Relation between Authentication and Authorisation

Authentication and authorisation are two of the most fundamental concepts in API security. They are often used interchangeably, but they're two very different terms. 

Authentication is the process of determining if a person is the same as the one he is claiming to be.

Authorisation is the process of determining if a person is allowed to perform a particular action or not.

We can ask ourselves these questions to help understand them better:

**Authentication** | *Q. Are you really X? Prove it.*
------------------ |   --------------------------
**Authorisation**  | *Q. Are you allowed to do X?*
