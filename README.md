# Backend/API Developer Code Challenge

## Overview

In this task, you will develop a commandline interface for working with a standard [todo.txt](http://todotxt.org/) file.  (`todo.txt` is a plaintext format for organizing tasks inspired by David Allen's [Getting Things Done (GTD)](https://en.wikipedia.org/wiki/Getting_Things_Done) methodology).  Since the purpose of this project is primarily to demonstrate your ability to quickly learn and modify an existing codebase, we ask that you use `1set`'s [todotxt](https://github.com/1set/todotxt) library, which you will need to extend to complete the list of requirements.

## Format

Your program should run on a standard \*nix (MacOS Terminal, Linux Shell, etc.) - or Windows PowerShell - commandline.  It should accept input from a user and make appropriate modifications to a file called `todo.txt` which should be stored somewhere "nearby" (i.e. in the same folder as or possibly a subfolder of the folder in which the person evaluating your code would run the program).  You do not need to rely on environment variables or config files to specify the location of this file, as it's really not important to the test.  It's OK to hardcode the path to the `todo.txt` file.

In a \*nix environment, the tester would expect to run your program in the following way:

```bash
prompt> ./todo add {task}
```

## Command List

Your CLI should support the following commands:

1. `ls` - lists all tasks.  It should accept the following optional filter parameters:
    1. `@context` - any string that starts with '@' should be interpreted as a "context tag," and the output should only include tasks with that context.  If more than one "context tag" param is included, the filter should accept any todo that has one or more of the given context tagss and reject any todo lacking all of the specified context tags
    1. `+project` - any string that starts with '+' should be interpreted as a "project tag," and the output should only include tasks associated with that project.  If more than one "project tag" param is included, the filter should accept any todo that has one or more of the given project tags and reject any todo lacking all of the specified project tags
    1. `(A)` - any string that is surrounded with `()` should be interpreted as a "priority" tag, and the output should include only tasks that have the associated priority.  If more than one priority tag is included in the params, the output should include all and only tasks that have *any* of the specified priority tags.  
    1. `tag:` - any string that ends in a `:` should filter out any tasks that do not contain the associated tag.  If more than one such string is included, the application should filter out any tasks that do not include *any* of the specified tags
    1. `>datestring`, `<datestring` - any string that starts with a `>` and is in `YYYY-MM-DD` format (e.g. "2021-01-01") should filter out any tasks that have due dates before the specified time.  Any string that starts with a `<` and is in `YYYY-MM-DD` format (e.g. "2021-01-01") should filter out any tasks that have due dates after the specified time.  If multiple such strings are included, it should respect the instructions of all of them (this may result in an empty list returned).
    1. `|order` Any param that starts with `|` should be interpreted as an ordering parameters.  You can define your own list of strings for the "order" param, but you should support all the orders defined by the `TaskSortByType` type in the codebase (defined [here](https://pkg.go.dev/github.com/1set/todotxt#TaskSortByType))
    1. `completed`  If this string is included, output completed tasks at the end of the list, subject to all the same filters, separated from the other tasks by an empty line
1. `completed` - lists all completed tasks.  It should accept the same filters as `ls`, with the difference that the `<datestring` and `>datestring` filters should apply to the "completed" date rather than the due date.
1. `add {task_string}` - adds a task to the file  
1. `rm {task_id}` - removes a task from the file
1. `do {task_id}` - mark the task specified by the id as completed
1. `tags` - list all *unique* tags in the file.  It is not enough to simply list the tags, you have to ensure that there are no duplicates in the output.
1. `projects` - list all *unique* projects in the file.  It is not enough to simply list the projects, you must ensure that there are no duplicates in the output.
1. `due` - output a list of *unique* due dates followed by the number of tasks due on that date.  At the end of the list, include the number of tasks that are missing a due date.  Example output:
    ```
    2021-03-01 16
    2021-03-02 5
    2021-03-15 1
    26
    ```
1. `extend {task_id} {quantity} {unit}` - Should extend the due date of the task identified by `{task_id}` by `{quantity}` number of `{unit}` time.  For example, `extend 5 1 day` would add one day to the due date on task 5.  `extend 6 -2 week` would shorten the due date of task 6 by two weeks.  So, if task 5 was due on `2021-03-03`, it would be due on `2021-03-04` as a result of the command.  If task 6 were due on `2021-03-03`, it would be due on `2021-02-17` as a result of the command.  If a task has no duedate when this call is made on it, it should behave as though the task were due today.  You should support the following units:
    1. `day`
    1. `week`
    1. `month`
    1. `year`

## Other Requirements

1. By default - i.e. in the absence of any other filters - output should be sorted primarily by priority, secondarily by due date, and tertiarily by created date.  Completed tasks should *always* be listed separately

## Bonus Requirements

1. Add support for `>=` and `<=` as due date filters
1. Add a "help" or "man" page instruction that prints out instructions on how to use the app.
1. For the `due` command, include the number of tasks due on the date that are *not* completed in parentheses.  E.g.:
    ```
    2021-03-01 16 (8)
    2021-03-02 5 (1) 
    2021-03-15 1 (0)
    26 (22)
    ```
1. Include a "TOTAL: {number}" tag at the end of the output of the `ls` command that shows the number of tasks in the output
1. Add terminal colors to the output.  So, for example, all tasks of priority `(A)` might be yellow, the tasks of priority `(B)` might be red, etc.
