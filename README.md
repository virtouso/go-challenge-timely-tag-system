# How many users? (Go Challenge)

Suppose there is a "User Segmentation Service" (USS) that segments users based on their activities.
For example, if a user visits sports news, USS classifies and tags "sport" to him.
So we have a pair: (the user_id, and the segment) for example, (u104010, "sports").

We want to develop an Estimation Service (ES) that interacts with USS directly.
ES receives the pair from USS as input and stores it.
The responsibility of ES is to answer a simple query: "How many users exist on a specific segment?".
For example, "how many users are in the sports segment?".

![](https://raw.githubusercontent.com/ArmanCreativeSolutions/go-challenge/main/Untitled%20Diagram.drawio.png?raw=true)

The query is simple, but two assumptions may make it a little challenging:
- A specific user remains just two weeks on a segment. After that,
we should not count "u104010" on the sports segment.
- There are millions of users and hundreds of segments. So your solution(s) must be scalable


## Requirements

- Implement a (REST API, RESTful API, soap, Graphql, RPC, gRPC, or whatever protocol you prefer)
interface to receive data (user_id, segment pair) from USS. 
- Implement a method to estimate the number of users in a specific segment. ( `func estimate(segment) -> number of users`)

## Implementation details

Try to write your code as reusable and readable as possible.
Also, don't forget to document your code and clear the reasons for all your decisions in the code.

If your solution is not simple enough for implementing fast, you can just describe it in your documents.

Use any tools that you prefer just explain the reason of choices in your documents.
For example explain why you choose REST API for receiving data.

It is more valuable to us that the project comes with unit tests.

Please fork this repository and add your code to that.
Don't forget that your commits are so important.
So be sure that you're committing your code often with a proper commit message.

# How Problem Solved

first of all. in my solution any user can participate in multiple segments or tags and is not limited to single segment

## general explanation

used redis `zset` to store users in segments. their score is the time that watched segment page or show. 
in time intervals i remove users from segment zset with lower score. 

### possible alternatives.
* using time series database like influx but need  index on user_id to make sure in given duration duplicate not exist. 
* certainly its slower compared to redis.

## steps

1. write scores. user score is now unix time.every segment is a seperate zset on redis. 
2. to count popularity just call count on redis zset. 
3. every time interval like every hour or every 3 hour remove the scores lower that given number. unix version of time is simply comparable like numbers. this is the golang code:

         time := time.Now().UTC().Add(-time.Duration(minutesCount) * time.Minute)
          unixTimeSeconds := time.Unix()
          count, err := repository.Redis.RemoveBelowScoreAll(float64(unixTimeSeconds))


## How to Test
just run the docker compose. 
for debugging and breakpoints use docker-compose-debug.yml that support delve.

## things could be better.

its just a test. many things could be better to make the code production ready

*  logs should not be  apart of code. it only makes the code a chaos.
*  configurations should be multi level. only reading from env is not best practice. 
* required time was 14 days but for testing i made it 4 minutes and made restart schedule every minute. just change it based on docker compose environment variables. 
* did not using jwt or auth. currently just sending segments and user names