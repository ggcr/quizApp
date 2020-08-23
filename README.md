# Quiz
A quiz app where you pass a .csv file with the problems and it will prompt them to the user with a limit time for each question. At the end it will give you the final score.

This was made as an introductory example program for _go routines_ and _channels_.

## Usage
To run the program follow the steps below:
```bash
git clone https://github.com/1063551/quizApp.git
cd quizApp/
```

To see the valid format of the arguments:
```bash
go run main.go -h
  -csv string
    	valid format is [question], [answer] (default "problems.csv")
  -timer int
    	seconds to answer each question (default 2)
```
And now run the program with your own __problems.csv__ and the __seconds limit for each question__ (in this case, 3):
```bash
go run main.go problems.csv 3
```

## Build your own questions .csv file
As long as the established format is followed, any kind of questions can be asked.

The format is as follows:
```
  __[question]__,__[answer]__
```

Here are some examples:

```
  What color is the sky,Blue
  What's 1 + 2,3
  What year did the WWII end,1945
  10/2,5
```
