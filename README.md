<div align="center">
<article style="display: flex; flex-direction: column; align-items: center; justify-content: center;">
    <p align="center"><img width="300" src="https://ik.imagekit.io/songlim/diglett.png?updatedAt=1693543872307" /></p>
    <h1 style="width: 100%; text-align: center;">Diglett</h1>
    <p>
        A port check tool written in Golang
    </p>
</article>

![Go][go-badge] ![Version][version-badge]

[go-badge]: https://img.shields.io/badge/Golang-1.20-blue
[version-badge]: https://img.shields.io/badge/release-1.0.0-powderblue

</div>

# 🏔️ What is Diglett?
**Diglett** is a minimalist, port check rest API service written in Golang. It is useful in identifying whether a port has been open/closed to the public network.

# 🌟 Getting Started
## Docker
```bash
docker run -d --name diglett -p 8080:8080 emokid327/diglett:latest
```

# Usage
#### Check port validity
<code>POST /check</code> Check if port is open or closed on specific IP Address</summary>

#### Body
```json
    {
        "address": "1.1.1.1",
        "port": "80",
    }
```


#### Responses

| http code     | content-type                      | response                                                            |
|---------------|-----------------------------------|---------------------------------------------------------------------|
| `200`         | `application/json`                | `{"message":"port 80 open on 1.1.1.1"}`                                             |
| `400`         | `application/json`                | `{"message":"invalid payload"}`                                     |
| `406`         | `application/json`                | `{"message":"65536 is not a valid port/IP address"}`                |

#### Example cURL

```shell
curl -X POST -H "Content-Type: application/json" --data '{"address":"1.1.1.1", "port": "80"}' http://localhost:8080/check
```



[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/songlim)