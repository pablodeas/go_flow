# GoFlow

Go script to automate git commands: add, commit and push.

## Installation
```bash
go build -o goflow main.go
```

## Usage
```bash
./goflow
```

The script will:
1. Show repository status
2. Ask for files to add (use `.` for all)
3. Request commit message
4. Push to `main` branch (or another of your choice)

## Requirements

- Go installed
- Git configured on system

## Example
```
=== Git Status ===
On branch main...

=== Git Add ===
Enter files to add (use '.' to add all): .

=== Git Commit ===
Enter commit message: Initial update

=== Git Push ===
Push will go to 'main' branch. Enter another branch name (or press Enter): 
Pushing commit to 'main' branch...
```