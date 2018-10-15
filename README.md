# Rails app demo using React and Go RPC server

## Gems used
* Go server is integrated using the quartz
* React uses react-rails and bootstrap-saas

## How to install
Make the Go server as follows
```sh
cd goserver
./buildit
```

Setup rails as follows
```sh
cd tagdemo
bundle install
```

Setup the test database as follows
```sh
go get github.com/DavidSantia/tag_api
cd $GOPATH/src/github.com/DavidSantia/tag_api/docker
docker build -t tagdemo/mysql ./data
```

## How to run
In one terminal, run the test database
```sh
docker run --name tagdemo-mysql --rm -p 3306:3306 tagdemo/mysql
```

In another terminal, run rails
```sh
cd tagdemo
rails s
```

Go to [localhost:3000](http://localhost:3000)
