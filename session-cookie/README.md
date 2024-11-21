# Session - Cookie Based Authentication & Authorization

__What is a Cookie ?__

A cookie is a file that holds information that a server can write to a client's machine if a client allows cookies to be written. For every request to a particular domain, the client's web browser first looks to see if there is a cookie from that domain on the client machine. With each request to that domain, the browser will send a cookie if a cookie has been written for that particular domain.

`Cookies are domain specific`

When we access different apps and websites, there are basically three important security steps occurring continuously:

`Identity`

`Authentication`

`Authorization`


All of these processes occur at API Gateway.

__Identity:__

User claims that he is Nani

__Authentication:__

User proves that he is Nani by providing password

__Authorization:__

Depending on the privilege of the user, access is given to some resources.


Session-cookie based authentication resolves an issue where HTTP basic access authentication fails to handle whether a user is logged in. 

To remember the status of the user during the visit, it generates a session ID. This session ID is stored on both the server-side and in the cookie of the client to authenticate. 

It is called a session-cookie because it's a cookie that has the session ID stored inside. The users still have to give their username and password once, but afterwards the server just creates a session for the user's visit. 

Following requests do include the cookie, hence the server can check the client-side against the server-side session ID.


# Let's break down how the session-based authentication process works in the context of the provided Go code:

# 1. Client Sends a Request to Access a Protected Resource

__Initial Request:__ The client attempts to access a protected resource (e.g., `/healthcheck`) without being authenticated.


__Login Prompt:__ If the client is not authenticated, they must first log in by sending their credentials `(username and password)` to the `/login` endpoint.

# 2. Server Verifies Credentials

__Credential Verification:__ The LoginHandler function receives the login request and decodes the credentials. It then calls `db.GetUser(creds.Username)` to retrieve the user from the database.

__Error Handling:__ If no user is found or if the password does not match, the server responds with an `"Invalid credentials"` message.

# 3. Session Creation

__Session ID Generation:__ Upon successful credential verification, a new session is created using `sessions.CookieStore`. A unique session ID is generated and stored in the session data.

__Session Storage:__ The session data, including the user ID and session ID, is stored server-side (in memory in this example).

# 4. Server Sends Session ID as a Cookie

__Set-Cookie Header:__ The server sends a Set-Cookie header in its response to the client, containing the session ID. This is handled by `session.Save(r, w)` in the `LoginHandler`.

# 5. Client Stores Session Cookie

__Cookie Storage:__ The client stores this session cookie, which will be used for subsequent requests to authenticate itself.

The client-side storage of the session cookie is not explicitly shown. However, this process is typically handled automatically by the client's browser when it receives the Set-Cookie header from the server.

In the `LoginHandler` function, the following line is responsible for sending the session cookie to the client:

`session.Save(r, w)`

# 6. Subsequent Requests with Session Cookie

__Request with Cookie:__ For any subsequent request to protected resources, the client includes the session cookie in the request headers.
For subsequent requests, the client's browser will automatically include this stored cookie in the request headers. This is handled by the browser itself, not explicitly in your server-side code.

The `AuthMiddleware` function then uses this cookie to authenticate the user:


`session, _ := sm.Store.Get(r, "session-name")`

`if session.Values["user_id"] == nil {`

`    http.Error(w, "Unauthorized", http.StatusUnauthorized)`
    
`    return`
    
`}`

# 7. Server Checks Session ID

__Session Validation:__ The `AuthMiddleware` function intercepts requests to protected endpoints. It retrieves and validates the session using `sm.Store.Get(r, "session-name")`.

__Access Granting:__ If the session ID is valid and corresponds to an authenticated user, access to the requested resource is granted.

# 8. Session Invalidation

__Logout or Expiration:__ When a user logs out via `/logout`, or when a session expires (based on MaxAge settings), the session is invalidated.

__Session Deletion:__ The server sets `session.Options.MaxAge = -1` to delete the session cookie, effectively logging out the user.





