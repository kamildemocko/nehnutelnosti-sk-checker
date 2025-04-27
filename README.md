# Nehnuteľnosti.sk Checker

This is a Python script that checks for new real estate listings on the nehnuteľnosti.sk website based on your specified criteria. It can notify you via email when new listings are found.

## Features

*   **Customizable Search:** Define your search criteria (location, price range, property type, etc.) in main.go file, more criteria groups can be specified
*   **Regular Checks:** The script can be scheduled to run at regular intervals (e.g., every hour) to check for new listings.
*   **Email Notifications:** Receive email notifications when new listings matching your criteria are found. (Gmail)
*   **Data Persistence:** The script keeps track of already seen listings to avoid duplicate notifications.

## Configuration

1.  Create a `.env` file in the project directory.
2.  Add required fields:
```
EMAIL_FROM=account@gmail.com
GMAIL_APP_PASSWORD=generated_app_password
EMAIL_TO=receiver@hotmail.com
SUBJECT=your subject!
```

## Running
```
go run .\src\cmd\app\
```

## Build release
Windows
```cmd
set GOOS=windows
go build -o .\dist\nehnutelnosti-checker.exe ./src/cmd/app
```
    
Linux
```bash
set GOOS=linux
go build -o ./dist/nehnutelnosti-checker ./src/cmd/app
```