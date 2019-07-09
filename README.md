# envStore

Store your environment variables from multiple sources (files, strings, JSON and the operating system) in one data structure, allowing easy access (potentially concurrent) from within your program.

## example
```golang
env, err := envStore.New(envStore.Config{
	FromFilePaths: []string{"data1.env", "data2.env"},
	FromStrings: []string{"x=128\ny=test\nz=/bin/bash"},
	FromJSONSlices: [][]byte{[]byte(`{"key":"value","key2":"value2"}`)},
	FromSystem: true,
	UseMutex: true,
	IgnoreEmptyLines: true,
})
if err != nil {
	panic(err)
}

env.Iterate(func(key, value string) {
	fmt.Printf("%v :: %v\n", key, value)
})
```
