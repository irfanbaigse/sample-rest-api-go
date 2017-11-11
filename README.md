# Sample REST Api in GO

Clone the repo: 

``
$ git clone https://github.com/irfanbaigse/sample-rest-api-go
``

Compile and install the code with:

``
$ go install
``

In the project directory, run the application:

``
$ PORT=8191 ./bin/rest-api
``

Make a HTTP request to the application:

[CURL](https://i.imgur.com/VHX4hTh.png)


# Deploying To Heroku

Create a Heroku account [here](https://signup.heroku.com/www-header)

``
$ brew install heroku
``


``
$ heroku login
``

``
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git sample-rest-api-go
``

``
$ git push heroku master
``
[Push](https://i.imgur.com/FeD5oED.png)


now goto https://sample-rest-api-go.herokuapp.com

By default, your app is deployed on a free dyno. Free dynos will sleep after a half hour of inactivity. To avoid that run below command and make sure your account is verified. 

``
$ heroku ps:scale web=1
``
