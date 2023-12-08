# Redirect Engine 

The RedirectEngine is a plugin for Traefik that adds redirection functionality to your routing configurations. With this plugin, you can easily set up complex redirects, customizing rules and behaviors

## Redirect rules example
```
    {
    "teste1.com.br": {
        "destiny": "teste2.com.br/v1",
        "uri": ["", ""]
    }, 

    "service1.com.br": {
        "destiny": "api.service.com",
        "uri": ["", ""]
    }
```

## Install
