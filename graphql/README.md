# GraphQL

## Queries

Queries are used to fetch data, just like GET requests in the REST API architecture.

## Schema

The schema describes our GraphQL service, what data it contains, and the format for that data. From our query, we’ve seen that we can specify what data will be sent to us and how we want that data presented.

## Mutations

GraphQL mutations is used to add new data to our data store or update existing data.

# Example with FastAPI

## Prerequisite

```bash
❯ cd demo

❯ python -m venv env

❯ source env/bin/activate

❯ pip install -r requirements.txt
```

## Run application

Start a simple application.

```bash
❯ uvicorn main:app --reload
INFO:     Uvicorn running on http://127.0.0.1:8000 (Press CTRL+C to quit)
INFO:     Started reloader process [47079] using statreload
INFO:     Started server process [47099]
INFO:     Waiting for application startup.
INFO:     Application startup complete.
```

Check the endpoint ([http://localhost:8000/](http://localhost:8000/)).

![image](https://user-images.githubusercontent.com/45956169/117667292-b9ebe500-b1df-11eb-9b53-7d6048723ed2.png)

## Try GraphQL

### Fetching all data

Let’s paste the following query on the left pane and make our API call by clicking on the run button:

```JSON
{
  getCourse {
    courseId
    title
    instructor
    publishDate
  }
}
```

The right pane looks like below.

![image](https://user-images.githubusercontent.com/45956169/117669488-e6a0fc00-b1e1-11eb-9cb2-fc021197b3da.png)

Also, you can check it with `curl` command.

```bash
❯ curl -X POST http://localhost:8000/ -H "Content-Type: application/json" -d '{"query": "{ getCourse { courseId\n title\n instructor\n publishDate\n} }"}' | jq -r
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   543  100   468  100    75   114k  18750 --:--:-- --:--:-- --:--:--  132k
{
  "data": {
    "getCourse": [
      {
        "courseId": "1",
        "title": "Python variables explained",
        "instructor": "Tracy Williams",
        "publishDate": "12th May 2020"
      },
      {
        "courseId": "2",
        "title": "How to use functions in Python",
        "instructor": "Jane Black",
        "publishDate": "9th April 2018"
      },
      {
        "courseId": "3",
        "title": "Asynchronous Python",
        "instructor": "Matthew Rivers",
        "publishDate": "10th July 2020"
      },
      {
        "courseId": "4",
        "title": "Build a REST API",
        "instructor": "Babatunde Mayowa",
        "publishDate": "3rd March 2016"
      }
    ]
  }
}
```

If you remove `instructor` and `publishDate` from the query, you can get `courseId` and title` as desired.

```bash
❯ curl -X POST http://localhost:8000/ -H "Content-Type: application/json" -d '{"query": "{ getCourse { courseId\n title} }"}' | jq -r
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   273  100   227  100    46  20636   4181 --:--:-- --:--:-- --:--:-- 24818
{
  "data": {
    "getCourse": [
      {
        "courseId": "1",
        "title": "Python variables explained"
      },
      {
        "courseId": "2",
        "title": "How to use functions in Python"
      },
      {
        "courseId": "3",
        "title": "Asynchronous Python"
      },
      {
        "courseId": "4",
        "title": "Build a REST API"
      }
    ]
  }
}
```

### Fetching only one course

Replace `Query` class like below.

```python
class Query(ObjectType):
  course_list = None
  get_course = Field(List(CourseType), id=String())
  async def resolve_get_course(self, info, id=None):
    with open("./courses.json") as courses:
      course_list = json.load(courses)
    if (id):
      for course in course_list:
        if course['id'] == id:
          return [course]
    return course_list
```

Let’s paste the following query on the left pane and make our API call by clicking on the run button:

```json
{
  getCourse(courseId: "2") {
    courseId
    title
    instructor
    publishDate
  }
}
```

The right pane looks like below.
![image](https://user-images.githubusercontent.com/45956169/117673190-8744eb00-b1e5-11eb-80d2-252184cb136d.png)

With `curl` command,

```bash
❯ curl -X POST http://localhost:8000/ -H "Content-Type: application/json" -d '{"query": "{ getCourse ( courseId: \"2\") {courseId\n title\n instructor\n publishDate\n } }"}' | jq -r
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   233  100   139  100    94  34750  23500 --:--:-- --:--:-- --:--:-- 58250
{
  "data": {
    "getCourse": [
      {
        "courseId": "2",
        "title": "How to use functions in Python",
        "instructor": "Jane Black",
        "publishDate": "9th April 2018"
      }
    ]
  }
}
```

### Mutation

We can now the mutation by running the following query in our GraphQL client.

```JSON
mutation {
  createCourse(
    courseId: "11"
    title: "Python Lists"
    instructor: "Jane Melody"
  ) {
    course {
      courseId
      title
      instructor
    }
  }
}
```

You can get the response like below

![image](https://user-images.githubusercontent.com/45956169/117679778-89aa4380-b1eb-11eb-9a5b-f93e66b76263.png)

Check that `courses.json` is updated or not.

```bash
❯ cat courses.json | jq -r
[
  {
    "course_id": "1",
    "title": "Python variables explained",
    "instructor": "Tracy Williams",
    "publish_date": "12th May 2020"
  },
  {
    "course_id": "2",
    "title": "How to use functions in Python",
    "instructor": "Jane Black",
    "publish_date": "9th April 2018"
  },
  {
    "course_id": "3",
    "title": "Asynchronous Python",
    "instructor": "Matthew Rivers",
    "publish_date": "10th July 2020"
  },
  {
    "course_id": "4",
    "title": "Build a REST API",
    "instructor": "Babatunde Mayowa",
    "publish_date": "3rd March 2016"
  },
  {
    "course_id": "11",
    "title": "Python Lists",
    "instructor": "Jane Melody"
  }
]
```

With `curl` command,

```bash
❯ curl -X POST http://localhost:8000/ -H "Content-Type: application/json" -d '{"query": "mutation { createCourse (courseId: \"123\"\n title: \"Golang\"\n instructor: \"John\") { course { courseId\n title\n instructor } } } "}' | jq -r
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   239  100    92  100   147   3833   6125 --:--:-- --:--:-- --:--:--  9958
{
  "data": {
    "createCourse": {
      "course": {
        "courseId": "123",
        "title": "Golang",
        "instructor": "John"
      }
    }
  }
}

❯ cat courses.json | jq -r '.[] | select(.course_id == "123")'
{
  "course_id": "123",
  "title": "Golang",
  "instructor": "John"
}
```

### Handling request errors

Let's change the application so that it can return an error. Add the following code before adding a new data.

```python
for course in course_list:
  if course['id'] == id:
    raise Exception('Course with provided id already exists!')
```

If you send a request like below,

```bash
mutation {
  createCourse(
    courseId: "1"
    title: "Python Lists"
    instructor: "Jane Melody"
  ) {
    course {
      courseId
      title
      instructor
    }
  }
}
```

then you can get the following response.

![image](https://user-images.githubusercontent.com/45956169/117981827-60ff8680-b370-11eb-83e0-1a9fc5593fca.png)

Check it with `curl` command.

```bash
❯ curl -X POST http://localhost:8000/ -H "Content-Type: application/json" -d '{"query": "mutation { createCourse (courseId: \"1\"\n title: \"Golang\"\n instructor: \"John\") { course { courseId\n title\n instructor } } } "}' | jq -r
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   308  100   163  100   145  40750  36250 --:--:-- --:--:-- --:--:--  100k
{
  "data": {
    "createCourse": null
  },
  "errors": [
    {
      "message": "Course with provided course_id already exists!",
      "locations": [
        {
          "line": 1,
          "column": 12
        }
      ],
      "path": [
        "createCourse"
      ]
    }
  ]
}
```
