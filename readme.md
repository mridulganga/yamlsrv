## yamlsrv
This is a simple read only api service which makes the yaml file available as set of APIs.
The resources can be accessed using the full path.

### Example
```yaml
fruits:
    - apple
    - name: mango
      season: spring
```

### Outputs
```
/fruits - ["apple",{"name":"mango","season":"spring"}]
/fruits/0 - "apple"
/fruits/1 - {"name":"mango","season":"spring"}  
/fruits/1/season - "spring"
/xyz - error 404
```

## Running the server
Runs on port 3000
```
go run main.go
```