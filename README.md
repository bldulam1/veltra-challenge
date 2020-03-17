# Short URL Architectural Design

## Introduction
This is a system design for a URL shortening service. Basically, the service transforms a long url string into a unique, fixed length (normally 6 to 8 characters) and shorter URL.

The system can provide to the user the following microservices:

* Create a shorter URL (POST request)
  * Input
    * Long URL
    * UserID (optional, for protected URL)
    * Expiration date (optional)
  * Output
    * Short URL (URL ID)
* Retrieve the original (long) URL using a short URL, then redirect to the original URL. (GET request)
  * Input
    * Short URL (URL ID)
    * UserID (for protected URLs)

Income Generating Services
In order that this service can generate income, it would be good to include a service for analyzing the most frequently visited urls. Having this information allows the the company to inform other businesses of the most frequently visited sites.


## Framework Choices
There are many frameworks that can be used to implement this system. I will be talking about the frameworks that I am very familiar.

**NodeJS (Express)**

This framework is easy to develop. It is one of the most popular frameworks nowadays. It has a lot of ready-made packages that can be used for other smaller tasks like reading and updating a database, and other data manipulation operations. Another advantage of this framework is that it can be easily deployed to cloud services like AWS, GCP, Azure, etc.

NodeJS has some disadvantages like being not able to perform multithreading, and not suitable for long running operations. But these disadvantages does not really affect the kind of application we are building.

**Echo (GO/Golang)**

GO has several popular frameworks like Gin, Gorilla Mux, but I find the Echo framework/library to be more developer-friendly. Go is famous for its speed being a compiled language, and its support for concurrency operations. The execution time of applications written in GO could be faster than NodeJS. Its syntax is also easy to learn.

Go is statically typed, it has less flexibility compared to dynamically typed languages. It has fewer third party libraries than NodeJS.

**Flask and Django (Python)**

Python's Flask and Django are also viable frameworks for this application. Being written in Python, makes these frameworks easy to understand and makes the project fast to develop. Python is also supported in major cloud services.

The execution time of Python applications could be slower than applications written in GO.

For the services of the system, I would like to implement this system using the AWS or Netlify cloud services. For this reason, I think I will be implementing using lambda functions.

## Deployment and infrastructure Choices

The reason why I want to use serverless or lambda services (Function as a Service) is that it is scalable. It can easily handle when the usage of the application is high and automatically reduces its resources when the usage is low. Thus, we only pay for the exact computing resources that the application used. Also, the management of the servers during the fluctuation of the network traffic is taken care by the cloud provider.

One disadvantage of using lambda services is that it takes time for cold-starting. Lambda functions are good for short-running short-lived operations. Using Lambda operation requires Database as a service as its partner. These disadvantages does not hurt much the kind of application we are building.

Other solutions for the infrastructure is using traditional physical servers. These servers could be cost-saving if the application has constant high-user demand and the operation of functions is time/resource consuming. It also has some disadvantages like the you will be paying for the servers that has been contracted even if the user demand is low. Also the management of the servers during the user demand fluctuations should be taken care by the company.

For this application, I think the lambda services would be the most cost-effective easier to maintain.

## Backend/ Preferred Communication Tools

For the client side, RESTful communication would be the suitable form of communication. RESTful APIs are easy to understand, and it already has been adapted to a lot of real life applications.

For the developers, GraphQL is also a viable option. It makes the documentation of APIs self-documenting. For the developers, there could be more information that will be fetched from the database like the statistics of the urls (number of times accessed, demography of who used the url, time of access, etc). Fetching customized set of data is easily done using GraphQL.

A mixture of GraphQL and RESTful API will be used as communication among the services.

## Issues to consider

Some major issues of the kind of application we are building are availability and scalability of the application and security for the kind of data that is sent by users (some URLs might include API_KEY of the user for some apps).

*Availability and Scalability*

During the initial launching of the product, the system will be using lambda services. Scalability and availability will be taken care by the cloud provider.

*Security*

User authentication might be required for confidential URL shortening purposes. Some URLs might be restricted to some group of people or individuals. So this will be taken care by having an authentication feature. User activities can also be tracked having a separate services for url shortening-related user activities.





