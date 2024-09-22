***

### list

#### desc
- lists all uncompleted tasks
- sorted in priority order
- tasks with the same priority are sorted alpabetically

#### output
`{id}, {priority}, {name}, ?{desc}, ?{marked}`

#### flags
`-a` lists ALL tasks (even completed)

***

### add

#### desc
- add a new task into the list
- gives the task a unique id

#### input (m=mandatory, o=optional)
- name = string
- priority = int, 0-100, (least-most)
- desc? = string
- marked? = bool

#### flags
`-p` insert everything in the command instead of using stdin

***

### mark

#### desc
- mark a task as completed or uncompleted

#### input
- id = int
    

***

### edit

#### desc
- edit a task

#### input
- id = int

***

### remove

#### desc
- remove a task

#### input
- id = int

***
