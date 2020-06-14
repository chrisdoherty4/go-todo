# RememberMe

> I need to remember _todo_ that

A todo list application written for academic purposes.

## Todo

- Create accessors/mutators for types. It will future proof the code so we can add any validation logic as needed and lay the foundation for LastUpdated and CreatedAt attributes. It will, however, introduce complexity around JSON marshalling.
- Change the todo.Repository interface to include a 'Save' and 'New' will create a new instance of something in the repository.
- Remove MarkComplete from repositories - it's not a function of a Repository instance
