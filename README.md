# Automatic-Birthday-Wisher

## Overview
**Automatic-Birthday-Wisher** is a command line utility to automatically wish birthdays to one, written in Go.


## Features
- Track birthdays from the `birthdays.xlsx`
- Automatically wish them happy birthday through email
- Could be deployed on a server


## Setup
A `birthdays.xlsx` in given in the repo. It consists of three columns, namely, *Name*, *Date* and *Email*.

The *Date* column must contain date values in `DD-MM` format. For example, if someone has his birthday on 24th January, then his birthday date in the excel file should be `24-01`. 

Fill the `birthdays.xlsx` excel file with data of all the people you would like to wish on their birthdays. You're then good to **GO**!

## Installation
### Using Git
Type the following command in your Git Bash:

- For SSH:
```git clone git@github.com:Shravan-1908/Automatic-Birthday-Wisher.git```
- For HTTPS: ```git clone https://github.com/Shravan-1908/Automatic-Birthday-Wisher.git```

The whole repository would be cloned in the directory you opened the Git Bash in.

### Using GitHub ZIP download
You can alternatively download the repository as a zip file using the GitHub **Download ZIP** feature by clicking [here]().


You can either use the *exe* build of the Automatic-Birthday-Wisher (named as `main.exe`) or if you have Go compiler installed on your system, you can run the program using the command ```go run .\main.go```.