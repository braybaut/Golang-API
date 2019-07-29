## Golang API

this is a minimal API to gettting books information and their Continuous integration workflow with Codebuild and ECR 


***Run Golang API:***

for test the golang api please, run the follow command, this API open the 8080 Port:

```console
go run *.go
```

***Build Docker image***

for build the docker image, run the follow commands:

   ```console 
    docker build -t libraryapi .
    ```

***Make the Codebuild and ECR***

add the variable to terraform.tfvars and run terrafom plan and terraform apply


