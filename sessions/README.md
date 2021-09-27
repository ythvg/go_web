## Sessions

This example will show how to store data in session cookies using the popular gorilla/session package io Go.

Cookies are small pieces of data stored in the browser of a user and are sent to our server on each request. In them, we can store e.g. whether or not a user is logged in into our website and figure out who he actually is (in our system).
In this example we will only allow authenticated users to view our secret message on the /secret page. To get access to it, the will first have to visit /login to get a valid session cookie, which logs him in. Additionally he can visit /logout to revoke his access to our secret message.
