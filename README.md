# todo ***[WIP]***

This project provide a CLI tool to manage [`.todo files`](#what-are-todo-files-) 

## What are `.todo` files ? 

Well, glad you asked, we all know it's not a great idea to use todo comments, they will probably stay in your code for a long time. 
In recent years, a lot of people tried to normalize the use of todo-like files ([Todo.md](https://github.com/todomd/todo.md),[todo-md](https://github.com/todo-md/todo-md/)).

I stumbled upon the work of [Matthew Palmer](https://github.com/matthewpalmer): [.todo, a text-based formatting/markdown language for todo lists](https://github.com/matthewpalmer/.todo).

It's clear and *simple*. In Github Flavoured Markdown there is lists that can be parsed as tasks: 

Exemple:
```md
- [] incomplete
- [x] completed
```

Let's reuse that syntax ! 
### Parsing `.todo` files
I want those task to be considered valid
```md 
- [] a task list item
- [] list _syntax_ required
- [] normal **formatting**
-[x]more relaxed about syntax than GFM
- [x]so inconsistencies should be ok
- [] incomplete
- [x] completed
```
<!--  -->
## CLI 
Here are the command I intend to support
```c <!-- I am aware this is not a c snippet, it's just here to make the syntax highlighting work -->
- todo ls // list all the todo items
- todo add <task> // add task
- todo rm <number(s)> // remove the tasks corresponding to the number(s) 
- todo init // create an empty todo.todo file

// eventually
- todo fmt // format the .todo file
- todo clean // remove all 
```

***GLOBAL FLAGS***
- `--path` 
  
  The path to your *`.todo`* file. (Exemple: `todo init --path todofolder/montodo.todo`)

***COMMAND SPECIFIC FLAGS***

- `ls`
  - `--done` 
    
    Only shows the ended tasks

  - `--todo` 
    
    Only show the waiting tasks