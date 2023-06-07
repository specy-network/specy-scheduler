# Regchain system manifest



## Deploy


```shell
$ yarn codegen && yarn build
$ yarn create-local && yarn deploy-local

```

For the list of supported networks, see the scripts in the [`package.json`](package.json) file.

## Querying the subgraph

### Query rules
You can query rules by name. 

```
 query{
    rules(where:{name:"ERC20"}){
      name
      content
      hash
      id
    }
  }
```
You also can query all rules.

```
 query{
    rules{
      name
      content
      hash
      id
    }
  }
```
### Query binding
It same like rule query.
```
query{
    bindings(where:{name:"ERC20-BINDING"}){
      name
      content
      hash
      id
    }
  }
```

### Query relation
```
 query{
    relations{
      binding{
        id
      }
      contractAddress
    }
  }
```

## Compliance engine query

You can query binding by contract address when need compliance a event. Use this case you will get binding content and rules.


```
  query{
    relations(where:{contractAddress:"ERC20-BINDING"}){
      binding{
        content
        rules{
          content
        }
      }
    }
  }
```