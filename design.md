## list
desc:
    lists all uncompleted tasks, sorted in priority order
    tasks with the same priority are sorted alpabetically

output:
    `[id], [priority], [name], ?[desc], ?[marked]`

flags:
    -u lists ALL tasks (even completed)



## add
desc:
    add a new task into the list
    also gives the task an unique id to be indentified by

input: (m=mandatory, o=optional)
    name[str,m], priority[int,m], desc[str,o] marked[bool,o]
    priority = 0(least) - 100(most)
    marked = true | false

flags:
    -p insert everything in the command instead of using stdin



## mark
desc:
    mark a task as completed or uncompleted

input:
    id[int,m]
    

## remove
desc:
    remove a flag completely from the list

input:
    id[int,m]
