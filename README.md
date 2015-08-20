# Daniel-Campos-Exercise
Short coding exercise as part of interview process - please use this repo to commit all of your work.

Additional question may be provided via issues to this repo. Good luck and have fun! :)

###Create HTTP Rest API:
1. Use echo or gin for web handler
2. Implement login endpoint with JWT token, and simple middleware that checks header for 'Authorization: Bearer %jwt_token%' in each request. Otherwise return 403 and json struct with error
3. Implement endpoint that will use oAuth2 authorization for FB, to login and issue access_token
3. Log each request including status code
4. Implement persistence with MySQL and Gorm (https://github.com/jinzhu/gorm)
5. Use Goose or other tool of choice for DB migration (https://bitbucket.org/liamstask/goose)
6. Implement save endpoint for Task object
7. Implement update endpoint for Task object
8. Implement get endpoint for Task object
9. Implement delete endpoint for Task object (just update IsDeleted field)  
10. Use CORS (reply with header Access-Control-Allow-Origin: *)
11. Add support for OPTION HTTP method for each endpoints  
12. Configure daemon over simple JSON config. Specify path as process flag for daemon. Required params: ListenAddress, DatabaseUri.
13. Put in comments below description of taken architectural decisions and


###Task:
```
type Task struct {
    Id          int64
    Title       string
    Description string
    Priority    int
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
    CompletedAt *time.Time
    IsDeleted   bool
    IsCompleted bool
}
```

## Architectural Decisions

- The system was separated in packages, so they can change in the future, serving just as wrappers to other services or even becoming individual micro-services (generate the server and client for each one), with much less impact to application. Including the motive to let the API CRUD implementation in the ***tasks*** package was thinking in separating it in the future (and maybe other API packages).
- The ***Auth*** functions stay on main package (***application***) due to simplification purposes, but we also could set a different package for that based on ***tasks*** package.
- My initial plan was to create a different package for the ***Logger*** too, with the same idea to transform it in a wrapper using Go Routine to Log and a micro-service in the future, but I keep using the Gin Logger for simplification purposes.
- A Test Client, a simple HTML page using angularjs has been created for testing purposes.
  - Running it with node server
  - How it works with the authentication
- Tried to use existent golang/community packages for all application features. So we can just update these packages from community when needed.
- Why choosing Gin instead of Echo:
- Names and considerations about Clean Code and OO Calisthenics:



#### Packages Descriptions

- application: Main appliation package
- tasks:
- config:


## Running the Application

- Maybe you need to compile the application before using it. For that, run ``` go install application ```

I used ***rerun*** to compile/restart/run/testing the application automatically at each code update, improving the coding performance.

Run it like: ``` rerun --test application ```

or with the Settings.json path: ``` rerun --test application settings.json ```
