[Preview the Swagger Spec](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/choicehealth/user-service/master/swagger.yaml)

rm -Rf /tmp/app.rsa*
openssl genrsa -out /tmp/app.rsa
openssl rsa -in /tmp/app.rsa -pubout > /tmp/app.rsa.pub


TODO:

setup proper permissions check for all endpoints and improve the pass reset handling

ability to set and change roles


implement context cancelling
testing ???
