I have made a attendance app which marks attendance using facial recognition api .In this I have used React ,Go ,Gin, Mongodb and face++ api to compare images
#How to run#
1.	First you have to clone this repository on your laptop  make sure you have installed and set up node golang and mongodb
2.	Then first run command (npm start) that will run your frontend on local host:3000 then go to backend folder and then run command(go run main.go) it will run your backend on localhost:9000
3.	Then as you don’t have an admin you will have to create a admin using postman or other services .You are creating a admin using postman because you have to create admin once you don’t have to create a admin on frontend  so admin will not created on frontend.
4.	Now to create a admin you have to make a post request to http://localhost:9000/admin/signup in the body of request you have to write email and password in json format now your admin will be created.
5.	Now you can open your localhost:3000 in which you will find login as admin page first in which you have to fill the admin details .then you will be directed to admin dashboard it will not be showing any user now , to create a user on the left side dashboard you can click create user then you will be redirected to create user where you can create a user with image
6.	Now you can click logout and you will see a login as user button on bottom click that and then login as user
7.	Till now it will be showing no attendance you can mark your attendance by clicking mark attendance on dashboard which clicks your picture and marks attendance for today it picture matches .
8.	On user dashboard you can see user profile also
This was the whole process of running
#Explanation of code#
1.You will see a backend folder in which I have my whole backend .it contains main.go which is the entry point of my backend and it runs my backend then you will see handlers.go which contains all my handler function for my routes then you can see middleware.go it is basically authentication middleware for authenticated routes .helper.go contains function to generate refdresh tokens.server.go contains function to open connection with mongodb and at the lest server.go contains all the routes
2.except the backend named folder is my frontend in src you can find my components folder which contains all the pages and app.js is the main entrypoint of my frontend
