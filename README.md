# REST Layer Null backend

This REST Layer resource storage backend doesn't store data.

## Usage

Simply create a resource handler per resource:

```go
import "github.com/omani/rest-layer-null"
```

```go
index.Bind("foo", foo, null.NewHandler(), resource.DefaultConf)
```

## What is this useful for?

If you want to have CRUD operations on something which is should not stored anywhere and kept in memory, this is what you want.
For example, I am using it for having a login Storer as CRUD (Insert mostly) to be able to login users in my rest-layer app.


Author: Hasan Pekdemir (omani)
License: MIT