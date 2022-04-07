
### Example Usage

```
    import "{{PACKAGE_PATH}}"

    ...

    apiUrl := "<actual url>"
    accessToken := "<actual token"

    cfg := &{{PACKAGE}}.Configuration{
        BasePath:      apiUrl,
        DefaultHeader: make(map[string]string),
        UserAgent:     "hpecli",
    }

    client = {{PACKAGE}}.NewAPIClient(cfg)

    ctx = context.WithValue(gocontext.Background(), {{PACKAGE}}.ContextAccessToken, accessToken)
    myThing, _, err := client.<XYZ>API.<Function>(ctx)
```
