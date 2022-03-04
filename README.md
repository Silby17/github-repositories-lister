# GitHub Repositories Lister

List all repositories of a given organization

## Usage

1. Create a GitHub token that allows access to all repositories
2. Compile binary: `GOOS=linux go build -o repo-lister`
3. Set environment variables
    1. `export TOKEN=*******`
    2. `export ORG=my-org-name`
4. Run `./repo-lister`

### Sample Console Output

```shell
INFO Starting...                                  
INFO Initializing GitHub Client                   
INFO Retrieving all repositories...               
INFO 509 repositories retieved                    
INFO Listing Active repositories...               
repo-01
repo-02
repo-03
repo-04
repo-05
...
INFO 504 Active repositories                      
INFO Listing Archived repositories...             
repo-06
repo-07
...
INFO 5 Archived repositories                      
INFO Listing Public repositories...               
repo-02
repo-03
repo-05
...
INFO 20 Public repositories                       
INFO Listing Private repositories...              
repo-01
repo-04
...
INFO 416 Private repositories 
```

## License

[MIT](LICENSE.md) © 2022-present Yossi Silberhaft