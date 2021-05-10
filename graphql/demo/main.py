import json

from fastapi import FastAPI
from graphene import Field, List, Mutation, ObjectType, Schema, String
from graphql.execution.executors.asyncio import AsyncioExecutor
from schemas import CourseType
from starlette.graphql import GraphQLApp

# Fetch all data
# class Query(ObjectType):
#     course_list = None
#     get_course = List(CourseType)

#     async def resolve_get_course(self, info):
#         with open("./courses.json") as courses:
#             course_list = json.load(courses)
#         return course_list


# Fetch only one data
class Query(ObjectType):
    course_list = None
    get_course = Field(List(CourseType), course_id=String())

    async def resolve_get_course(self, info, course_id=None):
        with open("./courses.json") as courses:
            course_list = json.load(courses)
        if course_id:
            for course in course_list:
                if course["course_id"] == course_id:
                    return [course]
        return course_list


class CreateCourse(Mutation):
    class Arguments:
        course_id = String(required=True)
        title = String(required=True)
        instructor = String(required=True)

    course = Field(CourseType)

    async def mutate(self, info, course_id, title, instructor):
        with open("./courses.json", "r+") as courses:
            course_list = json.load(courses)
            course_list.append(
                {
                    "course_id": course_id,
                    "title": title,
                    "instructor": instructor,
                }
            )
            courses.seek(0)
            json.dump(course_list, courses, indent=2)
        return CreateCourse(course=course_list[-1])


class Mutation(ObjectType):
    create_course = CreateCourse.Field()


app = FastAPI()
app.add_route(
    "/",
    GraphQLApp(
        schema=Schema(query=Query, mutation=Mutation),
        executor_class=AsyncioExecutor
    ),
)
