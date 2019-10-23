
Commands for generating client/server code with protoc

#go
```sh
export PATH=$PATH:/Users/{username}/go/bin

protoc --go_out=plugins=grpc:. *.proto 

go install
```

#web gRPC / js
```sh
git clone https://github.com/grpc/grpc-web.git

cd grpc-web

sudo make install-plugin

#or
sudo mv ~/Downloads/protoc-gen-grpc-web-1.0.6-darwin-x86_64 \
  /usr/local/bin/protoc-gen-grpc-web

chmod +x /usr/local/bin/protoc-gen-grpc-web

#then
protoc -I=. config-service.proto \
  --js_out=import_style=commonjs:web \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:web
  ```

add the following line to config-service.pb.go
```js
/* eslint-disable */
```
add dependencies to package.json
```js
    "@grpc/proto-loader": "^0.3.0",
    "google-protobuf": "^3.6.1",
    "grpc": "^1.15.0",
    "grpc-web": "^1.0.0"
```
```sh
npm install
```
example client code
```js
import { DeviceEUI, SetDesiredRequest } from '../../proto/config-service_pb.js';
import { ConfigServiceClient } from '../../proto/config-service_grpc_web_pb.js';

const client = new ConfigServiceClient(process.env.REACT_APP_GRPC_SERVICE,
        null, null);

try {
    var request = new Identifier();
    request.setIdentifier(docid);
    keycloak.updateToken();
    const metadata = { 'authorization': keycloak.token };
    const call = client.getAllConfig(request, metadata, (err, response) => {
        if (err) {
            console.log(err);
            alert(err);
        } else {
            let data = response.getFieldsList();
            //sort alphabetical order
            data.sort((a,b) => {
                if (a.getName() < b.getName()) {
                    return -1;
                }
                return 1;
            })
            setConfigFields(data);
        }
    });
    //this code is necessary to stop the above callback getting called 3 times
    call.on('status', function (status) {
        console.log(status.code);
        console.log(status.details);
    });
} catch (err) {
    console.log(err);
    alert(err);
}
```
  
#python
```sh
pip install grpcio-tools

python -m grpc_tools.protoc -I . --python_out=python --grpc_python_out=python config-service.proto
```