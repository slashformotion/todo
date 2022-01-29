# todo
This project provide a CLI tool to manage [`.todo`](https://github.com/slashformotion/.todo) files. Currently this project support the `.todo` Specification *v0.0.1*.


### Parsing `.todo` files
Those task to be considered valid even if some of them are not entirely complient with the [.todo file spec](https://github.com/slashformotion/.todo). You can mess your formating, the CLI will still be able to parse it and format it correctly
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
Here are the supported commands:
```c <!-- I am aware this is not a c snippet, it's just here to make the syntax highlighting work -->
- todo ls // list all the todo items
- todo add <task> // add task
- todo rm <number(s)> // remove the tasks corresponding to the number(s)
- todo init // create an empty todo.todo file
- todo clean // remove all tasks
- todo fmt // format the .todo file
```


***GLOBAL FLAGS***
- `--path`

  The path to your *`.todo`* file. (Exemple: `todo init --path todofolder/mytodo.todo`)

<!-- ***COMMAND SPECIFIC FLAGS***

- `ls`
  - `--done`

    Only shows the ended tasks

  - `--todo`

    Only show the waiting tasks -->
