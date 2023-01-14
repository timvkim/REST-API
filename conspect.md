Создаем сущности: `Course` `User`

Описываем endpoints:

|      Method      | URL Pattern | Handler| Action |
|:----------------:|:-----------:|:------:|:------:|
|     GET          |     /v1/courses     |  GetAllCourses   |  Show list of courses   |
|  POST  |     /v1/courses     |  CreateCourse   |  Adding new course   |
| PUT |     /v1/course/{id}      |  UpdateCourseById   |  Edit info about course   |  
| DELETE |     /v1/course/{id}      |  DeleteCourseById   |  Delete course   |  
