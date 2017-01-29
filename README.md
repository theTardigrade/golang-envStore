# envStore
## version 1.0.0

Store your environment variables from multiple sources (files, strings and the operating system) in one data structure, allowing easy access (potentially concurrent) from within your program.

### example
    env, err := envStore.New(envStore.Config{
        FromFilePaths: []string{data1.env, data2.env},
        FromStrings: []string{x=128ny=testnz=/bin/bash},
        FromSystem: true,
        UseMutex: false,
    })
    if err != nil {
        panic(err)
    }

    env.Iterate(func(key, value string) {
        fmt.Printf(%v :: %vn, key, value)
    })
