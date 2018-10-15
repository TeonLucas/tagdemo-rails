# Rails app demo using React and Go RPC server

## Gems used
* Go server is integrated using the quartz
* React uses react-rails and bootstrap-saas

## How to install
Setup rails as follows
```sh
git clone https://github.com/DavidSantia/tagdemo-rails
cd tagdemo-rails
bundle install

```
Make the Go server as follows
```sh
./build-goserver.sh
```

## How to run
In one terminal, setup the test database
```sh
go get github.com/DavidSantia/tag_api
cd $GOPATH/src/github.com/DavidSantia/tag_api/docker
docker build -t tagdemo/mysql ./data
```
And run the database container
```sh
docker run --name tagdemo-mysql --rm -p 3306:3306 tagdemo/mysql
```

In another terminal, run rails
```sh
cd tagdemo-rails
rails s
```

Go to [localhost:3000](http://localhost:3000)
