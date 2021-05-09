from graphene import ObjectType, String


class CourseType(ObjectType):
    course_id = String(required=True)
    title = String(required=True)
    instructor = String(required=True)
    publish_date = String()
